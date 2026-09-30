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

func TestRegisterDomainTools_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterDomainTools(srv.MCPServer(), client)

	expected := []string{
		"whmcs_register_domain",
		"whmcs_transfer_domain",
		"whmcs_renew_domain",
		"whmcs_get_domain_whois",
		"whmcs_get_domain_nameservers",
		"whmcs_update_domain_nameservers",
		"whmcs_get_domain_lock_status",
		"whmcs_update_domain_lock_status",
		"whmcs_get_tld_pricing",
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

func TestDomainTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "DomainRegister":
			_ = json.NewEncoder(w).Encode(model.RegisterDomainResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "DomainTransfer":
			_ = json.NewEncoder(w).Encode(model.TransferDomainResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "DomainRenew":
			_ = json.NewEncoder(w).Encode(model.RenewDomainResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "DomainWhois":
			_ = json.NewEncoder(w).Encode(model.GetDomainWhoisResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				Whois:          "Domain Name: example.com ...",
			})
		case "DomainGetNameservers":
			_ = json.NewEncoder(w).Encode(model.GetDomainNameserversResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				NS1:            "ns1.example.com",
			})
		case "DomainUpdateNameservers":
			_ = json.NewEncoder(w).Encode(model.UpdateDomainNameserversResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "DomainGetLockingStatus":
			_ = json.NewEncoder(w).Encode(model.GetDomainLockStatusResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				LockStatus:     "locked",
			})
		case "DomainUpdateLockingStatus":
			_ = json.NewEncoder(w).Encode(model.UpdateDomainLockStatusResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
			})
		case "GetTLDPricing":
			_ = json.NewEncoder(w).Encode(model.GetTLDPricingResponse{
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

	tools.RegisterDomainTools(srv.MCPServer(), client)

	tests := []struct {
		name      string
		toolName  string
		args      map[string]interface{}
		wantError bool
	}{
		{
			name:     "register domain success",
			toolName: "whmcs_register_domain",
			args: map[string]interface{}{
				"domain": "example.com",
			},
			wantError: false,
		},
		{
			name:      "transfer domain missing domainid",
			toolName:  "whmcs_transfer_domain",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "transfer domain success",
			toolName: "whmcs_transfer_domain",
			args: map[string]interface{}{
				"domainid": 88,
			},
			wantError: false,
		},
		{
			name:     "renew domain success",
			toolName: "whmcs_renew_domain",
			args: map[string]interface{}{
				"domainid": 88,
			},
			wantError: false,
		},
		{
			name:     "get domain whois success",
			toolName: "whmcs_get_domain_whois",
			args: map[string]interface{}{
				"domainid": 88,
			},
			wantError: false,
		},
		{
			name:     "get domain nameservers success",
			toolName: "whmcs_get_domain_nameservers",
			args: map[string]interface{}{
				"domainid": 88,
			},
			wantError: false,
		},
		{
			name:     "update domain nameservers success",
			toolName: "whmcs_update_domain_nameservers",
			args: map[string]interface{}{
				"domainid": 88,
				"ns1":      "ns1.example.com",
			},
			wantError: false,
		},
		{
			name:     "get domain lock status success",
			toolName: "whmcs_get_domain_lock_status",
			args: map[string]interface{}{
				"domainid": 88,
			},
			wantError: false,
		},
		{
			name:     "update domain lock status success",
			toolName: "whmcs_update_domain_lock_status",
			args: map[string]interface{}{
				"domainid":   88,
				"lockstatus": true,
			},
			wantError: false,
		},
		{
			name:      "get tld pricing success",
			toolName:  "whmcs_get_tld_pricing",
			args:      map[string]interface{}{},
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
