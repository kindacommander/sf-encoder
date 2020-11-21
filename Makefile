PKG := github.com/kindacommander/sf-encoder

go-build:
	@echo " > Building binary..."
	go build -o vendor/sf-encoder cmd/main.go

BIN=$(shell pwd)/vendor/sf-encoder
exec:
	@echo " > Executing..."
ifneq ("$(wildcard $(BIN))","")
	vendor/sf-encoder
else
	@echo " Error: No binary file"
endif