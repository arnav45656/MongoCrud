# Builder Stage
FROM golang:latest AS builder

WORKDIR /go/src/app
ADD . .

RUN go get -d -v ./... && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app/main ./main.go

# Runtime Stage
FROM alpine:3.1

WORKDIR /application
COPY --from=builder /go/bin/app/main /application/run
RUN chmod +x /application/run

EXPOSE 8000
CMD ["/application/run"]

LABEL Name=ring-master-service-go Version=0.0.1
