ANTLR_VERSION=antlr-4.10.1-complete

GO_TEST_MODULES = $(shell cd go && go list ./...)

# Mockery can cause a circular dependency with modules so disable them as they're not required
MOCKERY_CMD := GOFLAGS="" mockery
MOCKERY_ARGS := --all --keeptree

.PHONY: generate-parser
generate-parser: go-generate-parser

# lint: Runs linting on the library
.PHONY: lint
lint: go-lint

# lint: Runs formatting tools on the library
.PHONY: fmt
fmt: go-fmt

.PHONY: mocks
mocks: go-mocks

# lint: Runs pre-commit checks on the repo
.PHONY: checks
checks: go-tidy generate-parser fmt mocks

.PHONY: build
build: go-build

.PHONY: go-generate-parser
go-generate-parser:
	rm -rf go/internal/parser/antlr ; \
	java -jar /usr/bin/$(ANTLR_VERSION).jar -Dlanguage=Go -o ./go/internal/parser ./antlr/XMLLexer.g4 && \
	java -jar /usr/bin/$(ANTLR_VERSION).jar -Dlanguage=Go -o ./go/internal/parser -lib ./go/internal/parser/antlr ./antlr/XMLParser.g4

.PHONY: go-build
go-build: go-generate-parser go-download
	cd go && \
	go build -o ./dist/ ./cmd/...

.PHONY: go-unit
go-unit: go-generate-parser go-download
	cd go && \
	go test $(GO_TEST_MODULES) \
		-syntax-test-dir="../../../test-resources/syntax" \
		-semantic-test-dir="../../../test-resources/semantic" \
		-cover \
		-coverprofile=c.out \
		-count=1 && \
	awk 'BEGIN {cov=0; stat=0;} $$3!="" { cov+=($$3==1?$$2:0); stat+=$$2; } \
	END {printf("Total coverage: %.2f%% of statements\n", (cov/stat)*100);}' c.out && \
	go tool cover -html=c.out -o unit_test_coverage.html

.PHONY: go-tidy
go-tidy:
	cd go && \
	go mod tidy

.PHONY: go-download
go-download:
	cd go && \
	go mod download

.PHONY: go-lint
go-lint:
	cd go && \
	golangci-lint run

.PHONY: go-fmt
go-fmt:
	cd go && \
	gofmt -s -w -e ./internal && \
	gci --local github.hpe.com --write ./internal && \
	goimports -local github.hpe.com -w ./internal

.PHONY: go-mocks
go-mocks:
	cd go ; \
	rm -rf mocks_maketemp ; \
	# Mockery returns error code 0 on these errors but produces incorrect output \
	if $(MOCKERY_CMD) $(MOCKERY_ARGS) --output mocks_maketemp 2>&1 | grep ERR ; then \
		rm -rf mocks_maketemp ; \
		exit 1 ; \
	fi ; \
	rm -rf mocks ; \
	rm -rf internal/mocks ; \
	if [ -d mocks_maketemp ]; then \
		mv mocks_maketemp mocks ; \
	fi ; \
	if [ -d mocks/internal ]; then \
		mv mocks/internal internal/mocks ; \
	fi
