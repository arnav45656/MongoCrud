# 

# WORKDIR /app

# COPY . .

# RUN go mod tidy

# RUN go build -o main .

# EXPOSE 8000

# CMD ["./main"]

FROM golang:latest AS Builder


WORKDIR /app
ADD . /app/
RUN go get -d -v ./... # See Below for details

RUN go build -o /application/run ./main.go

FROM alpine:3.14
RUN apk add ca-certificates
WORKDIR /application
COPY --from=builder ./app /application/run
EXPOSE 8000
ENTRYPOINT /application/run
