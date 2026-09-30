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

func TestRegisterBatch1Prompts_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	prompts.RegisterBatch1Prompts(srv.MCPServer())

	promptMap := srv.MCPServer().ListPrompts()
	expectedPrompts := []string{
		"client-onboarding",
		"client-health-check",
	}

	for _, name := range expectedPrompts {
		if _, exists := promptMap[name]; !exists {
			t.Errorf("expected prompt %s to be registered", name)
		}
	}
}

func TestBatch1Prompts_Execution(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	prompts.RegisterBatch1Prompts(srv.MCPServer())

	t.Run("client-onboarding", func(t *testing.T) {
		promptItem := srv.MCPServer().ListPrompts()["client-onboarding"]
		if promptItem == nil {
			t.Fatal("prompt client-onboarding not found")
		}

		req := mcp.GetPromptRequest{}
		req.Params.Name = "client-onboarding"
		req.Params.Arguments = map[string]string{
			"clientName": "Alice Smith",
			"email":      "alice@example.com",
			"products":   "cPanel Basic",
		}

		res, err := promptItem.Handler(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected prompt error: %v", err)
		}
		if len(res.Messages) == 0 {
			t.Fatal("expected non-empty prompt messages")
		}

		txtContent, ok := res.Messages[0].Content.(mcp.TextContent)
		if !ok {
			t.Fatalf("expected TextContent, got %T", res.Messages[0].Content)
		}
		if !strings.Contains(txtContent.Text, "Alice Smith") {
			t.Errorf("expected prompt text to contain 'Alice Smith', got: %s", txtContent.Text)
		}
		if !strings.Contains(txtContent.Text, "cPanel Basic") {
			t.Errorf("expected prompt text to contain 'cPanel Basic', got: %s", txtContent.Text)
		}
	})

	t.Run("client-health-check", func(t *testing.T) {
		promptItem := srv.MCPServer().ListPrompts()["client-health-check"]
		if promptItem == nil {
			t.Fatal("prompt client-health-check not found")
		}

		req := mcp.GetPromptRequest{}
		req.Params.Name = "client-health-check"
		req.Params.Arguments = map[string]string{
			"clientId": "456",
		}

		res, err := promptItem.Handler(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected prompt error: %v", err)
		}
		if len(res.Messages) == 0 {
			t.Fatal("expected non-empty prompt messages")
		}

		txtContent, ok := res.Messages[0].Content.(mcp.TextContent)
		if !ok {
			t.Fatalf("expected TextContent, got %T", res.Messages[0].Content)
		}
		if !strings.Contains(txtContent.Text, "#456") {
			t.Errorf("expected prompt text to contain '#456', got: %s", txtContent.Text)
		}
	})
}
