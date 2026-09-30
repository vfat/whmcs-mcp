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

func TestRegisterBatch3Resources_AllPresent(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	resources.RegisterBatch3Resources(srv.MCPServer(), client)

	reg := srv.MCPServer().ListResources()
	expectedURIs := []string{
		"whmcs://tld-pricing",
	}

	for _, uri := range expectedURIs {
		if _, exists := reg[uri]; !exists {
			t.Errorf("expected resource URI %s to be registered", uri)
		}
	}
}

func TestBatch3Resources_Read(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(model.GetTLDPricingResponse{
			CommonResponse: model.CommonResponse{Result: "success"},
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

	resources.RegisterBatch3Resources(srv.MCPServer(), client)

	serverRes := srv.MCPServer().ListResources()["whmcs://tld-pricing"]
	if serverRes == nil {
		t.Fatal("resource whmcs://tld-pricing not found")
	}

	readReq := mcp.ReadResourceRequest{}
	readReq.Params.URI = "whmcs://tld-pricing"

	contents, err := serverRes.Handler(context.Background(), readReq)
	if err != nil || len(contents) == 0 {
		t.Fatalf("failed reading whmcs://tld-pricing: %v", err)
	}
}
