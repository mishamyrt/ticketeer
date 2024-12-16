VERSION = 0.1.0

GOLANGCI_LINT_VERSION = v1.62.2
REVIVE_VERSION = v1.5.1
GIT_CHGLOG_VERSION = v0.15.4

GO_BIN_PATH := $(shell go env GOPATH)/bin
TEST_MODULES := $(shell go list ./... | grep -v -e /cmd/)

.PHONY: build
build:
	go build -o build/ticketeer .

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -rf build
	rm -rf dist
	rm -f packaging/npm/ticketeer-*/ticketeer

.PHONY: release
release: clean
	@goreleaser release --snapshot --clean
    python3 packaging/publish.py $(VERSION)
    git add packaging/npm
    git add Makefile
    git commit -m "chore: release $(VERSION)"
    git tag -a $(VERSION) -m "release $(VERSION)"
    git push && git push --tags

.PHONY: test
test:
	@go test $(TEST_MODULES)

.PHONY: setup
setup:
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(GO_BIN_PATH) $(GOLANGCI_LINT_VERSION)
	go install github.com/mgechev/revive@$(REVIVE_VERSION)
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG_VERSION)

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	golangci-lint run ./...
	revive -config ./revive.toml  ./...

.PHONY: coverage
coverage:
	@mkdir -p coverage
	@go test -coverprofile=coverage/cover.out $(TEST_MODULES)
	@go tool cover -html coverage/cover.out -o coverage/cover.html

.PHONY: check
check:
	@make lint
	@make test
	@make build

.PHONY: changelog
changelog:
	git-chglog -o CHANGELOG.md
