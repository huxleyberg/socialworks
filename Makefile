OUTPUT = main/bootstrap
PACKAGED_TEMPLATE = packaged.yaml 
TEMPLATE = template.yaml
VERSION = 0.1
S3_BUCKET := $(S3_BUCKET)
ZIPFILE = lambda.zip
SERVICE_NAME = socialworks
MIGRATION_PATH=internal/db/migrations
POSTGRES_DB_URL := $(POSTGRES_DB_URL)

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
	# go get -u honnef.co/go/tools/cmd/megacheck
	go get -u golang.org/x/lint/golint

local-install:
	go get -u github.com/awslabs/aws-sam-local

.PHONY: lint
lint: install
	# megacheck -go $(VERSION)
	golint -set_exit_status

main:
	go build -tags lambda.norpc -o $(OUTPUT) ./cmd/$(SERVICE_NAME)-lambda/main.go

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(MAKE) main

# create a lambda deployment package
$(ZIPFILE): clean lambda
	zip -9 -r $(ZIPFILE) $(OUTPUT)

.PHONY: run-local
local-deploy: local-install
	aws-sam-local local start-api

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