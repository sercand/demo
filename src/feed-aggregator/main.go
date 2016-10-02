package main

import (
	"encoding/json"
	pb "feedpb"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/sercand/kuberesolver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	port       = flag.Int("port", 3010, "The server port")
	configFile = flag.String("config-file", "config.json", "A json file for configuration")
)

type serverConfig struct {
	Providers []string
}

type newsFeedServer struct {
	config   serverConfig
	clients  []pb.FeedProviderClient
	balancer *kuberesolver.Balancer
}
//Implement feedpb.NewsFeedServer
func (s *newsFeedServer) Get(ctx context.Context, in *pb.FeedGetRequest) (*pb.FeedGetResponse, error) {
	r := new(pb.FeedGetResponse)
	for _, c := range s.clients {
		resp, err := c.Get(ctx, &pb.ProviderGetRequest{Request: in})
		if err != nil {
			return nil, err
		}
		r.Items = append(r.Items, resp.Items...)
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
			grpclog.Fatalf("fail to dial to %s: %v", p, err)
		}
		s.clients = append(s.clients, pb.NewFeedProviderClient(conn))
	}
}

func newServer() *newsFeedServer {
	s := &newsFeedServer{
		balancer: kuberesolver.New(),
		clients:  make([]pb.FeedProviderClient, 0),
	}
	s.loadConfig(*configFile)
	s.connectProviders()
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNewsFeedServer(grpcServer, newServer())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
