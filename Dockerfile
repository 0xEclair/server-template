FROM alibaba-cloud-linux-3-registry.cn-hangzhou.cr.aliyuncs.com/alinux3/golang:1.19.4

ADD . /server-template

WORKDIR /server-template

RUN go build -o bitmap-server

ENTRYPOINT ["sh", "-c", "./bitmap-server"]