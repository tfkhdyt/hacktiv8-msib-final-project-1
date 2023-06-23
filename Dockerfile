# build stage
FROM docker.io/golang:1.21rc2-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o executable .

# deploy stage
FROM docker.io/alpine:latest

WORKDIR /

COPY --from=builder /app/executable /app

EXPOSE 8080

ENV GIN_MODE=release

ENTRYPOINT [ "/app" ]
