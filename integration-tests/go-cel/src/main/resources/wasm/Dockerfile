# Dockerfile to build CEL WASM module from Go source
#
# Build:
#   docker build -t go-cel-builder .

FROM golang:1.24-alpine AS builder

WORKDIR /build

# Copy dependency files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source
COPY main.go ./

# Build WASM (WASI target) with size optimization flags
RUN GOOS=wasip1 GOARCH=wasm go build -ldflags="-s -w" -o go-cel.wasm main.go

# Final stage: minimal alpine image with just the WASM file
FROM alpine:3.19

COPY --from=builder /build/go-cel.wasm /wasm/go-cel.wasm

# Output WASM file to stdout when container runs
ENTRYPOINT ["cat", "/wasm/go-cel.wasm"]
