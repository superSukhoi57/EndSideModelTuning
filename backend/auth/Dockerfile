FROM golang:1.25.9-alpine3.22 AS builder

WORKDIR /build
COPY . .
WORKDIR /build/auth
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o auth .

FROM alpine:3.22

WORKDIR /work

COPY --from=builder /build/auth/auth .
COPY --from=builder /build/auth/etc/ ./etc/

EXPOSE 12300

CMD ["./auth", "-f", "etc/auth.yaml"]
