FROM golang:1.25.9-alpine3.22 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth .

FROM alpine:3.22

WORKDIR /work

COPY --from=builder /build/auth .
COPY etc/ ./etc/

EXPOSE 12300

CMD ["./auth", "-f", "etc/auth.yaml"]
