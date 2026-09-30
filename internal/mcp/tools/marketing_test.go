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

func TestMarketingTools_Registration(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tools.RegisterMarketingTools(srv.MCPServer(), client)

	expected := []string{
		"whmcs_get_affiliates",
		"whmcs_activate_affiliate",
		"whmcs_get_promotions",
		"whmcs_log_activity",
		"whmcs_get_email_templates",
		"whmcs_delete_client",
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

func TestMarketingTools_Execution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		action := r.URL.Query().Get("action")
		if action == "" {
			action = r.FormValue("action")
		}

		switch action {
		case "GetAffiliates":
			_ = json.NewEncoder(w).Encode(model.GetAffiliatesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
				Affiliates: model.AffiliateListContainer{
					Affiliate: []model.AffiliateItem{{ID: 1, UserID: 10, Balance: "50.00"}},
				},
			})
		case "AffiliateActivate":
			_ = json.NewEncoder(w).Encode(model.AffiliateActivateResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				AffiliateID:    1,
			})
		case "GetPromotions":
			_ = json.NewEncoder(w).Encode(model.GetPromotionsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
				Promotions: model.PromotionListContainer{
					Promotion: []model.PromotionItem{{ID: 1, Code: "PROMO10", Type: "percentage"}},
				},
			})
		case "LogActivity":
			_ = json.NewEncoder(w).Encode(model.CommonResponse{Result: "success"})
		case "GetEmailTemplates":
			_ = json.NewEncoder(w).Encode(model.GetEmailTemplatesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
				EmailTemplates: model.EmailTemplateContainer{
					EmailTemplate: []model.EmailTemplateItem{{ID: 1, Name: "Welcome Email", Subject: "Welcome!"}},
				},
			})
		case "DeleteClient":
			_ = json.NewEncoder(w).Encode(model.CommonResponse{Result: "success"})
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

	tools.RegisterMarketingTools(srv.MCPServer(), client)

	tests := []struct {
		name      string
		toolName  string
		args      map[string]interface{}
		wantError bool
	}{
		{
			name:     "get affiliates success",
			toolName: "whmcs_get_affiliates",
			args: map[string]interface{}{
				"limitstart": 0,
				"limitnum":   10,
			},
			wantError: false,
		},
		{
			name:      "activate affiliate missing userid",
			toolName:  "whmcs_activate_affiliate",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "activate affiliate success",
			toolName: "whmcs_activate_affiliate",
			args: map[string]interface{}{
				"userid": 10,
			},
			wantError: false,
		},
		{
			name:     "get promotions success",
			toolName: "whmcs_get_promotions",
			args: map[string]interface{}{
				"code": "PROMO10",
			},
			wantError: false,
		},
		{
			name:      "log activity missing description",
			toolName:  "whmcs_log_activity",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "log activity success",
			toolName: "whmcs_log_activity",
			args: map[string]interface{}{
				"description": "User logged in",
				"userid":      10,
			},
			wantError: false,
		},
		{
			name:     "get email templates success",
			toolName: "whmcs_get_email_templates",
			args: map[string]interface{}{
				"type": "general",
			},
			wantError: false,
		},
		{
			name:      "delete client missing clientid",
			toolName:  "whmcs_delete_client",
			args:      map[string]interface{}{},
			wantError: true,
		},
		{
			name:     "delete client success",
			toolName: "whmcs_delete_client",
			args: map[string]interface{}{
				"clientid": 10,
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
