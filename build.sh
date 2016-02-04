#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

docker run --rm -it \
  -v "$(pwd):/go/src/github.com/karlkfi/atlas-meta" \
  -w /go/src/github.com/karlkfi/atlas-meta \
  golang:1.5.3-alpine \
  sh -ceux 'apk add --update git && go get github.com/golang/glog && go build .'

docker build -t karlkfi/atlas-meta .

rm ./atlas-meta
