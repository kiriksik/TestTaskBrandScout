FROM golang:alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

RUN apk add --no-cache postgresql-client
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
RUN go build ./internal/src/main.go

EXPOSE 8080

CMD ["./internal/src/main"]