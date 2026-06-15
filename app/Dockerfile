FROM golang:1.26.3-alpine AS builder
WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN CGO_ENALBLED=0 GOOS=linux go build -o http-server-projeto-korp .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/http-server-projeto-korp .
EXPOSE 8080
CMD ["./http-server-projeto-korp"]