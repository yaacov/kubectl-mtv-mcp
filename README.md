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

## Recommended Installation

The recommended way to use this MCP server is through the [kubectl-mtv](https://github.com/yaacov/kubectl-mtv) CLI tool, which integrates this server and provides a convenient interface.

### Quick Setup

**For Claude Desktop:**
```bash
claude mcp add kubectl-mtv kubectl mtv mcp-server
```

**For Cursor IDE:**

Add to your Cursor MCP settings (`~/.cursor/mcp.json`):

```json
{
  "mcpServers": {
    "kubectl-mtv": {
      "command": "kubectl",
      "args": ["mtv", "mcp-server"]
    }
  }
}
```

### Why Use kubectl-mtv CLI?

- **Single Installation**: Get kubectl-mtv and the MCP server together
- **Always Up-to-Date**: MCP server updates come with kubectl-mtv releases
- **Better Integration**: Seamless access to all kubectl-mtv functionality
- **Easier Setup**: No need to install a separate binary

### Documentation

For complete setup and usage instructions, see the [kubectl-mtv MCP Server Guide](https://github.com/yaacov/kubectl-mtv/blob/main/docs/README_mcp_server.md).

## About kubectl-mtv

[kubectl-mtv](https://github.com/yaacov/kubectl-mtv) is a kubectl plugin for migrating virtual machines to KubeVirt using Forklift. The MCP server integration enables AI assistants to help with all kubectl-mtv operations.

## Prerequisites

- [kubectl-mtv](https://github.com/yaacov/kubectl-mtv) installed and in your PATH
- Access to a Kubernetes cluster with MTV deployed
- Appropriate cluster permissions

## Advanced: Standalone Usage

While using the kubectl-mtv CLI is recommended, this server can also be built and run standalone for development purposes. See the documentation in the [docs](docs/) directory for details.

## License

Apache License 2.0
