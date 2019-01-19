# Build Stage
FROM golang:1.11 AS build-stage

LABEL app="build-emsci"
LABEL REPO="https://github.com/jmartin84/emsci"

ENV PROJPATH=/go/src/github.com/jmartin84/emsci

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/jmartin84/emsci
WORKDIR /go/src/github.com/jmartin84/emsci

