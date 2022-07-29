.PHONY: build 
build: clean
	go build -o build/scaffold main.go

run: build
	./build/scaffold

clean:
	rm -rf build/

install: build
	cp build/scaffold ~/bin/

