FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o blog-api ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/blog-api .

EXPOSE 8080

CMD ["./blog-api"]