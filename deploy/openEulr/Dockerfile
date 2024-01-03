FROM openeuler/openeuler:20.03

ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

RUN dnf -y update \
    && dnf -y clean all \
    && mkdir -p /app/static

WORKDIR /app

COPY ./dist/simple-file-server_linux_amd64_v1/simple-file-server app
COPY ./static/dist static

ENV WORK_HOME=/tmp PASSWD=123456 GIN_MODE=release

ENTRYPOINT [ "./app" ]

EXPOSE 8080