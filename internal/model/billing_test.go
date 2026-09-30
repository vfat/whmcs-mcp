package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestBillingModels_Serialization(t *testing.T) {
	req := model.CreateInvoiceRequest{
		UserID:        42,
		Status:        "Unpaid",
		PaymentMethod: "paypal",
		TaxRate:       10.5,
		DueDate:       "2026-10-15",
		ItemDescription: []string{
			"Cloud VPS Hosting - 1 Mo",
			"Dedicated IPv4",
		},
		ItemAmount: []float64{
			15.00,
			3.00,
		},
		ItemTaxed: []bool{
			true,
			false,
		},
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize CreateInvoiceRequest: %v", err)
	}

	if values.Get("userid") != "42" {
		t.Errorf("expected userid=42, got %s", values.Get("userid"))
	}
	if values.Get("status") != "Unpaid" {
		t.Errorf("expected status=Unpaid, got %s", values.Get("status"))
	}
	if values.Get("itemdescription[0]") != "Cloud VPS Hosting - 1 Mo" {
		t.Errorf("expected line item description at 0, got %s", values.Get("itemdescription[0]"))
	}
	if values.Get("itemamount[1]") != "3" {
		t.Errorf("expected line item amount at 1, got %s", values.Get("itemamount[1]"))
	}

	jsonBlob := `{
		"result": "success",
		"totalresults": 1,
		"startnumber": 0,
		"numreturned": 1,
		"invoices": {
			"invoice": [
				{
					"id": 101,
					"userid": 42,
					"firstname": "Jane",
					"lastname": "Doe",
					"date": "2026-09-30",
					"duedate": "2026-10-15",
					"total": "18.00",
					"status": "Unpaid"
				}
			]
		}
	}`

	var resp model.GetInvoicesResponse
	if err := json.Unmarshal([]byte(jsonBlob), &resp); err != nil {
		t.Fatalf("failed to unmarshal GetInvoicesResponse: %v", err)
	}

	if resp.Result != "success" {
		t.Errorf("expected success result, got %s", resp.Result)
	}
	if resp.TotalResults != 1 {
		t.Errorf("expected totalresults=1, got %d", resp.TotalResults)
	}
	if len(resp.Invoices.Invoice) != 1 || resp.Invoices.Invoice[0].ID != 101 {
		t.Errorf("unexpected invoice list: %+v", resp.Invoices)
	}
}

func TestProductModels_Serialization(t *testing.T) {
	req := model.GetProductsRequest{
		PID:        10,
		GID:        2,
		ModuleName: "cpanel",
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize GetProductsRequest: %v", err)
	}

	if values.Get("pid") != "10" {
		t.Errorf("expected pid=10, got %s", values.Get("pid"))
	}
	if values.Get("gid") != "2" {
		t.Errorf("expected gid=2, got %s", values.Get("gid"))
	}
	if values.Get("module") != "cpanel" {
		t.Errorf("expected module=cpanel, got %s", values.Get("module"))
	}
}
