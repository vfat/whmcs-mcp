package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterOrderTools mendaftarkan seluruh 10 tool Order & Quote management ke MCP server.
func RegisterOrderTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_orders
	s.AddTool(mcp.NewTool("whmcs_get_orders",
		mcp.WithDescription("Get client orders with optional filtering by status or client ID."),
		mcp.WithNumber("limitstart", mcp.Description("Starting offset (default: 0)")),
		mcp.WithNumber("limitnum", mcp.Description("Number of records to return")),
		mcp.WithNumber("id", mcp.Description("Specific order ID")),
		mcp.WithNumber("userid", mcp.Description("Filter by client ID")),
		mcp.WithString("status", mcp.Description("Filter by status (Pending, Active, Fraud, Cancelled)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetOrdersRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetOrdersResponse
		if err := client.Execute(ctx, "GetOrders", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_accept_order
	s.AddTool(mcp.NewTool("whmcs_accept_order",
		mcp.WithDescription("Accept and process a pending order, optionally provisioning products or registering domains."),
		mcp.WithNumber("orderid", mcp.Required(), mcp.Description("Order ID to accept")),
		mcp.WithNumber("serverid", mcp.Description("Server ID to provision products on")),
		mcp.WithString("serviceusername", mcp.Description("Username for hosting service")),
		mcp.WithString("servicepassword", mcp.Description("Password for hosting service")),
		mcp.WithString("registrar", mcp.Description("Domain registrar module name")),
		mcp.WithBoolean("autosetup", mcp.Description("Trigger automatic product setup")),
		mcp.WithBoolean("sendemail", mcp.Description("Send welcome/order confirmation email")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.AcceptOrderRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.OrderID == 0 {
			return mcp.NewToolResultError("orderid is required"), nil
		}
		var resp model.AcceptOrderResponse
		if err := client.Execute(ctx, "AcceptOrder", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_cancel_order
	s.AddTool(mcp.NewTool("whmcs_cancel_order",
		mcp.WithDescription("Cancel an order in WHMCS."),
		mcp.WithNumber("orderid", mcp.Required(), mcp.Description("Order ID to cancel")),
		mcp.WithBoolean("cancelsub", mcp.Description("Cancel recurring payment subscription")),
		mcp.WithBoolean("noemail", mcp.Description("Do not send email notification")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.CancelOrderRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.OrderID == 0 {
			return mcp.NewToolResultError("orderid is required"), nil
		}
		var resp model.CancelOrderResponse
		if err := client.Execute(ctx, "CancelOrder", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_delete_order
	s.AddTool(mcp.NewTool("whmcs_delete_order",
		mcp.WithDescription("Permanently delete an order from WHMCS (destructive operation)."),
		mcp.WithNumber("orderid", mcp.Required(), mcp.Description("Order ID to delete")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.DeleteOrderRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.OrderID == 0 {
			return mcp.NewToolResultError("orderid is required"), nil
		}
		var resp model.DeleteOrderResponse
		if err := client.Execute(ctx, "DeleteOrder", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_fraud_order
	s.AddTool(mcp.NewTool("whmcs_fraud_order",
		mcp.WithDescription("Mark an order as fraudulent and cancel associated subscriptions."),
		mcp.WithNumber("orderid", mcp.Required(), mcp.Description("Order ID to mark as fraud")),
		mcp.WithBoolean("cancelsub", mcp.Description("Cancel payment subscription")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.FraudOrderRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.OrderID == 0 {
			return mcp.NewToolResultError("orderid is required"), nil
		}
		var resp model.FraudOrderResponse
		if err := client.Execute(ctx, "FraudOrder", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_pending_order
	s.AddTool(mcp.NewTool("whmcs_pending_order",
		mcp.WithDescription("Reset an order status back to Pending."),
		mcp.WithNumber("orderid", mcp.Required(), mcp.Description("Order ID to set to pending")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.PendingOrderRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.OrderID == 0 {
			return mcp.NewToolResultError("orderid is required"), nil
		}
		var resp model.PendingOrderResponse
		if err := client.Execute(ctx, "PendingOrder", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 7. whmcs_get_quotes
	s.AddTool(mcp.NewTool("whmcs_get_quotes",
		mcp.WithDescription("Get list of price quotes with optional filters."),
		mcp.WithNumber("quoteid", mcp.Description("Specific quote ID")),
		mcp.WithNumber("clientid", mcp.Description("Filter by client ID")),
		mcp.WithString("subject", mcp.Description("Filter by quote subject")),
		mcp.WithString("stage", mcp.Description("Filter by quote stage (Draft, Delivered, On Hold, Accepted, Lost, Dead)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetQuotesRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetQuotesResponse
		if err := client.Execute(ctx, "GetQuotes", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 8. whmcs_create_quote
	s.AddTool(mcp.NewTool("whmcs_create_quote",
		mcp.WithDescription("Create a new price quote for a client or prospect."),
		mcp.WithString("subject", mcp.Required(), mcp.Description("Quote subject line")),
		mcp.WithString("stage", mcp.Required(), mcp.Description("Quote stage (Draft, Delivered, On Hold, Accepted, Lost, Dead)")),
		mcp.WithString("date", mcp.Description("Quote date (YYYY-MM-DD)")),
		mcp.WithString("validuntil", mcp.Description("Quote expiration date (YYYY-MM-DD)")),
		mcp.WithNumber("userid", mcp.Description("Associated client ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.CreateQuoteRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.Subject == "" || req.Stage == "" {
			return mcp.NewToolResultError("subject and stage are required"), nil
		}
		var resp model.CreateQuoteResponse
		if err := client.Execute(ctx, "CreateQuote", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 9. whmcs_accept_quote
	s.AddTool(mcp.NewTool("whmcs_accept_quote",
		mcp.WithDescription("Accept a quote and automatically convert it into an invoice."),
		mcp.WithNumber("quoteid", mcp.Required(), mcp.Description("Quote ID to accept")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.AcceptQuoteRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.QuoteID == 0 {
			return mcp.NewToolResultError("quoteid is required"), nil
		}
		var resp model.AcceptQuoteResponse
		if err := client.Execute(ctx, "AcceptQuote", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 10. whmcs_delete_quote
	s.AddTool(mcp.NewTool("whmcs_delete_quote",
		mcp.WithDescription("Permanently delete a price quote."),
		mcp.WithNumber("quoteid", mcp.Required(), mcp.Description("Quote ID to delete")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.DeleteQuoteRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.QuoteID == 0 {
			return mcp.NewToolResultError("quoteid is required"), nil
		}
		var resp model.DeleteQuoteResponse
		if err := client.Execute(ctx, "DeleteQuote", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
