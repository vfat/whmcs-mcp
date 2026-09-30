package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterBatch4Resources mendaftarkan resource skema whmcs:// untuk server dan promosi Batch 4.
func RegisterBatch4Resources(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs://servers
	s.AddResource(
		mcp.NewResource(
			"whmcs://servers",
			"Servers",
			mcp.WithResourceDescription("Hosting servers information and active status"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetServersResponse
			if err := client.Execute(ctx, "GetServers", map[string]interface{}{"fetchStatus": true}, &resp); err != nil {
				return formatResourceJSON("whmcs://servers", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://servers", resp)
		},
	)

	// 2. whmcs://promotions
	s.AddResource(
		mcp.NewResource(
			"whmcs://promotions",
			"Promotions",
			mcp.WithResourceDescription("Active promotion codes and discount coupons"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetPromotionsResponse
			if err := client.Execute(ctx, "GetPromotions", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://promotions", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://promotions", resp)
		},
	)
}
