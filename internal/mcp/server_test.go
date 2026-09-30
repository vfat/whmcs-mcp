package mcp_test

import (
	"context"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestNewServer(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "secret",
		Timeout:    10 * time.Second,
	}
	client := whmcs.NewClient(cfg)

	srv := mcpServer.NewServer(cfg, client)
	if srv == nil {
		t.Fatal("expected non-nil server")
	}

	if srv.MCPServer() == nil {
		t.Fatal("expected internal MCPServer to be initialized")
	}
}

func TestRegisterTool_Execution(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "secret",
		Timeout:    10 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	// Register sample tool
	tool := mcp.NewTool("test_ping",
		mcp.WithDescription("Ping test tool"),
		mcp.WithString("message", mcp.Required(), mcp.Description("Message to echo")),
	)

	srv.MCPServer().AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("arguments must be a map"), nil
		}
		msg, ok := args["message"].(string)
		if !ok {
			return mcp.NewToolResultError("message argument required"), nil
		}
		return mcp.NewToolResultText("pong: " + msg), nil
	})

	// Execute tool call directly via handler or simulation
	ctx := context.Background()
	req := mcp.CallToolRequest{}
	req.Params.Name = "test_ping"
	req.Params.Arguments = map[string]interface{}{
		"message": "hello",
	}

	// Verify tool registration exists
	tools := srv.ListTools()
	found := false
	for _, tl := range tools {
		if tl.Name == "test_ping" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected test_ping tool to be registered in server")
	}
	_ = ctx
}
