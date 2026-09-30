package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

// Config menyimpan seluruh konfigurasi runtime WHMCS MCP Server.
type Config struct {
	URL          string
	Identifier   string
	Secret       string
	AccessKey    string
	HTTPUsername string
	HTTPPassword string
	Timeout      time.Duration
	Debug        bool
}

const (
	// DefaultTimeout adalah durasi batas waktu standar pemanggilan HTTP WHMCS.
	DefaultTimeout = 30 * time.Second
)

// Load memuat variabel lingkungan dan memvalidasi kelayakan parameter koneksi.
func Load() (*Config, error) {
	rawURL := strings.TrimSpace(os.Getenv("WHMCS_URL"))
	if rawURL == "" {
		rawURL = strings.TrimSpace(os.Getenv("WHMCS_API_URL"))
	}
	if rawURL == "" {
		return nil, errors.New("WHMCS_URL is required")
	}

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return nil, fmt.Errorf("WHMCS_URL must be a valid HTTP or HTTPS URL: %s", rawURL)
	}

	identifier := strings.TrimSpace(os.Getenv("WHMCS_API_IDENTIFIER"))
	if identifier == "" {
		return nil, errors.New("WHMCS_API_IDENTIFIER is required")
	}

	secret := strings.TrimSpace(os.Getenv("WHMCS_API_SECRET"))
	if secret == "" {
		return nil, errors.New("WHMCS_API_SECRET is required")
	}

	accessKey := strings.TrimSpace(os.Getenv("WHMCS_API_ACCESS_KEY"))
	if accessKey == "" {
		accessKey = strings.TrimSpace(os.Getenv("WHMCS_ACCESS_KEY"))
	}

	// Dukungan HTTP Basic Authentication (misal untuk proteksi Nginx htpasswd / Cloudflare)
	httpUser := strings.TrimSpace(os.Getenv("WHMCS_HTTP_USERNAME"))
	httpPass := strings.TrimSpace(os.Getenv("WHMCS_HTTP_PASSWORD"))

	// Jika tidak ditentukan di env terpisah, ekstrak kredensial dari URL jika ada (misal: https://user:pass@host/)
	if parsedURL.User != nil {
		if httpUser == "" {
			httpUser = parsedURL.User.Username()
		}
		if httpPass == "" {
			if p, ok := parsedURL.User.Password(); ok {
				httpPass = p
			}
		}
	}

	timeout := DefaultTimeout
	if timeoutStr := strings.TrimSpace(os.Getenv("WHMCS_TIMEOUT")); timeoutStr != "" {
		if d, err := time.ParseDuration(timeoutStr); err == nil && d > 0 {
			timeout = d
		}
	}

	debug := false
	if debugStr := strings.ToLower(strings.TrimSpace(os.Getenv("WHMCS_DEBUG"))); debugStr == "true" || debugStr == "1" {
		debug = true
	}

	return &Config{
		URL:          rawURL,
		Identifier:   identifier,
		Secret:       secret,
		AccessKey:    accessKey,
		HTTPUsername: httpUser,
		HTTPPassword: httpPass,
		Timeout:      timeout,
		Debug:        debug,
	}, nil
}
