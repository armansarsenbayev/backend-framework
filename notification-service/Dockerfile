FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/notification-service
RUN go build -o notification-main .

CMD ["./notification-main"]
