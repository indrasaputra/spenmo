GO_UNIT_TEST_FILES	= $(shell go list ./... | grep -v /feature)
PROTOGEN_IMAGE 		= indrasaputra/protogen:2021-09-07
ENT_SCHEMA_DIR		= ./internal/repository/model/ent/schema

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: format
format:
	bin/format.sh

.PHONY: lint.cleancache
lint.cleancache:
	golangci-lint cache clean

.PHONY: lint
lint: lint.cleancache
	buf lint
	golangci-lint run ./...

.PHONY: pretty
pretty: tidy format lint

.PHONY: check.import
check.import:
	bin/check-import.sh

.PHONY: test.cleancache
test.cleancache:
	go clean -testcache

.PHONY: test.unit
test.unit: test.cleancache
	go test -v -race $(GO_UNIT_TEST_FILES)

.PHONY: test.cover
test.cover:
	go test -v -race $(GO_UNIT_TEST_FILES) -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out

.PHONY: test.coverhtml
test.coverhtml:
	go test -v -race $(GO_UNIT_TEST_FILES) -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: test.integration
test.integration:
	bin/godog.sh

.PHONY: gen.mock
gen.mock:
	bin/generate-mock.sh

.PHONY: gen.proto
gen.proto:
	bin/generate-proto.sh

.PHONY: gen.proto.docker
gen.proto.docker:
	docker run -it --rm \
    --mount "type=bind,source=$(PWD),destination=/work" \
    --mount "type=volume,source=spenmo-go-mod-cache,destination=/go,consistency=cached" \
    --mount "type=volume,source=spenmo-buf-cache,destination=/home/.cache,consistency=cached" \
    -w /work $(PROTOGEN_IMAGE) make -e -f Makefile gen.proto pretty

.PHONY: init.model
init.model:
	ent init --target $(ENT_SCHEMA_DIR) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: gen.model
gen.model:
	ent generate $(ENT_SCHEMA_DIR)

.PHONY: compile
compile:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o spenmo cmd/api/main.go

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

.PHONY: validate.migration
validate.migration:
	bin/validate-migration.sh

%:
	@: