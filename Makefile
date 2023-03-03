#!make
.SILENT:
.DEFAULT_GOAL := help

help: ## Show this help
	@echo "Usage:\n  make <target>\n"
	@echo "Targets:"
	@grep -h -E '^[a-zA-Z_-].+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'


compose.run: ## Start with docker-compose
	docker-compose up --build

compose.stop: ## Stop docker-compose
	docker-compose down

deps: ## Download dependencies
	go mod download && go mod tidy

lint: ## Check code (used golangci-lint)
	GO111MODULE=on golangci-lint run


test: ## Run tests
	go clean --testcache
	go test ./...

release: ## Git tag create and push
	git tag -s -a v${tag} -m 'chore(release): v$(tag) [skip ci]'
	git push origin v${tag}

release.revert: ## Revert git release tag
	git tag -d v${tag}
	git push --delete origin v${tag}

# logger config
export LOGGER_CALLER=false
export LOGGER_STACKTRACE=true
export LOGGER_LEVEL=debug
# database config
export DB_HOST=localhost:3306
export DB_USER=test_mysql
export DB_PASSWORD=a2s_kjlasjd
export DB_BASE=test
# exchange config
export EXCHANGE_URL=wss://ws-feed.exchange.coinbase.com
export EXCHANGE_ORIGIN=https://coinbase.com
export EXCHANGE_PROTOCOL=
export EXCHANGE_SYMBOLS=ETH-BTC,BTC-USD,BTC-EUR
export EXCHANGE_CHANNELS=ticker

run: ## Run application local
	go run cmd/app/main.go