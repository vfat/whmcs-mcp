package prompts

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch2Prompts mendaftarkan templat Prompt MCP untuk domain Finansial dan Layanan Tiket Batch 2.
func RegisterBatch2Prompts(s *server.MCPServer) {
	// 1. ticket-response
	s.AddPrompt(
		mcp.NewPrompt(
			"ticket-response",
			mcp.WithPromptTitle("Support Ticket Response"),
			mcp.WithPromptDescription("Generate a professional response to a support ticket"),
			mcp.WithArgument("ticketId", mcp.RequiredArgument(), mcp.ArgumentDescription("The ticket ID to respond to")),
			mcp.WithArgument("issueType", mcp.RequiredArgument(), mcp.ArgumentDescription("Type of issue (billing, technical, sales, general)")),
			mcp.WithArgument("tone", mcp.ArgumentDescription("Tone of the response (formal, friendly, apologetic)")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			ticketID := request.Params.Arguments["ticketId"]
			issueType := request.Params.Arguments["issueType"]
			tone := request.Params.Arguments["tone"]
			if tone == "" {
				tone = "friendly"
			}

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("Please help me respond to support ticket #%s.\n\n", ticketID))
			sb.WriteString("First, use whmcs_get_ticket to retrieve the ticket details.\n\n")
			sb.WriteString(fmt.Sprintf("Then craft a %s response appropriate for a %s issue that:\n", tone, issueType))
			sb.WriteString("- Acknowledges the customer's concern\n")
			sb.WriteString("- Provides a clear solution or next steps\n")
			sb.WriteString("- Offers additional assistance if needed\n\n")
			sb.WriteString("After reviewing the ticket, draft the response and I'll confirm before sending with whmcs_add_ticket_reply.")

			msg := mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(sb.String()))
			return mcp.NewGetPromptResult("Support Ticket Response Guide", []mcp.PromptMessage{msg}), nil
		},
	)

	// 2. revenue-report
	s.AddPrompt(
		mcp.NewPrompt(
			"revenue-report",
			mcp.WithPromptTitle("Revenue Analysis Report"),
			mcp.WithPromptDescription("Generate a comprehensive revenue analysis from WHMCS data"),
			mcp.WithArgument("period", mcp.RequiredArgument(), mcp.ArgumentDescription("Time period for analysis (today, week, month, year)")),
			mcp.WithArgument("includeForecasting", mcp.ArgumentDescription("Include revenue forecasting (true/false)")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			period := request.Params.Arguments["period"]
			includeForecasting := request.Params.Arguments["includeForecasting"] == "true"

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("Generate a revenue analysis report for the %s.\n\n", period))
			sb.WriteString("Please gather the following data:\n")
			sb.WriteString("1. Use whmcs_get_stats to get overall system statistics\n")
			sb.WriteString("2. Use whmcs_get_transactions to analyze payment transactions\n")
			sb.WriteString("3. Use whmcs_get_invoices to review invoice status\n\n")
			sb.WriteString("Provide insights on:\n")
			sb.WriteString("- Total revenue collected\n")
			sb.WriteString("- Outstanding invoices\n")
			sb.WriteString("- Payment method breakdown\n")
			sb.WriteString("- Top revenue-generating products\n")
			if includeForecasting {
				sb.WriteString("- Revenue forecast for next period based on trends\n")
			}
			sb.WriteString("\nPresent the data in a clear, executive summary format.")

			msg := mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(sb.String()))
			return mcp.NewGetPromptResult("Revenue Analysis Guide", []mcp.PromptMessage{msg}), nil
		},
	)

	// 3. bulk-invoice-reminder
	s.AddPrompt(
		mcp.NewPrompt(
			"bulk-invoice-reminder",
			mcp.WithPromptTitle("Bulk Invoice Reminder"),
			mcp.WithPromptDescription("Send payment reminders to clients with overdue invoices"),
			mcp.WithArgument("daysOverdue", mcp.RequiredArgument(), mcp.ArgumentDescription("Minimum days overdue to include")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			daysOverdue := request.Params.Arguments["daysOverdue"]

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("Help me send payment reminders to clients with invoices overdue by %s+ days.\n\n", daysOverdue))
			sb.WriteString("Steps:\n")
			sb.WriteString("1. Use whmcs_get_invoices with status=Overdue to find overdue invoices\n")
			sb.WriteString(fmt.Sprintf("2. Filter for invoices overdue by at least %s days\n", daysOverdue))
			sb.WriteString("3. Group by client to avoid multiple emails to the same client\n")
			sb.WriteString("4. For each client, use whmcs_send_email to send a payment reminder\n\n")
			sb.WriteString("Before sending any emails, show me:\n")
			sb.WriteString("- List of clients to contact\n")
			sb.WriteString("- Total amount outstanding per client\n")
			sb.WriteString("- Number of overdue invoices per client\n\n")
			sb.WriteString("I'll confirm before you proceed with sending reminders.")

			msg := mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(sb.String()))
			return mcp.NewGetPromptResult("Bulk Invoice Reminder Guide", []mcp.PromptMessage{msg}), nil
		},
	)
}
