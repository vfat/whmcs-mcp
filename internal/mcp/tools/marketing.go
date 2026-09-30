package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterMarketingTools mendaftarkan tools untuk afiliasi, promosi, email templates, audit logging, dan administrasi klien.
func RegisterMarketingTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_affiliates
	s.AddTool(mcp.NewTool("whmcs_get_affiliates",
		mcp.WithDescription("Get list of affiliates with referral and earnings statistics."),
		mcp.WithNumber("limitstart", mcp.Description("Starting offset for results")),
		mcp.WithNumber("limitnum", mcp.Description("Number of results to return")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetAffiliatesRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		var resp model.GetAffiliatesResponse
		if err := client.Execute(ctx, "GetAffiliates", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_activate_affiliate
	s.AddTool(mcp.NewTool("whmcs_activate_affiliate",
		mcp.WithDescription("Activate a client as an affiliate partner."),
		mcp.WithNumber("userid", mcp.Required(), mcp.Description("Client ID to activate")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.AffiliateActivateRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.UserID <= 0 {
			return mcp.NewToolResultError("userid is required"), nil
		}
		var resp model.AffiliateActivateResponse
		if err := client.Execute(ctx, "AffiliateActivate", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_get_promotions
	s.AddTool(mcp.NewTool("whmcs_get_promotions",
		mcp.WithDescription("Get list of promotions and discount coupon codes."),
		mcp.WithString("code", mcp.Description("Filter by specific promotion code")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetPromotionsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		var resp model.GetPromotionsResponse
		if err := client.Execute(ctx, "GetPromotions", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_log_activity
	s.AddTool(mcp.NewTool("whmcs_log_activity",
		mcp.WithDescription("Add a custom entry to the WHMCS system activity log."),
		mcp.WithString("description", mcp.Required(), mcp.Description("Activity description")),
		mcp.WithNumber("userid", mcp.Description("Associated user ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.LogActivityRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.Description == "" {
			return mcp.NewToolResultError("description is required"), nil
		}
		var resp model.CommonResponse
		if err := client.Execute(ctx, "LogActivity", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_get_email_templates
	s.AddTool(mcp.NewTool("whmcs_get_email_templates",
		mcp.WithDescription("Get list of email templates configured in WHMCS."),
		mcp.WithString("type", mcp.Description("Template type (general, product, domain, invoice, support, affiliate)")),
		mcp.WithString("language", mcp.Description("Template language")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetEmailTemplatesRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		var resp model.GetEmailTemplatesResponse
		if err := client.Execute(ctx, "GetEmailTemplates", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_delete_client
	s.AddTool(mcp.NewTool("whmcs_delete_client",
		mcp.WithDescription("Delete a client account permanently from WHMCS (use with extreme caution)."),
		mcp.WithNumber("clientid", mcp.Required(), mcp.Description("The client ID to delete")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.DeleteClientRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.ClientID <= 0 {
			return mcp.NewToolResultError("clientid is required"), nil
		}
		var resp model.CommonResponse
		if err := client.Execute(ctx, "DeleteClient", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
