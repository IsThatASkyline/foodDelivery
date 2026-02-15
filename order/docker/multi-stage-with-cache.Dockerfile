# --------- Stage 1: Build ----------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
    -o app ./cmd

# --------- Stage 2: Runtime ----------
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/app /app
COPY --from=builder /app/migrations /migrations

USER 1001

ENTRYPOINT ["/app"]
