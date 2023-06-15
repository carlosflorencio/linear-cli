# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Name of your Go CLI app
APP_NAME = linear
OUTPUT_DIR = bin

all: build

build:
	$(GOBUILD) -o $(OUTPUT_DIR)/$(APP_NAME)

clean:
	$(GOCLEAN)
	rm -f $(OUTPUT_DIR)/$(APP_NAME)

test:
	$(GOTEST) -v ./...

run:
	$(GOBUILD) -o $(OUTPUT_DIR)/$(APP_NAME)
	./$(OUTPUT_DIR)/$(APP_NAME)

.PHONY: all build clean test run
