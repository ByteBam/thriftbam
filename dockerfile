FROM golang:1.22.3-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY . .
RUN go mod tidy
RUN go build -o output/bin/thriftbam

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/output .
COPY --from=builder /app/biz/config ./biz/config

EXPOSE 8888

CMD ["./bin/thriftbam"]