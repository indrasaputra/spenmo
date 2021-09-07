GO_UNIT_TEST_FILES	= $(shell go list ./... | grep -v /feature)

.PHONY: test.cleancache
test.cleancache:
	go clean -testcache

.PHONY: test.unit
test.unit: test.cleancache
	go test -v -race $(GO_UNIT_TEST_FILES)
