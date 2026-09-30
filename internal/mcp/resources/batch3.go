package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterBatch3Resources mendaftarkan resource skema whmcs:// untuk domain Siklus Hidup Domain.
func RegisterBatch3Resources(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs://tld-pricing
	s.AddResource(
		mcp.NewResource(
			"whmcs://tld-pricing",
			"TLD Pricing",
			mcp.WithResourceDescription("Domain TLD pricing information"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetTLDPricingResponse
			if err := client.Execute(ctx, "GetTLDPricing", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://tld-pricing", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://tld-pricing", resp)
		},
	)
}
