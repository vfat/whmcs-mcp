package whmcs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/vfat/whmcs-mcp/internal/config"
)

// Client adalah HTTP API client untuk memanggil endpoint WHMCS (includes/api.php).
type Client struct {
	config     *config.Config
	httpClient *http.Client
	endpoint   string
}

// NewClient menginisialisasi Client baru berdasarkan konfigurasi yang diberikan.
func NewClient(cfg *config.Config) *Client {
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     90 * time.Second,
	}

	endpoint := strings.TrimRight(cfg.URL, "/") + "/includes/api.php"

	return &Client{
		config: cfg,
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   timeout,
		},
		endpoint: endpoint,
	}
}

// ApiBaseResponse adalah representasi respon dasar WHMCS untuk pemeriksaan status keberhasilan.
type ApiBaseResponse struct {
	Result  string `json:"result"`
	Message string `json:"message,omitempty"`
}

// Execute mengirimkan request HTTP POST terautentikasi ke WHMCS dan meng-unmarshal hasil JSON.
func (c *Client) Execute(ctx context.Context, action string, params interface{}, result interface{}) error {
	// Serialize parameter opsional
	values, err := SerializeParams(params)
	if err != nil {
		return fmt.Errorf("failed to serialize params: %w", err)
	}

	// Inject parameter wajib & otentikasi
	values.Set("action", action)
	values.Set("responsetype", "json")
	if c.config.Identifier != "" {
		values.Set("identifier", c.config.Identifier)
	}
	if c.config.Secret != "" {
		values.Set("secret", c.config.Secret)
	}
	if c.config.AccessKey != "" {
		values.Set("accesskey", c.config.AccessKey)
	}

	bodyReader := strings.NewReader(values.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	if c.config.HTTPUsername != "" || c.config.HTTPPassword != "" {
		req.SetBasicAuth(c.config.HTTPUsername, c.config.HTTPPassword)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("whmcs http error: status code %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Cek respon dasar WHMCS untuk result: error
	var baseResp ApiBaseResponse
	if err := json.Unmarshal(bodyBytes, &baseResp); err == nil {
		if strings.ToLower(baseResp.Result) == "error" {
			msg := baseResp.Message
			if msg == "" {
				msg = "unknown error"
			}
			return fmt.Errorf("whmcs api error: %s", msg)
		}
	}

	if result != nil {
		if err := json.Unmarshal(bodyBytes, result); err != nil {
			return fmt.Errorf("failed to unmarshal json response: %w", err)
		}
	}

	return nil
}
