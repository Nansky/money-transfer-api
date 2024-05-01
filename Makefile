DOCKER_COMPOSE_FILE ?= docker-compose.yml

#========================#
#== DATABASE MIGRATION ==#
#========================#
migrate-up: ## Run migrations UP
migrate-up:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

migrate-down: ## Rollback migrations against non test DB
migrate-down:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

run:
run:
	docker compose -f ${DOCKER_COMPOSE_FILE} up
