OUT_DIR = bin
ifndef FILES
	FILES = $(shell ls -d cmd/*/ | cut -d/ -f2)
endif
FILES_OUT = $(addprefix ${OUT_DIR}/,${FILES})
UNAME_M := $(shell uname -m)

VERSION_FILE = version/version.json
VERSION_FILE_DEFAULT =version/version.default.json

ifeq ($(OS),Windows_NT)
	PLATFORM ?= windows
	DEST ?= windows
else ifeq ($(UNAME_M),x86_64)
	PLATFORM ?= linux
	DEST ?= linux
else ifeq ($(UNAME_M),armv7l)
	PLATFORM ?= linux
	DEST ?= arm
endif

ARCH =

all:
ifeq ($(OS),Windows_NT)
	make win
else ifeq ($(UNAME_M),x86_64)
	make linux
else ifeq ($(UNAME_M),armv7l)
	make arm32
endif

linux: PLATFORM = linux
linux: ${OUT_DIR } codegen ${FILES_OUT} 
linux: version

win: PLATFORM = windows
win: ${OUT_DIR} codegen ${FILES_OUT:=.exe}
win: version

arm32: PLATFORM = linux
arm32: ARCH = arm
arm32: ${OUT_DIR} ${FILES_OUT}

.FORCE:
${OUT_DIR}/%: .FORCE
	@echo compiling $(@)...
	GOOS=$(PLATFORM) GOARCH=$(ARCH) go build -o $(@) -tags $(PLATFORM),$(TAGS) ./cmd/$(basename ${@F})

go-generate:
ifeq (, $(shell which genny))
	@echo installing genny...
	@go get github.com/cheekybits/genny
else
	@echo genny is already installed
endif
ifeq (, $(shell which protoc-gen-go))
	@echo installing protoc-gen-go...
	@go get github.com/golang/protobuf/protoc-gen-go
else
	@echo protoc-gen-go is already installed
endif
	go generate ./...

clean:
	rm ${OUT_DIR} -rf

${OUT_DIR}:
	@echo create output dir...
	@mkdir ${OUT_DIR}

codegen: go-generate

version: 
	sed -e 's/"BuildTime": [^\]*/"BuildTime": "${shell date +%Y%m%d%H | cut -c3-}"/' ${VERSION_FILE_DEFAULT} | tee ${VERSION_FILE} 

.PHONY: all clean win linux arm32 .FORCE version codegen