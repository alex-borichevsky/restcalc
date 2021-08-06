BIN_FOLDER=bin
IGNORED_FOLDER=.ignore
COVERAGE_FILE=$(IGNORED_FOLDER)/coverage.out


all: tools install lint swag test build ## Start all

.PHONY: all


##
## Quality Code
##

lint: ## Lint
	@golangci-lint run


test: ## Test
	@mkdir -p ${IGNORED_FOLDER}
	@go test -count=1 -race -coverprofile=${COVERAGE_FILE} -covermode=atomic ./...

cover: ## Cover
	@if [ ! -e ${COVERAGE_FILE} ]; then \
		echo "Error: ${COVERAGE_FILE} doesn't exists. Please run \`make test\` then retry."; \
		exit 1; \
	fi
	@go tool cover -func=${COVERAGE_FILE}

.PHONY: lint test cover


##
## Tools
##

tools-lint: ## Install go lint dependencies
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

tools-docs: ## Install go docs dependencies
	@go install github.com/swaggo/swag/cmd/swag@latest

tools: tools-lint tools-docs  ## Install tools

.PHONY: tools tools-lint tools-docs


##
## Docs
##
swag: ## Generate swagger files
	@swag init --parseDependency --parseDepth=2 -g main.go -o ./docs

.PHONY: swag


##
## Building
##

install: ## Download and install go mod
	@go mod download

build: ## Build app
	@go build -a -o ${BIN_FOLDER}/app

run:
	@go run ${BIN_FOLDER}/app

.PHONY: install build run

##
## Files management
##

clean: ## Clean
	@rm -rf ${IGNORED_FOLDER}
	@rm -rf ${COVERAGE_FILE}
	@rm -rf ${BIN_FOLDER}