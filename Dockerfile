FROM ubuntu:14.04

ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /root/go

RUN apt-get -qq install -y wget \
 && cd /root \
 && wget -q https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz \
 && tar -C /usr/local/ -zxvf go1.4.2.linux-amd64.tar.gz \
 && rm -f go1.4.2.linux-amd64.tar.gz \
 && apt-get remove --purge -y wget git-core \
 && apt-get remove --purge -y `apt-mark showauto`\
 && rm -rf /var/lib/apt/lists/*

WORKDIR /root/go/src/hoard-collector/
ADD src/* /root/go/src/hoard-collector/

EXPOSE 80
CMD go run *.go
