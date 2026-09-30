package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// RegisterProvisioningTools mendaftarkan 6 tool provisioning server dan modul WHMCS.
func RegisterProvisioningTools(s *server.MCPServer, client *whmcs.Client) {
	// 1. whmcs_get_servers
	s.AddTool(mcp.NewTool("whmcs_get_servers",
		mcp.WithDescription("Get list of configured servers in WHMCS with status and utilization details."),
		mcp.WithBoolean("fetchStatus", mcp.Description("Fetch live server status and account count")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.GetServersRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		var resp model.GetServersResponse
		if err := client.Execute(ctx, "GetServers", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 2. whmcs_module_create
	s.AddTool(mcp.NewTool("whmcs_module_create",
		mcp.WithDescription("Execute module create command to provision a service account on the remote server."),
		mcp.WithNumber("accountid", mcp.Required(), mcp.Description("The service ID to create/provision")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.ModuleCreateRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.AccountID <= 0 {
			return mcp.NewToolResultError("accountid is required"), nil
		}
		var resp model.ModuleResponse
		if err := client.Execute(ctx, "ModuleCreate", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 3. whmcs_module_suspend
	s.AddTool(mcp.NewTool("whmcs_module_suspend",
		mcp.WithDescription("Suspend a service account on the remote server."),
		mcp.WithNumber("accountid", mcp.Required(), mcp.Description("The service ID to suspend")),
		mcp.WithString("suspendreason", mcp.Description("Reason for suspension")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.ModuleSuspendRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.AccountID <= 0 {
			return mcp.NewToolResultError("accountid is required"), nil
		}
		var resp model.ModuleResponse
		if err := client.Execute(ctx, "ModuleSuspend", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 4. whmcs_module_unsuspend
	s.AddTool(mcp.NewTool("whmcs_module_unsuspend",
		mcp.WithDescription("Unsuspend a service account on the remote server."),
		mcp.WithNumber("accountid", mcp.Required(), mcp.Description("The service ID to unsuspend")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.ModuleUnsuspendRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.AccountID <= 0 {
			return mcp.NewToolResultError("accountid is required"), nil
		}
		var resp model.ModuleResponse
		if err := client.Execute(ctx, "ModuleUnsuspend", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 5. whmcs_module_terminate
	s.AddTool(mcp.NewTool("whmcs_module_terminate",
		mcp.WithDescription("Terminate and delete a service account permanently from the remote server."),
		mcp.WithNumber("accountid", mcp.Required(), mcp.Description("The service ID to terminate")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.ModuleTerminateRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.AccountID <= 0 {
			return mcp.NewToolResultError("accountid is required"), nil
		}
		var resp model.ModuleResponse
		if err := client.Execute(ctx, "ModuleTerminate", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})

	// 6. whmcs_module_change_password
	s.AddTool(mcp.NewTool("whmcs_module_change_password",
		mcp.WithDescription("Change the password for a service account on the remote server."),
		mcp.WithNumber("accountid", mcp.Required(), mcp.Description("The service ID")),
		mcp.WithString("servicepassword", mcp.Description("The new service password")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req model.ModuleChangePasswordRequest
		if err := parseArguments(request.Params.Arguments, &req); err != nil {
			return mcp.NewToolResultError("invalid arguments: " + err.Error()), nil
		}
		if req.AccountID <= 0 {
			return mcp.NewToolResultError("accountid is required"), nil
		}
		var resp model.ModuleResponse
		if err := client.Execute(ctx, "ModuleChangePassword", req, &resp); err != nil {
			return mcp.NewToolResultError("WHMCS API error: " + err.Error()), nil
		}
		return formatJSONResult(resp)
	})
}
