GOCMD        = go
GOBUILD      = $(GOCMD) build
GOCLEAN      = $(GOCMD) clean
GOTEST       = $(GOCMD) test
GOVET        = $(GOCMD) vet
GOGET        = $(GOCMD) get
GOX          = $(GOPATH)/bin/gox
GOGET        = $(GOCMD) get
GOINSTALL    = $(GOCMD) install 

GOX_ARGS     = -output="$(BUILD_DIR)/{{.Dir}}-{{.OS}}-{{.Arch}}" -osarch="linux/amd64 darwin/amd64"

BUILD_DIR    = build
EMOJIDB_DIR  = emojidb
BINARY_NAME  = emojify-ipv6

all: clean vet test build

build:
	${GOGET} -u github.com/go-bindata/go-bindata/...
	go-bindata -pkg emojidb -o emojidb/emoji.go emojidb/emoji.json
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v -ldflags "-X main.gitTag=`git describe --tags --abbrev=0`"

vet:
	${GOVET} ./...

test:
	${GOTEST} ./...

coverage:
	${GOTEST} -coverprofile=coverage.txt -covermode=atomic ./...

clean:
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/*
	rm -f $(EMOJIDB_DIR)/emojidb.go

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

install: 
	$(GOINSTALL) -ldflags "-X main.gitTag=`git describe --tags --abbrev=0`"

uninstall:
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

release:
	${GOGET} -u github.com/mitchellh/gox
	${GOX} -ldflags "${LD_FLAGS}" ${GOX_ARGS}
	shasum -a 512 build/* > build/sha512sums.txt

docker:
	docker build --rm --force-rm --no-cache -t kkh913/emojify-ipv6 .

.PHONY: all vet test coverage clean build run release docker
