FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main /app/server/main.go
# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/server/users.json .

EXPOSE 8080
CMD [ "/app/main" ]