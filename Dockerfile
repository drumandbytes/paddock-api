FROM golang:1.26-alpine AS builder
RUN apk add --no-cache ca-certificates

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o paddock-api .
RUN mkdir /data

# scratch has no CA bundle or zoneinfo: CAs copied below, tzdata embedded (main.go).
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /app
COPY --from=builder /build/paddock-api /app/paddock-api
COPY static ./static

# Durable cache for finished sessions; non-root-owned so a mounted volume is writable.
COPY --from=builder --chown=65532:65532 /data /data
ENV CACHE_DIR=/data

EXPOSE 4463
USER 65532:65532

ENTRYPOINT ["/app/paddock-api"]
