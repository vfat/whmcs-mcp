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

func TestRegisterOrderAndQuoteTools_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterOrderTools(srv.MCPServer(), client)

	expected := []string{
		"whmcs_get_orders",
		"whmcs_accept_order",
		"whmcs_cancel_order",
		"whmcs_delete_order",
		"whmcs_fraud_order",
		"whmcs_pending_order",
		"whmcs_get_quotes",
		"whmcs_create_quote",
		"whmcs_accept_quote",
		"whmcs_delete_quote",
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

func TestOrderAndQuoteTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "GetOrders":
			_ = json.NewEncoder(w).Encode(model.GetOrdersResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			})
		case "AcceptOrder":
			_ = json.NewEncoder(w).Encode(model.AcceptOrderResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "CancelOrder":
			_ = json.NewEncoder(w).Encode(model.CancelOrderResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "DeleteOrder":
			_ = json.NewEncoder(w).Encode(model.DeleteOrderResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "FraudOrder":
			_ = json.NewEncoder(w).Encode(model.FraudOrderResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "PendingOrder":
			_ = json.NewEncoder(w).Encode(model.PendingOrderResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "GetQuotes":
			_ = json.NewEncoder(w).Encode(model.GetQuotesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			})
		case "CreateQuote":
			_ = json.NewEncoder(w).Encode(model.CreateQuoteResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				QuoteID:        15,
			})
		case "AcceptQuote":
			_ = json.NewEncoder(w).Encode(model.AcceptQuoteResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				InvoiceID:      105,
			})
		case "DeleteQuote":
			_ = json.NewEncoder(w).Encode(model.DeleteQuoteResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
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

	tools.RegisterOrderTools(srv.MCPServer(), client)

	tests := []struct {
		name      string
		toolName  string
		args      map[string]interface{}
		wantError bool
	}{
		{
			name:     "get orders success",
			toolName: "whmcs_get_orders",
			args: map[string]interface{}{
				"status": "Pending",
			},
			wantError: false,
		},
		{
			name:      "accept order missing orderid",
			toolName:  "whmcs_accept_order",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "accept order success",
			toolName: "whmcs_accept_order",
			args: map[string]interface{}{
				"orderid":   77,
				"autosetup": true,
			},
			wantError: false,
		},
		{
			name:      "cancel order missing orderid",
			toolName:  "whmcs_cancel_order",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "cancel order success",
			toolName: "whmcs_cancel_order",
			args: map[string]interface{}{
				"orderid": 77,
			},
			wantError: false,
		},
		{
			name:     "delete order success",
			toolName: "whmcs_delete_order",
			args: map[string]interface{}{
				"orderid": 77,
			},
			wantError: false,
		},
		{
			name:     "fraud order success",
			toolName: "whmcs_fraud_order",
			args: map[string]interface{}{
				"orderid": 77,
			},
			wantError: false,
		},
		{
			name:     "pending order success",
			toolName: "whmcs_pending_order",
			args: map[string]interface{}{
				"orderid": 77,
			},
			wantError: false,
		},
		{
			name:     "get quotes success",
			toolName: "whmcs_get_quotes",
			args: map[string]interface{}{
				"stage": "Delivered",
			},
			wantError: false,
		},
		{
			name:      "create quote missing subject",
			toolName:  "whmcs_create_quote",
			args:      map[string]interface{}{"stage": "Draft"},
			wantError: true,
		},
		{
			name:     "create quote success",
			toolName: "whmcs_create_quote",
			args: map[string]interface{}{
				"subject": "Web Hosting Quote",
				"stage":   "Delivered",
			},
			wantError: false,
		},
		{
			name:     "accept quote success",
			toolName: "whmcs_accept_quote",
			args: map[string]interface{}{
				"quoteid": 15,
			},
			wantError: false,
		},
		{
			name:     "delete quote success",
			toolName: "whmcs_delete_quote",
			args: map[string]interface{}{
				"quoteid": 15,
			},
			wantError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			toolItem := srv.MCPServer().GetTool(tc.toolName)
			if toolItem == nil {
				t.Fatalf("tool %s not registered", tc.toolName)
			}
			req := mcp.CallToolRequest{}
			req.Params.Name = tc.toolName
			req.Params.Arguments = tc.args

			res, err := toolItem.Handler(context.Background(), req)
			if err != nil {
				t.Fatalf("unexpected handler err: %v", err)
			}
			if tc.wantError && !res.IsError {
				t.Errorf("expected error result but got success")
			}
			if !tc.wantError && res.IsError {
				t.Errorf("expected success result but got error: %v", res.Content)
			}
		})
	}
}
