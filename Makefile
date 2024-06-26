# Here you can reformat, check or build the binary
.PHONY: fmt vet tidy download test coverage doc run build development

BINARY_NAME=mobyspulse
APP_PATH=cmd/${BINARY_NAME}/main.go

fmt:
	@go fmt ./...

vet:
	@go vet ./...

tidy:
	@go mod tidy

download:
	@go mod download

test:
	@go test -v ./... -bench=.

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060

run:
	@go run ${APP_PATH}

build:
	go mod download
	go build -o ${BINARY_NAME} ${APP_PATH}

development:
	docker compose watch