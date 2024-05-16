FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o ./out/go-url-backend ./cmd/webserver

CMD ["./out/go-url-backend"]
