GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
GORUN=$(GO) run
BINAME=size
SRCNAME=main.go

all: clean build tests run

tests: test cover

test:
	$(GOTEST)  -v ./Size/
cover:
	#$(GOTEST) -cover -v ./Size/
	$(GOTEST) -coverprofile=coverage.out -v ./Size/
	$(GO) tool cover -func=coverage.out

build:
	GOOS=linux $(GOBUILD) -o $(BINAME) $(SRCNAME)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINNAMES)
	rm -f $(BINAME)
run: build
	./$(BINAME)

default: all
