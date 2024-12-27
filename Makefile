VERSION = 0.1.5

GOLANGCI_LINT_VERSION = v1.62.2
REVIVE_VERSION = v1.5.1
GIT_CHGLOG_VERSION = v0.15.4

VENV_PATH = ./venv
PYTHON = python3
VENV = . "$(VENV_PATH)/bin/activate";

GO_BIN_DIR := $(shell go env GOPATH)/bin
TEST_MODULES := $(shell go list ./... | grep -v -e /cmd/)
COVERAGE_DIR := $(shell pwd)/coverage

.PHONY: build
build:
	go build -ldflags "-s -w" -o ticketeer

.PHONY: build-release
build-release:
	@goreleaser release --snapshot --clean

.PHONY: build-coverage
build-coverage:
	go build -cover -ldflags "-s -w" -o ticketeer

.PHONY: install
install:
	go install

.PHONY: install-coverage
install-coverage:
	go install -cover

.PHONY: clean
clean:
	rm -rf build
	rm -rf dist
	rm -f packaging/npm/ticketeer-*/ticketeer \
		packaging/npm/ticketeer-*/ticketeer.exe
	rm -rf \
		packaging/pypi/**/*.egg-info \
		packaging/pypi/**/build \
		packaging/pypi/**/dist
	rm -f \
		packaging/pypi/ticketeer_*/ticketeer_*/ticketeer_* \
		packaging/pypi/ticketeer_*/ticketeer_*/ticketeer_*.exe \
		packaging/pypi/**/.version \
		packaging/pypi/**/README.md
	git restore packaging/npm
	git restore packaging/pypi

.PHONY: release
release: clean
	make build-release
	python3 scripts/publish.py $(VERSION)
	make changelog
	git tag -d v$(VERSION)
	git add \
		packaging/npm \
		packaging/pypi \
		Makefile \
		CHANGELOG.md
	git commit -m "chore: release $(VERSION)"
	git tag -a v$(VERSION) -m "release $(VERSION)"
	@git push && git push --tags

.PHONY: test
test:
	@go test $(TEST_MODULES)

test-e2e: install
	@go test \
		-race \
		-count=1 \
		-timeout=30s \
		-tags=e2e \
		e2e_test.go

.PHONY: coverage
coverage: install-coverage
	@rm -rf "$(COVERAGE_DIR)"
	@mkdir "$(COVERAGE_DIR)"
	@mkdir "$(COVERAGE_DIR)/e2e"
	@go test -coverprofile=coverage/unit.cover.out $(TEST_MODULES)
	@GOCOVERDIR="$(COVERAGE_DIR)/e2e" \
		go test \
			-race -count=1 -timeout=30s \
			-tags=e2e \
			e2e_test.go
	@go tool covdata percent -i="$(COVERAGE_DIR)/e2e" -o=coverage/e2e.cover.out
	@python3 scripts/combine_coverage.py \
		--output "$(COVERAGE_DIR)/cover.out" \
		"$(COVERAGE_DIR)/e2e.cover.out" \
		"$(COVERAGE_DIR)/unit.cover.out"

.PHONY: coverage-html
coverage-html: coverage
	@covreport -i coverage/cover.out -o coverage/cover.html

.PHONY: setup
setup:
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(GO_BIN_DIR) $(GOLANGCI_LINT_VERSION)
	go install github.com/mgechev/revive@$(REVIVE_VERSION)
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG_VERSION)
	go install github.com/cancue/covreport@latest
	rm -rf "$(VENV_PATH)"
	python3 -m venv "$(VENV_PATH)"
	$(VENV) pip install -r requirements.txt

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	golangci-lint run ./...
	revive -config ./revive.toml  ./...

.PHONY: report-coverage
report-coverage:
	@coveralls report \
		--repo-token=$(COVERALLS_TICKETEER_TOKEN) \
		coverage/cover.out

.PHONY: check
check:
	@make lint
	@make test
	@make build

.PHONY: changelog
changelog:
	git-chglog -o CHANGELOG.md
