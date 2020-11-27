build: 
	go build -o movies-app app/main.go

test: 
	@go test -v -cover -covermode=atomic ./...

testcoverage:
	@go test -v -cover -covermode=atomic ./... --coverprofile=coverage.out ./... && \
	go tool cover -func=coverage.out

testcoveragehtml:
	@go test -v -cover -covermode=atomic ./... --coverprofile=coverage.out ./... && \
	go tool cover -html=coverage.out

compose_up:
	@docker-compose up -d

compose_down:
	@docker-compose down && docker image prune -f

.PHONY: build test testcoverage testcoveragehtml compose_up compose_down