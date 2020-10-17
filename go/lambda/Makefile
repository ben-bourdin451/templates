#--------------------------------------------------
# Variables
#--------------------------------------------------
PACKAGES := $$(go list ./...)
SOURCES := $(shell find . -type f -name "*.go" -a ! -name "*_test.go")
TEST_SOURCES := $(shell find . -type f -name "*_test.go")
GOFMT_FILES := $$(gofmt -l -s $(SOURCES) $(TEST_SOURCES))

#--------------------------------------------------
# Targets
#--------------------------------------------------
fmt:
	@echo "==> Fixing source code with gofmt..."
	@gofmt -s -w $(SOURCES)

fmtcheck:
	@echo "==> Checking that code complies with gofmt requirements..."
	@if [[ -n $(GOFMT_FILES) ]]; then \
		echo 'gofmt needs running on the following files:'; \
		echo "$(GOFMT_FILES)"; \
		echo "You can use the command: \`make fmt\` to reformat code."; \
		exit 1; \
	fi

test:
	@go test -race $(PACKAGES)

testacc:
	@go test -race -tags=integration $(PACKAGES)

coverage: $(SOURCES) $(TEST_SOURCES)
	@go test -cover -coverprofile=coverage $(PACKAGES)
	@go tool cover -html=coverage

clean:
	@rm coverage

.PHONY: clean fmt fmtcheck package test
