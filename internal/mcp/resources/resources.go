package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func formatResourceJSON(uri string, v interface{}) ([]mcp.ResourceContents, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		errJSON := fmt.Sprintf(`{"error": "%s"}`, err.Error())
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      uri,
				MIMEType: "application/json",
				Text:     errJSON,
			},
		}, nil
	}
	return []mcp.ResourceContents{
		mcp.TextResourceContents{
			URI:      uri,
			MIMEType: "application/json",
			Text:     string(b),
		},
	}, nil
}

// RegisterBatch1Resources mendaftarkan 5 MCP Resource inti Batch 1.
func RegisterBatch1Resources(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs://stats
	s.AddResource(
		mcp.NewResource(
			"whmcs://stats",
			"WHMCS Statistics",
			mcp.WithResourceDescription("Current WHMCS system statistics including client counts, revenue, and service status"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetStatsResponse
			if err := client.Execute(ctx, "GetStats", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://stats", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://stats", resp)
		},
	)

	// 2. whmcs://admin-users
	s.AddResource(
		mcp.NewResource(
			"whmcs://admin-users",
			"Admin Users",
			mcp.WithResourceDescription("List of administrator users in WHMCS"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetAdminUsersResponse
			if err := client.Execute(ctx, "GetAdminUsers", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://admin-users", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://admin-users", resp)
		},
	)

	// 3. whmcs://currencies
	s.AddResource(
		mcp.NewResource(
			"whmcs://currencies",
			"Currencies",
			mcp.WithResourceDescription("Currency configuration in WHMCS"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetCurrenciesResponse
			if err := client.Execute(ctx, "GetCurrencies", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://currencies", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://currencies", resp)
		},
	)

	// 4. whmcs://payment-methods
	s.AddResource(
		mcp.NewResource(
			"whmcs://payment-methods",
			"Payment Methods",
			mcp.WithResourceDescription("Available payment methods configured in WHMCS"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetPaymentMethodsResponse
			if err := client.Execute(ctx, "GetPaymentMethods", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://payment-methods", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://payment-methods", resp)
		},
	)

	// 5. whmcs://admin/todo
	s.AddResource(
		mcp.NewResource(
			"whmcs://admin/todo",
			"Admin To-Do Items",
			mcp.WithResourceDescription("Administrative to-do items and tasks"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			var resp model.GetToDoItemsResponse
			if err := client.Execute(ctx, "GetToDoItems", map[string]interface{}{}, &resp); err != nil {
				return formatResourceJSON("whmcs://admin/todo", map[string]string{"error": err.Error()})
			}
			return formatResourceJSON("whmcs://admin/todo", resp)
		},
	)
}
