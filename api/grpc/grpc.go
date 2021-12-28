package apiGRPC

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/TiregeRRR/fibApi/config"
	f "github.com/TiregeRRR/fibApi/fibonacci"
	"google.golang.org/grpc"
)

// GRPCSrv реализует интерфейс FibServer
type GRPCSrv struct {
	UnimplementedFibServer
}

// GetFib возвращает слайс uint64
func (g *GRPCSrv) GetFib(ctx context.Context, fr *FibRequest) (*FibResponse, error) {
	resp, _, err := f.GetFibSlice(fmt.Sprint(fr.GetX()), fmt.Sprint(fr.GetY()))
	if err != nil {
		return nil, err
	}
	return &FibResponse{FibList: resp}, nil
}

// StartGRPC запускает grpc сервер на 8080 порту
func StartGRPC() {
	cfg := config.GetConfig()
	port := cfg.GetString("grpc_port")
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	RegisterFibServer(grpcServer, &GRPCSrv{})
	log.Printf("Starting GRPC server on %v port\n", port)
	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatalln(err)
	}
}
