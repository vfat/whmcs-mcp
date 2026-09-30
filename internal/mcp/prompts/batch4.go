package prompts

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterBatch4Prompts mendaftarkan templat Prompt MCP untuk konfigurasi produk hosting baru.
func RegisterBatch4Prompts(s *server.MCPServer) {
	// 1. new-product-setup
	s.AddPrompt(
		mcp.NewPrompt(
			"new-product-setup",
			mcp.WithPromptTitle("New Product Setup Guide"),
			mcp.WithPromptDescription("Guide through creating a new product/service including pricing and server module configuration"),
			mcp.WithArgument("productName", mcp.RequiredArgument(), mcp.ArgumentDescription("Name of the new product")),
			mcp.WithArgument("productType", mcp.RequiredArgument(), mcp.ArgumentDescription("Type of product (hostingaccount, reselleraccount, server, other)")),
			mcp.WithArgument("monthlyPrice", mcp.ArgumentDescription("Monthly price (e.g., '9.99')")),
			mcp.WithArgument("serverType", mcp.ArgumentDescription("Server module type (e.g., 'cpanel', 'plesk')")),
		),
		func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
			productName := request.Params.Arguments["productName"]
			productType := request.Params.Arguments["productType"]
			monthlyPrice := request.Params.Arguments["monthlyPrice"]
			serverType := request.Params.Arguments["serverType"]

			var sb strings.Builder
			sb.WriteString(fmt.Sprintf("I need to set up a new product in WHMCS with the following details:\n\n"))
			sb.WriteString(fmt.Sprintf("**Product Name:** %s\n", productName))
			sb.WriteString(fmt.Sprintf("**Product Type:** %s\n", productType))
			if monthlyPrice != "" {
				sb.WriteString(fmt.Sprintf("**Monthly Price:** $%s\n", monthlyPrice))
			}
			if serverType != "" {
				sb.WriteString(fmt.Sprintf("**Server Module:** %s\n", serverType))
			}
			sb.WriteString("\nPlease guide me through:\n")
			sb.WriteString("1. Recommended product group assignment\n")
			sb.WriteString("2. Pricing structure setup (monthly, annual, setup fees)\n")
			sb.WriteString("3. Module settings and automation options\n")
			sb.WriteString("4. Custom fields or configurable options that might be useful\n")
			sb.WriteString("5. Welcome email template selection\n\n")
			sb.WriteString("Provide recommendations for best practices.")

			return &mcp.GetPromptResult{
				Description: "Step-by-step guide to configure a new hosting package in WHMCS",
				Messages: []mcp.PromptMessage{
					{
						Role: mcp.RoleUser,
						Content: mcp.TextContent{
							Type: "text",
							Text: sb.String(),
						},
					},
				},
			}, nil
		},
	)
}
