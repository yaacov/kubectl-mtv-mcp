# kubectl-mtv MCP Server

Model Context Protocol (MCP) server that provides AI assistants with comprehensive tools to interact with Migration Toolkit for Virtualization (MTV) through kubectl-mtv commands.

## What it does

This MCP server enables AI assistants to help with MTV operations by providing:

- **Read operations**: Safe operations for monitoring, troubleshooting, and discovering MTV resources and provider inventories
- **Write operations**: **USE WITH CAUTION** - Full lifecycle management including creating, modifying, and deleting MTV resources

## Prerequisites

- `kubectl-mtv` binary installed and available in PATH (for MTV operations)
- Access to a Kubernetes cluster with MTV deployed

## Setup & Usage

For detailed setup instructions with MCP clients like Cursor and Claude Desktop, see **[MCP_SETUP.md](MCP_SETUP.md)**.

### Quick Start with Claude Code

```bash
# Add the kubectl-mtv MCP server (includes both read and write operations)
claude mcp add kubectl-mtv kubectl-mtv-mcp
```

## Server Overview

### kubectl-mtv-mcp (19 tools)

Comprehensive MTV operations including both read and write capabilities:

#### Read Operations (Safe for all users)

| Tool | Description |
|------|-------------|
| `ListResources` | List MTV resources (providers, plans, mappings, hosts, hooks) |
| `ListInventory` | Query provider inventories with SQL-like syntax |
| `GetLogs` | Retrieve logs from MTV controller and importer pods |
| `GetMigrationStorage` | Access migration PVCs and DataVolumes |
| `GetVersion` | Get kubectl-mtv and MTV operator version information |
| `GetPlanVms` | Get VMs and their status from migration plans |

#### Write Operations (USE WITH CAUTION)

| Tool | Description |
|------|-------------|
| `ManagePlanLifecycle` | Start, cancel, cutover, archive, unarchive plans |
| `CreateProvider` | Create source virtualization provider connections |
| `ManageMapping` | Create, delete, patch network and storage mappings |
| `CreatePlan` | Create comprehensive migration plans |
| `CreateHost` | Create migration hosts for direct data transfer |
| `CreateHook` | Create migration hooks for custom automation |
| `DeleteProvider` | Delete one or more providers |
| `DeletePlan` | Delete one or more migration plans |
| `DeleteHost` | Delete one or more migration hosts |
| `DeleteHook` | Delete one or more migration hooks |
| `PatchProvider` | Patch/modify existing providers |
| `PatchPlan` | Patch/modify migration plans |
| `PatchPlanVm` | Patch VM-specific fields in plans |

## Building

### Quick Build

```bash
# Build all servers for current platform
make build

# Build individual server
make build-kubectl-mtv-mcp
```

### Multi-Architecture Builds

```bash
# Build for specific platforms
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
# Install development tools (first time)
make install-tools

# Format code
make fmt

# Run linters
make lint

# Run all checks
make verify
```

### Distribution

```bash
# Create release archives with checksums
make dist-all

# Create local distribution
make dist
```

## Development

### Running Locally

```bash
# Run server
make run-kubectl-mtv-mcp
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage
```

### Build Optimizations

The Makefile includes production-ready optimizations:

- **Binary size**: 5.2MB (33% smaller than unoptimized)
- **Stripped symbols**: Debug info removed for smaller binaries
- **Version embedding**: Git tags automatically embedded
- **Reproducible builds**: Consistent flags across all platforms
- **Silent output**: Minimal output following Unix philosophy

### Installation

```bash
# Install to GOPATH/bin
make install

# Clean build artifacts
make clean

# Deep clean (includes Go cache)
make clean-all
```

## MCP Client Configuration

### Cursor IDE

Add to your Cursor settings:

```json
{
  "mcpServers": {
    "kubectl-mtv": {
      "command": "kubectl-mtv-mcp",
      "args": []
    }
  }
}
```

### Claude Desktop

Edit `~/.config/claude/claude_desktop_config.json` (Linux) or `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS):

```json
{
  "mcpServers": {
    "kubectl-mtv": {
      "command": "kubectl-mtv-mcp",
      "args": []
    }
  }
}
```

See [MCP_SETUP.md](MCP_SETUP.md) for complete configuration instructions.

## Security Considerations

- The MCP server executes kubectl-mtv commands with your current Kubernetes permissions
- **Read operations**: Safe to use - only performs read operations
- **Write operations**: **USE WITH CAUTION** - can create, modify, and delete resources
- Ensure your MCP client is running in a secure environment
- Consider using dedicated service accounts for production environments
- By default, servers use stdio transport (no network exposure).
- When started with --sse, an HTTP server is exposed on the configured host:port.
  Prefer binding to 127.0.0.1, and restrict access via firewall when necessary.

## Available Platforms

Pre-built binaries available for:

- Linux AMD64
- Linux ARM64
- macOS Intel (AMD64)
- macOS Apple Silicon (ARM64)
- Windows AMD64

## Troubleshooting

### Server Not Found

```bash
# If binaries are in PATH, find their locations
which kubectl-mtv-mcp

# Check common installation locations
ls -l ~/.local/bin/kubectl-mtv-mcp
ls -l /usr/local/bin/kubectl-mtv-mcp

# Search your system for the binaries
find ~ -name "kubectl-mtv-mcp" -type f 2>/dev/null

# Add to PATH if needed
export PATH="$HOME/.local/bin:$PATH"

# Or use absolute paths in configuration
```

### Permission Issues

```bash
# Make binaries executable
chmod +x ~/.local/bin/kubectl-mtv-mcp
```

### Testing Server Connectivity

```bash
# Each server should start and wait for input
kubectl-mtv-mcp
# Press Ctrl+C to exit
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `make verify` to check code quality
5. Submit a pull request

## License

Apache License 2.0 - See the main kubectl-mtv project for details.
