## The makefile for golang.
##
## Author: Uberate
## Email: <ubserate@gmail.com>
##
## Search "CHANGE_ME" and replace value.
##
## To build for all arch and paltform
## =====================================================================================================================
.DEFAULT_GOAL:=help

.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make target \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
	@echo
	@echo Platform and arch support: $(PLATFORM_LIST)

CMD_PATH := cmd# the cmd path "CHANGE_ME"
OUTPUT_PATH := output

GIT_VERSION := $(shell git describe --tags || echo "unknown")
GO_FDFLAGS := -ldflags "-w -s -X 'main.Version=${VERSION}' -X 'main.HashTag=`git rev-parse HEAD`' -X 'main.BranchName=`git rev-parse --abbrev-ref HEAD`' -X 'main.BuildDate=`date -u '+%Y-%m-%d_%I:%M:%S%p'`' -X 'main.GoVersion=`go version`'"

GO_BUILD_CMD := CGO_ENABLED=0 go build $(GO_FDFLAGS)

GOOS = $(shell echo $@ | awk -F_ '{print $$2}')
ARCH = $(shell echo $@ | awk -F_ '{print $$3}')

BIN_FILE_NAME_GOM := gom# the output name of bin. "CHANGE_ME"
BIN_GOM := cmd/stdcmd/main.go#"CHANGE_ME"

gom_%: ## build the specify os and arch bin.
	GOOS=$(GOOS) GOARCH=$(ARCH) $(GO_BUILD_CMD) -o $(OUTPUT_PATH)/$@_$(GIT_VERSION)/$(BIN_FILE_NAME_GOM) $(BIN_GOM)

PLATFORM_LIST := \
    darwin_arm64 \
    linux_arm64 \
    windows_arm64 \
    darwin_amd64 \
    linux_amd64 \
    windows_amd64 \

ALL_GOM_ARCH = $(addprefix gom_, $(PLATFORM_LIST))

gom-all-arch: $(ALL_GOM_ARCH) ## Build all binary for all platform and arch.

output: ## Create the output dir.
	mkdir -p $(OUTPUT_PATH)

.PHONY: clean
clean: ## Clean the output dir.
	rm -rf $(OUTPUT_PATH)

RELEASES = \
	$(addsuffix .tar.gz, $(ALL_GOM_ARCH))

$(RELEASES): %.tar.gz: %
	tar czf $(OUTPUT_PATH)/$@ $(OUTPUT_PATH)/$<_$(GIT_VERSION)/
	rm -rf $(OUTPUT_PATH)/$<_$(GIT_VERSION)

releases: $(RELEASES) ## Tar the project release.

.PHONY: test
test: ## Test the project.
	go test ./...