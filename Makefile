
BINARY="gtp"
VERSION=0.0.1
BUILD=`date +%F`
SHELL := /bin/bash
SYSTEMOS="linux"
ARCH="amd64"

versionDir="gtp/extend"
gitTag=$(shell git log --pretty=format:'%h' -n 1)
gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
buildDate=$(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit=$(shell git rev-parse --short HEAD)

ldflags="-s -w -X ${versionDir}.version=${VERSION} -X ${versionDir}.gitBranch=${gitBranch} -X '${versionDir}.gitTag=${gitTag}' -X '${versionDir}.gitCommit=${gitCommit}' -X '${versionDir}.buildDate=${buildDate}'"

default:
	@echo "build the ${BINARY}"
	@GOOS=${SYSTEMOS} GOARCH=${ARCH} CGO_ENABLED=0 go build  -a -installsuffix cgo -v $(pwd) -ldflags ${ldflags} -o build/${BINARY}.${SYSTEMOS} -x
# @GOOS=${SYSTEMOS} GOARCH=${ARCH} CGO_ENABLED=1 go build -v $(pwd) -ldflags ${ldflags} -o build/${BINARY}.${SYSTEMOS} -tags=jsoniter -x 
# @go build -ldflags ${ldflags} -o  build/${BINARY}.${SYSTEMOS}  -tags=jsoniter
	@echo "build done."