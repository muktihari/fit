.PHONY: build run pprof trace took

build:
	go build -gcflags="-m=2" main.go

run:
	go build main.go -path $(path)
	./main

pprof:
	go build main.go
	./main -path $(path) -opt $(opt)
	go tool pprof -http=:8080 $(opt).pprof

trace:
	go build main.go
	./main -path $(path) -opt $(opt)
	go tool trace $(opt).out
	
took:
	go build main.go
	./main -path $(path) -opt took