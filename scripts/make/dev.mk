.PHONY: status logs start stop clean

status: ## Get status of containers
	docker compose ps

logs: ## Get logs of containers	## docker compose logs --tail=0 --follow
	docker compose logs --follow

start:build ## Build and start docker containers
	docker compose up -d

stop: ## Stop docker containers
	docker compose stop

clean:stop ## Stop docker containers, clean data and workspace
	docker compose down -v --remove-orphans