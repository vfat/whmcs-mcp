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

func TestProvisioningTools_Registration(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterProvisioningTools(srv.MCPServer(), client)

	expected := []string{
		"whmcs_get_servers",
		"whmcs_module_create",
		"whmcs_module_suspend",
		"whmcs_module_unsuspend",
		"whmcs_module_terminate",
		"whmcs_module_change_password",
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

func TestProvisioningTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "GetServers":
			_ = json.NewEncoder(w).Encode(model.GetServersResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				Servers: model.ServerListContainer{
					Server: []model.ServerItem{{ID: 1, Name: "srv1", Hostname: "srv1.example.com"}},
				},
			})
		case "ModuleCreate":
			_ = json.NewEncoder(w).Encode(model.ModuleResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				Message:        "Service provisioned",
			})
		case "ModuleSuspend":
			_ = json.NewEncoder(w).Encode(model.ModuleResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				Message:        "Service suspended",
			})
		case "ModuleUnsuspend":
			_ = json.NewEncoder(w).Encode(model.ModuleResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				Message:        "Service unsuspended",
			})
		case "ModuleTerminate":
			_ = json.NewEncoder(w).Encode(model.ModuleResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				Message:        "Service terminated",
			})
		case "ModuleChangePassword":
			_ = json.NewEncoder(w).Encode(model.ModuleResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				Message:        "Password changed",
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

	tools.RegisterProvisioningTools(srv.MCPServer(), client)

	tests := []struct {
		name      string
		toolName  string
		args      map[string]interface{}
		wantError bool
	}{
		{
			name:     "get servers success",
			toolName: "whmcs_get_servers",
			args: map[string]interface{}{
				"fetchStatus": true,
			},
			wantError: false,
		},
		{
			name:      "module create missing accountid",
			toolName:  "whmcs_module_create",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "module create success",
			toolName: "whmcs_module_create",
			args: map[string]interface{}{
				"accountid": 101,
			},
			wantError: false,
		},
		{
			name:     "module suspend success",
			toolName: "whmcs_module_suspend",
			args: map[string]interface{}{
				"accountid":     101,
				"suspendreason": "Non-payment",
			},
			wantError: false,
		},
		{
			name:     "module unsuspend success",
			toolName: "whmcs_module_unsuspend",
			args: map[string]interface{}{
				"accountid": 101,
			},
			wantError: false,
		},
		{
			name:     "module terminate success",
			toolName: "whmcs_module_terminate",
			args: map[string]interface{}{
				"accountid": 101,
			},
			wantError: false,
		},
		{
			name:     "module change password success",
			toolName: "whmcs_module_change_password",
			args: map[string]interface{}{
				"accountid":       101,
				"servicepassword": "newpass",
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
