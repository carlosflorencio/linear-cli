# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Name of your Go CLI app
APP_NAME = linear

all: build

build:
	$(GOBUILD) -o $(APP_NAME)

clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(APP_NAME)
	./$(APP_NAME)

.PHONY: all build clean test run
