# Kubernetes setup

setting up the application to run in k8s via minikube

## Running the project in a kubernetes pod

```bash
# root directory
skaffold dev
```

visit `http://127.0.0.1:3000/`

with the application running in kubernetes, we've completed Milestone 3, now we'll look into creating an nginx ingress deployment as the load balancer & http traffic router