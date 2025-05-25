FROM golang:1.21 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ingestlog main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ingestlog .
RUN apk --no-cache add ca-certificates
ENV S3_BUCKET=ingestlog
ENTRYPOINT ["/app/ingestlog"]