FROM golang:1.20.4 AS builder

WORKDIR /build
RUN adduser -u 10001 -disabled-password -gecos '' music-api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o music-api cmd/api/main.go

FROM alpine:3.17 AS final

WORKDIR /app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/music-api /app/music-api
COPY config /app/config

EXPOSE 3000

USER music-api
ENTRYPOINT ["/app/music-api"]
