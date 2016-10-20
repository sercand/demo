package main

import (
	"encoding/json"
	pb "feedpb"
	"flag"
	"flagutil"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sort"
	"time"

	"github.com/sercand/kuberesolver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	port          = flag.Int("port", 3010, "The server port")
	configFile    = flag.String("config-file", "config.json", "A json file for configuration")
	colorEndpoint = flag.String("color", "localhost:4000", "ColorGenerator endpoint")
)

type Items []*pb.FeedItem

func (s Items) Len() int {
	return len(s)
}
func (s Items) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Items) Less(i, j int) bool {
	return s[j].Score < s[i].Score
}

type serverConfig struct {
	Providers    []string
	DefaultLimit int
}

type newsFeedServer struct {
	config   serverConfig
	clients  []pb.FeedProviderClient
	balancer *kuberesolver.Balancer
	color    string
	name     string
}

func getFromRemote(ctx context.Context, in *pb.FeedGetRequest, c pb.FeedProviderClient, resChan chan<- []*pb.FeedItem) {
	if resp, err := c.Get(ctx, &pb.ProviderGetRequest{Request: in}); err != nil {
		grpclog.Printf("failed to get items from remote, %v", err)
		resChan <- nil
	} else {
		resChan <- resp.Items
	}
}

//Implement feedpb.NewsFeedServer
func (s *newsFeedServer) Get(ctx context.Context, in *pb.FeedGetRequest) (*pb.FeedGetResponse, error) {
	grpclog.Printf("newsFeedServer.Get(_):%v", in)
	r := new(pb.FeedGetResponse)
	r.Items = append(r.Items, &pb.FeedItem{Text: "feed-aggregator", Title: s.name, ProviderName: s.name, Score: 1000, Color: s.color})
	resChan := make(chan []*pb.FeedItem, len(s.clients))

	if in.Limit == 0 {
		in.Limit = int32(s.config.DefaultLimit)
	}
	ct, _ := context.WithTimeout(ctx, time.Second)
	for _, c := range s.clients {
		go getFromRemote(ct, in, c, resChan)
	}

	for i := 0; i < len(s.clients); i++ {
		iis := <-resChan
		if len(iis) > 0 {
			r.Items = append(r.Items, iis...)
		}
	}
	fmt.Println(r.Items)
	sort.Sort(Items(r.Items))
	close(resChan)
	if len(r.Items) > int(in.Limit) {
		r.Items = r.Items[:in.Limit]
	}
	return r, nil
}

func (s *newsFeedServer) loadConfig(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		grpclog.Fatalf("failed to load config: %v", err)
	}
	if err := json.Unmarshal(file, &s.config); err != nil {
		grpclog.Fatalf("failed to unmarshal config file: %v", err)
	}
}

func (s *newsFeedServer) connectProviders() {
	for _, p := range s.config.Providers {
		conn, err := s.balancer.Dial(p, grpc.WithInsecure())
		if err != nil {
			grpclog.Fatalf("failed to dial to %s: %v", p, err)
		}
		s.clients = append(s.clients, pb.NewFeedProviderClient(conn))
	}
}

func newServer() *newsFeedServer {
	s := &newsFeedServer{
		balancer: kuberesolver.New(),
		clients:  make([]pb.FeedProviderClient, 0),
		name:     os.Getenv("HOSTNAME"),
	}
	//get color from color generator server
	colorConn, err := s.balancer.Dial(*colorEndpoint, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}
	defer colorConn.Close()
	cg := pb.NewColorGeneratorClient(colorConn)
	if c, err := cg.Next(context.TODO(), &pb.NewColorRequest{}); err == nil {
		s.color = c.Hex
	} else {
		grpclog.Fatalf("failed to get color: %v", err)
	}
	s.loadConfig(*configFile)
	s.connectProviders()
	return s
}

func main() {
	flag.Parse()
	if err := flagutil.SetFlagsFromEnv(flag.CommandLine, "DEMO"); err != nil {
		grpclog.Fatal(err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNewsFeedServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Fatal(err)
	}
}
