FROM alpine:3.2

ENV GOPATH /root/go

ADD src/* /root/go/src/hoard-collector/src/
RUN apk update && apk add git go \
 && cd /root/go/src/hoard-collector/src \
 && go get \
 && cd /root/go/src/hoard-collector \
 && go build -o bin/hoard-collector src/*.go \
 && mv bin/hoard-collector /usr/local/bin/hoard-collector \
 && apk del git go \
 && rm -rf /var/cache/apk/* \
 && rm -rf /root/go/src/hoard-collector/src \
 && rm -rf $GOPATH

EXPOSE 80
CMD ["hoard-collector"]
