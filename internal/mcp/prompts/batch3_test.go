package prompts_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
	"github.com/vfat/whmcs-mcp/internal/mcp/prompts"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestRegisterBatch3Prompts_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	prompts.RegisterBatch3Prompts(srv.MCPServer())

	promptMap := srv.MCPServer().ListPrompts()
	expectedPrompts := []string{
		"domain-expiry-audit",
		"fraud-investigation",
	}

	for _, name := range expectedPrompts {
		if _, exists := promptMap[name]; !exists {
			t.Errorf("expected prompt %s to be registered", name)
		}
	}
}

func TestBatch3Prompts_Execution(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	prompts.RegisterBatch3Prompts(srv.MCPServer())

	t.Run("domain-expiry-audit", func(t *testing.T) {
		promptItem := srv.MCPServer().ListPrompts()["domain-expiry-audit"]
		if promptItem == nil {
			t.Fatal("prompt domain-expiry-audit not found")
		}

		req := mcp.GetPromptRequest{}
		req.Params.Name = "domain-expiry-audit"
		req.Params.Arguments = map[string]string{
			"daysUntilExpiry": "30",
		}

		res, err := promptItem.Handler(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected prompt error: %v", err)
		}
		txtContent := res.Messages[0].Content.(mcp.TextContent)
		if !strings.Contains(txtContent.Text, "30") {
			t.Errorf("expected prompt text to contain '30', got: %s", txtContent.Text)
		}
	})

	t.Run("fraud-investigation", func(t *testing.T) {
		promptItem := srv.MCPServer().ListPrompts()["fraud-investigation"]
		if promptItem == nil {
			t.Fatal("prompt fraud-investigation not found")
		}

		req := mcp.GetPromptRequest{}
		req.Params.Name = "fraud-investigation"
		req.Params.Arguments = map[string]string{
			"orderId": "991",
		}

		res, err := promptItem.Handler(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected prompt error: %v", err)
		}
		txtContent := res.Messages[0].Content.(mcp.TextContent)
		if !strings.Contains(txtContent.Text, "#991") {
			t.Errorf("expected prompt text to contain '#991', got: %s", txtContent.Text)
		}
	})
}
