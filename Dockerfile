FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 go build -o accuscraper .



FROM alpine:latest

WORKDIR /

COPY --from=builder /app/accuscraper .

EXPOSE 8080

USER nobody

CMD ["./accuscraper"]