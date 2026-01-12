FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o scanner ./cmd/scan/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/scanner .
ENTRYPOINT ["./scanner"]