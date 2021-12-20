# Disable built-in rules and variables because we do not need them.
# - https://www.gnu.org/software/make/manual/html_node/Catalogue-of-Rules.html#Catalogue-of-Rules
# - https://www.gnu.org/software/make/manual/html_node/Implicit-Variables.html#Implicit-Variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables

all: build

PROJECT_DIR 	= $(shell pwd)
BIN_DIR				= $(PROJECT_DIR)/bin

.PHONY: clean ## Clean the project.
clean:
	@echo "Deleting files ..."
	rm -rfv $(BIN_DIR)

.PHONY: build
build: ## Build the project
	@echo "Building project ..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/battleship $(PROJECT_DIR)/cmd/battleship/main.go
	go build -o $(BIN_DIR)/server $(PROJECT_DIR)/cmd/server/main.go

	@echo "Done!"
	@echo "Final binary is located in $(BIN_DIR)"

.PHONY: run-battleship
run-battleship: ## Run the project's scanner binary.
	go run $(PROJECT_DIR)/cmd/battleship/main.go.

PHONY: run-server
run-server: ## Run the project's scanner binary.
	go run $(PROJECT_DIR)/cmd/server/main.go

.PHONY: test
test: build ## Run the unit tests.
	@echo "Testing project ..."
	go test -v ./...

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
