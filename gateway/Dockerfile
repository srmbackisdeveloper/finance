FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build -o gateway-service .

EXPOSE 8080

CMD ["./gateway-service"]
