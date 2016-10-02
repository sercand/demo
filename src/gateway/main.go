package main

import (
	gw "feedpb"
	"flag"
	"fmt"

	"github.com/codegangsta/negroni"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sercand/kuberesolver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	endpoint = flag.String("newsfeed-endpoint", "localhost:3010", "endpoint of NewsFeedService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	balancer := kuberesolver.New()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if balancer.IsInCluster() {
		opts = append(opts, balancer.DialOption())
	}
	if err := gw.RegisterNewsFeedHandlerFromEndpoint(ctx, mux, *endpoint, opts); err != nil {
		return err
	}

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(mux)

	grpclog.Printf("binding :%d for server", 8080)
	n.Run(fmt.Sprintf(":%d", 8080))
	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
