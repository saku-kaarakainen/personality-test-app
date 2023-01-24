test-api:
	cd api && go test -v ./... ./internal/...

run:
	sh api/run-backend.sh &
	sh app/run-frontend.sh &