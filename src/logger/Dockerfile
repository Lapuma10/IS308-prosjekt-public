FROM golang:1.12.0-alpine3.9

RUN apk update
RUN apk add git

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/streadway/amqp
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
CMD ["./main"]
