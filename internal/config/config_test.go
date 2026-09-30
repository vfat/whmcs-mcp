package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/vfat/whmcs-mcp/internal/config"
)

func clearEnv() {
	os.Unsetenv("WHMCS_URL")
	os.Unsetenv("WHMCS_API_IDENTIFIER")
	os.Unsetenv("WHMCS_API_SECRET")
	os.Unsetenv("WHMCS_API_ACCESS_KEY")
	os.Unsetenv("WHMCS_HTTP_USERNAME")
	os.Unsetenv("WHMCS_HTTP_PASSWORD")
	os.Unsetenv("WHMCS_TIMEOUT")
	os.Unsetenv("WHMCS_DEBUG")
}

func TestLoad_Success(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_URL", "https://billing.example.com")
	os.Setenv("WHMCS_API_IDENTIFIER", "test-identifier")
	os.Setenv("WHMCS_API_SECRET", "test-secret")
	os.Setenv("WHMCS_API_ACCESS_KEY", "test-access-key")
	os.Setenv("WHMCS_TIMEOUT", "15s")
	os.Setenv("WHMCS_DEBUG", "true")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if cfg.URL != "https://billing.example.com" {
		t.Errorf("expected URL https://billing.example.com, got %s", cfg.URL)
	}
	if cfg.Identifier != "test-identifier" {
		t.Errorf("expected Identifier test-identifier, got %s", cfg.Identifier)
	}
	if cfg.Secret != "test-secret" {
		t.Errorf("expected Secret test-secret, got %s", cfg.Secret)
	}
	if cfg.AccessKey != "test-access-key" {
		t.Errorf("expected AccessKey test-access-key, got %s", cfg.AccessKey)
	}
	if cfg.Timeout != 15*time.Second {
		t.Errorf("expected Timeout 15s, got %v", cfg.Timeout)
	}
	if !cfg.Debug {
		t.Errorf("expected Debug true, got false")
	}
}

func TestLoad_DefaultTimeout(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_URL", "https://billing.example.com")
	os.Setenv("WHMCS_API_IDENTIFIER", "test-identifier")
	os.Setenv("WHMCS_API_SECRET", "test-secret")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if cfg.Timeout != 30*time.Second {
		t.Errorf("expected default timeout 30s, got %v", cfg.Timeout)
	}
}

func TestLoad_MissingURL(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_API_IDENTIFIER", "test-identifier")
	os.Setenv("WHMCS_API_SECRET", "test-secret")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for missing WHMCS_URL, got nil")
	}
}

func TestLoad_InvalidURL(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_URL", "not-a-valid-url")
	os.Setenv("WHMCS_API_IDENTIFIER", "test-identifier")
	os.Setenv("WHMCS_API_SECRET", "test-secret")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for invalid WHMCS_URL scheme, got nil")
	}
}

func TestLoad_MissingIdentifier(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_URL", "https://billing.example.com")
	os.Setenv("WHMCS_API_SECRET", "test-secret")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for missing WHMCS_API_IDENTIFIER, got nil")
	}
}

func TestLoad_MissingSecret(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_URL", "https://billing.example.com")
	os.Setenv("WHMCS_API_IDENTIFIER", "test-identifier")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error for missing WHMCS_API_SECRET, got nil")
	}
}

func TestLoad_HTTPBasicAuth(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_URL", "https://billing.example.com")
	os.Setenv("WHMCS_API_IDENTIFIER", "test-identifier")
	os.Setenv("WHMCS_API_SECRET", "test-secret")
	os.Setenv("WHMCS_HTTP_USERNAME", "admin")
	os.Setenv("WHMCS_HTTP_PASSWORD", "secretpass")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if cfg.HTTPUsername != "admin" {
		t.Errorf("expected HTTPUsername admin, got %s", cfg.HTTPUsername)
	}
	if cfg.HTTPPassword != "secretpass" {
		t.Errorf("expected HTTPPassword secretpass, got %s", cfg.HTTPPassword)
	}
}

func TestLoad_HTTPBasicAuthInURL(t *testing.T) {
	clearEnv()
	defer clearEnv()

	os.Setenv("WHMCS_URL", "https://admin:Makanan%402@home-n-stuff.net/myqueen_7/")
	os.Setenv("WHMCS_API_IDENTIFIER", "test-identifier")
	os.Setenv("WHMCS_API_SECRET", "test-secret")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if cfg.HTTPUsername != "admin" {
		t.Errorf("expected HTTPUsername admin, got %s", cfg.HTTPUsername)
	}
	if cfg.HTTPPassword != "Makanan@2" {
		t.Errorf("expected HTTPPassword Makanan@2, got %s", cfg.HTTPPassword)
	}
}

