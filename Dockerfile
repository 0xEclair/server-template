FROM alibaba-cloud-linux-3-registry.cn-hangzhou.cr.aliyuncs.com/alinux3/golang:1.19.4

ADD . /server-template

WORKDIR /server-template

ENTRYPOINT ["sh", "-c", "go run main.go"]