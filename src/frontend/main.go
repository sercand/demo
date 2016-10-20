package main

import (
	pb "feedpb"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"flagutil"

	"log"

	"github.com/codegangsta/negroni"
	"github.com/sercand/kuberesolver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	port = flag.Int("port", 3030, "The server port")
	endpoint = flag.String("endpoint", "localhost:3010", "NewsFeedService endpoint")
	colorEndpoint = flag.String("color", "localhost:4000", "ColorGenerator endpoint")
	templateFile = flag.String("template", "index.html", "html file template")
	certFile = flag.String("tls-cert", "", "tls cert file")
	keyFile = flag.String("tls-key", "", "tls key file")
)

type server struct {
	client pb.NewsFeedClient
	htpl   *template.Template
	name   string
	color  string
}

func newServer() *server {
	s := new(server)
	s.name = os.Getenv("HOSTNAME")
	balancer := kuberesolver.New()
	conn, err := balancer.Dial(*endpoint, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("failed to dial: %v", err)
	}
	s.client = pb.NewNewsFeedClient(conn)
	if s.htpl, err = template.ParseFiles(*templateFile); err != nil {
		grpclog.Fatalf("failed pase index.html: %v", err)
	}
	//get color from color generator server
	colorConn, err := balancer.Dial(*colorEndpoint, grpc.WithInsecure())
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
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lmt := 0
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		var err error
		if lmt, err = strconv.Atoi(limit); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"BadRequest"}`))
			return
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 2)
	resp, err := s.client.Get(ctx, &pb.FeedGetRequest{Limit: int32(lmt)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error":"%v"}`, err.Error())
		return
	}
	grpclog.Printf("[frontend] get %d items from remote", len(resp.Items))
	data := &pb.FeedGetResponse{Items: make([]*pb.FeedItem, 1)}
	data.Items[0] = &pb.FeedItem{Title: "frontend", Text: s.name, ProviderName: s.name, Color: s.color}
	data.Items = append(data.Items, resp.Items...)
	if err := s.htpl.Execute(w, data); err != nil {
		grpclog.Printf("failed to render template:%v", err)
	}
}
func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://demo.sercand.com" + r.RequestURI, http.StatusMovedPermanently)
}
func main() {
	flag.Parse()
	if err := flagutil.SetFlagsFromEnv(flag.CommandLine, "DEMO"); err != nil {
		grpclog.Fatal(err)
	}
	n := negroni.Classic()
	n.UseHandler(newServer())
	addr := fmt.Sprintf(":%d", *port)
	if *certFile != "" && *keyFile != "" {
		l := log.New(os.Stdout, "[negroni] ", 0)
		l.Printf("listening on %s", addr)
		go http.ListenAndServe(":8080", redirectToHttps)
		l.Fatal(http.ListenAndServeTLS(addr, *certFile, *keyFile, n))
	} else {
		n.Run(addr)
	}
}
