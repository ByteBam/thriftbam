FROM golang:1.22.3-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

COPY . .
RUN go mod tidy
RUN mkdir -p output
RUN cp script/* output 2>/dev/null
RUN go build -o output/bin/thriftbam

FROM alpine:latest AS runner

WORKDIR /app

ENV NODE_ENV production

RUN addgroup -g 1001 -S gopher
RUN adduser -S thriftbam -u 1001

COPY --from=builder /app/output .
COPY --from=builder /app/biz/config ./biz/config

USER thriftbam

EXPOSE 8888

ENV PORT 8888

CMD ["sh", "./bootstrap.sh"]