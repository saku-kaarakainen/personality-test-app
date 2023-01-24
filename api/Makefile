.PHONY: test
test:
	go test -v ./... ./internal/...

.PHONY: run
run:
	docker build -t personality-test-api . && docker-compose up