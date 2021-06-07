# docker build -t plugchain .
# Server:

#从 golang:alpine 映像开始构建镜像。
FROM golang:alpine AS build-env

# 安装最低限度的必要依赖项
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev
RUN sed -i 's/https/http/' /etc/apk/repositories \
&& apk add --no-cache $PACKAGES 

#设置工作目录
WORKDIR /code

#将 . 项目中的当前目录复制到 . 镜像中的工作目录。
COPY . .

RUN GOPROXY=https://goproxy.io make install

# CMD ["plugchaind"]