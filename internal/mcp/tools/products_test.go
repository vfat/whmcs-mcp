package tools_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
	"github.com/vfat/whmcs-mcp/internal/mcp/tools"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestRegisterProductTools_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterProductTools(srv.MCPServer(), client)

	expected := []string{
		"whmcs_get_products",
		"whmcs_get_product_groups",
	}

	toolMap := make(map[string]bool)
	for _, tl := range srv.ListTools() {
		toolMap[tl.Name] = true
	}

	for _, name := range expected {
		if !toolMap[name] {
			t.Errorf("expected tool %s to be registered", name)
		}
	}
}

func TestProductTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "GetProducts":
			_ = json.NewEncoder(w).Encode(model.GetProductsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			})
		case "GetProductGroups":
			_ = json.NewEncoder(w).Encode(model.GetProductGroupsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			})
		default:
			_ = json.NewEncoder(w).Encode(map[string]string{"result": "error"})
		}
	}))
	defer mockServer.Close()

	cfg := &config.Config{
		URL:        mockServer.URL,
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterProductTools(srv.MCPServer(), client)

	t.Run("whmcs_get_products", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_products")
		req := mcp.CallToolRequest{}
		req.Params.Arguments = map[string]interface{}{
			"pid": 10,
		}
		res, err := toolItem.Handler(context.Background(), req)
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	t.Run("whmcs_get_product_groups", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_product_groups")
		res, err := toolItem.Handler(context.Background(), mcp.CallToolRequest{})
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})
}
