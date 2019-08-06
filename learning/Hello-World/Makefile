GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
GORUN=$(GO) run
BINNAMES=hello-stripped
BINAME=hello-normal
SRCNAME=hello.go

all: clean build test run
test:
	$(GOTEST)
build:
	GOOS=linux $(GOBUILD) -o $(BINNAMES) -ldflags="-s -w" $(SRCNAME)
	GOOS=linux $(GOBUILD) -o $(BINAME) $(SRCNAME)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINNAMES)
	rm -f $(BINAME)
run: build
	./$(BINAME)

default: all
