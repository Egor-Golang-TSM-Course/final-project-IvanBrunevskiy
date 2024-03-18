FROM golang:1.22 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway-service ./cmd/gateway

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /build/gateway-service .

EXPOSE 8080

CMD ["./gateway-service"]