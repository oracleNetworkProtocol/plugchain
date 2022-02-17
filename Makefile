#!/usr/bin/make -f

AppVersion ?= $(shell echo $(shell git describe --tags `git rev-list --tags="v*" --max-count=1`) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
TMVERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
PLUGCHAIN_BINARY= plugchaind
PLUGCHAIN_DIR = plugchain
HTTPS_GIT = https://github.com/oracleNetworkProtocol/plugchain.git
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf

export GO111MODULE = on

# Default target executed when no arguments are given to make.
default_target: all

.PHONY: default_target

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))




# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Version=$(AppVersion) \
 		  	  -X github.com/cosmos/cosmos-sdk/version.Name=onp \
		  	  -X github.com/cosmos/cosmos-sdk/version.AppName=$(PLUGCHAIN_BINARY) \
		  	  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
			  -X github.com/tharsis/ethermint/version.AppVersion=$(PLUGCHAIN_BINARY) \
			  -X github.com/tharsis/ethermint/version.GitCommit=$(AppVersion) \
			  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TMVERSION) \
		 	    -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"
# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif


###############################################################################
###                                    Build                               ###
###############################################################################
BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/
build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

all: build

.PHONY: all build  
###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################
go.sum: go.mod
	@echo "Ensure dependencies have not been modified ..."
	go mod verify
	go mod tidy

contract-tools:
ifeq (, $(shell which stringer))
	@echo "Installing stringer..."
	@go get golang.org/x/tools/cmd/stringer
else
	@echo "stringer already installed; skipping..."
endif

ifeq (, $(shell which go-bindata))
	@echo "Installing go-bindata..."
	@go get github.com/kevinburke/go-bindata/go-bindata
else
	@echo "go-bindata already installed; skipping..."
endif

ifeq (, $(shell which gencodec))
	@echo "Installing gencodec..."
	@go get github.com/fjl/gencodec
else
	@echo "gencodec already installed; skipping..."
endif

ifeq (, $(shell which protoc-gen-go))
	@echo "Installing protoc-gen-go..."
	@go get github.com/fjl/gencodec github.com/golang/protobuf/protoc-gen-go
else
	@echo "protoc-gen-go already installed; skipping..."
endif

ifeq (, $(shell which protoc))
	@echo "Please istalling protobuf according to your OS"
	@echo "macOS: brew install protobuf"
	@echo "linux: apt-get install -f -y protobuf-compiler"
else
	@echo "protoc already installed; skipping..."
endif

ifeq (, $(shell which solcjs))
	@echo "Installing solcjs..."
	@npm install -g solc@0.5.11
else
	@echo "solcjs already installed; skipping..."
endif


tools: tools-stamp
tools-stamp: contract-tools proto-tools
	# Create dummy file to satisfy dependency and avoid
	# rebuilding when this Makefile target is hit twice
	# in a row.
	touch $@

tools-clean:
	rm -f tools-stamp

.PHONY: tools contract-tools proto-tools tools-stamp tools-clean

###############################################################################
###                               Localnet                                  ###
###############################################################################

# Run a single testnet locally
localnet: 
	@echo "start make install and ./scripts/setup.sh"
	@make install 
	./scripts/setup.sh

.PHONY: localnet


###############################################################################
###                            Documentation                                ###
###############################################################################
docs-rely: 
	@echo "install vuepress rely ..."
	@cd docs/ && sudo cnpm i

build-docs: 
	@echo "vuepress build ..."
	@cd docs/ && ./deploy.sh 
	

.PHONY: npm-vue vuepress


###############################################################################
###                                Protobuf                                 ###
###############################################################################
PREFIX ?= /usr/local
BIN ?= $(PREFIX)/bin
UNAME_S ?= $(shell uname -s)
UNAME_M ?= $(shell uname -m)

BUF_VERSION="1.0.0-rc10"
BINARY_NAME="buf"
PROTOC_VERSION ?= 3.13.0
PROTOC_GRPC_GATEWAY_VERSION = 1.14.7
ifeq ($(UNAME_S),Linux)
  PROTOC_ZIP ?= protoc-3.13.0-linux-x86_64.zip
  PROTOC_GRPC_GATEWAY_BIN ?= protoc-gen-grpc-gateway-v1.14.7-linux-x86_64
endif
ifeq ($(UNAME_S),Darwin)
  PROTOC_ZIP ?= protoc-3.13.0-osx-x86_64.zip
  PROTOC_GRPC_GATEWAY_BIN ?= protoc-gen-grpc-gateway-v1.14.7-darwin-x86_64
endif

proto-tools: buf
ifeq (, $(shell which protoc))
	@echo "Installing protoc compiler..."
	@(cd /tmp; \
	curl -OL "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}"; \
	unzip -o ${PROTOC_ZIP} -d $(PREFIX) bin/protoc; \
	unzip -o ${PROTOC_ZIP} -d $(PREFIX) 'include/*'; \
	rm -f ${PROTOC_ZIP})
else
	@echo "protoc already installed; skipping..."
endif

ifeq (, $(shell which protoc-gen-gocosmos))
	@echo "Installing protoc-gen-gocosmos..."
	@go install github.com/regen-network/cosmos-proto/protoc-gen-gocosmos
else
	@echo "protoc-gen-gocosmos already installed; skipping..."
endif

ifeq (, $(shell which protoc-gen-grpc-gateway))
	@echo "Installing protoc-gen-grpc-gateway..."
	@curl -o "${BIN}/protoc-gen-grpc-gateway" -L "https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${PROTOC_GRPC_GATEWAY_VERSION}/${PROTOC_GRPC_GATEWAY_BIN}"
	@chmod +x "${BIN}/protoc-gen-grpc-gateway"
else
	@echo "protoc-gen-grpc-gateway already installed; skipping..."
endif

ifeq (, $(shell which protoc-gen-swagger))
	@echo "Installing protoc-gen-swagger..."
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@npm install -g swagger-combine
else
	@echo "protoc-gen-grpc-gateway already installed; skipping..."
endif

buf: buf-stamp

buf-stamp:
	@echo "Installing buf..."
	@curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/${BINARY_NAME}-${UNAME_S}-${UNAME_M}" \
    -o "${BIN}/${BINARY_NAME}" && \
	chmod +x "${BIN}/${BINARY_NAME}"

proto-gen: 
	@./scripts/protocgen.sh

proto-swagger-gen: 
	@./scripts/protoc-swagger-gen.sh

.PHONY: buf buf-stamp proto-gen proto-swagger-gen 


###############################################################################
###                                Releasing                                ###
###############################################################################
PACKAGE_NAME:=github.com/oracleNetworkProtocol/plugchain
GOLANG_CROSS_VERSION  = v1.17.6
GOPATH ?= '$(HOME)/go'
release-dry-run:
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-v ${GOPATH}/pkg:/go/pkg \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/troian/golang-cross:${GOLANG_CROSS_VERSION} \
		--skip-validate  --snapshot

release:
	@if [ ! -f ".release-env" ]; then \
		echo "\033[91m.release-env is required for release\033[0m";\
		exit 1;\
	fi
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		--env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/troian/golang-cross:${GOLANG_CROSS_VERSION} \
		release --rm-dist --skip-validate

.PHONY: release-dry-run release