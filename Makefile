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
	go build src/main.go

# server
run: 
	reflex -r "\.go$" -s -- sh -c "go run src/main.go" 
