FROM tinygo/tinygo:0.34.0

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN tinygo build -o qr-generator.wasm -target=wasi main.go

ENTRYPOINT ["cat", "qr-generator.wasm"]
