package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
)

func TestSystemModels_GetStatsResponse(t *testing.T) {
	jsonPayload := `{
		"result": "success",
		"income_today": "150.00",
		"income_thismonth": "4500.00",
		"income_thisyear": "54000.00",
		"orders_pending": 3,
		"tickets_open": 5,
		"services_active": 42
	}`

	var resp model.GetStatsResponse
	if err := json.Unmarshal([]byte(jsonPayload), &resp); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if resp.Result != "success" {
		t.Errorf("expected result success, got %s", resp.Result)
	}
	if resp.IncomeToday != "150.00" {
		t.Errorf("expected income_today 150.00, got %s", resp.IncomeToday)
	}
	if resp.OrdersPending != 3 {
		t.Errorf("expected orders_pending 3, got %d", resp.OrdersPending)
	}
}

func TestSystemModels_SendEmailSerialization(t *testing.T) {
	req := model.SendEmailRequest{
		CustomType:    "general",
		CustomSubject: "Pemberitahuan Sistem",
		CustomMessage: "<p>Halo, ini adalah notifikasi pemeliharaan server.</p>",
		ID:            101,
	}

	if req.CustomSubject != "Pemberitahuan Sistem" {
		t.Errorf("expected subject matching, got %s", req.CustomSubject)
	}
}
