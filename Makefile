obu:
	@go build -o bin/obu.exe obu/main.go
	@./bin/obu.exe

receiver:
	@go build -o bin/receiver.exe receiver/main.go
	@./bin/receiver.exe

test:
	@echo testttttt

.PHONY: obu receiver test
