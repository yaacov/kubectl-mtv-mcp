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

# The binary will be in the current directory
./kubectl-mtv-mcp --version
```

## Development Commands

### Building

```bash
# Build for current platform
make build

# Install to GOPATH/bin
make install

# Build for specific platform
make build-linux-amd64      # Linux AMD64
make build-linux-arm64      # Linux ARM64
make build-darwin-amd64     # macOS Intel
make build-darwin-arm64     # macOS Apple Silicon
make build-windows-amd64    # Windows AMD64

# Build for all platforms
make build-all
```

### Code Quality

```bash
# Install development tools (first time only)
make install-tools

# Format code
make fmt

# Run linters
make lint

# Run all checks (format + lint)
make verify
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# Run tests with verbose output
go test -v ./...
```

### Running Locally

```bash
# Run server in stdio mode
make run-kubectl-mtv-mcp

# Or run directly
go run ./cmd/kubectl-mtv-mcp

# Run in SSE mode for testing
go run ./cmd/kubectl-mtv-mcp --sse --host 127.0.0.1 --port 8080
```

### Distribution

```bash
# Create release archives with checksums
make dist-all

# Create archive for current platform
make dist

# Clean build artifacts
make clean

# Deep clean (includes Go cache)
make clean-all
```

## Testing with MCP Clients

### Testing with Claude Desktop

1. Build and install:
   ```bash
   make build
   make install
   ```

2. Configure Claude Desktop to use your local build

3. Restart Claude Desktop

4. Test your new tool

### Testing with Cursor

1. Build and install:
   ```bash
   make build
   make install
   ```

2. Update Cursor's MCP config to point to your binary

3. Reload Cursor's MCP configuration

4. Test in Cursor

## License

Apache License 2.0 - See LICENSE file for details.

