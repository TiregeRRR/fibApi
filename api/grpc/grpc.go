package grpc

import (
	context "context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	f "github.com/TiregeRRR/fibApi/fibonacci"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// GRPCSrv реализует интерфейс FibServer
type GRPCSrv struct {
	UnimplementedFibServer
}

// GetFib возвращает слайс string
func (g *GRPCSrv) GetFib(ctx context.Context, fr *FibRequest) (*FibResponse, error) {
	resp, _, err := f.GetFibSlice(fmt.Sprint(fr.GetX()), fmt.Sprint(fr.GetY()))
	if err != nil {
		return nil, err
	}
	return &FibResponse{FibList: resp}, nil
}

// StartGRPC запускает grpc сервер и REST gateway
func StartGRPC() {
	go func() {
		port := os.Getenv("rest_port")
		if port == "" {
			port = "8000"
		}
		mux := runtime.NewServeMux()
		RegisterFibHandlerServer(context.Background(), mux, &GRPCSrv{})
		log.Printf("Starting REST gateway on %v port\n", port)
		log.Fatalln(http.ListenAndServe(":"+port, mux))
	}()
	port := os.Getenv("grpc_port")
	if port == "" {
		port = "8080"
	}
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
