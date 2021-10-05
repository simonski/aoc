default_target: build
.PHONY : default_target upload

usage:
	@echo "The aoc2020 go Makefile"
	@echo ""
	@echo "Usage : make <command> "
	@echo ""
	@echo "commands"
	@echo ""
	@echo "  clean                 - cleans temp files"
	@echo "  test                  - builds and runs tests"
	@echo "  build                 - creates binary"
	@echo "  install               - builds and installs"
	@echo ""
	@echo "  all                   - all of the above"
	@echo ""

# init:
# 	go get github.com/rakyll/statik

clean:
	go clean
	
# statik:
# 	statik -src=data -include=*.txt

build:
	go build
	
test:
	go test

install:
	go install

all: clean build test
	go install
