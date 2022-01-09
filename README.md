# My gRPC sample app

If you update `echo.proto`, please execute the following to re-generate code:

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    echo.proto
```

## Run server

```sh
cd server
go run main.go
```
