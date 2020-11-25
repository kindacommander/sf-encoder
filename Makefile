PKG := github.com/kindacommander/sf-encoder

go-build-encoder:
	@echo " > Building binary: encoder..."
	go build -o bin/sf-encoder cmd/encoder/main.go

go-build-benchmark:
	@echo " > Building binary: benchmark..."
	go build -o bin/encoder-benchmark cmd/benchmark/main.go

BIN=$(shell pwd)/bin/sf-encoder
exec:
	@echo " > Executing..."
ifneq ("$(wildcard $(BIN))","")
	bin/sf-encoder
else
	@echo " Error: No binary file"
endif

go-test-race:
	@echo " > Testing for data race..."
	go test test/race_test.go -race

go-test-bench:
	@echo " > Benchmarking timing..."
	go test -bench=. test/bench_test.go

go-test-benchmem:
	@echo " > Benchmarking memory usage..."
	go test -bench=. test/bench_test.go -benchmem