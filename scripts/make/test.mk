.PHONY: status-test logs-test start-test stop-test clean-test exec-test test

status-test: ## Get status of containers
	docker compose -f docker-compose.test.yaml ps

logs-test: ## Get logs of containers	## docker compose logs --tail=0 --follow
	docker compose -f docker-compose.test.yaml logs --follow

start-test:build-test ## Build and start docker containers
	docker compose -f docker-compose.test.yaml up -d
	docker container ls -la

stop-test: ## Stop docker containers
	docker compose -f docker-compose.test.yaml stop

clean-test:stop-test ## Stop docker containers, clean data and workspace
	docker compose -f docker-compose.test.yaml down -v --remove-orphans

exec-test: ## Execute test suite
	docker compose -f docker-compose.test.yaml exec weather_service_test gotest -v -race -coverprofile=/tmp/coverage.out ./...		

test: start-test exec-test clean-test ## Run test suite	