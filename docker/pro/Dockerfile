FROM golang:latest AS build
LABEL pavle=523260513@qq.com

WORKDIR $GOPATH/src/

RUN mkdir -p $GOPATH/src/git.vodjk.com/yst/ \
    && cd git.vodjk.com/yst/ \
    && git clone ssh://git@git.vodjk.com:10022/yst/protobuf.git \
    && git clone ssh://git@git.vodjk.com:10022/yst/rpc-server.git


RUN cd $GOPATH/src/git.vodjk.com/golang/common \
    && git checkout master \
    && git pull origin master \
    && cd $GOPATH/src/git.vodjk.com/yst/protobuf \
    && git checkout master \
    && git pull origin master \
    && cd $GOPATH/src/git.vodjk.com/yst/rpc-server \
    && git checkout master \
    && git pull origin master \
    && CGO_ENABLED=0 go build -a

FROM alpine:latest
LABEL pavle=523260513@qq.com

RUN mkdir -p /usr/local/server
WORKDIR /usr/local/server

COPY --from=build /go/src/git.vodjk.com/yst/rpc-server/conf ./conf/
COPY --from=build /go/src/git.vodjk.com/yst/rpc-server/data ./data/
COPY --from=build /go/src/git.vodjk.com/yst/rpc-server/rpc-server .

CMD ["/usr/local/server/enc"]
