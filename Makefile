build:
	GOOS=linux go build -o bin/main

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

run:
	sls invoke -f hello
