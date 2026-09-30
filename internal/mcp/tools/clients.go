package tools

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func parseArguments(args any, target any) error {
	if args == nil {
		return nil
	}
	b, err := json.Marshal(args)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, target)
}

func formatJSONResult(v interface{}) (*mcp.CallToolResult, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to format response: %v", err)), nil
	}
	return mcp.NewToolResultText(string(b)), nil
}

// RegisterClientTools mendaftarkan seluruh 8 tools Client Management ke server MCP.
func RegisterClientTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_clients
	s.AddTool(mcp.NewTool("whmcs_get_clients",
		mcp.WithDescription("Get a list of clients from WHMCS with optional filtering, search, and pagination."),
		mcp.WithString("search", mcp.Description("Search term for client name, email, or company name")),
		mcp.WithString("status", mcp.Description("Filter by client status (Active, Inactive, Closed)")),
		mcp.WithNumber("limitstart", mcp.Description("Offset for pagination (default: 0)")),
		mcp.WithNumber("limitnum", mcp.Description("Number of records to return (default: 25)")),
		mcp.WithString("sorting", mcp.Description("Sorting order (ASC or DESC)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetClientsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetClientsResponse
		if err := client.Execute(ctx, "GetClients", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_get_client_details
	s.AddTool(mcp.NewTool("whmcs_get_client_details",
		mcp.WithDescription("Get detailed profile information and account statistics for a specific client."),
		mcp.WithNumber("clientid", mcp.Description("The ID of the client to retrieve")),
		mcp.WithString("email", mcp.Description("The email address of the client to retrieve")),
		mcp.WithBoolean("stats", mcp.Description("Whether to include account statistics (invoices, products, tickets count)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetClientDetailsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.ClientID == 0 && req.Email == "" {
			return mcp.NewToolResultError("either clientid or email must be provided"), nil
		}
		var resp model.GetClientDetailsResponse
		if err := client.Execute(ctx, "GetClientsDetails", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_add_client
	s.AddTool(mcp.NewTool("whmcs_add_client",
		mcp.WithDescription("Create a new client account in WHMCS."),
		mcp.WithString("firstname", mcp.Required(), mcp.Description("Client first name")),
		mcp.WithString("lastname", mcp.Required(), mcp.Description("Client last name")),
		mcp.WithString("email", mcp.Required(), mcp.Description("Client email address")),
		mcp.WithString("address1", mcp.Required(), mcp.Description("Street address line 1")),
		mcp.WithString("city", mcp.Required(), mcp.Description("City")),
		mcp.WithString("state", mcp.Required(), mcp.Description("State or region")),
		mcp.WithString("postcode", mcp.Required(), mcp.Description("Postal / ZIP code")),
		mcp.WithString("country", mcp.Required(), mcp.Description("Two-letter country code (ISO 3166-1 alpha-2, e.g. US, ID)")),
		mcp.WithString("companyname", mcp.Description("Company name")),
		mcp.WithString("phonenumber", mcp.Description("Phone number")),
		mcp.WithString("password2", mcp.Description("Initial account password")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.AddClientRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.AddClientResponse
		if err := client.Execute(ctx, "AddClient", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_update_client
	s.AddTool(mcp.NewTool("whmcs_update_client",
		mcp.WithDescription("Update an existing client profile in WHMCS."),
		mcp.WithNumber("clientid", mcp.Required(), mcp.Description("The ID of the client to update")),
		mcp.WithString("firstname", mcp.Description("Client first name")),
		mcp.WithString("lastname", mcp.Description("Client last name")),
		mcp.WithString("companyname", mcp.Description("Company name")),
		mcp.WithString("email", mcp.Description("Client email address")),
		mcp.WithString("address1", mcp.Description("Street address")),
		mcp.WithString("city", mcp.Description("City")),
		mcp.WithString("state", mcp.Description("State or region")),
		mcp.WithString("postcode", mcp.Description("Postal / ZIP code")),
		mcp.WithString("country", mcp.Description("Two-letter country code")),
		mcp.WithString("phonenumber", mcp.Description("Phone number")),
		mcp.WithString("status", mcp.Description("Status (Active, Inactive, Closed)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.UpdateClientRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.UpdateClientResponse
		if err := client.Execute(ctx, "UpdateClient", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_close_client (Destructive)
	s.AddTool(mcp.NewTool("whmcs_close_client",
		mcp.WithDescription("Close a client account in WHMCS (sets status to Closed)."),
		mcp.WithNumber("clientid", mcp.Required(), mcp.Description("The ID of the client to close")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.CloseClientRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.CloseClientResponse
		if err := client.Execute(ctx, "CloseClient", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_get_client_products
	s.AddTool(mcp.NewTool("whmcs_get_client_products",
		mcp.WithDescription("Get hosting services and products owned by a client."),
		mcp.WithNumber("clientid", mcp.Description("The client ID to filter products for")),
		mcp.WithNumber("serviceid", mcp.Description("Specific service ID to retrieve")),
		mcp.WithNumber("pid", mcp.Description("Product ID to filter")),
		mcp.WithString("domain", mcp.Description("Domain name to filter")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetClientProductsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetClientProductsResponse
		if err := client.Execute(ctx, "GetClientsProducts", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 7. whmcs_get_client_domains
	s.AddTool(mcp.NewTool("whmcs_get_client_domains",
		mcp.WithDescription("Get domains registered by a client in WHMCS."),
		mcp.WithNumber("clientid", mcp.Description("The client ID to filter domains for")),
		mcp.WithNumber("domainid", mcp.Description("Specific domain ID to retrieve")),
		mcp.WithString("domain", mcp.Description("Domain name to search/filter")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetClientDomainsRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		var resp model.GetClientDomainsResponse
		if err := client.Execute(ctx, "GetClientsDomains", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 8. whmcs_get_client_invoices
	s.AddTool(mcp.NewTool("whmcs_get_client_invoices",
		mcp.WithDescription("Get invoice history for a specific client."),
		mcp.WithNumber("clientid", mcp.Required(), mcp.Description("The client ID to retrieve invoices for")),
		mcp.WithString("status", mcp.Description("Filter by invoice status (Paid, Unpaid, Cancelled, Overdue)")),
		mcp.WithNumber("limitnum", mcp.Description("Number of invoices to return (default: 25)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetClientInvoicesRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid arguments: %v", err)), nil
		}
		if req.ClientID == 0 {
			return mcp.NewToolResultError("clientid is required"), nil
		}

		var resp map[string]interface{}
		if err := client.Execute(ctx, "GetInvoices", req, &resp); err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
