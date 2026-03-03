# Fine Art Gallery (FAG)

A microservices-based e-commerce backend built in Go, demonstrating GraphQL APIs, gRPC inter-service communication, and Kubernetes-native deployment.

---

## Architecture Overview

```
Client
  │
  ▼
order-gateway          ← HTTP/GraphQL entry point
  │
  ├──► order-rpc-service      (gRPC)
  ├──► payment-rpc-service    (gRPC)
  └──► shipping-rpc-service   (gRPC)

order-graphql-stream   ← GraphQL subscription / streaming layer
```

The system follows a **gateway + backend-for-frontend** pattern. The `order-gateway` exposes a public-facing API and fans out to downstream gRPC services. Each service owns its own domain logic and is independently deployable.

---

## Services

| Service | Description |
|---|---|
| `order-gateway` | HTTP API gateway — routes and orchestrates client requests |
| `order-graphql-stream` | GraphQL endpoint with real-time streaming support |
| `order-rpc-service` | Manages order lifecycle via gRPC |
| `payment-rpc-service` | Handles payment processing via gRPC |
| `shipping-rpc-service` | Manages shipping and fulfillment via gRPC |

---

## Tech Stack

- **Language**: Go (standard library-first, Go workspace with `go.work`)
- **API layer**: GraphQL + HTTP
- **Inter-service**: gRPC / Protocol Buffers
- **Containerisation**: Docker, Docker Compose
- **Orchestration**: Kubernetes with Helm charts (`charts/`)
- **Observability**: Kiali (service mesh visualisation)

---

## Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- `make`

### Run locally with Docker Compose

```bash
git clone https://github.com/sklrsn/FAG.git
cd FAG
docker-compose up --build
```

This starts all services with their dependencies wired together.

### Build all services

```bash
make build
```

### Run tests

```bash
make test
```

---

## Kubernetes Deployment

Helm charts are located under `charts/`. Deploy to a running cluster:

```bash
helm install fag ./charts
```

To visualise the service mesh topology, apply the Kiali configuration:

```bash
kubectl apply -f Kiali.yaml
```

---

## Project Structure

```
FAG/
├── order-gateway/          # HTTP API gateway
├── order-graphql-stream/   # GraphQL + streaming
├── order-rpc-service/      # Order domain (gRPC)
├── payment-rpc-service/    # Payment domain (gRPC)
├── shipping-rpc-service/   # Shipping domain (gRPC)
├── charts/                 # Helm charts for Kubernetes
├── Images/                 # Architecture diagrams
├── docker-compose.yaml     # Local development stack
├── Kiali.yaml              # Service mesh observability config
├── Makefile                # Build / test targets
└── go.work                 # Go workspace definition
```

---

## Development Notes

Each service is an independent Go module within the workspace. Changes to shared types are propagated via the `go.work` workspace without publishing intermediate versions.

See [`notes.md`](notes.md) for design decisions and implementation notes.

---

## License

MIT
