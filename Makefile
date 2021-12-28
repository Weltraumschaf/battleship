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
	rm -rfv $(BIN_DIR)

.PHONY: build
build: dependencies ## Build the project
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/battleship $(PROJECT_DIR)/cmd/battleship/main.go
	go build -o $(BIN_DIR)/server $(PROJECT_DIR)/cmd/server/main.go

	@echo "Done! Final binary is located in $(BIN_DIR)"

.PHONY: dependencies
dependencies: ## Get the Go dependencies
	go get -t -d -v ./...

.PHONY: run-battleship
run-battleship: ## Run the project's scanner binary.
	go run $(PROJECT_DIR)/cmd/battleship/main.go.

PHONY: run-server
run-server: ## Run the project's scanner binary.
	go run $(PROJECT_DIR)/cmd/server/main.go

.PHONY: test
test: ## Run the unit tests.
	go test -v ./...

.PHONY: vet
vet: ## Vet the project.
	go vet -v ./...

.PHONY: image
image: ## Build the Docker image.
	docker image build  -t "weltraumschaf/battleship:1.0.0" .

.PHONY: run-container
run-container: image ## Run the Docker container.
	docker container run -p 10000:10000 weltraumschaf/battleship:1.0.0

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
