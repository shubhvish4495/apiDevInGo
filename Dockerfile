FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY ./ ./

RUN go build -o /docker-go-api

CMD ["/docker-go-api"]