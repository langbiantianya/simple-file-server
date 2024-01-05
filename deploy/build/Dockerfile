FROM goreleaser/goreleaser-cross:latest

RUN mkdir -p /usr/local/loongson \
    && sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list \
    && apt-get update \
    && apt-get upgrade -y\
    && apt-get install tar gzip bash -y

COPY ./static/loongson-gnu-toolchain-8.3-x86_64-loongarch64-linux-gnu-rc1.2.tar.xz /usr/local/loongson/loongson-gnu-toolchain-8.3-x86_64-loongarch64-linux-gnu-rc1.2.tar.xz

RUN tar -xvf /usr/local/loongson/loongson-gnu-toolchain-8.3-x86_64-loongarch64-linux-gnu-rc1.2.tar.xz -C /usr/local/loongson \
    && rm /usr/local/loongson/loongson-gnu-toolchain-8.3-x86_64-loongarch64-linux-gnu-rc1.2.tar.xz 
    
ENV PATH=/usr/local/loongson/loongson-gnu-toolchain-8.3-x86_64-loongarch64-linux-gnu-rc1.2/bin:$PATH
