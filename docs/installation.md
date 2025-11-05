# Installation Guide

Complete installation instructions for kubectl-mtv MCP server across different platforms and MCP clients.

## Prerequisites

Before installing the MCP server, ensure you have:

1. **kubectl-mtv** installed and available in your PATH
2. Access to a Kubernetes cluster with MTV deployed
3. Appropriate cluster permissions to manage MTV resources

## Installing the Binary

Download the appropriate binary for your platform from the releases page:

- Linux AMD64: `kubectl-mtv-mcp-linux-amd64.tar.gz`
- Linux ARM64: `kubectl-mtv-mcp-linux-arm64.tar.gz`
- macOS Intel: `kubectl-mtv-mcp-darwin-amd64.tar.gz`
- macOS Apple Silicon: `kubectl-mtv-mcp-darwin-arm64.tar.gz`
- Windows: `kubectl-mtv-mcp-windows-amd64.zip`

## MCP Client Configuration

### Claude Code (Desktop App)

The simplest way to install with Claude Desktop:

```bash
claude mcp add kubectl-mtv kubectl-mtv-mcp
```

This automatically configures Claude Desktop to use the server.

**Manual Configuration:**

Edit the Claude Desktop config file:
- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Linux**: `~/.config/claude/claude_desktop_config.json`

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

After editing, restart Claude Desktop.

### Cursor IDE

Add the server to your Cursor MCP configuration:

1. Open Cursor settings
2. Navigate to MCP settings
3. Edit `~/.cursor/mcp.json`:

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

**Using Absolute Path:**

If the binary is not in PATH, use the full path:

```json
{
  "mcpServers": {
    "kubectl-mtv": {
      "command": "/Users/yourname/.local/bin/kubectl-mtv-mcp",
      "args": []
    }
  }
}
```

After editing, reload Cursor's MCP configuration.

### Other MCP Clients

For other MCP-compatible clients, configure them to run:

```bash
kubectl-mtv-mcp
```

The server uses stdio transport by default, which is the standard for MCP servers.

## Advanced Configuration

### SSE Mode (HTTP Server)

To run the server in SSE mode for remote access:

```bash
kubectl-mtv-mcp --sse --host 127.0.0.1 --port 8080
```

Configure your MCP client to connect to:
```
http://127.0.0.1:8080/sse
```

**Security Warning:** When using SSE mode, restrict access to localhost or use appropriate firewall rules.

### Running as a Service

For production environments, consider running the server as a systemd service (Linux) or launchd service (macOS).

**Example systemd service** (`/etc/systemd/system/kubectl-mtv-mcp.service`):

```ini
[Unit]
Description=kubectl-mtv MCP Server
After=network.target

[Service]
Type=simple
User=youruser
ExecStart=/usr/local/bin/kubectl-mtv-mcp --sse --host 127.0.0.1 --port 8080
Restart=on-failure
Environment="KUBECONFIG=/home/youruser/.kube/config"

[Install]
WantedBy=multi-user.target
```
