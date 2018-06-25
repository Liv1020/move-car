FROM golang:latest AS build
LABEL pavle=523260513@qq.com

WORKDIR $GOPATH/src/

COPY .ssh /root/.ssh
RUN chmod 0600 /root/.ssh/id_rsa

RUN mkdir -p $GOPATH/src/github.com/Liv1020/ \
    && cd github.com/Liv1020/ \
    && git clone git@github.com:Liv1020/move-car.git

RUN cd $GOPATH/src/github.com/Liv1020/move-car/ \
    && git checkout master \
    && git pull origin master \
    && CGO_ENABLED=0 go build -a

FROM alpine:latest
LABEL pavle=523260513@qq.com

RUN mkdir -p /usr/local/server
WORKDIR /usr/local/server

COPY --from=build /go/src/github.com/Liv1020/move-car/conf ./conf/
COPY --from=build /go/src/github.com/Liv1020/move-car/move-car .

CMD ["/usr/local/server/move-car"]
