FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache bash postgresql-client

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY migrations/ ./migrations/
COPY wait-for-it.sh ./
COPY init.sh ./
RUN chmod +x wait-for-it.sh init.sh

RUN go build -o server ./cmd/server
RUN go build -o seed ./cmd/seed

EXPOSE 8080
CMD ["./init.sh"]
