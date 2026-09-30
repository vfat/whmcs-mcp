package resources

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterBatch2Resources mendaftarkan resource skema whmcs:// untuk domain Produk dan Tiket Bantuan.
func RegisterBatch2Resources(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs://products
	s.AddResource(
		mcp.NewResource(
			"whmcs://products",
			"WHMCS Products Catalog",
			mcp.WithResourceDescription("List of all products and services available in WHMCS"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetProductsResponse
			if err := client.Execute(ctx, "GetProducts", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://products", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://products", resp)
		},
	)

	// 2. whmcs://support/departments
	s.AddResource(
		mcp.NewResource(
			"whmcs://support/departments",
			"Support Departments",
			mcp.WithResourceDescription("List of all support departments configured in WHMCS"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetSupportDepartmentsResponse
			if err := client.Execute(ctx, "GetSupportDepartments", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://support/departments", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://support/departments", resp)
		},
	)

	// 3. whmcs://support/statuses
	s.AddResource(
		mcp.NewResource(
			"whmcs://support/statuses",
			"Ticket Statuses",
			mcp.WithResourceDescription("Available support ticket statuses"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetSupportStatusesResponse
			if err := client.Execute(ctx, "GetSupportStatuses", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://support/statuses", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://support/statuses", resp)
		},
	)
}
