# stage of building
FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

# stage of publishing
FROM alpine:latest AS publisher
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
