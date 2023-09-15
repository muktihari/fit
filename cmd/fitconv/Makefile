.PHONY: build run pprof trace took

build:
	go build -gcflags="-m=2" main.go

run:
	go build main.go
	./main

pprof:
	go build main.go
	./main -opt $(lastword $(MAKECMDGOALS))
	go tool pprof -http=:8080 $(lastword $(MAKECMDGOALS)).pprof

trace:
	go build main.go
	./main -opt $(lastword $(MAKECMDGOALS))
	go tool trace $(lastword $(MAKECMDGOALS)).out
	
took:
	go build main.go
	./main -opt took