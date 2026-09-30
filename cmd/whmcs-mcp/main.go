package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
	"github.com/vfat/whmcs-mcp/internal/mcp/prompts"
	"github.com/vfat/whmcs-mcp/internal/mcp/resources"
	"github.com/vfat/whmcs-mcp/internal/mcp/tools"
	"github.com/vfat/whmcs-mcp/internal/version"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func main() {
	// Logger WAJIB menggunakan os.Stderr untuk menjaga os.Stdout murni protokol JSON-RPC 2.0 (L-001).
	logLevel := slog.LevelInfo
	if os.Getenv("WHMCS_DEBUG") == "true" || os.Getenv("WHMCS_DEBUG") == "1" {
		logLevel = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load configuration", "error", err)
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	slog.Info("initializing WHMCS MCP Server", "version", version.Version, "url", cfg.URL, "timeout", cfg.Timeout)

	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	// Registrasi komponen Batch 1
	tools.RegisterClientTools(srv.MCPServer(), client)
	tools.RegisterSystemTools(srv.MCPServer(), client)
	resources.RegisterBatch1Resources(srv.MCPServer(), client)
	prompts.RegisterBatch1Prompts(srv.MCPServer())

	// Registrasi komponen Batch 2
	tools.RegisterProductTools(srv.MCPServer(), client)
	tools.RegisterBillingTools(srv.MCPServer(), client)
	tools.RegisterTicketTools(srv.MCPServer(), client)
	resources.RegisterBatch2Resources(srv.MCPServer(), client)
	prompts.RegisterBatch2Prompts(srv.MCPServer())

	// Registrasi komponen Batch 3
	tools.RegisterDomainTools(srv.MCPServer(), client)
	tools.RegisterOrderTools(srv.MCPServer(), client)
	resources.RegisterBatch3Resources(srv.MCPServer(), client)
	prompts.RegisterBatch3Prompts(srv.MCPServer())

	// Registrasi komponen Batch 4 (Final Batch)
	tools.RegisterProvisioningTools(srv.MCPServer(), client)
	tools.RegisterMarketingTools(srv.MCPServer(), client)
	resources.RegisterBatch4Resources(srv.MCPServer(), client)
	prompts.RegisterBatch4Prompts(srv.MCPServer())

	slog.Info("registered MCP components",
		"tools", len(srv.MCPServer().ListTools()),
		"resources", len(srv.MCPServer().ListResources()),
		"prompts", len(srv.MCPServer().ListPrompts()),
	)

	slog.Info("starting stdio server loop...")
	if err := srv.ServeStdio(); err != nil {
		slog.Error("stdio server terminated with error", "error", err)
		os.Exit(1)
	}
}
