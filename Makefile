build:
	go build -o ./bin/example ./_example/main.go

run: build
	./bin/example
