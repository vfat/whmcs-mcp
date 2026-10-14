package tools_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
	"github.com/vfat/whmcs-mcp/internal/mcp/tools"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestRegisterClientTools_AllToolsPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterClientTools(srv.MCPServer(), client)

	registered := srv.ListTools()
	expectedTools := []string{
		"whmcs_get_clients",
		"whmcs_get_client_details",
		"whmcs_add_client",
		"whmcs_update_client",
		"whmcs_close_client",
		"whmcs_get_client_products",
		"whmcs_get_client_domains",
		"whmcs_get_client_invoices",
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

func TestClientTools_GetClientsCall(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := model.GetClientsResponse{
			TotalResults: 1,
		}
		resp.Result = "success"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
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

	tools.RegisterClientTools(srv.MCPServer(), client)

	toolItem := srv.MCPServer().GetTool("whmcs_get_clients")
	if toolItem == nil {
		t.Fatal("tool whmcs_get_clients not found on server")
	}

	callReq := mcp.CallToolRequest{}
	callReq.Params.Name = "whmcs_get_clients"
	callReq.Params.Arguments = map[string]interface{}{
		"limitnum": 10,
	}

	res, err := toolItem.Handler(context.Background(), callReq)
	if err != nil {
		t.Fatalf("unexpected handler error: %v", err)
	}

	if res.IsError {
		t.Errorf("expected success tool result, got error: %v", res.Content)
	}
}

func TestClientTools_GetClientDetails(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := model.GetClientDetailsResponse{
			ClientDetails: model.ClientDetails{
				ClientID: 123,
				Email:    "test@example.com",
			},
		}
		resp.Result = "success"
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
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
	tools.RegisterClientTools(srv.MCPServer(), client)

	toolItem := srv.MCPServer().GetTool("whmcs_get_client_details")
	if toolItem == nil {
		t.Fatal("tool whmcs_get_client_details not found")
	}

	// Missing clientid and email -> error
	callReqMissing := mcp.CallToolRequest{}
	callReqMissing.Params.Name = "whmcs_get_client_details"
	callReqMissing.Params.Arguments = map[string]interface{}{}
	resMissing, _ := toolItem.Handler(context.Background(), callReqMissing)
	if !resMissing.IsError {
		t.Errorf("expected error when both clientid and email missing")
	}

	// Success call with clientid
	callReq := mcp.CallToolRequest{}
	callReq.Params.Name = "whmcs_get_client_details"
	callReq.Params.Arguments = map[string]interface{}{
		"clientid": 123,
	}
	res, err := toolItem.Handler(context.Background(), callReq)
	if err != nil || res.IsError {
		t.Fatalf("expected success, got error: %v, content: %v", err, res.Content)
	}
}

func TestClientTools_CRUDAndSubResources(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"result":   "success",
			"clientid": 999,
		})
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
	tools.RegisterClientTools(srv.MCPServer(), client)

	tests := []struct {
		name      string
		toolName  string
		args      map[string]interface{}
		wantError bool
	}{
		{
			name:     "add client success",
			toolName: "whmcs_add_client",
			args: map[string]interface{}{
				"firstname": "John",
				"lastname":  "Doe",
				"email":     "john@example.com",
				"address1":  "123 Main St",
				"city":      "Anytown",
				"state":     "CA",
				"postcode":  "90210",
				"country":   "US",
			},
			wantError: false,
		},
		{
			name:     "update client success",
			toolName: "whmcs_update_client",
			args: map[string]interface{}{
				"clientid":  999,
				"firstname": "Johnny",
			},
			wantError: false,
		},
		{
			name:     "close client success",
			toolName: "whmcs_close_client",
			args: map[string]interface{}{
				"clientid": 999,
			},
			wantError: false,
		},
		{
			name:     "get client products success",
			toolName: "whmcs_get_client_products",
			args: map[string]interface{}{
				"clientid": 999,
			},
			wantError: false,
		},
		{
			name:     "get client domains success",
			toolName: "whmcs_get_client_domains",
			args: map[string]interface{}{
				"clientid": 999,
			},
			wantError: false,
		},
		{
			name:      "get client invoices missing clientid",
			toolName:  "whmcs_get_client_invoices",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "get client invoices success",
			toolName: "whmcs_get_client_invoices",
			args: map[string]interface{}{
				"clientid": 999,
				"status":   "Paid",
				"limitnum": 10,
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
			callReq := mcp.CallToolRequest{}
			callReq.Params.Name = tc.toolName
			callReq.Params.Arguments = tc.args

			res, err := toolItem.Handler(context.Background(), callReq)
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

func TestClientTools_GetClientProducts_VMID(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		serviceID := r.FormValue("serviceid")
		if serviceID == "77273" {
			_, _ = w.Write([]byte(`{
				"result": "success",
				"totalresults": 1,
				"startnumber": 0,
				"numreturned": 1,
				"products": {
					"product": [
						{
							"id": 77273,
							"clientid": 109524,
							"name": "Classic Standard",
							"status": "Active",
							"serverip": "157.66.55.179",
							"customfields": {
								"customfield": [
									{"id": 155, "name": "vmid", "value": "101"},
									{"id": 156, "name": "node", "value": ""}
								]
							}
						}
					]
				}
			}`))
			return
		}

		// Service without vmid
		_, _ = w.Write([]byte(`{
			"result": "success",
			"totalresults": 1,
			"startnumber": 0,
			"numreturned": 1,
			"products": {
				"product": [
					{
						"id": 77274,
						"clientid": 109524,
						"name": "Shared Hosting",
						"status": "Active",
						"serverip": "157.66.55.180",
						"customfields": ""
					}
				]
			}
		}`))
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
	tools.RegisterClientTools(srv.MCPServer(), client)

	toolItem := srv.MCPServer().GetTool("whmcs_get_client_products")
	if toolItem == nil {
		t.Fatal("tool whmcs_get_client_products not found")
	}

	// 1. Test dengan serviceid 77273 -> mengharapkan vmid: "101"
	callReq1 := mcp.CallToolRequest{}
	callReq1.Params.Name = "whmcs_get_client_products"
	callReq1.Params.Arguments = map[string]interface{}{
		"serviceid": 77273,
	}

	res1, err := toolItem.Handler(context.Background(), callReq1)
	if err != nil {
		t.Fatalf("unexpected handler error: %v", err)
	}
	if res1.IsError {
		t.Fatalf("expected success, got error: %v", res1.Content)
	}

	content1, ok := res1.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", res1.Content[0])
	}

	if !strings.Contains(content1.Text, `"vmid": "101"`) {
		t.Errorf("expected output to contain \"vmid\": \"101\", got: %s", content1.Text)
	}
	if strings.Contains(content1.Text, `"customfields"`) {
		t.Errorf("expected output to NOT contain nested customfields, got: %s", content1.Text)
	}

	// 2. Test dengan service tanpa vmid -> mengharapkan vmid: "000"
	callReq2 := mcp.CallToolRequest{}
	callReq2.Params.Name = "whmcs_get_client_products"
	callReq2.Params.Arguments = map[string]interface{}{
		"serviceid": 77274,
	}

	res2, err := toolItem.Handler(context.Background(), callReq2)
	if err != nil {
		t.Fatalf("unexpected handler error: %v", err)
	}
	if res2.IsError {
		t.Fatalf("expected success, got error: %v", res2.Content)
	}

	content2, ok := res2.Content[0].(mcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", res2.Content[0])
	}

	if !strings.Contains(content2.Text, `"vmid": "000"`) {
		t.Errorf("expected output to contain \"vmid\": \"000\", got: %s", content2.Text)
	}
}

