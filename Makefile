SHELL = bash
default: help

HELP_FORMAT="    \033[36m%-25s\033[0m %s\n"

.PHONY: help
help: ## Display this usage information
	@echo "Valid targets:"
	@grep -E '^[^ ]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
			{printf $(HELP_FORMAT), $$1, $$2}'
	@echo ""

.PHONY: build
build:
	go build -o bin/ntui main.go
