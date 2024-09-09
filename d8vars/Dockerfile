# syntax=docker/dockerfile:1

FROM ubuntu:22.04 AS build-stage

# Install Go 1.22

RUN apt-get update && apt-get install -y ca-certificates \
                                                       wget \
                                                       git \
                                                       gcc


RUN wget -P /tmp "https://dl.google.com/go/go1.22.5.linux-amd64.tar.gz"
RUN rm -rf /usr/local/go
RUN tar -C /usr/local -xzf "/tmp/go1.22.5.linux-amd64.tar.gz"
RUN rm "/tmp/go1.22.5.linux-amd64.tar.gz"
                                                       
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN go version

# Install gnuCobol 3.2

WORKDIR /download

RUN wget -P /download "https://ftp.gnu.org/gnu/gnucobol/gnucobol-3.2.tar.gz"
RUN tar -zxvf "gnucobol-3.2.tar.gz"

RUN apt-get install -y \
       build-essential \
            libgmp-dev \
             libdb-dev \
        libncurses-dev \
          libcjson-dev \
           libxml2-dev 

WORKDIR /download/gnucobol-3.2

RUN ./configure --prefix=/opt/gnucobol-3.2 \
                                   && make \ 
                             && make install 
 
# Compile Go

WORKDIR /app

COPY . /app/
RUN go mod download

WORKDIR /app/cmd

RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 \
    CC=aarch64-linux-gnu-gcc \
    CGO_CFLAGS="-I/opt/gnucobol-3.2/include" \
    CGO_LDFLAGS="-L/opt/gnucobol-3.2/lib -lcob" \
    go build main.go

# Compile COBOL

WORKDIR /app/cobol

ENV PATH=$PATH:/opt/gnucobol-3.2/bin
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/gnucobol-3.2/lib  

RUN cobc -m *.cbl -fbinary-byteorder=native -fbinary-size=2-4-8

# Deploy the application binary into a lean image
FROM ubuntu:22.04 AS build-release-stage

WORKDIR /

COPY --from=build-stage /opt/gnucobol-3.2 /opt/gnucobol-3.2 
RUN apt-get update && apt-get install -yq --no-install-recommends \
#                                                  gcc \
                                                  libgmp-dev \
                                                  libdb-dev \
                                                  libncurses-dev \
                                                  libcjson-dev \
                                                  libxml2-dev \
                                                  && apt-get clean \
                                      && rm -rf /var/lib/apt/lists/* 


ENV PATH=$PATH:/opt/gnucobol-3.2/bin
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/gnucobol-3.2/lib  


COPY --from=build-stage /app/cmd/main /app/cmd/
COPY --from=build-stage /app/cobol/*.so /app/cmd/
COPY --from=build-stage /app/conf/. /app/conf/

#USER root

CMD ["/bin/bash"]


 