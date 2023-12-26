FROM openeuler/openeuler:20.03

ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

RUN dnf -y update && dnf -y clean all

WORKDIR /app

COPY ./dist/simple-file-server_linux_amd64_v1/simple-file-server app

ENV WORK_HOME=/tmp PASSWD=123456

ENTRYPOINT [ "./app" ]

EXPOSE 8080