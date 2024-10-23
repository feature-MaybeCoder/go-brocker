run:
	@echo false
	go build -o bin/main cmd/cli/main.go
	./bin/main