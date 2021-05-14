build: clean
	go build -o build/arctic main.go

run: build
	./build/server

dev:
	air -c .air.toml

clean:
	rm -rf build/

install: build
	cp build/arctic ~/bin/

