package version

// Version mendefinisikan versi rilis WHMCS MCP Server.
// Nilai ini dapat di-override saat kompilasi biner via flag ldflags:
// -ldflags="-X github.com/vfat/whmcs-mcp/internal/version.Version=v0.1.0"
var Version = "v0.1.0"
