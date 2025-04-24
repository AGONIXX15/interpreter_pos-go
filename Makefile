
build:
	go build -o test/interpreter


install: build
	sudo mv test/interpreter  /usr/local/bin/interpreter

test: build
	test/interpreter test/main.pos
