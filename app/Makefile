.PHONY: test
test:
	npm ci && CI=true npm run test

.PHONY: run
run:
	docker build -t personality-test-app . && docker-compose up