FROM golang:1.25.9-alpine3.22 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o iterative .

FROM alpine:3.22

WORKDIR /work

COPY --from=builder /build/iterative .

EXPOSE 12301

CMD ["./iterative"]
