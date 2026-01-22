# Building from Source

Instructions for building, developing, and contributing to kubectl-mtv MCP server.

## Prerequisites

- Go 1.21 or later
- make
- git

## Quick Build

Build for your current platform:

```bash
# Clone the repository
git clone https://github.com/yaacov/kubectl-mtv-mcp.git
cd kubectl-mtv-mcp

# Build the server
make build

# The binary will be in the bin directory
./bin/kubectl-mtv-mcp --version
```

## Development Commands

Run `make help` to see all available targets:

```
kubectl-mtv-mcp Makefile

Usage: make [target]

Targets:
  help             Show this help message
  build            Build the binary for current platform
  install          Install binary to GOPATH/bin
  run              Build and run the MCP server
  test             Run tests with coverage
  lint             Run linters (go vet + golangci-lint)
  fmt              Format Go source files
  tidy             Tidy Go modules
  verify           Run fmt, lint, and test
  clean            Remove build artifacts
  install-tools    Install development tools (golangci-lint)
  dist             Build release archive for linux-amd64
```

### Building

```bash
# Build for current platform
make build

# Install to GOPATH/bin
make install
```

### Code Quality

```bash
# Install development tools (first time only)
make install-tools

# Format code
make fmt

# Run linters
make lint

# Run all checks (format + lint + test)
make verify
```

### Testing

```bash
# Run tests
make test

# Run tests with verbose output
go test -v ./...
```

### Running Locally

```bash
# Build and run server in stdio mode
make run

# Or run directly
go run .

# Run in SSE mode for testing
go run . --sse --host 127.0.0.1 --port 8080
```

### Cleanup

```bash
# Clean build artifacts
make clean
```

## Testing with MCP Clients

### Testing with Claude Desktop

1. Build and install:
   ```bash
   make install
   ```

2. Configure Claude Desktop to use your local build

3. Restart Claude Desktop

4. Test your new tool

### Testing with Cursor

1. Build and install:
   ```bash
   make install
   ```

2. Update Cursor's MCP config to point to your binary

3. Reload Cursor's MCP configuration

4. Test in Cursor

## License

Apache License 2.0 - See LICENSE file for details.
