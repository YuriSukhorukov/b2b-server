FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go get -d -v /go/src/app

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN export PATH=$(go env GOPATH)/bin:$PATH
RUN swag init

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

ENV PORT 8080
EXPOSE 8080

FROM scratch AS final

COPY --from=builder /go/src/app/main .
ENTRYPOINT ["./main"]