FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY . .

# init module kalau belum ada
RUN if [ ! -f go.mod ]; then \
      go mod init app; \
    fi

# ambil dependency yang kamu butuhkan
RUN go get github.com/spf13/viper github.com/lib/pq

# rapihin deps
RUN go mod tidy

# build
RUN go build -o server .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server /app/server
EXPOSE 8080
CMD ["/app/server"]
