FROM golang:1.21 AS builder

WORKDIR /build
RUN adduser -u 10001 -disabled-password -gecos '' music-api

RUN wget https://github.com/swagger-api/swagger-ui/archive/refs/tags/v5.9.1.tar.gz -O swagger-ui-5.9.1.tar.gz
RUN mkdir -p dist
RUN tar -C dist --strip-components=2 -xzvf swagger-ui-5.9.1.tar.gz swagger-ui-5.9.1/dist
RUN sed -i 's|https://petstore.swagger.io/v2/swagger.json|./swagger.json|g' dist/swagger-initializer.js

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o music-api cmd/api/main.go


FROM alpine:3.18 AS final

WORKDIR /app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/music-api /app/music-api
COPY --from=builder /build/dist /app/docs/dist
COPY config /app/config
COPY docs/swagger.json /app/docs/dist

EXPOSE 3000

USER music-api
ENTRYPOINT ["/app/music-api"]
