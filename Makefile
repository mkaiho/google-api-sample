BIN_DIR:=bin
ROOT_PACKAGE:=github.com/mkaiho/google-api-sample
COMMAND_PACKAGES:=$(shell go list ./cmd/...)
BINARIES := $(COMMAND_PACKAGES:$(ROOT_PACKAGE)/cmd/%=$(BIN_DIR)/%)

.PHONY: build
build: $(BINARIES)

$(BINARIES): $(GO_FILES)
	@go build -o $@ $(@:$(BIN_DIR)/%=$(ROOT_PACKAGE)/cmd/%)

.PHONY: dev-deps
dev-deps:
	go get gotest.tools/gotestsum@v1.7.0
	go get github.com/vektra/mockery/v2/.../
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: gen-mock
gen-mock:
	make dev-deps
	mockery --all --inpackage --case underscore

.PHONY: test
test:
	gotestsum

.PHONY: clean
clean:
	@rm -rf ./bin
