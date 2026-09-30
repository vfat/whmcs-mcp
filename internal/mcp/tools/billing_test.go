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

func TestRegisterBillingTools_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterBillingTools(srv.MCPServer(), client)

	expected := []string{
		"whmcs_get_invoices",
		"whmcs_get_invoice",
		"whmcs_create_invoice",
		"whmcs_update_invoice",
		"whmcs_add_payment",
		"whmcs_apply_credit",
		"whmcs_get_transactions",
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

func TestBillingTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "GetInvoices":
			_ = json.NewEncoder(w).Encode(model.GetInvoicesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
			})
		case "GetInvoice":
			_ = json.NewEncoder(w).Encode(model.GetInvoiceResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				InvoiceID:      101,
			})
		case "CreateInvoice":
			_ = json.NewEncoder(w).Encode(model.CreateInvoiceResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				InvoiceID:      102,
			})
		case "UpdateInvoice":
			_ = json.NewEncoder(w).Encode(model.UpdateInvoiceResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				InvoiceID:      102,
			})
		case "AddInvoicePayment":
			_ = json.NewEncoder(w).Encode(model.AddInvoicePaymentResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "ApplyCredit":
			_ = json.NewEncoder(w).Encode(model.ApplyCreditResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				InvoiceID:      102,
				Amount:         "10.00",
			})
		case "GetTransactions":
			_ = json.NewEncoder(w).Encode(model.GetTransactionsResponse{
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

	tools.RegisterBillingTools(srv.MCPServer(), client)

	tests := []struct {
		name      string
		toolName  string
		args      map[string]interface{}
		wantError bool
	}{
		{
			name:     "get invoices success",
			toolName: "whmcs_get_invoices",
			args: map[string]interface{}{
				"status": "Unpaid",
			},
			wantError: false,
		},
		{
			name:      "get invoice missing invoiceid",
			toolName:  "whmcs_get_invoice",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "get invoice success",
			toolName: "whmcs_get_invoice",
			args: map[string]interface{}{
				"invoiceid": 101,
			},
			wantError: false,
		},
		{
			name:      "create invoice missing userid",
			toolName:  "whmcs_create_invoice",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "create invoice success",
			toolName: "whmcs_create_invoice",
			args: map[string]interface{}{
				"userid": 42,
				"status": "Unpaid",
			},
			wantError: false,
		},
		{
			name:      "update invoice missing invoiceid",
			toolName:  "whmcs_update_invoice",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "update invoice success",
			toolName: "whmcs_update_invoice",
			args: map[string]interface{}{
				"invoiceid": 102,
				"status":    "Paid",
			},
			wantError: false,
		},
		{
			name:      "add payment missing fields",
			toolName:  "whmcs_add_payment",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "add payment success",
			toolName: "whmcs_add_payment",
			args: map[string]interface{}{
				"invoiceid": 102,
				"transid":   "TRX-999",
				"gateway":   "paypal",
				"amount":    18.5,
			},
			wantError: false,
		},
		{
			name:      "apply credit missing invoiceid",
			toolName:  "whmcs_apply_credit",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "apply credit success",
			toolName: "whmcs_apply_credit",
			args: map[string]interface{}{
				"invoiceid": 102,
				"amount":    10.0,
			},
			wantError: false,
		},
		{
			name:     "get transactions success",
			toolName: "whmcs_get_transactions",
			args: map[string]interface{}{
				"invoiceid": 102,
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
