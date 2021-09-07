GO_UNIT_TEST_FILES	= $(shell go list ./... | grep -v /feature)

.PHONY: test.cleancache
test.cleancache:
	go clean -testcache

.PHONY: test.unit
test.unit: test.cleancache
	go test -v -race $(GO_UNIT_TEST_FILES)

.PHONY: migration
migration:
	migrate create -ext sql -dir db/migrations $(name)

.PHONY: migrate
migrate:
	migrate -path db/migrations -database "$(url)?sslmode=disable" -verbose up

.PHONY: rollback
rollback:
	migrate -path db/migrations -database "$(url)?sslmode=disable" -verbose down 1

.PHONY: rollback.all
rollback.all:
	migrate -path db/migrations -database "$(url)?sslmode=disable" -verbose down -all

.PHONY: migrate.force
migrate.force:
	migrate -path db/migrations -database "$(url)?sslmode=disable" -verbose force $(version)