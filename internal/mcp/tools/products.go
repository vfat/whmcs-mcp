package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterProductTools mendaftarkan tool katalog produk ke MCP server.
func RegisterProductTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_products
	s.AddTool(mcp.NewTool("whmcs_get_products",
		mcp.WithDescription("Get available products/services from WHMCS with optional filtering."),
		mcp.WithNumber("pid", mcp.Description("Specific product ID")),
		mcp.WithNumber("gid", mcp.Description("Filter by product group ID")),
		mcp.WithString("module", mcp.Description("Filter by server module")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetProductsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetProductsResponse
		if err := client.Execute(ctx, "GetProducts", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_get_product_groups
	s.AddTool(mcp.NewTool("whmcs_get_product_groups",
		mcp.WithDescription("Get all product groups from WHMCS."),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetProductGroupsRequest
		var resp model.GetProductGroupsResponse
		if err := client.Execute(ctx, "GetProductGroups", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
