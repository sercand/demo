package main

import (
	gw "feedpb"
	"flag"
	"flagutil"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sercand/kuberesolver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	port     = flag.Int("port", 8080, "The server port")
	endpoint = flag.String("newsfeed-endpoint", "localhost:3010", "endpoint of NewsFeedService")
	certFile = flag.String("tls-cert", "", "tls cert file")
	keyFile  = flag.String("tls-key", "", "tls key file")
)

func run() {
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
		grpclog.Fatal(err)
	}

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(mux)

	addr := fmt.Sprintf(":%d", *port)
	if *certFile != "" && *keyFile != "" {
		l := log.New(os.Stdout, "[negroni] ", 0)
		l.Printf("listening on %s", addr)
		l.Fatal(http.ListenAndServeTLS(addr, *certFile, *keyFile, n))
	} else {
		n.Run(addr)
	}
}

func main() {
	flag.Parse()
	if err := flagutil.SetFlagsFromEnv(flag.CommandLine, "DEMO"); err != nil {
		grpclog.Fatal(err)
	}
	run()
}
