all: run

build:
	@go build -o ./bin/

run: build
	@./bin/unit_converter --watch