INSTALL_DIR ?= ~/.local/bin
BIN_NAME ?= chatlog-sanitizer

build: ## Build the binary
	go build -o $(BIN_NAME) .

install: build ## Install the binary to $(INSTALL_DIR)
	cp $(BIN_NAME) $(INSTALL_DIR)/

uninstall: ## Remove the binary from $(INSTALL_DIR)
	rm -f $(INSTALL_DIR)/$(BIN_NAME)

help: ## Show this help
	@echo ""
	@echo "Specify a command. The choices are:"
	@echo ""
	@grep -hE '^[0-9a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[0;36m%-20s\033[m %s\n", $$1, $$2}'
	@echo ""
.PHONY: help

.DEFAULT_GOAL := help
