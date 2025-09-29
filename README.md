# Spotify Microservice (Go + Clean Architecture)

## Run locally (Docker)
```bash
make up
# open http://localhost:8080/api/tracks
```

## Endpoints
- GET /api/health
- GET /api/tracks
- GET /api/tracks/{id}

## Project Structure
- cmd/server: HTTP bootstrap
- internal/domain: entities
- internal/usecase: business cases
- internal/adapter/http/handler: HTTP handlers (BFF-like aggregation)
- internal/infra/cache: in-memory cache
- internal/infra/repo: in-memory repository
- loadtest: vegeta targets and instructions

## AWS ECS (Fargate) sketch
1. docker build -t <acct>.dkr.ecr.<region>.amazonaws.com/spotify-ms-clean:latest .
2. docker push <acct>.dkr.ecr.<region>.amazonaws.com/spotify-ms-clean:latest
3. Create ECS service (desiredCount=3) behind ALB, target path /api/*
4. Enable auto-scaling by CPU ~50% and health checks on /api/health
