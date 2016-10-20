package main

import (
	pb "feedpb"
	"flag"
	"flagutil"
	"fmt"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	port       = flag.Int("port", 4000, "The server port")
	colorIndex = 0
	mu         = sync.Mutex{}
	colors     = []string{
		"#F44336",
		"#E91E63",
		"#9C27B0",
		"#673AB7",
		"#3F51B5",
		"#2196F3",
		"#03A9F4",
		"#00BCD4",
		"#009688",
		"#4CAF50",
		"#8BC34A",
		"#CDDC39",
		"#FFEB3B",
		"#FFC107",
		"#FF9800",
		"#FF5722",
		"#795548",
		"#9E9E9E",
		"#607D8B",
	}
)

type colorServer struct {
}

func (s *colorServer) Next(ctx context.Context, in *pb.NewColorRequest) (*pb.Color, error) {
	mu.Lock()
	ci := colorIndex % len(colors)
	colorIndex += 1
	mu.Unlock()
	color := colors[ci]
	grpclog.Printf("colorServer.Next: %s", color)
	return &pb.Color{Hex: color}, nil
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
	pb.RegisterColorGeneratorServer(grpcServer, &colorServer{})
	grpclog.Printf("serving on :%d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		grpclog.Fatal(err)
	}
}
