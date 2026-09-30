package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterSystemTools mendaftarkan 9 tools administrasi sistem ke MCP server.
func RegisterSystemTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_stats
	s.AddTool(mcp.NewTool("whmcs_get_stats",
		mcp.WithDescription("Get WHMCS system statistics including income, order counts, tickets, and active services."),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetStatsRequest
		var resp model.GetStatsResponse
		if err := client.Execute(ctx, "GetStats", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_get_activity_log
	s.AddTool(mcp.NewTool("whmcs_get_activity_log",
		mcp.WithDescription("Get system audit activity log with optional filters."),
		mcp.WithNumber("limitstart", mcp.Description("Starting offset (default: 0)")),
		mcp.WithNumber("limitnum", mcp.Description("Number of records to return")),
		mcp.WithNumber("userid", mcp.Description("Filter by user/client ID")),
		mcp.WithString("date", mcp.Description("Filter by date (YYYY-MM-DD)")),
		mcp.WithString("user", mcp.Description("Filter by username/admin")),
		mcp.WithString("description", mcp.Description("Filter by log description text")),
		mcp.WithString("ipaddress", mcp.Description("Filter by IP address")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetActivityLogRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetActivityLogResponse
		if err := client.Execute(ctx, "GetActivityLog", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_get_admin_users
	s.AddTool(mcp.NewTool("whmcs_get_admin_users",
		mcp.WithDescription("Get list of staff administrator users."),
		mcp.WithNumber("roleid", mcp.Description("Filter by admin role ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetAdminUsersRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetAdminUsersResponse
		if err := client.Execute(ctx, "GetAdminUsers", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_get_todo_items
	s.AddTool(mcp.NewTool("whmcs_get_todo_items",
		mcp.WithDescription("Get administrative checklist and staff to-do items."),
		mcp.WithNumber("limitstart", mcp.Description("Starting offset (default: 0)")),
		mcp.WithNumber("limitnum", mcp.Description("Number of items to return")),
		mcp.WithString("status", mcp.Description("Filter by status (Incomplete, Complete, Pending)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetToDoItemsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetToDoItemsResponse
		if err := client.Execute(ctx, "GetToDoItems", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_get_todo_item_statuses
	s.AddTool(mcp.NewTool("whmcs_get_todo_item_statuses",
		mcp.WithDescription("Get list of available to-do item statuses and item counts."),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var resp model.GetToDoStatusesResponse
		if err := client.Execute(ctx, "GetToDoItemStatuses", map[string]interface{}{}, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_update_todo_item
	s.AddTool(mcp.NewTool("whmcs_update_todo_item",
		mcp.WithDescription("Update administrative to-do item attributes, assignment, or completion status."),
		mcp.WithNumber("itemid", mcp.Required(), mcp.Description("The ID of the to-do item to update")),
		mcp.WithString("status", mcp.Description("New status of the to-do item")),
		mcp.WithString("date", mcp.Description("Date of the item (YYYY-MM-DD)")),
		mcp.WithString("title", mcp.Description("Title of the item")),
		mcp.WithString("description", mcp.Description("Detailed description of the item")),
		mcp.WithNumber("adminid", mcp.Description("Assigned administrator ID")),
		mcp.WithString("duedate", mcp.Description("Due date (YYYY-MM-DD)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.UpdateToDoItemRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.ID == 0 {
			return mcp.NewToolResultError("itemid is required"), nil
		}
		var resp model.UpdateToDoItemResponse
		if err := client.Execute(ctx, "UpdateToDoItem", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 7. whmcs_get_currencies
	s.AddTool(mcp.NewTool("whmcs_get_currencies",
		mcp.WithDescription("Get configured billing currencies and exchange rates."),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var resp model.GetCurrenciesResponse
		if err := client.Execute(ctx, "GetCurrencies", map[string]interface{}{}, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 8. whmcs_get_payment_methods
	s.AddTool(mcp.NewTool("whmcs_get_payment_methods",
		mcp.WithDescription("Get active payment gateway methods."),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var resp model.GetPaymentMethodsResponse
		if err := client.Execute(ctx, "GetPaymentMethods", map[string]interface{}{}, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 9. whmcs_send_email
	s.AddTool(mcp.NewTool("whmcs_send_email",
		mcp.WithDescription("Send an email or trigger an email template to a client."),
		mcp.WithString("messagename", mcp.Description("Email template name to send")),
		mcp.WithNumber("id", mcp.Description("Related entity ID (client ID, invoice ID, etc.)")),
		mcp.WithString("customtype", mcp.Description("Custom type (general, product, domain, invoice, support, affiliate)")),
		mcp.WithString("customsubject", mcp.Description("Custom subject for direct email")),
		mcp.WithString("custommessage", mcp.Description("Custom message body for direct email")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.SendEmailRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.SendEmailResponse
		if err := client.Execute(ctx, "SendEmail", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
