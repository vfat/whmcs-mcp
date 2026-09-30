package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterDomainTools mendaftarkan 9 tool manajemen domain ke MCP server.
func RegisterDomainTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_register_domain
	s.AddTool(mcp.NewTool("whmcs_register_domain",
		mcp.WithDescription("Send domain registration command to registrar module."),
		mcp.WithNumber("domainid", mcp.Description("Domain ID in WHMCS")),
		mcp.WithString("domain", mcp.Description("Domain name to register")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.RegisterDomainRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.RegisterDomainResponse
		if err := client.Execute(ctx, "DomainRegister", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_transfer_domain
	s.AddTool(mcp.NewTool("whmcs_transfer_domain",
		mcp.WithDescription("Send domain transfer command to registrar module."),
		mcp.WithNumber("domainid", mcp.Required(), mcp.Description("Domain ID in WHMCS")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.TransferDomainRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DomainID == 0 {
			return mcp.NewToolResultError("domainid is required"), nil
		}
		var resp model.TransferDomainResponse
		if err := client.Execute(ctx, "DomainTransfer", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_renew_domain
	s.AddTool(mcp.NewTool("whmcs_renew_domain",
		mcp.WithDescription("Send domain renewal command to registrar module."),
		mcp.WithNumber("domainid", mcp.Required(), mcp.Description("Domain ID in WHMCS")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.RenewDomainRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DomainID == 0 {
			return mcp.NewToolResultError("domainid is required"), nil
		}
		var resp model.RenewDomainResponse
		if err := client.Execute(ctx, "DomainRenew", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_get_domain_whois
	s.AddTool(mcp.NewTool("whmcs_get_domain_whois",
		mcp.WithDescription("Get WHOIS lookup information for a registered domain."),
		mcp.WithNumber("domainid", mcp.Required(), mcp.Description("Domain ID in WHMCS")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetDomainWhoisRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DomainID == 0 {
			return mcp.NewToolResultError("domainid is required"), nil
		}
		var resp model.GetDomainWhoisResponse
		if err := client.Execute(ctx, "DomainWhois", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_get_domain_nameservers
	s.AddTool(mcp.NewTool("whmcs_get_domain_nameservers",
		mcp.WithDescription("Get current authoritative nameservers for a domain from the registrar."),
		mcp.WithNumber("domainid", mcp.Required(), mcp.Description("Domain ID in WHMCS")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetDomainNameserversRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DomainID == 0 {
			return mcp.NewToolResultError("domainid is required"), nil
		}
		var resp model.GetDomainNameserversResponse
		if err := client.Execute(ctx, "DomainGetNameservers", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_update_domain_nameservers
	s.AddTool(mcp.NewTool("whmcs_update_domain_nameservers",
		mcp.WithDescription("Update authoritative nameservers for a domain at the registrar."),
		mcp.WithNumber("domainid", mcp.Required(), mcp.Description("Domain ID in WHMCS")),
		mcp.WithString("ns1", mcp.Description("Primary nameserver")),
		mcp.WithString("ns2", mcp.Description("Secondary nameserver")),
		mcp.WithString("ns3", mcp.Description("Tertiary nameserver")),
		mcp.WithString("ns4", mcp.Description("Fourth nameserver")),
		mcp.WithString("ns5", mcp.Description("Fifth nameserver")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.UpdateDomainNameserversRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DomainID == 0 {
			return mcp.NewToolResultError("domainid is required"), nil
		}
		var resp model.UpdateDomainNameserversResponse
		if err := client.Execute(ctx, "DomainUpdateNameservers", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 7. whmcs_get_domain_lock_status
	s.AddTool(mcp.NewTool("whmcs_get_domain_lock_status",
		mcp.WithDescription("Get domain registrar transfer lock status (locked or unlocked)."),
		mcp.WithNumber("domainid", mcp.Required(), mcp.Description("Domain ID in WHMCS")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetDomainLockStatusRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DomainID == 0 {
			return mcp.NewToolResultError("domainid is required"), nil
		}
		var resp model.GetDomainLockStatusResponse
		if err := client.Execute(ctx, "DomainGetLockingStatus", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 8. whmcs_update_domain_lock_status
	s.AddTool(mcp.NewTool("whmcs_update_domain_lock_status",
		mcp.WithDescription("Update domain registrar transfer lock status (lock or unlock)."),
		mcp.WithNumber("domainid", mcp.Required(), mcp.Description("Domain ID in WHMCS")),
		mcp.WithBoolean("lockstatus", mcp.Description("True to lock, false to unlock")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.UpdateDomainLockStatusRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.DomainID == 0 {
			return mcp.NewToolResultError("domainid is required"), nil
		}
		var resp model.UpdateDomainLockStatusResponse
		if err := client.Execute(ctx, "DomainUpdateLockingStatus", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 9. whmcs_get_tld_pricing
	s.AddTool(mcp.NewTool("whmcs_get_tld_pricing",
		mcp.WithDescription("Get domain TLD pricing information across all extensions."),
		mcp.WithNumber("currencyid", mcp.Description("Currency ID for pricing display")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetTLDPricingRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetTLDPricingResponse
		if err := client.Execute(ctx, "GetTLDPricing", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
