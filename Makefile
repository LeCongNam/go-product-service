run:
	go run cmd/server/main.go

dev:
	air

wire_build:
	wire ./internal/di/wire.go