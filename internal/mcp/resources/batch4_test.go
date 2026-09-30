package resources_test

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
	"github.com/vfat/whmcs-mcp/internal/mcp/resources"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestRegisterBatch4Resources_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	resources.RegisterBatch4Resources(srv.MCPServer(), client)

	reg := srv.MCPServer().ListResources()
	expectedURIs := []string{
		"whmcs://servers",
		"whmcs://promotions",
	}

	for _, uri := range expectedURIs {
		if _, exists := reg[uri]; !exists {
			t.Errorf("expected resource URI %s to be registered", uri)
		}
	}
}

func TestBatch4Resources_Read(t *testing.T) {
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
					Server: []model.ServerItem{{ID: 1, Name: "Server 1"}},
				},
			})
		case "GetPromotions":
			_ = json.NewEncoder(w).Encode(model.GetPromotionsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   1,
				Promotions: model.PromotionListContainer{
					Promotion: []model.PromotionItem{{ID: 1, Code: "PROMO2026"}},
				},
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

	resources.RegisterBatch4Resources(srv.MCPServer(), client)

	tests := []string{
		"whmcs://servers",
		"whmcs://promotions",
	}

	for _, uri := range tests {
		t.Run(uri, func(t *testing.T) {
			serverRes := srv.MCPServer().ListResources()[uri]
			if serverRes == nil {
				t.Fatalf("resource %s not found", uri)
			}

			readReq := mcp.ReadResourceRequest{}
			readReq.Params.URI = uri

			contents, err := serverRes.Handler(context.Background(), readReq)
			if err != nil || len(contents) == 0 {
				t.Fatalf("failed reading %s: %v", uri, err)
			}
			textContents, ok := contents[0].(mcp.TextResourceContents)
			if !ok {
				t.Fatalf("expected TextResourceContents for %s", uri)
			}
			if textContents.MIMEType != "application/json" {
				t.Errorf("expected application/json MIME type, got %s", textContents.MIMEType)
			}
		})
	}
}
