SHELL			=	/bin/bash
NAME			:= ntui
PACKAGE			:= github.com/shappy0/$(NAME)
UNAME			:= $(shell uname -s)
DIRNAME			:= .$(NAME)
CONFIGFILE		:= config.toml
HOMEDIR			:= ""
HELPFORMAT		:= "  \033[36m%-25s\033[0m %s\n"
VERSION			?= v1.0.0
GIT_REV 		:= $(shell git rev-parse --short HEAD)
GO_LDFLAGS 		:= "$(GO_LDFLAGS) -X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.Commit=$(GIT_REV)"
OUTPUT			:= ./bin/$(NAME)

default: help

.PHONY: help

help: ## Display help options
	@echo "Valid targets:"
	@grep -E '^[^ ]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
			{printf $(HELPFORMAT), $$1, $$2}'
	@echo ""

.PHONY: install

install: ## Install ntui
	@echo "==> Installing ntui"

ifeq ($(UNAME), Darwin)
	@if [ ! -d $(HOME)/$(DIRNAME) ]; then \
		mkdir -p $(HOME)/$(DIRNAME); \
	fi \

	@if [ ! -f $(HOME)/$(DIRNAME)/$(CONFIGFILE) ]; then \
		cp config.toml $(HOME)/$(DIRNAME)/$(CONFIGFILE); \
	fi \

else ifeq ($(UNAME), Linux)
	@if [ ! -d $(HOME)/$(DIRNAME) ]; then \
		mkdir -p $(HOME)/$(DIRNAME); \
	fi \

	@if [ ! -f $(HOME)/$(DIRNAME)/$(CONFIGFILE) ]; then \
		cp config.toml $(HOME)/$(DIRNAME)/$(CONFIGFILE); \
	fi \

else
	@if [ ! -d $(HOME)/$(DIRNAME) ]; then \
		mkdir -p $(HOME)/$(DIRNAME); \
	fi \

	@if [ ! -f $(HOME)/$(DIRNAME)/$(CONFIGFILE) ]; then \
		cp config.toml $(HOME)/$(DIRNAME)/$(CONFIGFILE); \
	fi \

endif
	@echo "==> Installation Done!"



.PHONY: build

build: ## Build ntui
	@echo "==> Building ntui"
	@go build -ldflags $(GO_LDFLAGS) -o $(OUTPUT)

	@if echo $(PATH) | grep -q "/usr/local/bin"; then \
		if cp $(OUTPUT) /usr/local/bin ; then \
			echo ""; \
			echo "Done!"; \
			exit 0; \
		else \
			echo ""; \
			if sudo cp $(OUTPUT) /usr/local/bin ; then \
				echo ""; \
				echo "Done!"; \
				exit 0; \
			fi \
		fi; \
	fi \

.PHONY: clean

clean: ## Clean build
	@echo "==> Cleaning"
	rm -f $(OUTPUT)
	@echo "Cleaning Done!"
