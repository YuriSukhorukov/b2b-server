# FROM alpine:latest
FROM golang:alpine

# install
# RUN apk add --no-cache go

WORKDIR /go/src/app

COPY . .

# build
RUN go build main.go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["./main"]