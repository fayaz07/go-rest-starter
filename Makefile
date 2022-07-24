setup:
	bash repo_setup/scripts/setup_project.sh

dev:
	go run src/main.go

all: install

install:
	echo "Installing go modules..." && \
	go mod download && \
	echo "Completed" 

build:
	go build -v ./...

# server
run: 
	reflex -r "\.go$" -s -- sh -c "go run src/main.go" 

clean_cache:
	go clean -testcache	

test: clean_cache
	go test ./src/... -coverprofile=coverage.out

test_cov:  clean_cache
	go test ./src/... -coverprofile=coverage.out; go tool cover -html=coverage.out
