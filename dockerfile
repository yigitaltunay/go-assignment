
FROM golang:latest

LABEL maintainer="yigit"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go version

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]