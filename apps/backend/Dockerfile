FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

RUN go work sync
RUN go build -o main apps/backend/cmd/main.go 

FROM alpine:latest as runner

WORKDIR /app/book-shop
COPY --from=builder /app/main ./
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./main"]