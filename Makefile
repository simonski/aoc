default_target: build
.PHONY : default_target upload

help:
	@echo "The aoc go Makefile"
	@echo ""
	@echo "Usage : make <command> "
	@echo ""
	@echo "commands"
	@echo ""
	@echo "  clean                 - cleans temp files"
	@echo "  test                  - builds and runs tests"
	@echo "  build                 - creates binary"

	@echo "  install               - builds and installs"
	@echo "  docker                - creates aoc docker image"
	@echo "  publish               - pushes aoc image to dockerhub"
	@echo "  release               - gorelaser create binary releases"
	@echo ""
	@echo "  server                - runs the aoc server in production mode"
	@echo "  devserver             - runs the aoc server in developent mode"	
	@echo "                          (serves files from disk)."
	@echo ""

setup:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go get github.com/simonski/bn

format:
	staticcheck ./...

clean:
	go clean
	rm -rf dist
	
build:
	mkdir -p bin
	bn revision
	go fmt ./...
	go build -o bin/aoc

server: build
	./bin/aoc server

devserver: build
	./bin/aoc server -fs ./api

test:
	# go test ./app/aoc2020 -timeout 10s
	go test ./... -timeout 10s

install:
	go install

release:
	goreleaser --snapshot --skip-publish --rm-dist

docker: build
	GOOS=linux GOARCH=amd64 go build -o bin/aoc_linux
	docker build -t aoc .

publish: 
	docker tag aoc simongauld/aoc:latest
	docker push simongauld/aoc:latest
