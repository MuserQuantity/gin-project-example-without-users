FROM golang:1.19-alpine as build-env

WORKDIR /go/src/github.com/MuserQuantity/gin-project-example-without-users

ENV GOPROXY=https://goproxy.cn,direct

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0
RUN go mod download -x

COPY . .
RUN go mod vendor
RUN go build -o main ./main.go

# Final image
FROM alpine:latest
WORKDIR /root
# Main App
COPY --from=build-env /go/src/github.com/MuserQuantity/gin-project-example-without-users/main /usr/bin/main
# Timezone
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add -U tzdata
ENV TZ=Asia/Shanghai
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8899

CMD ["main"]
