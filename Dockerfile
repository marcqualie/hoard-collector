FROM ubuntu:14.04

ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /root/go

ADD src/* /root/go/src/hoard-collector/src/
RUN apt-get -qq install -y wget git-core \
 && cd /root \
 && wget -q https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz \
 && tar -C /usr/local/ -zxf go1.4.2.linux-amd64.tar.gz \
 && rm -f go1.4.2.linux-amd64.tar.gz \
 && cd /root/go/src/hoard-collector/src \
 && go get \
 && cd /root/go/src/hoard-collector \
 && go build -o bin/hoard-collector src/*.go \
 && apt-get remove --purge -y wget git-core \
 && apt-get remove --purge -y `apt-mark showauto`\
 && rm -rf /var/lib/apt/lists/* \
 && rm -rf /root/go/src/hoard-collector/src \
 && rm -rf /usr/local/go

EXPOSE 80
CMD /root/go/src/hoard-collector/bin/hoard-collector
