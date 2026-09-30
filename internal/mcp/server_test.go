package mcp_test

import (
	"context"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
	"github.com/vfat/whmcs-mcp/internal/mcp/prompts"
	"github.com/vfat/whmcs-mcp/internal/mcp/resources"
	"github.com/vfat/whmcs-mcp/internal/mcp/tools"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestNewServer(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "secret",
		Timeout:    10 * time.Second,
	}
	client := whmcs.NewClient(cfg)

	srv := mcpServer.NewServer(cfg, client)
	if srv == nil {
		t.Fatal("expected non-nil server")
	}

	if srv.MCPServer() == nil {
		t.Fatal("expected internal MCPServer to be initialized")
	}
}

func TestRegisterTool_Execution(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "secret",
		Timeout:    10 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	tool := mcp.NewTool("test_ping",
		mcp.WithDescription("Ping test tool"),
		mcp.WithString("message", mcp.Required(), mcp.Description("Message to echo")),
	)

	srv.MCPServer().AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("arguments must be a map"), nil
		}
		msg, ok := args["message"].(string)
		if !ok {
			return mcp.NewToolResultError("message argument required"), nil
		}
		return mcp.NewToolResultText("pong: " + msg), nil
	})

	toolsList := srv.ListTools()
	found := false
	for _, tl := range toolsList {
		if tl.Name == "test_ping" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected test_ping tool to be registered in server")
	}
}

func TestFullServerParity(t *testing.T) {
	cfg := &config.Config{
		URL:        "https://billing.example.com",
		Identifier: "id",
		Secret:     "secret",
		Timeout:    10 * time.Second,
	}
	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	// Registrasi seluruh Batch komponen
	tools.RegisterClientTools(srv.MCPServer(), client)
	tools.RegisterSystemTools(srv.MCPServer(), client)
	resources.RegisterBatch1Resources(srv.MCPServer(), client)
	prompts.RegisterBatch1Prompts(srv.MCPServer())

	tools.RegisterProductTools(srv.MCPServer(), client)
	tools.RegisterBillingTools(srv.MCPServer(), client)
	tools.RegisterTicketTools(srv.MCPServer(), client)
	resources.RegisterBatch2Resources(srv.MCPServer(), client)
	prompts.RegisterBatch2Prompts(srv.MCPServer())

	tools.RegisterDomainTools(srv.MCPServer(), client)
	tools.RegisterOrderTools(srv.MCPServer(), client)
	resources.RegisterBatch3Resources(srv.MCPServer(), client)
	prompts.RegisterBatch3Prompts(srv.MCPServer())

	tools.RegisterProvisioningTools(srv.MCPServer(), client)
	tools.RegisterMarketingTools(srv.MCPServer(), client)
	resources.RegisterBatch4Resources(srv.MCPServer(), client)
	prompts.RegisterBatch4Prompts(srv.MCPServer())

	// 1. Verifikasi Seluruh 62 Tool Paritas
	expectedTools := []string{
		"whmcs_get_clients",
		"whmcs_get_client_details",
		"whmcs_add_client",
		"whmcs_update_client",
		"whmcs_delete_client",
		"whmcs_get_client_products",
		"whmcs_get_client_domains",
		"whmcs_get_products",
		"whmcs_get_product_groups",
		"whmcs_get_invoices",
		"whmcs_get_invoice",
		"whmcs_create_invoice",
		"whmcs_update_invoice",
		"whmcs_add_payment",
		"whmcs_apply_credit",
		"whmcs_get_transactions",
		"whmcs_get_tickets",
		"whmcs_get_ticket",
		"whmcs_open_ticket",
		"whmcs_add_ticket_reply",
		"whmcs_add_ticket_note",
		"whmcs_update_ticket",
		"whmcs_delete_ticket",
		"whmcs_get_support_departments",
		"whmcs_get_support_statuses",
		"whmcs_register_domain",
		"whmcs_transfer_domain",
		"whmcs_renew_domain",
		"whmcs_get_domain_whois",
		"whmcs_get_domain_nameservers",
		"whmcs_update_domain_nameservers",
		"whmcs_get_domain_lock_status",
		"whmcs_update_domain_lock_status",
		"whmcs_get_tld_pricing",
		"whmcs_get_orders",
		"whmcs_accept_order",
		"whmcs_cancel_order",
		"whmcs_delete_order",
		"whmcs_fraud_order",
		"whmcs_pending_order",
		"whmcs_get_servers",
		"whmcs_module_create",
		"whmcs_module_suspend",
		"whmcs_module_unsuspend",
		"whmcs_module_terminate",
		"whmcs_module_change_password",
		"whmcs_get_stats",
		"whmcs_get_admin_users",
		"whmcs_get_payment_methods",
		"whmcs_get_currencies",
		"whmcs_get_activity_log",
		"whmcs_log_activity",
		"whmcs_get_email_templates",
		"whmcs_send_email",
		"whmcs_get_todo_items",
		"whmcs_get_affiliates",
		"whmcs_activate_affiliate",
		"whmcs_get_promotions",
		"whmcs_get_quotes",
		"whmcs_create_quote",
		"whmcs_accept_quote",
		"whmcs_delete_quote",
	}

	registeredTools := make(map[string]bool)
	for _, tl := range srv.ListTools() {
		registeredTools[tl.Name] = true
	}

	for _, toolName := range expectedTools {
		if !registeredTools[toolName] {
			t.Errorf("parity error: tool %s is missing from registered tools", toolName)
		}
	}

	// 2. Verifikasi Seluruh 11 Resources Paritas
	expectedResources := []string{
		"whmcs://stats",
		"whmcs://admin-users",
		"whmcs://currencies",
		"whmcs://payment-methods",
		"whmcs://admin/todo",
		"whmcs://products",
		"whmcs://support/departments",
		"whmcs://support/statuses",
		"whmcs://tld-pricing",
		"whmcs://servers",
		"whmcs://promotions",
	}

	registeredResources := srv.MCPServer().ListResources()
	for _, uri := range expectedResources {
		if _, exists := registeredResources[uri]; !exists {
			t.Errorf("parity error: resource %s is missing from registered resources", uri)
		}
	}

	// 3. Verifikasi Seluruh 8 Prompts Paritas
	expectedPrompts := []string{
		"client-onboarding",
		"client-health-check",
		"ticket-response",
		"revenue-report",
		"bulk-invoice-reminder",
		"domain-expiry-audit",
		"fraud-investigation",
		"new-product-setup",
	}

	registeredPrompts := srv.MCPServer().ListPrompts()
	for _, promptName := range expectedPrompts {
		if _, exists := registeredPrompts[promptName]; !exists {
			t.Errorf("parity error: prompt %s is missing from registered prompts", promptName)
		}
	}
}
