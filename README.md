# kubectl-mtv MCP Server

A Model Context Protocol (MCP) server that gives AI assistants the ability to manage Migration Toolkit for Virtualization (MTV) operations.

## What is it?

An MCP server that wraps `kubectl-mtv` commands, allowing AI assistants like Claude and Cursor to help you with VM migration tasks.

## What does it do?

- **Query** MTV resources (providers, plans, VMs, mappings)
- **Monitor** migration status and logs
- **Create and manage** migration plans
- **Configure** providers, networks, and storage mappings
- **Control** migration lifecycle (start, cancel, cutover)

## Quick Install

### Claude Code (Desktop)

```bash
claude mcp add kubectl-mtv kubectl-mtv-mcp
```

### Cursor

Add to your Cursor MCP settings (`~/.cursor/mcp.json`):

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

## Prerequisites

- `kubectl-mtv` installed and in your PATH
- Access to a Kubernetes cluster with MTV deployed
- Appropriate cluster permissions

## Documentation

- [Installation Guide](docs/installation.md) - Detailed setup for all platforms
- [Building from Source](docs/building.md) - Development and build instructions

## Security Note

This server can modify cluster resources. The AI assistant executes commands with your current Kubernetes permissions. Use with caution in production environments.

## License

Apache License 2.0
