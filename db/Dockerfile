FROM golang:1.18.3 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o gen ./cmd/gen/
RUN /build/gen -output=seed.sql

FROM postgres:14 AS final

ENV POSTGRES_DB music
ENV POSTGRES_USER music
ENV POSTGRES_PASSWORD music

COPY db/schema.sql /docker-entrypoint-initdb.d/00-schema.sql
COPY --from=builder /build/seed.sql /docker-entrypoint-initdb.d/01-seed.sql