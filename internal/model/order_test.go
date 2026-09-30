package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestOrderModels_Serialization(t *testing.T) {
	req := model.AcceptOrderRequest{
		OrderID:   77,
		AutoSetup: true,
		SendEmail: true,
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize AcceptOrderRequest: %v", err)
	}

	if values.Get("orderid") != "77" {
		t.Errorf("expected orderid=77, got %s", values.Get("orderid"))
	}
	if values.Get("autosetup") != "true" {
		t.Errorf("expected autosetup=true, got %s", values.Get("autosetup"))
	}

	jsonBlob := `{
		"result": "success",
		"totalresults": 1,
		"orders": {
			"order": [
				{
					"id": 77,
					"ordernum": "9876543210",
					"userid": 42,
					"date": "2026-09-30",
					"status": "Pending"
				}
			]
		}
	}`

	var resp model.GetOrdersResponse
	if err := json.Unmarshal([]byte(jsonBlob), &resp); err != nil {
		t.Fatalf("failed to unmarshal GetOrdersResponse: %v", err)
	}

	if resp.Result != "success" || len(resp.Orders.Order) != 1 {
		t.Errorf("unexpected orders response: %+v", resp)
	}
}

func TestQuoteModels_Serialization(t *testing.T) {
	req := model.CreateQuoteRequest{
		Subject: "Dedicated Server Setup Quote",
		Stage:   "Delivered",
		UserID:  42,
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize: %v", err)
	}

	if values.Get("subject") != "Dedicated Server Setup Quote" {
		t.Errorf("expected subject, got %s", values.Get("subject"))
	}
	if values.Get("stage") != "Delivered" {
		t.Errorf("expected stage, got %s", values.Get("stage"))
	}
}
