FROM golang:1.23.3-alpine AS builder
WORKDIR /app
RUN apk add --no-cache build-base sqlite-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates sqlite-libs tzdata \
    && adduser -D appuser
WORKDIR /home/appuser/app
COPY --from=builder /app/main .
RUN mkdir -p shared && chown -R appuser:appuser /home/appuser/app
USER appuser
EXPOSE 3333
CMD ["./main"]
