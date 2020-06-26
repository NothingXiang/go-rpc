# go-rpc
go-rpc-learning

# step

1. generate grpc and grpc-gateway code
```shell script
    protoc -I=$PATH\include -I=. --grpc-gateway_out=.  --go_out=plugins=grpc:. api/*.proto
```
2. run demo code
```shell script
    go run main.go
```