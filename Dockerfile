FROM golang:1.21 as builder

WORKDIR /build

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn/,direct
ENV compress 1
COPY . .
RUN mv ./pkg/upx upx
RUN chmod +x ./upx
RUN go mod download
RUN go mod tidy
RUN bash ./scripts/build.sh

FROM alpine:3.19
# use aliyun
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add tzdata
ENV TZ Asia/Shanghai
WORKDIR /project

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder  /output .
COPY --from=builder  /build/scripts/run.sh .
