FROM golang:1.11-alpine3.7 AS build
LABEL pavle=523260513@qq.com

RUN apk add --no-cache git openssh

COPY .ssh /root/.ssh
RUN chmod 0600 /root/.ssh/id_rsa

WORKDIR $GOPATH/src/

RUN mkdir -p $GOPATH/src/github.com/Liv1020/ \
    && cd github.com/Liv1020/ \
    && git clone git@github.com:Liv1020/move-car-api.git

RUN cd $GOPATH/src/github.com/Liv1020/move-car-api/ \
    && git checkout master \
    && git pull origin master \
    && CGO_ENABLED=0 go build -a -ldflags '-w -s'

FROM alpine:latest
LABEL pavle=523260513@qq.com

RUN apk add --no-cache ca-certificates

RUN mkdir -p /usr/local/server
WORKDIR /usr/local/server

VOLUME /usr/local/server/conf
COPY --from=build /go/src/github.com/Liv1020/move-car-api/data ./data/
COPY --from=build /go/src/github.com/Liv1020/move-car-api/move-car-api .

CMD ["/usr/local/server/move-car-api"]
