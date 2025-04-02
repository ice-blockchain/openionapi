# 🛠 Build stage
FROM golang:1.23.0-bookworm AS gobuild

# Install build dependencies
RUN apt-get update && apt-get install -y \
    libssl-dev libsecp256k1-dev libsodium-dev gcc g++ make \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /build

# Go module cache layer
COPY go.mod go.sum ./
RUN go mod download

# Copy full source
COPY . .

# Build the Go API binary
RUN go build -o /tmp/opentonapi ./cmd/api

# Copy OpenAPI specs
RUN mkdir -p /tmp/openapi && \
    cp ./api/openapi.* /tmp/openapi/


# 🚀 Runtime stage
FROM ubuntu:22.04 AS runner

# Install runtime deps + nginx
RUN apt-get update && apt-get install -y \
    wget openssl ca-certificates libsecp256k1-dev libsodium-dev nginx \
 && rm -rf /var/lib/apt/lists/*

# 🧠 Emulator lib (runtime native dependency)
RUN mkdir -p /app/lib && \
    wget -q -O /app/lib/libemulator.so https://github.com/ton-blockchain/ton/releases/download/v2024.08/libemulator-linux-x86_64.so

# Setup dynamic linker
ENV LD_LIBRARY_PATH=/app/lib

# Copy Go API binary and OpenAPI
COPY --from=gobuild /tmp/opentonapi /usr/bin/
COPY --from=gobuild /tmp/openapi /app/openapi

# Copy NGINX config
COPY nginx.conf /etc/nginx/nginx.conf

# Expose ONLY NGINX externally
EXPOSE 8090

# Start both Go app and NGINX
CMD /usr/bin/opentonapi & nginx -g "daemon off;"