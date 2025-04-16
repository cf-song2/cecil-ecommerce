FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache bash postgresql-client

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY migrations ./migrations
COPY wait-for-it.sh init.sh ./
RUN chmod +x wait-for-it.sh init.sh
RUN go build -o server ./cmd/server
RUN go build -o seed ./cmd/seed

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache bash postgresql-client

COPY --from=builder /app/server .
COPY --from=builder /app/seed .
COPY --from=builder /app/init.sh .
COPY --from=builder /app/wait-for-it.sh .
COPY --from=builder /app/migrations ./migrations

RUN chmod +x init.sh wait-for-it.sh

EXPOSE 8080
CMD ["./init.sh"]
