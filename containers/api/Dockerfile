FROM golang:1.21 AS builder

ARG SWAGGER_UI_VERSION=5.10.0
ARG SOUNDFONT_VERSION=3.1

WORKDIR /build
RUN adduser -u 10001 -disabled-password -gecos '' music-api

RUN wget https://github.com/swagger-api/swagger-ui/archive/refs/tags/v${SWAGGER_UI_VERSION}.tar.gz -O swagger-ui-${SWAGGER_UI_VERSION}.tar.gz
RUN mkdir -p dist
RUN tar -C dist --strip-components=2 -xzvf swagger-ui-${SWAGGER_UI_VERSION}.tar.gz swagger-ui-${SWAGGER_UI_VERSION}/dist
RUN sed -i 's|https://petstore.swagger.io/v2/swagger.json|./swagger.json|g' dist/swagger-initializer.js
RUN wget https://github.com/jenskutilek/free-fonts/raw/master/Droid/Droid%20Sans%20Fallback/DroidSansFallback.ttf
RUN wget https://deb.debian.org/debian/pool/main/f/fluid-soundfont/fluid-soundfont_${SOUNDFONT_VERSION}.orig.tar.gz
RUN tar xzf fluid-soundfont_${SOUNDFONT_VERSION}.orig.tar.gz
RUN cp fluid-soundfont-${SOUNDFONT_VERSION}/FluidR3_GM.sf2 .
RUN chmod 644 FluidR3_GM.sf2

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
COPY --from=builder /build/DroidSansFallback.ttf /app/DroidSansFallback.ttf
COPY --from=builder /build/FluidR3_GM.sf2 /app/soundfonts/FluidR3_GM.sf2
COPY config /app/config
COPY docs/swagger.json /app/docs/dist

EXPOSE 3000

USER music-api
ENTRYPOINT ["/app/music-api"]
