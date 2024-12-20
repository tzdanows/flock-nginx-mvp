# Go ConnectRPC Ping Service

Milestone one involves creating a simple ConnectRPC client-server

## Running the backend locally

```bash
cd backend
go run cmd/server/main.go
# in another terminal
go run cmd/client/main.go
```

You should see a response like:
```
Server Response: Server received: Hello from client!
Server Timestamp: 2024-12-20 14:30:45 PST
```

## Testing with curl

```bash
curl \
  --header "Content-Type: application/json" \
  --header "Connect-Protocol-Version: 1" \
  --data '{"message": "Hello from curl"}' \
  http://localhost:8080/v1.PingService/Ping
```

Expected response:
```json
{
  "message": "Server received: Hello from curl",
  "timestamp": "1703116800"
}
```

This will have a different purpose later, but this completes Milestone 1