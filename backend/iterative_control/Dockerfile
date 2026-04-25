FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o iterative .

FROM alpine:latest

WORKDIR /work

COPY --from=builder /build/iterative .

EXPOSE 12301

CMD ["./iterative"]
