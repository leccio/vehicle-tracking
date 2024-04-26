obu:
	@go build -o bin/obu.exe obu/main.go
	@./bin/obu.exe

receiver:
	@go build -o bin/receiver.exe receiver/main.go
	@./bin/receiver.exe

distance:
	@go build -o bin/distance.exe distance_calculator/main.go
	@./bin/distance.exe

test:
	@echo testttttt

.PHONY: obu receiver test
