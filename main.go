package main

import (
	g "github.com/TiregeRRR/fibApi/api/grpc"
	r "github.com/TiregeRRR/fibApi/api/rest"
)

func main() {
	go r.StartREST()
	g.StartGRPC()
}
