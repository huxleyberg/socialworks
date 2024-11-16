OUTPUT = main/bootstrap
PACKAGED_TEMPLATE = packaged.yaml
TEMPLATE = template.yaml
VERSION = 0.1
S3_BUCKET := $(S3_BUCKET)
ZIPFILE = lambda.zip
SERVICE_NAME = social-works-service
MIGRATION_PATH = internal/db/migrations
POSTGRES_DB_URL := $(POSTGRES_DB_URL)
DOCKER_COMPOSE_FILE = docker-compose.yml

create-migration:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

.PHONY: ci
ci: install lint test

.PHONY: test
test:
	go test ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o ./testCoverage.html

.PHONY: clean
clean:
	rm -f $(OUTPUT)
	rm -f $(ZIPFILE)
	rm -f coverage.out
	rm -f testCoverage.html
	rm -f packaged.yaml

.PHONY: install
install:
	go get -t ./...

.PHONY: lint
lint: install
	golint -set_exit_status ./...

main:
	go build -tags lambda.norpc -o $(OUTPUT) ./cmd/$(SERVICE_NAME)-lambda/main.go

.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(MAKE) main

$(ZIPFILE): clean lambda
	zip -9 -r $(ZIPFILE) $(OUTPUT)

.PHONY: build
build: clean lambda

build-local:
	go build -o $(OUTPUT) ./cmd/$(SERVICE_NAME)/main.go

.PHONY: package
package:
	aws s3 cp open-api-integrated.yaml s3://$(S3_BUCKET)/open-api/$(SERVICE_NAME)/open-api-integrated.yaml
	aws cloudformation package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

update:
	go get -u ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o ./testCoverage.html
	open ./testCoverage.html

run: build-local
	@echo ">> Running application ..."
	POSTGRES_DB_URL=$(POSTGRES_DB_URL) \
	./$(OUTPUT)

### Docker Commands ###
.PHONY: docker-build
docker-build:
	@echo ">> Building Docker image..."
	docker build -t $(SERVICE_NAME):latest .

.PHONY: docker-up
docker-up:
	@echo ">> Starting Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: docker-down
docker-down:
	@echo ">> Stopping Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: docker-clean
docker-clean:
	@echo ">> Removing Docker containers, images, and volumes..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v
	docker rmi $(SERVICE_NAME):latest || true

.PHONY: docker-logs
docker-logs:
	@echo ">> Showing logs for Docker containers..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

.PHONY: migrate-up
migrate-up:
	@echo ">> Running migrations..."
	migrate -path $(MIGRATION_PATH) -database $(POSTGRES_DB_URL) up

.PHONY: migrate-down
migrate-down:
	@echo ">> Rolling back migrations..."
	migrate -path $(MIGRATION_PATH) -database $(POSTGRES_DB_URL) down
