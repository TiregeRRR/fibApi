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



update-modfile:
	go mod download
	go mod tidy

ENVFILE = ./config/conf.env
run-docker:
	sudo docker-compose --env-file $(ENVFILE) up 

build-and-run-docker:
	sudo docker-compose --env-file $(ENVFILE) up --build

run-tests:
	sudo docker run --rm --name redis-test-instance -p 6379:6379 -d redis
	go test -v fibonacci/*.go
	go test -v api/grpc/*.go
	sudo docker stop redis-test-instance