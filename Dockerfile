# Build stage
FROM golang:1.17.6-alpine3.15 as builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/app/main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY /config/app.env /app/config/

EXPOSE 8080
EXPOSE 8081
ENTRYPOINT [ "/app/main" ]
