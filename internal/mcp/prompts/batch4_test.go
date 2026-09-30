package prompts_test

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/mcp/prompts"
)

func TestRegisterBatch4Prompts_AllPresent(t *testing.T) {
	s := server.NewMCPServer("test", "1.0.0")
	prompts.RegisterBatch4Prompts(s)

	reg := s.ListPrompts()
	if _, exists := reg["new-product-setup"]; !exists {
		t.Errorf("expected prompt new-product-setup to be registered")
	}
}

func TestBatch4Prompts_Execution(t *testing.T) {
	s := server.NewMCPServer("test", "1.0.0")
	prompts.RegisterBatch4Prompts(s)

	promptItem := s.ListPrompts()["new-product-setup"]
	if promptItem == nil {
		t.Fatal("prompt new-product-setup not found")
	}

	req := mcp.GetPromptRequest{}
	req.Params.Name = "new-product-setup"
	req.Params.Arguments = map[string]string{
		"productName":  "Starter Hosting",
		"productType":  "hostingaccount",
		"monthlyPrice": "4.99",
		"serverType":   "cpanel",
	}

	res, err := promptItem.Handler(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error rendering prompt: %v", err)
	}
	if len(res.Messages) == 0 {
		t.Fatalf("expected at least 1 message")
	}

	msg := res.Messages[0]
	textMsg, ok := msg.Content.(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent")
	}

	if !strings.Contains(textMsg.Text, "Starter Hosting") {
		t.Errorf("expected productName in output prompt")
	}
	if !strings.Contains(textMsg.Text, "hostingaccount") {
		t.Errorf("expected productType in output prompt")
	}
	if !strings.Contains(textMsg.Text, "4.99") {
		t.Errorf("expected monthlyPrice in output prompt")
	}
	if !strings.Contains(textMsg.Text, "cpanel") {
		t.Errorf("expected serverType in output prompt")
	}
}
