FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o app .

FROM alpine

WORKDIR /app

COPY --from=builder /build/app /app/app

CMD ["./app"]