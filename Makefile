UTESTS?=$$(go list ./... | egrep -v "vendor|integration")
ITESTS?=$$(go list ./... | egrep "integration")
BINARY=journal
ENTRY=main.go

default: build

all: clean build_all install

build:
	go build -o ${BINARY} ${ENTRY}

build-all:
	GOOS=darwin GOARCH=amd64 go build -v -o ${BINARY}-mac-amd64 ${ENTRY}
	GOOS=linux GOARCH=386 go build -v -o ${BINARY}-linux-386 ${ENTRY}
	GOOS=linux GOARCH=amd64 go build -v -o ${BINARY}-linux-amd64 ${ENTRY}
	GOOS=windows GOARCH=386 go build -v -o ${BINARY}-windows-386.exe ${ENTRY}
	GOOS=windows GOARCH=amd64 go build -v -o ${BINARY}-windows-amd64.exe ${ENTRY}

install:
	go install

clean:
	go clean
	find ${ROOT_DIR} -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete
	find ${ROOT_DIR} -name '${BINARY}' -delete

test: utest itest

itest:
	go test $(VERBOSE) $(ITESTS) -failfast -cover

utest:
	go test $(VERBOSE) $(UTESTS) -failfast -cover