GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
GORUN=$(GO) run
BINAME=math
SRCNAME=math.go

.PHONY: all
all: clean build tests run

.PHONY: tests
tests: test cover

.PHONY: test
test:
	@echo "### test:  run tests  ###"
	$(GOTEST)  -v ./...

.PHONY: cover
cover:
	@echo "### cover:  run coverage  ###"
	#$(GOTEST) -cover -v ./Size/
	@echo "### go test - set - did each statement run? ###"
	$(GOTEST) -covermode=set -coverprofile=coverage.out -v ./...
	@echo "### go test - count - how many times did each statement run? ###"
	$(GOTEST) -covermode=count -coverprofile=coverage.out -v ./...
	@echo "### go test - atomic - like count but counts precisely in parallel programs -did each statement run? ###"
	$(GOTEST) -covermode=atomic -coverprofile=coverage.out -v ./...
	$(GO) tool cover -func=coverage.out
	$(GO) tool cover -html=coverage.out

.PHONY: build
build:
	@echo "### build:  the binary  ###"
	GOOS=linux $(GOBUILD) -o $(BINAME) $(SRCNAME)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BINNAMES)
	rm -f $(BINAME)
.PHONY: run
run: build
	@echo "### run:  And here is the output  ###"
	./$(BINAME)

default: all
