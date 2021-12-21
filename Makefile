AppVersion=$(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BUILDDIR ?= $(CURDIR)

export GO111MODULE = on


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

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))




# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Version=$(AppVersion) \
 		  	  -X github.com/cosmos/cosmos-sdk/version.Name=plugchain \
		  	  -X github.com/cosmos/cosmos-sdk/version.AppName=plugchaind \
		  	  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
			  -X github.com/tharsis/ethermint/version.AppVersion=plugchaind \
			  -X github.com/tharsis/ethermint/version.GitCommit=$(AppVersion) \
		 	    -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

all: install build build-linux

###############################################################################
###                                Documentation                            ###
###############################################################################
BUILD_TARGETS := build install

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) ./cmd/plugchaind

build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

build-window:
	GOOS=windows GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################
go.sum: go.mod
	echo "Ensure dependencies have not been modified ..."
	go mod verify
	go mod tidy

.PHONY: all build-linux build-window install

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
###                                   Docs                                  ###
###############################################################################
npm-vue: 
	@echo "install vuepress rely ..."
	@cd docs/ && sudo npm install vue && sudo npm install -D vuepress

vuepress: 
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
    "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-${UNAME_S}-${UNAME_M}" \
    -o "${BIN}/buf" && \
	chmod +x "${BIN}/buf"

	touch $@

proto-gen: 
	@./scripts/protocgen.sh

proto-swagger-gen: 
	@./scripts/protoc-swagger-gen.sh

.PHONY: buf buf-stamp proto-gen proto-swagger-gen 