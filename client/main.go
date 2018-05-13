package main

import (
	"log"
	"net"
	"time"

	"github.com/dustin-decker/grpc-firewall-bypass/api"
	"google.golang.org/grpc"
)

// TCP client and GRPC server

// custom net.Listener allows gRPC to serve over TCP connection initiated from the server
type listener struct {
	conn net.Conn
}

func (l listener) Accept() (net.Conn, error) {
	log.Println("connecting with TCP client...")
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8081", time.Second*5)
	l.conn = conn
	return conn, err
}

func (l listener) Close() error {
	err := l.conn.Close()
	log.Println("closed connection")
	return err
}

func (l listener) Addr() net.Addr {
	return &net.TCPAddr{IP: net.IP{0, 0, 0, 0}, Port: 7777}
}

func main() {
	lis := listener{}

	// create a server instance
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the gRPC erver
	log.Println("launching gRPC server over TCP connection...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
