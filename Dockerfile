FROM docker.io/library/golang:1.24

COPY . /code
WORKDIR /code

RUN go test ./...
