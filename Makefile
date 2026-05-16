.PHONY: build install-local test clean fmt deps ci

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

# Run the same checks as CI
ci:
	go mod download
	@test -z "$$(gofmt -l .)"
	@mkdir -p ./.bin
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./.bin v1.60.3; \
		PATH="$$PWD/.bin:$$PATH"; \
	fi; \
	PATH="$$PWD/.bin:$$PATH" golangci-lint run ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test ./... -count=1
