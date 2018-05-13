# grpc-firewall-bypass

Example of connecting to gRPC servers that live behind a firewall.

The solution is to establish a TCP connection from the endpoint behind the firewall to a public endpoint, and then dial the gRPC server behind the firewall over that TCP connection from the publicly accessible endpoint.

### Components

#### client

The client is the TCP client and gRPC server that lives behind the firwall.

The TCP client dials the publically accessable TCP server.
Its gRPC server listens on that TCP connection.

#### server 

The server is the TCP server and the gRPC client that lives on a publically accessable server.

The TCP server listens to incoming connections from the TCP clients behind the firwalls.
Once there is an established incoming TCP connection, the gRPC client dials the the gRPC server that is listening on that connection.

### Security

It is expected that you would use mutually authenticated TLS (mTLS) either on the TCP or gRPC layers.

### Run it

#### Deps

`dep ensure`

#### Build

```
go build -o bin/client client/main.go
go build -o bin/server server/main.go
```

#### Run

```
./bin/server
```

```
./bin/client
```

#### Generate API code from proto definitions

The proto go code is already generated for you, but if you make changes to the proto definition, use this to generate new code:

```
protoc -I api/ \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:api \
    api/api.proto
```

### TODO

- figure out why it repeats
- add mTLS