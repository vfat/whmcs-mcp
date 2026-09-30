package prompts

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch1Prompts mendaftarkan templat Prompt MCP untuk domain Klien Batch 1.
func RegisterBatch1Prompts(s *server.MCPServer) {
	// 1. client-onboarding
	s.AddPrompt(
		mcp.NewPrompt(
			"client-onboarding",
			mcp.WithPromptTitle("Client Onboarding Assistant"),
			mcp.WithPromptDescription("Guide through onboarding a new client including account setup, product assignment, and welcome communication"),
			mcp.WithArgument("clientName", mcp.RequiredArgument(), mcp.ArgumentDescription("Name of the new client")),
			mcp.WithArgument("email", mcp.RequiredArgument(), mcp.ArgumentDescription("Client email address")),
			mcp.WithArgument("products", mcp.ArgumentDescription("Comma-separated list of products to assign")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			clientName := request.Params.Arguments["clientName"]
			email := request.Params.Arguments["email"]
			products := request.Params.Arguments["products"]

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("I need to onboard a new client with the following details:\n\n"))
			sb.WriteString(fmt.Sprintf("**Client Name:** %s\n", clientName))
			sb.WriteString(fmt.Sprintf("**Email:** %s\n", email))
			if products != "" {
				sb.WriteString(fmt.Sprintf("**Products to Assign:** %s\n", products))
			}
			sb.WriteString("\nPlease help me:\n")
			sb.WriteString("1. Create the client account in WHMCS using whmcs_add_client\n")
			sb.WriteString("2. Set up any requested products/services\n")
			sb.WriteString("3. Generate a welcome invoice if needed\n")
			sb.WriteString("4. Suggest a welcome email template\n\n")
			sb.WriteString("Walk me through each step and confirm before proceeding.")

			msg := mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(sb.String()))
			return mcp.NewGetPromptResult("Client Onboarding Guide", []mcp.PromptMessage{msg}), nil
		},
	)

	// 2. client-health-check
	s.AddPrompt(
		mcp.NewPrompt(
			"client-health-check",
			mcp.WithPromptTitle("Client Account Health Check"),
			mcp.WithPromptDescription("Perform a comprehensive health check on a client account"),
			mcp.WithArgument("clientId", mcp.RequiredArgument(), mcp.ArgumentDescription("Client ID to analyze")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			clientID := request.Params.Arguments["clientId"]

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("Perform a comprehensive health check on client #%s.\n\n", clientID))
			sb.WriteString("Please gather and analyze:\n")
			sb.WriteString("1. Use whmcs_get_client_details with stats=true for client overview\n")
			sb.WriteString("2. Use whmcs_get_client_products to review active services\n")
			sb.WriteString("3. Use whmcs_get_client_domains to check domain status\n")
			sb.WriteString("4. Use whmcs_get_invoices filtered by clientid to review payment history\n")
			sb.WriteString("5. Use whmcs_get_tickets filtered by clientid to review support history\n\n")
			sb.WriteString("Provide a summary including:\n")
			sb.WriteString("- Account standing (good, at-risk, churning)\n")
			sb.WriteString("- Services overview and renewal dates\n")
			sb.WriteString("- Payment history and any overdue amounts\n")
			sb.WriteString("- Recent support interactions\n")
			sb.WriteString("- Recommendations for account management")

			msg := mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(sb.String()))
			return mcp.NewGetPromptResult("Client Health Check Guide", []mcp.PromptMessage{msg}), nil
		},
	)
}
