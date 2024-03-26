FROM golang:1.22 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hashing-service ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /build/hashing-service .

EXPOSE 50052

CMD ["./hashing-service"]