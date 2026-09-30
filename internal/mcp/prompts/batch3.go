package prompts

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch3Prompts mendaftarkan templat Prompt MCP untuk domain Domain & Order Batch 3.
func RegisterBatch3Prompts(s *server.MCPServer) {
	// 1. domain-expiry-audit
	s.AddPrompt(
		mcp.NewPrompt(
			"domain-expiry-audit",
			mcp.WithPromptTitle("Domain Expiry Audit"),
			mcp.WithPromptDescription("Audit domains expiring soon and take action"),
			mcp.WithArgument("daysUntilExpiry", mcp.RequiredArgument(), mcp.ArgumentDescription("Days until expiry to check")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			daysUntilExpiry := request.Params.Arguments["daysUntilExpiry"]

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("Audit all domains expiring within the next %s days.\n\n", daysUntilExpiry))
			sb.WriteString("Please:\n")
			sb.WriteString("1. Use whmcs_get_clients to get all clients\n")
			sb.WriteString("2. For each client, use whmcs_get_client_domains to check domain expiry dates\n")
			sb.WriteString(fmt.Sprintf("3. Compile a list of domains expiring within %s days\n\n", daysUntilExpiry))
			sb.WriteString("Report should include:\n")
			sb.WriteString("- Domain name\n")
			sb.WriteString("- Client name and contact\n")
			sb.WriteString("- Expiry date\n")
			sb.WriteString("- Auto-renew status\n")
			sb.WriteString("- Recommended action (renew, let expire, contact client)\n\n")
			sb.WriteString("After the audit, I may ask you to:\n")
			sb.WriteString("- Send renewal reminders via whmcs_send_email\n")
			sb.WriteString("- Process renewals via whmcs_renew_domain")

			msg := mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(sb.String()))
			return mcp.NewGetPromptResult("Domain Expiry Audit Guide", []mcp.PromptMessage{msg}), nil
		},
	)

	// 2. fraud-investigation
	s.AddPrompt(
		mcp.NewPrompt(
			"fraud-investigation",
			mcp.WithPromptTitle("Fraud Investigation"),
			mcp.WithPromptDescription("Investigate a potentially fraudulent order or client"),
			mcp.WithArgument("orderId", mcp.ArgumentDescription("Order ID to investigate")),
			mcp.WithArgument("clientId", mcp.ArgumentDescription("Client ID to investigate")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			orderID := request.Params.Arguments["orderId"]
			clientID := request.Params.Arguments["clientId"]

			target := "target"
			if orderID != "" {
				target = fmt.Sprintf("order #%s", orderID)
			} else if clientID != "" {
				target = fmt.Sprintf("client #%s", clientID)
			}

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("I need to investigate a potentially fraudulent %s.\n\n", target))
			sb.WriteString("Please help me gather evidence:\n")
			if orderID != "" {
				sb.WriteString(fmt.Sprintf("1. Use whmcs_get_orders to get order details for order #%s\n", orderID))
			}
			if clientID != "" {
				sb.WriteString(fmt.Sprintf("1. Use whmcs_get_client_details to review client #%s\n", clientID))
			}
			sb.WriteString("2. Check the activity log (whmcs_get_activity_log) for suspicious patterns\n")
			sb.WriteString("3. Review any related support tickets (whmcs_get_tickets)\n")
			sb.WriteString("4. Check payment history and chargebacks (whmcs_get_transactions)\n\n")
			sb.WriteString("Red flags to look for:\n")
			sb.WriteString("- Mismatched billing/service locations\n")
			sb.WriteString("- Multiple failed payment attempts\n")
			sb.WriteString("- Rapid order submissions\n")
			sb.WriteString("- Known proxy/VPN usage\n")
			sb.WriteString("- Disposable email addresses\n\n")
			sb.WriteString("Based on findings, recommend action:\n")
			sb.WriteString("- Approve order\n")
			sb.WriteString("- Request verification\n")
			sb.WriteString("- Mark as fraudulent (whmcs_fraud_order)\n")
			sb.WriteString("- Suspend account")

			msg := mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent(sb.String()))
			return mcp.NewGetPromptResult("Fraud Investigation Guide", []mcp.PromptMessage{msg}), nil
		},
	)
}
