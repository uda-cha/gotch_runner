#!/bin/bash

docker run --rm -it -v ${PWD}:/go/app --workdir /go/app golang:1.15.6-buster $@