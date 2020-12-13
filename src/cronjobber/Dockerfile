FROM golang:1.15.5-alpine as builder
RUN mkdir /build
ADD . /build/

WORKDIR /build
# needed to go get from github
RUN apk add git
# needed in services that communicate with rabbitMQ
RUN go get github.com/streadway/amqp
# needed to communicate with mysql databases
RUN go get github.com/go-sql-driver/mysql

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app

CMD ["./main"]