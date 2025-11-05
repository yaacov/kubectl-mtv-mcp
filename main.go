package main

import (
	"log"

	cmd "github.com/yaacov/kubectl-mtv-mcp/cmd/kubectl-mtv-mcp"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
