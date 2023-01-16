# Build gqlc in a stock Go builder container
FROM golang:1.16.3-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

COPY . /qlcchain/go-sonata-server
RUN cd /qlcchain/go-sonata-server && make clean build

# Pull gqlc into a second stage deploy alpine container
FROM alpine:3.17.1
LABEL maintainer="developers@qlink.mobi"

ENV QLCHOME /qlcchain

RUN apk --no-cache add ca-certificates && \
    addgroup qlcchain && \
    adduser -S -G qlcchain qlcchain -s /bin/sh -h "$QLCHOME" && \
    chown -R qlcchain:qlcchain "$QLCHOME"

USER qlcchain

WORKDIR $QLCHOME

COPY --from=builder /qlcchain/go-sonata-server/build/gsonata /usr/local/bin/gsonata

EXPOSE 9999

ENTRYPOINT [ "gsonata"]

VOLUME [ "$QLCHOME" ]
