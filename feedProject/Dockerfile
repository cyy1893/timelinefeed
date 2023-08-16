# 构建阶段
FROM golang:1.19 AS builder

WORKDIR /app

COPY . .

# 设置环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# 构建项目
RUN CGO_ENABLED=0 GOOS=linux go build -o feedProject

# 最终阶段
FROM ubuntu:20.04

COPY --from=builder /app/feedProject /app/feedProject

ENV PORT=8080

EXPOSE 8080

CMD ["/app/feedProject"]