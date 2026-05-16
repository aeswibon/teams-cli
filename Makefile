.PHONY: build install-local install-help test clean fmt

# Build the CLI binary
build:
	go build -o teams-cli .

# Build and install to ~/.local/bin
install-local: build install-help
	@mkdir -p ~/.local/bin
	@cp teams-cli ~/.local/bin/teams-cli
	@chmod +x ~/.local/bin/teams-cli
	@echo "✅ Installed teams-cli to ~/.local/bin/teams-cli"
	@echo "🔧 Make sure ~/.local/bin is in your PATH"

install-help:
	@cp teams-cli-token-help.sh ~/.local/bin/teams-cli-token-help
	@chmod +x ~/.local/bin/teams-cli-token-help
	@echo "✅ Installed teams-cli-token-help to ~/.local/bin/"

# Run tests
test:
	go test ./... -v

# Format code
fmt:
	go fmt ./...

# Clean build artifacts
clean:
	rm -f teams-cli

# Download dependencies
deps:
	go mod download
	go mod tidy
