FROM golang:1.21 as builder

WORKDIR /build

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn/,direct

COPY . .
RUN tar -xvf ./pkg/upx.4.2.4.tar.xz
RUN mv ./pkg/upx upx
RUN chmod +x ./upx
RUN go mod download
RUN go mod tidy
RUN bash ./scripts/build.sh

FROM alpine:3.19
RUN apk update && apk add tzdata # 时区 https://segmentfault.com/a/1190000040524996
ENV TZ Asia/Shanghai  #panic: unknown time zone Asia/Shanghai

WORKDIR /project

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder  /output .
COPY --from=builder  /build/scripts/run.sh .
CMD ["top"]
