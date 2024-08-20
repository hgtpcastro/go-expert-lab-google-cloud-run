.PHONY: build build-test build-prod

build: ## Build docker image
	docker compose build

build-test: ## Build docker image (test)
	docker compose -f docker-compose.test.yaml build	

build-prod: ## Build docker image (production)
	docker compose -f docker-compose.prod.yaml build	