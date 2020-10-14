FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go get -d -v /go/src/app
RUN go build main.go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["./main"]