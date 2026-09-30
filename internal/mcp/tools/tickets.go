package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterTicketTools mendaftarkan seluruh 9 tool tiket bantuan pelanggan ke MCP server.
func RegisterTicketTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_tickets
	s.AddTool(mcp.NewTool("whmcs_get_tickets",
		mcp.WithDescription("Get support tickets with optional filtering by status, client, or department."),
		mcp.WithNumber("limitstart", mcp.Description("Starting offset (default: 0)")),
		mcp.WithNumber("limitnum", mcp.Description("Number of records to return")),
		mcp.WithNumber("deptid", mcp.Description("Filter by department ID")),
		mcp.WithNumber("clientid", mcp.Description("Filter by client ID")),
		mcp.WithString("email", mcp.Description("Filter by contact email address")),
		mcp.WithString("status", mcp.Description("Filter by ticket status")),
		mcp.WithString("subject", mcp.Description("Filter by ticket subject text")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetTicketsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetTicketsResponse
		if err := client.Execute(ctx, "GetTickets", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_get_ticket
	s.AddTool(mcp.NewTool("whmcs_get_ticket",
		mcp.WithDescription("Get detailed information about a specific support ticket including conversation replies and staff notes."),
		mcp.WithNumber("ticketid", mcp.Required(), mcp.Description("The ticket ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetTicketRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.TicketID == 0 {
			return mcp.NewToolResultError("ticketid is required"), nil
		}
		var resp model.GetTicketResponse
		if err := client.Execute(ctx, "GetTicket", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_open_ticket
	s.AddTool(mcp.NewTool("whmcs_open_ticket",
		mcp.WithDescription("Create and open a new support ticket in WHMCS."),
		mcp.WithNumber("deptid", mcp.Required(), mcp.Description("Target department ID")),
		mcp.WithString("subject", mcp.Required(), mcp.Description("Ticket subject line")),
		mcp.WithString("message", mcp.Required(), mcp.Description("Ticket body message")),
		mcp.WithNumber("clientid", mcp.Description("Associated client ID (if existing client)")),
		mcp.WithNumber("contactid", mcp.Description("Associated contact ID")),
		mcp.WithString("name", mcp.Description("Contact name (if not a client)")),
		mcp.WithString("email", mcp.Description("Contact email (if not a client)")),
		mcp.WithString("priority", mcp.Description("Priority (Low, Medium, High)")),
		mcp.WithNumber("serviceid", mcp.Description("Related service/product ID")),
		mcp.WithNumber("domainid", mcp.Description("Related domain ID")),
		mcp.WithBoolean("admin", mcp.Description("Flag if ticket is created on behalf of admin")),
		mcp.WithBoolean("markdown", mcp.Description("Whether the message body contains markdown formatting")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.OpenTicketRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DeptID == 0 || req.Subject == "" || req.Message == "" {
			return mcp.NewToolResultError("deptid, subject, and message are required"), nil
		}
		var resp model.OpenTicketResponse
		if err := client.Execute(ctx, "OpenTicket", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_add_ticket_reply
	s.AddTool(mcp.NewTool("whmcs_add_ticket_reply",
		mcp.WithDescription("Add a reply to an existing support ticket thread."),
		mcp.WithNumber("ticketid", mcp.Required(), mcp.Description("Ticket ID to reply to")),
		mcp.WithString("message", mcp.Required(), mcp.Description("Reply message body")),
		mcp.WithNumber("clientid", mcp.Description("Client ID if replying as client")),
		mcp.WithNumber("contactid", mcp.Description("Contact ID")),
		mcp.WithString("name", mcp.Description("Display name")),
		mcp.WithString("email", mcp.Description("Contact email")),
		mcp.WithString("adminusername", mcp.Description("Admin username if replying as staff")),
		mcp.WithString("status", mcp.Description("Update ticket status on reply")),
		mcp.WithBoolean("noemail", mcp.Description("Do not send email notification")),
		mcp.WithBoolean("markdown", mcp.Description("Whether the message contains markdown")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.AddTicketReplyRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.TicketID == 0 || req.Message == "" {
			return mcp.NewToolResultError("ticketid and message are required"), nil
		}
		var resp model.AddTicketReplyResponse
		if err := client.Execute(ctx, "AddTicketReply", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_add_ticket_note
	s.AddTool(mcp.NewTool("whmcs_add_ticket_note",
		mcp.WithDescription("Add an admin-only internal note to a ticket. Notes are hidden from clients and do not trigger emails."),
		mcp.WithNumber("ticketid", mcp.Required(), mcp.Description("Ticket ID to add the note to")),
		mcp.WithString("message", mcp.Required(), mcp.Description("The internal note content")),
		mcp.WithBoolean("markdown", mcp.Description("Whether the message contains markdown formatting")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.AddTicketNoteRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.TicketID == 0 || req.Message == "" {
			return mcp.NewToolResultError("ticketid and message are required"), nil
		}
		var resp model.AddTicketNoteResponse
		if err := client.Execute(ctx, "AddTicketNote", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_update_ticket
	s.AddTool(mcp.NewTool("whmcs_update_ticket",
		mcp.WithDescription("Update ticket properties such as department, priority, status, or assignment flag."),
		mcp.WithNumber("ticketid", mcp.Required(), mcp.Description("Ticket ID to update")),
		mcp.WithNumber("deptid", mcp.Description("Reassign to department ID")),
		mcp.WithString("subject", mcp.Description("Update subject line")),
		mcp.WithNumber("userid", mcp.Description("Reassign to client ID")),
		mcp.WithString("name", mcp.Description("Contact name")),
		mcp.WithString("email", mcp.Description("Contact email")),
		mcp.WithString("priority", mcp.Description("Priority (Low, Medium, High)")),
		mcp.WithString("status", mcp.Description("New ticket status")),
		mcp.WithNumber("flag", mcp.Description("Flag/assign to admin ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.UpdateTicketRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.TicketID == 0 {
			return mcp.NewToolResultError("ticketid is required"), nil
		}
		var resp model.UpdateTicketResponse
		if err := client.Execute(ctx, "UpdateTicket", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 7. whmcs_delete_ticket
	s.AddTool(mcp.NewTool("whmcs_delete_ticket",
		mcp.WithDescription("Permanently delete a support ticket (destructive operation)."),
		mcp.WithNumber("ticketid", mcp.Required(), mcp.Description("Ticket ID to permanently delete")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.DeleteTicketRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.TicketID == 0 {
			return mcp.NewToolResultError("ticketid is required"), nil
		}
		var resp model.DeleteTicketResponse
		if err := client.Execute(ctx, "DeleteTicket", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 8. whmcs_get_support_departments
	s.AddTool(mcp.NewTool("whmcs_get_support_departments",
		mcp.WithDescription("Get list of support departments and open ticket counts."),
		mcp.WithBoolean("ignore_dept_assignments", mcp.Description("Ignore administrator department assignments")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetSupportDepartmentsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetSupportDepartmentsResponse
		if err := client.Execute(ctx, "GetSupportDepartments", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 9. whmcs_get_support_statuses
	s.AddTool(mcp.NewTool("whmcs_get_support_statuses",
		mcp.WithDescription("Get configured support ticket statuses and ticket counts."),
		mcp.WithNumber("deptid", mcp.Description("Filter counts by department ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetSupportStatusesRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetSupportStatusesResponse
		if err := client.Execute(ctx, "GetSupportStatuses", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
