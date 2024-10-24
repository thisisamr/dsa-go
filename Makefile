build:
	@echo "building project ... please wait"
	@go build -o ./bin/main.exe ./cmd/main.go


run:
	@./bin/main.exe
