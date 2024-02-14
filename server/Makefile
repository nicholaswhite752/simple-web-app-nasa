generate-docs:
	swag init -g ./cmd/api/main.go

build-docker: generate-docs
	docker build -t web-api:latest -f Dockerfile .

build-docker-no-docs:
	docker build -t web-api:latest -f Dockerfile .

# You need to run this before you can run generate-docs or build-docker. It 
# installs the command line tool used to generate the Swagger docs
install-go-swag:
	go install github.com/swaggo/swag/cmd/swag@latest