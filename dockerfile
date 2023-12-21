FROM debian:stable-slim

ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
    && apt-get update \
    && apt-get upgrade -y \
    && rm -rf /var/lib/apt/lists/* \
    && mkdir /app

WORKDIR /app

COPY ./dist/simple-file-server_linux_amd64_v1/simple-file-server app

ENV WORK_HOME=/tmp PASSWD=123456

ENTRYPOINT [ "./app" ]

EXPOSE 8080

