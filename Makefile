MAIN_PATH := ./cmd

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(dir $(MKFILE_PATH))

GO      := $(shell which go)
GOPATH  := $(shell $(GO) env GOPATH)

MOCKGEN := $(GOPATH)/bin/mockgen
MOCK_SRC_FILES := $(shell grep --include=\*.go -lrw '$(CURRENT_DIR)' -e 'mockgen' --exclude-dir=mocks)
MOCK_GEN_FILES := $(MOCK_SRC_FILES:.go=_mock.go)

COVERAGE_FILE := coverage.out
REPORT_FILE   := report.json

.PHONY: help
.SILENT: help
help: ## show this help
	echo "usage: make [target]"
	echo ""
	grep -E "^(.+)\:\ .*##\ (.+)" $(MAKEFILE_LIST) | sed 's/:.*##/#/' | column -t -c 2 -s '#'

.SILENT: $(MOCKGEN)
$(MOCKGEN):
	echo "Instaling mockgen dependency"
	$(GO) install github.com/golang/mock/mockgen@latest

.SILENT: $(MOCK_GEN_FILES)
$(MOCK_GEN_FILES): %_mock.go: %.go $(MOCKGEN)
	echo "Generating file $@"
	$(GO) generate -run mockgen $<

.PHONY: mock
mock: ## generate mocks
mock: $(MOCK_SRC_FILES) $(MOCK_GEN_FILES)

.PHONY: rmmock
rmmock: ## remove generated mocks
rmmock:
	rm -f $(MOCK_GEN_FILES)

.PHONY: test
test: ## run tests
test: mock
	$(GO) test -coverprofile=$(COVERAGE_FILE) ./... -covermode=count -json > $(REPORT_FILE)

.PHONY: cover
cover: ## see coverage in browser
cover: test
	$(GO) tool cover -html=$(COVERAGE_FILE)

.PHONY: release
release: ## build a release
release:
	$(GO) build -tags=release

.PHONY: run
run: ## build and run
run:
	$(GO) run $(MAIN_PATH)/
