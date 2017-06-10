NAME=noface
SOURCE=cmd/noface/main.go

GOBUILD=go build

DEPEND=github.com/Masterminds/glide

# Command to get glide, you need to run it only once
.PHONY: get_glide
get_glide:
	go get -u -v $(DEPEND)
	$(GOPATH)/bin/glide install

# Command to install dependencies using glide
.PHONY: install_dependencies
install_dependencies:
	glide install

# Run tests in verbose mode with race detector and display coverage
.PHONY: test
test:
	go test -v -cover -race $(shell glide novendor)

# Removing artifacts
.PHONY: clean
clean:
	$(info * Cleaning build folder)
	@rm -rf build/*

# Building linux binaries
.PHONY: build_linux
build_linux:
	$(info * Cleaning build folder)
	@rm -rf build/*
	$(info * Building executable for linux x64 [$(SOURCE) -> build/linux_x64/$(NAME)])
	@GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/linux_x64/$(NAME) $(SOURCE)

# Building osx binaries
.PHONY: build_osx
build_osx:
	$(info * Cleaning build folder)
	@rm -rf build/*
	$(info * Building executable for osx x64 [$(SOURCE) -> build/darwin_amd64/$(NAME)])
	@GOOS=darwin GOARCH=amd64 $(GOBUILD) -o build/darwin_amd64/$(NAME) $(SOURCE)

# Run the application
.PHONY: run
run:
	go run cmd/noface/main.go
