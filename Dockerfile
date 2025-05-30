# Dockerfile
FROM golang:latest

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

#EXPOSE the port
EXPOSE 8080

CMD ["./main"]
