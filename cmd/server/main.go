package main

import (
	"context"
	"log"
	"net/http"
	"time"

	v1 "tzdanows/flock-nginx-mvp/gen/go/proto/v1"
	"tzdanows/flock-nginx-mvp/gen/go/proto/v1/v1connect"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type PingServer struct{}

func (s *PingServer) Ping(
	ctx context.Context,
	req *connect.Request[v1.PingRequest],
) (*connect.Response[v1.PingResponse], error) {
	log.Printf("Received ping: %s", req.Msg.Message)

	return connect.NewResponse(&v1.PingResponse{
		Message:   "Server received: " + req.Msg.Message,
		Timestamp: time.Now().Unix(),
	}), nil
}

func main() {
	server := &PingServer{}
	path, handler := v1connect.NewPingServiceHandler(server)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	log.Printf("Starting server on :8080...")
	if err := http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
