package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterBillingTools mendaftarkan seluruh tool penagihan dan faktur ke MCP server.
func RegisterBillingTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_invoices
	s.AddTool(mcp.NewTool("whmcs_get_invoices",
		mcp.WithDescription("Get invoices with optional filtering by status, user, and pagination."),
		mcp.WithNumber("limitstart", mcp.Description("Starting offset (default: 0)")),
		mcp.WithNumber("limitnum", mcp.Description("Number of records to return")),
		mcp.WithNumber("userid", mcp.Description("Filter by client ID")),
		mcp.WithString("status", mcp.Description("Filter by status (Paid, Unpaid, Cancelled, Refunded, Collections, Draft, Overdue)")),
		mcp.WithString("orderby", mcp.Description("Field to order by (id, date, duedate, total, status)")),
		mcp.WithString("order", mcp.Description("Sort order (asc, desc)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetInvoicesRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetInvoicesResponse
		if err := client.Execute(ctx, "GetInvoices", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_get_invoice
	s.AddTool(mcp.NewTool("whmcs_get_invoice",
		mcp.WithDescription("Get detailed information about a specific invoice including line items and transactions."),
		mcp.WithNumber("invoiceid", mcp.Required(), mcp.Description("The ID of the invoice to retrieve")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetInvoiceRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.InvoiceID == 0 {
			return mcp.NewToolResultError("invoiceid is required"), nil
		}
		var resp model.GetInvoiceResponse
		if err := client.Execute(ctx, "GetInvoice", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_create_invoice
	s.AddTool(mcp.NewTool("whmcs_create_invoice",
		mcp.WithDescription("Create a new invoice for a client with line items."),
		mcp.WithNumber("userid", mcp.Required(), mcp.Description("Client ID to bill")),
		mcp.WithString("status", mcp.Description("Invoice status (Draft, Unpaid, Paid)")),
		mcp.WithBoolean("sendinvoice", mcp.Description("Send invoice notification email to client")),
		mcp.WithString("paymentmethod", mcp.Description("Payment method gateway")),
		mcp.WithNumber("taxrate", mcp.Description("Tax rate percentage")),
		mcp.WithNumber("taxrate2", mcp.Description("Second tax rate percentage")),
		mcp.WithString("date", mcp.Description("Invoice creation date (YYYY-MM-DD)")),
		mcp.WithString("duedate", mcp.Description("Invoice due date (YYYY-MM-DD)")),
		mcp.WithString("notes", mcp.Description("Internal staff notes")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.CreateInvoiceRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.UserID == 0 {
			return mcp.NewToolResultError("userid is required"), nil
		}
		var resp model.CreateInvoiceResponse
		if err := client.Execute(ctx, "CreateInvoice", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_update_invoice
	s.AddTool(mcp.NewTool("whmcs_update_invoice",
		mcp.WithDescription("Update an existing invoice attributes, status, or dates."),
		mcp.WithNumber("invoiceid", mcp.Required(), mcp.Description("Invoice ID to update")),
		mcp.WithString("status", mcp.Description("New status (Draft, Unpaid, Paid, Cancelled, Refunded, Collections)")),
		mcp.WithString("paymentmethod", mcp.Description("Payment method")),
		mcp.WithString("date", mcp.Description("Invoice date (YYYY-MM-DD)")),
		mcp.WithString("duedate", mcp.Description("Due date (YYYY-MM-DD)")),
		mcp.WithString("notes", mcp.Description("Invoice notes")),
		mcp.WithBoolean("publish", mcp.Description("Publish draft invoice")),
		mcp.WithBoolean("publishandsendemail", mcp.Description("Publish and send email")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.UpdateInvoiceRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.InvoiceID == 0 {
			return mcp.NewToolResultError("invoiceid is required"), nil
		}
		var resp model.UpdateInvoiceResponse
		if err := client.Execute(ctx, "UpdateInvoice", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_add_payment
	s.AddTool(mcp.NewTool("whmcs_add_payment",
		mcp.WithDescription("Record a payment on an invoice."),
		mcp.WithNumber("invoiceid", mcp.Required(), mcp.Description("Invoice ID to apply payment to")),
		mcp.WithString("transid", mcp.Required(), mcp.Description("Unique transaction identifier")),
		mcp.WithString("gateway", mcp.Required(), mcp.Description("Payment gateway module name")),
		mcp.WithNumber("amount", mcp.Description("Payment amount (defaults to balance due)")),
		mcp.WithNumber("fees", mcp.Description("Transaction processing fee")),
		mcp.WithBoolean("noemail", mcp.Description("Do not send payment confirmation email")),
		mcp.WithString("date", mcp.Description("Payment date (YYYY-MM-DD)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.AddInvoicePaymentRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.InvoiceID == 0 || req.TransID == "" || req.Gateway == "" {
			return mcp.NewToolResultError("invoiceid, transid, and gateway are required"), nil
		}
		var resp model.AddInvoicePaymentResponse
		if err := client.Execute(ctx, "AddInvoicePayment", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_apply_credit
	s.AddTool(mcp.NewTool("whmcs_apply_credit",
		mcp.WithDescription("Apply client account credit balance to pay an invoice."),
		mcp.WithNumber("invoiceid", mcp.Required(), mcp.Description("Invoice ID")),
		mcp.WithNumber("amount", mcp.Required(), mcp.Description("Amount of credit to apply")),
		mcp.WithBoolean("noemail", mcp.Description("Do not send payment confirmation email")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.ApplyCreditRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.InvoiceID == 0 || req.Amount <= 0 {
			return mcp.NewToolResultError("invoiceid and positive amount are required"), nil
		}
		var resp model.ApplyCreditResponse
		if err := client.Execute(ctx, "ApplyCredit", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 7. whmcs_get_transactions
	s.AddTool(mcp.NewTool("whmcs_get_transactions",
		mcp.WithDescription("Get payment and ledger transactions with optional filters."),
		mcp.WithNumber("invoiceid", mcp.Description("Filter by invoice ID")),
		mcp.WithNumber("clientid", mcp.Description("Filter by client ID")),
		mcp.WithString("transid", mcp.Description("Filter by transaction ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetTransactionsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetTransactionsResponse
		if err := client.Execute(ctx, "GetTransactions", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
