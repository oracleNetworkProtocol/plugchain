# docker build -t plugchain .

#从 golang:alpine 映像开始构建镜像。
FROM golang:alpine AS build-env

#指定维护者信息
# MAINTAINER

# 安装最低限度的必要依赖项
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev jq
RUN sed -i 's/https/http/' /etc/apk/repositories \
&& apk add --no-cache $PACKAGES 
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

#设置工作目录
WORKDIR /plugchain

#将 . 项目中的当前目录复制到 . 镜像中的工作目录。
COPY . .

#设置代理
ENV GOPROXY=https://proxy.golang.com.cn,direct

#编译程序
RUN make install

RUN cd cosmovisor && make install && cd -

#暴露端口
# EXPOSE 26656

# CMD ["plugchaind"]