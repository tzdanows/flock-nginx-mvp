# NextJS Frontend

a basic nextjs app and dockerfiles for the backend/frontend

![pingService](pingService.png)

## Running the project locally:

```bash
cd backend
go run cmd/server/main.go

# in another terminal
cd frontend
npm run dev
```

navigate to `localhost:3000`

## Running the project with docker:

```bash
# from the root directory
docker-compose -f docker-compose.yaml up backend

# in another terminal
docker-compose -f docker-compose.yaml up frontend
```

navigate to `localhost:3000`

This completes Milestone 2 and the overall product, now to incorporate k8s and ingress
