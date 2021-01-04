#!/bin/bash -ux

TAG=$1

docker run --rm -it -v ${PWD}:/go/app --workdir /go/app golang:1.15.6-buster bash -c "\
  go get github.com/mitchellh/gox/... && \
  go mod tidy && \
  mkdir -p pkg && \
  cd pkg && \
  gox -os='linux darwin windows' -arch='amd64' ../... && \
  go get github.com/tcnksm/ghr && \
  GITHUB_TOKEN=${GITHUB_TOKEN} ghr ${TAG} ."

docker run --rm -it -v ${PWD}:/go/app --workdir /go/app golang:1.15.6-buster bash -c "\
  rm -rf pkg && \
  go mod tidy"
