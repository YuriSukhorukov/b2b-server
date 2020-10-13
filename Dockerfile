# FROM alpine:latest
FROM golang:alpine

# install
RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY . .

# build
RUN go build main.go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["./main"]