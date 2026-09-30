package whmcs_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/vfat/whmcs-mcp/internal/config"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

type MockClientResponse struct {
	Result       string `json:"result"`
	Message      string `json:"message,omitempty"`
	TotalResults int    `json:"totalresults,omitempty"`
}

func TestClient_ExecuteSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/includes/api.php" {
			t.Errorf("expected /includes/api.php, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		values, err := url.ParseQuery(string(body))
		if err != nil {
			t.Fatalf("failed to parse body: %v", err)
		}

		if values.Get("action") != "GetClients" {
			t.Errorf("expected action GetClients, got %s", values.Get("action"))
		}
		if values.Get("identifier") != "test-id" {
			t.Errorf("expected identifier test-id, got %s", values.Get("identifier"))
		}
		if values.Get("secret") != "test-sec" {
			t.Errorf("expected secret test-sec, got %s", values.Get("secret"))
		}
		if values.Get("accesskey") != "test-acc" {
			t.Errorf("expected accesskey test-acc, got %s", values.Get("accesskey"))
		}
		if values.Get("responsetype") != "json" {
			t.Errorf("expected responsetype json, got %s", values.Get("responsetype"))
		}
		if values.Get("search") != "john" {
			t.Errorf("expected search john, got %s", values.Get("search"))
		}

		resp := MockClientResponse{
			Result:       "success",
			TotalResults: 5,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	cfg := &config.Config{
		URL:        server.URL,
		Identifier: "test-id",
		Secret:     "test-sec",
		AccessKey:  "test-acc",
		Timeout:    5 * time.Second,
	}

	client := whmcs.NewClient(cfg)

	var result MockClientResponse
	err := client.Execute(context.Background(), "GetClients", map[string]string{"search": "john"}, &result)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result.Result != "success" {
		t.Errorf("expected result success, got %s", result.Result)
	}
	if result.TotalResults != 5 {
		t.Errorf("expected TotalResults 5, got %d", result.TotalResults)
	}
}

func TestClient_ExecuteApiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"result":  "error",
			"message": "Invalid API Credentials",
		})
	}))
	defer server.Close()

	cfg := &config.Config{
		URL:        server.URL,
		Identifier: "wrong-id",
		Secret:     "wrong-sec",
		Timeout:    5 * time.Second,
	}

	client := whmcs.NewClient(cfg)

	var result MockClientResponse
	err := client.Execute(context.Background(), "GetClients", nil, &result)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expectedMsg := "whmcs api error: Invalid API Credentials"
	if err.Error() != expectedMsg {
		t.Errorf("expected error '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestClient_ExecuteHttpError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte("Bad Gateway"))
	}))
	defer server.Close()

	cfg := &config.Config{
		URL:        server.URL,
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}

	client := whmcs.NewClient(cfg)

	var result MockClientResponse
	err := client.Execute(context.Background(), "GetClients", nil, &result)
	if err == nil {
		t.Fatal("expected error for HTTP 502, got nil")
	}
}

func TestClient_ExecuteContextCanceled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &config.Config{
		URL:        server.URL,
		Identifier: "id",
		Secret:     "sec",
		Timeout:    5 * time.Second,
	}

	client := whmcs.NewClient(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	var result MockClientResponse
	err := client.Execute(ctx, "GetClients", nil, &result)
	if err == nil {
		t.Fatal("expected error for canceled context, got nil")
	}
}
