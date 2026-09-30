package mcp

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/vfat/whmcs-mcp/internal/config"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

// Server membungkus instance MCPServer dan dependensi HTTP Client WHMCS.
type Server struct {
	config    *config.Config
	client    *whmcs.Client
	mcpServer *server.MCPServer
}

// NewServer menginisialisasi Server MCP dengan kapabilitas tools, prompts, dan resources.
func NewServer(cfg *config.Config, client *whmcs.Client) *Server {
	mcpSrv := server.NewMCPServer(
		"whmcs-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithPromptCapabilities(true),
		server.WithResourceCapabilities(true, true),
	)

	return &Server{
		config:    cfg,
		client:    client,
		mcpServer: mcpSrv,
	}
}

// MCPServer mengembalikan pointer instance MCPServer internal.
func (s *Server) MCPServer() *server.MCPServer {
	return s.mcpServer
}

// Client mengembalikan WHMCS HTTP API client.
func (s *Server) Client() *whmcs.Client {
	return s.client
}

// ListTools mengembalikan daftar seluruh Tool yang terdaftar di dalam server.
func (s *Server) ListTools() []mcp.Tool {
	var tools []mcp.Tool
	for _, serverTool := range s.mcpServer.ListTools() {
		tools = append(tools, serverTool.Tool)
	}
	return tools
}

// ServeStdio menjalankan listen stream JSON-RPC 2.0 melalui standard I/O (os.Stdin / os.Stdout).
func (s *Server) ServeStdio() error {
	return server.ServeStdio(s.mcpServer)
}
