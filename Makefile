.PHONY: run
GO := $(shell which go)

run: 
	@$(GO) run main.go
