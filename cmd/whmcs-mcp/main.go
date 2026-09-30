package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/vfat/whmcs-mcp/internal/config"
	mcpServer "github.com/vfat/whmcs-mcp/internal/mcp"
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

	slog.Info("initializing WHMCS MCP Server", "url", cfg.URL, "timeout", cfg.Timeout)

	client := whmcs.NewClient(cfg)
	srv := mcpServer.NewServer(cfg, client)

	slog.Info("starting stdio server loop...")
	if err := srv.ServeStdio(); err != nil {
		slog.Error("stdio server terminated with error", "error", err)
		os.Exit(1)
	}
}
