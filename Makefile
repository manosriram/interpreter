default: buildrun

build:
	go build

run:
	./interpreter

buildrun:
	go build && ./interpreter

clean:
	rm ./interpreter
