.PHONY: build build-ea run pprof pprof-cpu pprof-mem trace took

build:
	go build main.go

build-ea: # escape analysis
	go build -gcflags="-m=2" main.go

run:
	make build
	./main $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

pprof:
	make build
	./main -opt $(opt) $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
	go tool pprof -http=:8080 $(opt).pprof

pprof-cpu:
	make pprof opt=cpu $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

pprof-mem:
	make pprof opt=mem $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

trace:
	make build
	./main -opt trace $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
	go tool trace trace.out
	
took:
	make build
	./main -opt took $(wordlist 2, $(words $(MAKECMDGOALS)),$(MAKECMDGOALS))