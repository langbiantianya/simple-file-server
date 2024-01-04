FROM opensuse/tumbleweed:latest
ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

RUN zypper mr -da && \
    zypper ar -fcg 'https://mirrors.ustc.edu.cn/opensuse/tumbleweed/repo/oss' USTC:OSS && \
    zypper ar -fcg 'https://mirrors.ustc.edu.cn/opensuse/tumbleweed/repo/non-oss' USTC:NON-OSS && \
    zypper ar -fcg 'https://mirrors.ustc.edu.cn/opensuse/update/tumbleweed' USTC:UPDATE && \
    zypper dup -y && zypper clean --all \
    && mkdir -p /app/static

WORKDIR /app

COPY ./dist/linux-amd64_linux_amd64_v1/simple-file-server app
COPY ./static/dist static

ENV WORK_HOME=/tmp PASSWD=123456 GIN_MODE=release

ENTRYPOINT [ "./app" ]

EXPOSE 8080

