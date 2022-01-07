create-proto:
	protoc --proto_path=proto proto/*.proto --go_out=api/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=api/
	protoc -I . --grpc-gateway_out api/ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt generate_unbound_methods=true \
    proto/fib.proto

clear-proto:
	rm api/grpc/fib_grpc.pb.go
	rm api/grpc/fib.pb.go

run-tests:
	go test api/grpc/*.go
	go test fibonacci/*.go


ENVFILE = ./config/conf.env
run-docker:
	docker-compose --env-file $(ENVFILE)  up --build