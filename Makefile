SHELL := /bin/bash
SERVER_OUT := "cmd/freelance/server"
API_OUT := "api/api.pb.go"
API_REST_OUT := "api/api.pb.gw.go"
SERVER_PKG_BUILD := "server/server.go""
GOPATH=$(shell go env GOPATH)

.PHONY: all api server

api/api.pb.go: 
	protoc -I/usr/local/include -I api/ \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --go_out=google/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:api \
    api/api.proto

api/api.pb.gw.go:
	protoc -I/usr/local/include -I api/ \
    -I${GOPATH}/src \
    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:api \
    api/api.proto

api: api/api.pb.go api/api.pb.gw.go

clean:
	rm $(SERVER_OUT) $(API_OUT) $(API_REST_OUT)

mysqlenv:
	export DATABASE_URL="dbuser_f:dbpass_f@tcp(localhost:3306)/freelance?multiStatements=true"

pgsqlenv:
	export DATABASE_URL="postgres://dbuser_f:dbpass_f@localhost:5432/freelance?query"

