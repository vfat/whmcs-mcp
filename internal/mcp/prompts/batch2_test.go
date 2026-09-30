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

func TestRegisterBatch2Prompts_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	prompts.RegisterBatch2Prompts(srv.MCPServer())

	promptMap := srv.MCPServer().ListPrompts()
	expectedPrompts := []string{
		"ticket-response",
		"revenue-report",
		"bulk-invoice-reminder",
	}

	for _, name := range expectedPrompts {
		if _, exists := promptMap[name]; !exists {
			t.Errorf("expected prompt %s to be registered", name)
		}
	}
}

func TestBatch2Prompts_Execution(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	prompts.RegisterBatch2Prompts(srv.MCPServer())

	t.Run("ticket-response", func(t *testing.T) {
		promptItem := srv.MCPServer().ListPrompts()["ticket-response"]
		if promptItem == nil {
			t.Fatal("prompt ticket-response not found")
		}

		req := mcp.GetPromptRequest{}
		req.Params.Name = "ticket-response"
		req.Params.Arguments = map[string]string{
			"ticketId":  "12345",
			"issueType": "technical",
			"tone":      "friendly",
		}

		res, err := promptItem.Handler(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected prompt error: %v", err)
		}
		txtContent := res.Messages[0].Content.(mcp.TextContent)
		if !strings.Contains(txtContent.Text, "#12345") {
			t.Errorf("expected prompt text to contain '#12345', got: %s", txtContent.Text)
		}
	})

	t.Run("revenue-report", func(t *testing.T) {
		promptItem := srv.MCPServer().ListPrompts()["revenue-report"]
		if promptItem == nil {
			t.Fatal("prompt revenue-report not found")
		}

		req := mcp.GetPromptRequest{}
		req.Params.Name = "revenue-report"
		req.Params.Arguments = map[string]string{
			"period": "month",
		}

		res, err := promptItem.Handler(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected prompt error: %v", err)
		}
		txtContent := res.Messages[0].Content.(mcp.TextContent)
		if !strings.Contains(txtContent.Text, "month") {
			t.Errorf("expected prompt text to contain 'month', got: %s", txtContent.Text)
		}
	})

	t.Run("bulk-invoice-reminder", func(t *testing.T) {
		promptItem := srv.MCPServer().ListPrompts()["bulk-invoice-reminder"]
		if promptItem == nil {
			t.Fatal("prompt bulk-invoice-reminder not found")
		}

		req := mcp.GetPromptRequest{}
		req.Params.Name = "bulk-invoice-reminder"
		req.Params.Arguments = map[string]string{
			"daysOverdue": "14",
		}

		res, err := promptItem.Handler(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected prompt error: %v", err)
		}
		txtContent := res.Messages[0].Content.(mcp.TextContent)
		if !strings.Contains(txtContent.Text, "14") {
			t.Errorf("expected prompt text to contain '14', got: %s", txtContent.Text)
		}
	})
}
