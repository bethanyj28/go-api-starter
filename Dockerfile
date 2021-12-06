FROM golang:1.16-alpine AS builder

WORKDIR /go/src/github.com/bethanyj28/go-api-starter
COPY . /go/src/github.com/bethanyj28/go-api-starter
RUN go mod vendor && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/bethanyj28/go-api-starter/app .
COPY --from=builder /go/src/github.com/bethanyj28/go-api-starter/internal/db/migrations ./migrations

# a bit of magic to 'wait' for dependencies to be available for testing
ENV WAIT_VERSION 2.7.3
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

ENTRYPOINT /wait && ./app
