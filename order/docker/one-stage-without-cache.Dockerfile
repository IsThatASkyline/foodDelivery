FROM golang:1.25.5-alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" \
    -o app ./cmd

ENTRYPOINT ["/app"]
