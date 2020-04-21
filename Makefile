.PHONY: deps clean format lint build

all: clean format lint build

lint:
	    golangci-lint run ./...

format:
	    go fmt ./...

deps:
	    go get -u ./...

clean: 
	    rm -rf ./bin/
		    
build:
	    go build -o ./bin/np ./

