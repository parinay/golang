GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOGET=$(GO) get
BINNAMES=hello-stripped
BINAME=hello-normal
all: clean build

build: 
	GOOS=linux $(GOBUILD) -o $(BINNAMES) -ldflags="-s -w" hello.go
	GOOS=linux $(GOBUILD) -o $(BINAME) hello.go
clean:
	$(GOCLEAN)
	rm -f $(BINNAMES)
	rm -f $(BINAME)

