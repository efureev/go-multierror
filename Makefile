TEST?=./...

default: test

# test runs the test suite and vets the code.
test: generate
	@echo "==> Running tests..."
	@go list $(TEST) \
		| grep -v "/vendor/" \
		| xargs -n1 go test -timeout=60s -parallel=10 ${TESTARGS}

# testrace runs the race checker
testrace: generate
	@echo "==> Running tests (race)..."
	@go list $(TEST) \
		| grep -v "/vendor/" \
		| xargs -n1 go test -timeout=60s -race ${TESTARGS}

# generate runs `go generate` to build the dynamically generated source files.
generate:
	@echo "==> Generating..."
	@find . -type f -name '.DS_Store' -delete
	@go list ./... \
		| grep -v "/vendor/" \
		| xargs -n1 go generate

.PHONY: default test testrace updatedeps generate
