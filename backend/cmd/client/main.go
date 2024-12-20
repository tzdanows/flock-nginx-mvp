package main

import (
	"context"
	"log"
	"net/http"
	"time"

	v1 "tzdanows/flock-nginx-mvp/gen/go/proto/v1"
	"tzdanows/flock-nginx-mvp/gen/go/proto/v1/v1connect"

	"github.com/bufbuild/connect-go"
)

func main() {
	client := v1connect.NewPingServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := connect.NewRequest(&v1.PingRequest{
		Message: "Client says hello!",
	})

	res, err := client.Ping(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server Response: %s", res.Msg.Message)
	log.Printf("Server Timestamp: %s", time.Unix(res.Msg.Timestamp, 0).Format(time.RFC3339))
}
