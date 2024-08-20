.PHONY: logs-prod start-prod

logs-prod: ## Get logs of containers	
	docker compose logs -f

start-prod:build-prod ## Build and start docker containers
	docker compose -f docker-compose.prod.yaml up -d