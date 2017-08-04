FROM golang:latest

COPY ./cmd/rest-api/rest-api /go/bin/rest-api

CMD ["rest-api"]