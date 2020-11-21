PKG := github.com/kindacommander/sf-encoder

go-build:
	@echo " > Building binary..."
	go build -o bin/sf-encoder cmd/main.go

BIN=$(shell pwd)/bin/sf-encoder
exec:
	@echo " > Executing..."
ifneq ("$(wildcard $(BIN))","")
	bin/sf-encoder
else
	@echo " Error: No binary file"
endif