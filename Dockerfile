# build as static binary
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-saas

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go-saas /go-saas
EXPOSE 8080
ENTRYPOINT ["/go-saas"]