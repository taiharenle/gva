BUILD_ENV := CGO_ENABLED=0
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

TARGET_EXEC := gva
# Git parameters
VERSION := $(shell git describe --tags --always)
BUILDTIME := $(shell date -u '+%Y_%m_%d_%H_%M_%SZ')

.PHONY: all clean setup build-l build-o build-w

export GO111MODULE=on
export GOPROXY=https://goproxy.cn

all: clean setup build-l build-o build-w

clean:
	rm -rf export

setup:
	mkdir -p export/linux
	mkdir -p export/osx
	mkdir -p export/windows

build-l: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=linux go build ${LDFLAGS} -o export/linux/${TARGET_EXEC}_$(VERSION)_$(BUILDTIME)_linux

build-o: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=darwin go build ${LDFLAGS} -o export/osx/${TARGET_EXEC}_$(VERSION)_$(BUILDTIME)_osx

build-w: setup
	${BUILD_ENV} GOARCH=amd64 GOOS=windows go build ${LDFLAGS} -o export/windows/${TARGET_EXEC}_$(VERSION)_$(BUILDTIME).exe