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

func TestRegisterSystemTools_AllToolsPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterSystemTools(srv.MCPServer(), client)

	registered := srv.ListTools()
	expectedTools := []string{
		"whmcs_get_stats",
		"whmcs_get_activity_log",
		"whmcs_get_admin_users",
		"whmcs_get_todo_items",
		"whmcs_get_todo_item_statuses",
		"whmcs_update_todo_item",
		"whmcs_get_currencies",
		"whmcs_get_payment_methods",
		"whmcs_send_email",
	}

	foundMap := make(map[string]bool)
	for _, tl := range registered {
		foundMap[tl.Name] = true
	}

	for _, expected := range expectedTools {
		if !foundMap[expected] {
			t.Errorf("expected tool %s to be registered", expected)
		}
	}
}

func TestSystemTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "GetStats":
			_ = json.NewEncoder(w).Encode(model.GetStatsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				IncomeToday:    "100.00",
			})
		case "GetActivityLog":
			resp := model.GetActivityLogResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			}
			_ = json.NewEncoder(w).Encode(resp)
		case "GetAdminUsers":
			resp := model.GetAdminUsersResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			}
			_ = json.NewEncoder(w).Encode(resp)
		case "GetToDoItems":
			resp := model.GetToDoItemsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			}
			_ = json.NewEncoder(w).Encode(resp)
		case "GetToDoItemStatuses":
			resp := model.GetToDoStatusesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			}
			_ = json.NewEncoder(w).Encode(resp)
		case "UpdateToDoItem":
			resp := model.UpdateToDoItemResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TodoID:         42,
			}
			_ = json.NewEncoder(w).Encode(resp)
		case "GetCurrencies":
			resp := model.GetCurrenciesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			}
			_ = json.NewEncoder(w).Encode(resp)
		case "GetPaymentMethods":
			resp := model.GetPaymentMethodsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			}
			_ = json.NewEncoder(w).Encode(resp)
		case "SendEmail":
			resp := model.SendEmailResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			}
			_ = json.NewEncoder(w).Encode(resp)
		default:
			_ = json.NewEncoder(w).Encode(map[string]string{"result": "error", "message": "Unknown action"})
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
	tools.RegisterSystemTools(srv.MCPServer(), client)

	// 1. GetStats
	t.Run("whmcs_get_stats", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_stats")
		if toolItem == nil {
			t.Fatal("tool not found")
		}
		res, err := toolItem.Handler(context.Background(), mcp.CallToolRequest{})
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 2. GetActivityLog
	t.Run("whmcs_get_activity_log", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_activity_log")
		callReq := mcp.CallToolRequest{}
		callReq.Params.Arguments = map[string]interface{}{
			"limitnum": 10,
		}
		res, err := toolItem.Handler(context.Background(), callReq)
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 3. GetAdminUsers
	t.Run("whmcs_get_admin_users", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_admin_users")
		res, err := toolItem.Handler(context.Background(), mcp.CallToolRequest{})
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 4. GetToDoItems
	t.Run("whmcs_get_todo_items", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_todo_items")
		callReq := mcp.CallToolRequest{}
		callReq.Params.Arguments = map[string]interface{}{
			"status": "Incomplete",
		}
		res, err := toolItem.Handler(context.Background(), callReq)
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 5. GetToDoItemStatuses
	t.Run("whmcs_get_todo_item_statuses", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_todo_item_statuses")
		res, err := toolItem.Handler(context.Background(), mcp.CallToolRequest{})
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 6. UpdateToDoItem - validation & success
	t.Run("whmcs_update_todo_item", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_update_todo_item")

		// Missing itemid -> error
		resMissing, _ := toolItem.Handler(context.Background(), mcp.CallToolRequest{})
		if !resMissing.IsError {
			t.Errorf("expected error when itemid is missing")
		}

		// Success call
		callReq := mcp.CallToolRequest{}
		callReq.Params.Arguments = map[string]interface{}{
			"itemid": 42,
			"status": "Completed",
		}
		res, err := toolItem.Handler(context.Background(), callReq)
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 7. GetCurrencies
	t.Run("whmcs_get_currencies", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_currencies")
		res, err := toolItem.Handler(context.Background(), mcp.CallToolRequest{})
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 8. GetPaymentMethods
	t.Run("whmcs_get_payment_methods", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_get_payment_methods")
		res, err := toolItem.Handler(context.Background(), mcp.CallToolRequest{})
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})

	// 9. SendEmail
	t.Run("whmcs_send_email", func(t *testing.T) {
		toolItem := srv.MCPServer().GetTool("whmcs_send_email")
		callReq := mcp.CallToolRequest{}
		callReq.Params.Arguments = map[string]interface{}{
			"messagename": "Client Signup Email",
			"id":          123,
		}
		res, err := toolItem.Handler(context.Background(), callReq)
		if err != nil || res.IsError {
			t.Fatalf("unexpected error: %v, content: %v", err, res.Content)
		}
	})
}
