FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 8080

CMD ["go", "run", "main.go"]