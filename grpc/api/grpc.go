package apiGRPC

import (
	context "context"
	"fmt"
	"log"
	"net"

	f "github.com/TiregeRRR/fibApi/fibonacci"

	"google.golang.org/grpc"
)

type GrpcSrv struct {
	UnimplementedFibServer
}

func (g *GrpcSrv) GetFib(ctx context.Context, fr *FibRequest) (*FibResponse, error) {
	resp, err := f.GetFibSlice(fmt.Sprint(fr.GetX()), fmt.Sprint(fr.GetY()))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &FibResponse{FibList: resp}, nil
}

func StartGRPC() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	RegisterFibServer(grpcServer, &GrpcSrv{})
	log.Println("Starting server")
	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatalln(err)
	}
}
