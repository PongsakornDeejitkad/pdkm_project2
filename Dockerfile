FROM golang:1.22.1-alpine3.19 AS builder
WORKDIR /intern-order-management
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o ./intern-order-management .

ENTRYPOINT [ "./intern-order-management" ]