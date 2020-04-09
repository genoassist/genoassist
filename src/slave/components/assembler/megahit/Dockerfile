FROM ubuntu:18.04

# TODO: this does not work yet, fix it
RUN apt-get update
RUN apt-get install -y --no-install-recommends apt-transport-https \
    wget

# first, we download and untar the Linux x86 executable of Megahit
RUN wget --no-check-certificate https://github.com/voutcn/megahit/releases/download/v1.2.9/MEGAHIT-1.2.9-Linux-x86_64-static.tar.gz
RUN tar zvxf MEGAHIT-1.2.9-Linux-x86_64-static.tar.gz

COPY . /root/megahit
WORKDIR /root/megahit

RUN apt-get install -y --no-install-recommends \
    bzip2 \
    cmake \
    gzip \
    g++ \
    libgomp1 \
    make \
    python \
    zlib1g-dev && \
    rm -rf build && \
    mkdir -p build && \
    cd build && \
    cmake -DCMAKE_BUILD_TYPE=Release .. && \
    make -j4 install && \
    apt-get autoremove --purge -y \
    cmake \
    g++ \
    make \
    zlib1g-dev

RUN megahit --test && megahit --test --kmin-1pass
ENTRYPOINT ["megahit"]