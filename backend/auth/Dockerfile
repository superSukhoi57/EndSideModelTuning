FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth .

FROM alpine:latest

WORKDIR /work

COPY --from=builder /build/auth .
COPY etc/ ./etc/

EXPOSE 12300

CMD ["./auth", "-f", "etc/auth.yaml"]
