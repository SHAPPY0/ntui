SHELL 		 = bash
NAME 		:= ntui11
PACKAGE 	:= github.com/shappy0/$(NAME)
UNAME 		:= $(shell uname -s)
DIR_NAME	 := .$(NAME)
CONFIG_FILE	 := config1.json
HOME_DIR	 := ""
HELP_FORMAT	 := "  \033[36m%-25s\033[0m %s\n"

default: help

.PHONY: help
help: ## Display help options
	@echo "Valid targets:"
	@grep -E '^[^ ]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
			{printf $(HELP_FORMAT), $$1, $$2}'
	@echo ""

.PHONY: install
install: ## Install ntui
	@echo "==> Installing ntui"
ifeq ($(UNAME), Darwin)
	@echo Mac
else ifeq ($(UNAME), Linux)
	HOMEDIR := $(HOME)/$(DIR_NAME)
	if [ ! -f "$(HOMEDIR)/$(CONFIG_FILE)" ]; then \
		@mkdir -p $(HOMEDIR)/$(CONFIG_FILE) \
		@echo okas > "$(HOMEDIR)/$(CONFIG_FILE)"; \
	fi
else

endif



.PHONY: build
build: ## Build ntui
	go build -o bin/ntui main.go