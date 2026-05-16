.PHONY: build install-local test clean fmt

# Build the CLI binary
build:
	go build -o teams-cli .

# Build and install to ~/.local/bin
install-local: build
	@mkdir -p ~/.local/bin
	@cp teams-cli ~/.local/bin/teams-cli
	@chmod +x ~/.local/bin/teams-cli
	@echo "✅ Installed teams-cli to ~/.local/bin/teams-cli"
	@echo "🔧 Make sure ~/.local/bin is in your PATH"

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
