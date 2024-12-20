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
    return connect.NewResponse(&v1.PingResponse{
        Message: "hello, I am a project using nginx",
        Timestamp: time.Now().Unix(),
    }), nil
}

func main() {
    server := &PingServer{}
    path, handler := v1connect.NewPingServiceHandler(server)
    
    // Create a new mux
    mux := http.NewServeMux()
    
    // Add CORS middleware
    corsHandler := func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, Connect-Protocol-Version")
            
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }
            
            h.ServeHTTP(w, r)
        })
    }
    
    mux.Handle(path, corsHandler(handler))
    
    log.Printf("Starting server on :8080...")
    if err := http.ListenAndServe(
        ":8080",
        h2c.NewHandler(mux, &http2.Server{}),
    ); err != nil {
        log.Fatal(err)
    }
}