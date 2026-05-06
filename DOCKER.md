## Dockerfile Analysis

The Dockerfile uses a multi-stage build with two stages.

### Stage 1: Builder Stage

```dockerfile
FROM golang:1.26-alpine AS builder
```

The first stage uses a Go image based on Alpine Linux. This stage is only used to build the application.

It sets the working directory to `/app`, copies `go.mod` and `go.sum`, downloads the dependencies and then copies the full source code.

```dockerfile
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-server ./cmd/api
```

This command compiles the Go application and creates the binary `/api-server`.

### Stage 2: Runtime Stage

```dockerfile
FROM alpine:3.19
```

The second stage uses a small Alpine Linux image. It only contains what is needed to run the compiled application.

```dockerfile
COPY --from=builder /api-server .
ENTRYPOINT ["./api-server"]
```

The container exposes port `8080` and starts the API server.

## What does CGO_ENABLED=0 do?

`CGO_ENABLED=0` disables CGO.

CGO allows Go programs to call C code, but this can require additional system libraries inside the container.

By disabling CGO, Go builds a statically linked binary. This is important because the final runtime image can stay small and does not need the full Go toolchain or extra C libraries.

## Final Image Size vs Single-Stage Build

A single-stage Docker build would include the Go compiler, source code and build dependencies in the final image.

The multi-stage build avoids this. The final image only contains:

- Alpine Linux
- CA certificates
- the compiled API binary

This makes the final image smaller, cleaner and more suitable for running the application.

## CRUD Test Results

Screenshots of the results are included in the submission PDF.

### Create Products

Three products were successfully created using curl:

- Laptop
- Mouse
- Keyboard

Example:

```bash
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d "{\"name\":\"Laptop\",\"price\":999.99}"
```

### Read Products

Products were successfully retrieved using:

```bash
curl http://localhost:8080/products
```

### Update Product

A product was successfully updated using:

```bash
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d "{\"name\":\"Gaming Laptop\",\"price\":1299.99}"
```

### Delete Product

A product was successfully deleted using:

```bash
curl -X DELETE http://localhost:8080/products/2
```

## Persistence Test

The containers were restarted with:

```bash
docker compose down
docker compose up
```

The products still existed after restarting the containers, confirming PostgreSQL data persistence.