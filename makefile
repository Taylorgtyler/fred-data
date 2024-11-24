build:
	@echo "Building fred-data"

	@go build -o fred-data cmd/api/main.go

run:
	@go run cmd/api/main.go
