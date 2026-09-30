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

func TestRegisterBatch2Resources_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	resources.RegisterBatch2Resources(srv.MCPServer(), client)

	reg := srv.MCPServer().ListResources()
	expectedURIs := []string{
		"whmcs://products",
		"whmcs://support/departments",
		"whmcs://support/statuses",
	}

	for _, uri := range expectedURIs {
		if _, exists := reg[uri]; !exists {
			t.Errorf("expected resource URI %s to be registered", uri)
		}
	}
}

func TestBatch2Resources_Read(t *testing.T) {
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
				TotalResults:   3,
			})
		case "GetSupportDepartments":
			_ = json.NewEncoder(w).Encode(model.GetSupportDepartmentsResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   2,
			})
		case "GetSupportStatuses":
			_ = json.NewEncoder(w).Encode(model.GetSupportStatusesResponse{
				CommonResponse: model.CommonResponse{Result: "success"},
				TotalResults:   4,
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

	resources.RegisterBatch2Resources(srv.MCPServer(), client)

	testURIs := []string{
		"whmcs://products",
		"whmcs://support/departments",
		"whmcs://support/statuses",
	}

	for _, uri := range testURIs {
		t.Run(uri, func(t *testing.T) {
			serverRes := srv.MCPServer().ListResources()[uri]
			if serverRes == nil {
				t.Fatalf("resource %s not found", uri)
			}

			readReq := mcp.ReadResourceRequest{}
			readReq.Params.URI = uri

			contents, err := serverRes.Handler(context.Background(), readReq)
			if err != nil {
				t.Fatalf("failed reading resource: %v", err)
			}
			if len(contents) == 0 {
				t.Fatalf("expected non-empty contents for %s", uri)
			}
		})
	}
}
