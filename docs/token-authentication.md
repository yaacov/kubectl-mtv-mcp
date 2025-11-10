# Token Authentication

The kubectl-mtv MCP server supports token-based authentication for kubectl and kubectl-mtv commands. This allows clients to provide Kubernetes authentication tokens with each MCP request.

## How It Works

In **SSE mode**, the server extracts bearer tokens from HTTP `Authorization` headers and passes them to kubectl/kubectl-mtv commands using the `--token` flag. If no token is provided, commands fall back to using the default kubeconfig.

**Note:** Token authentication is **only available in SSE mode**. Stdio mode uses the default kubeconfig authentication.

## Usage

### Start the Server in SSE Mode

```bash
kubectl-mtv-mcp --sse --host 127.0.0.1 --port 8080
```

### Send Requests with Bearer Token

Include the `Authorization` header with your Kubernetes token:

```http
Authorization: Bearer <your-kubernetes-token>
```

### Example: Get a Token

```bash
# Create a service account
kubectl create serviceaccount mcp-client -n default

# Generate a token (valid for 1 hour)
TOKEN=$(kubectl create token mcp-client -n default --duration=1h)

# Or use your current token (OpenShift)
TOKEN=$(oc whoami -t)
```

### Example: Connect with curl

```bash
curl -N -H "Authorization: Bearer $TOKEN" \
     -H "Accept: text/event-stream" \
     http://127.0.0.1:8080/sse
```

### Example: Python Client

```python
import requests

token = "your-kubernetes-token"
headers = {
    'Authorization': f'Bearer {token}',
    'Accept': 'text/event-stream'
}

response = requests.get('http://127.0.0.1:8080/sse', 
                        headers=headers, 
                        stream=True)
```

## Security Notes

- Tokens are stored in request context and never logged in full
- Tokens are sanitized in command output (displayed as `****`)
- Tokens are passed via the `--token` flag to kubectl commands
- Ensure your token has appropriate RBAC permissions for the operations needed
- Use time-limited tokens and rotate them regularly

## Stdio Mode

Token authentication is **not supported in stdio mode**. When running without `--sse`, the server uses the default kubeconfig:

```bash
# Uses default kubeconfig (~/.kube/config)
kubectl-mtv-mcp
```

