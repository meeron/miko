build:
	go build -o ./bin/miko .

run: build
	./bin/miko