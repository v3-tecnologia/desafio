DATABASE_HOST ?= localhost
DATABASE_PORT ?= 5432
DATABASE_USER ?= postgres
DATABASE_PASSWORD ?= postgres
DATABASE_SSL ?= disable
DATABASE_DATABASE = g3-db
DATABASE_DSN := "postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DATABASE}?sslmode=${DATABASE_SSL}"
MIGRATIONS_PATH="db/migrations"

###################
# Database        #
###################
.PHONY: mig-up
mig-up: ## Runs the migrations up
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} up

.PHONY: mig-down
mig-down: ## Runs the migrations down
	migrate -path ${MIGRATIONS_PATH} -database ${DATABASE_DSN} down

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(NAME)
