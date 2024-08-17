# Builder
FROM golang:1.22-alpine AS builder

COPY . /app

WORKDIR /app

RUN go mod download
RUN go build -o /go/bin/todo cmd/main.go

# Runner
FROM alpine:latest AS runner

COPY --from=builder /go/bin/todo /go/bin/todo

EXPOSE 50051

CMD ["/go/bin/todo"]
