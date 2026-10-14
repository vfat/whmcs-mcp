package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestClientModels_GetClientsSerialization(t *testing.T) {
	req := model.GetClientsRequest{
		Search:     "john",
		Status:     "Active",
		LimitStart: 0,
		LimitNum:   10,
		Sorting:    "ASC",
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize params: %v", err)
	}

	if values.Get("search") != "john" {
		t.Errorf("expected search john, got %s", values.Get("search"))
	}
	if values.Get("status") != "Active" {
		t.Errorf("expected status Active, got %s", values.Get("status"))
	}
	if values.Get("limitnum") != "10" {
		t.Errorf("expected limitnum 10, got %s", values.Get("limitnum"))
	}

	// Test Response deserialization
	jsonPayload := `{
		"result": "success",
		"totalresults": 1,
		"startnumber": 0,
		"numreturned": 1,
		"clients": {
			"client": [
				{
					"id": 101,
					"firstname": "John",
					"lastname": "Doe",
					"companyname": "Example Corp",
					"email": "john@example.com",
					"status": "Active"
				}
			]
		}
	}`

	var resp model.GetClientsResponse
	if err := json.Unmarshal([]byte(jsonPayload), &resp); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if resp.Result != "success" {
		t.Errorf("expected result success, got %s", resp.Result)
	}
	if resp.TotalResults != 1 {
		t.Errorf("expected totalresults 1, got %d", resp.TotalResults)
	}
	if len(resp.Clients.Client) != 1 {
		t.Fatalf("expected 1 client, got %d", len(resp.Clients.Client))
	}
	if resp.Clients.Client[0].Email != "john@example.com" {
		t.Errorf("expected email john@example.com, got %s", resp.Clients.Client[0].Email)
	}
}

func TestClientModels_AddClientSerialization(t *testing.T) {
	req := model.AddClientRequest{
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane@example.com",
		Password2: "SecretPassword123!",
		Address1:  "123 Street",
		City:      "Jakarta",
		State:     "DKI Jakarta",
		Postcode:  "12345",
		Country:   "ID",
		CustomFields: []string{
			"Field1",
			"Field2",
		},
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize params: %v", err)
	}

	if values.Get("firstname") != "Jane" {
		t.Errorf("expected firstname Jane, got %s", values.Get("firstname"))
	}
	if values.Get("customfields[0]") != "Field1" {
		t.Errorf("expected customfields[0] Field1, got %s", values.Get("customfields[0]"))
	}

	jsonPayload := `{"result":"success","clientid":102}`
	var resp model.AddClientResponse
	if err := json.Unmarshal([]byte(jsonPayload), &resp); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if resp.ClientID != 102 {
		t.Errorf("expected clientid 102, got %d", resp.ClientID)
	}
}

func TestExtractVMID(t *testing.T) {
	tests := []struct {
		name     string
		cf       model.RawCustomFields
		expected string
	}{
		{
			name: "vmid found lowercase",
			cf: model.RawCustomFields{
				CustomField: []model.CustomFieldItem{
					{ID: 155, Name: "vmid", Value: "101"},
					{ID: 156, Name: "node", Value: "pve1"},
				},
			},
			expected: "101",
		},
		{
			name: "vmid found uppercase case-insensitive",
			cf: model.RawCustomFields{
				CustomField: []model.CustomFieldItem{
					{ID: 1, Name: "VMID", Value: "205"},
				},
			},
			expected: "205",
		},
		{
			name: "vmid with whitespace trimmed",
			cf: model.RawCustomFields{
				CustomField: []model.CustomFieldItem{
					{ID: 1, Name: "vmid", Value: "  302  "},
				},
			},
			expected: "302",
		},
		{
			name: "vmid empty value defaults to 000",
			cf: model.RawCustomFields{
				CustomField: []model.CustomFieldItem{
					{ID: 1, Name: "vmid", Value: ""},
					{ID: 2, Name: "node", Value: "node1"},
				},
			},
			expected: "000",
		},
		{
			name: "vmid only whitespace defaults to 000",
			cf: model.RawCustomFields{
				CustomField: []model.CustomFieldItem{
					{ID: 1, Name: "vmid", Value: "   "},
				},
			},
			expected: "000",
		},
		{
			name: "vmid missing defaults to 000",
			cf: model.RawCustomFields{
				CustomField: []model.CustomFieldItem{
					{ID: 1, Name: "other_field", Value: "val"},
				},
			},
			expected: "000",
		},
		{
			name:     "empty customfields defaults to 000",
			cf:       model.RawCustomFields{},
			expected: "000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := model.ExtractVMID(tt.cf)
			if got != tt.expected {
				t.Errorf("ExtractVMID() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestGetClientProductsResponse_Deserialization(t *testing.T) {
	// 1. JSON dengan customfields berisi vmid (array)
	jsonWithVMID := `{
		"result": "success",
		"totalresults": 1,
		"startnumber": 0,
		"numreturned": 1,
		"products": {
			"product": [
				{
					"id": 77273,
					"clientid": 109524,
					"orderid": 554,
					"pid": 12,
					"name": "Classic Standard",
					"status": "Active",
					"serverip": "157.66.55.179",
					"serverhostname": "vps.example.com",
					"customfields": {
						"customfield": [
							{"id": 155, "name": "vmid", "value": "101"},
							{"id": 156, "name": "node", "value": "pve1"}
						]
					}
				}
			]
		}
	}`

	var resp1 model.GetClientProductsResponse
	if err := json.Unmarshal([]byte(jsonWithVMID), &resp1); err != nil {
		t.Fatalf("failed to unmarshal jsonWithVMID: %v", err)
	}
	if len(resp1.Products.Product) != 1 {
		t.Fatalf("expected 1 product, got %d", len(resp1.Products.Product))
	}
	prod1 := resp1.Products.Product[0]
	if model.ExtractVMIDFromPtr(prod1.CustomFields) != "101" {
		t.Errorf("expected VMID 101, got %s", model.ExtractVMIDFromPtr(prod1.CustomFields))
	}

	// 2. JSON tanpa customfield / customfields string kosong ""
	jsonWithoutVMID := `{
		"result": "success",
		"totalresults": 1,
		"startnumber": 0,
		"numreturned": 1,
		"products": {
			"product": [
				{
					"id": 77274,
					"clientid": 109524,
					"name": "Shared Hosting",
					"status": "Active",
					"serverip": "157.66.55.180",
					"customfields": ""
				}
			]
		}
	}`

	var resp2 model.GetClientProductsResponse
	if err := json.Unmarshal([]byte(jsonWithoutVMID), &resp2); err != nil {
		t.Fatalf("failed to unmarshal jsonWithoutVMID: %v", err)
	}
	if len(resp2.Products.Product) != 1 {
		t.Fatalf("expected 1 product, got %d", len(resp2.Products.Product))
	}
	prod2 := resp2.Products.Product[0]
	if model.ExtractVMIDFromPtr(prod2.CustomFields) != "000" {
		t.Errorf("expected VMID 000, got %s", model.ExtractVMIDFromPtr(prod2.CustomFields))
	}

	// 3. JSON dengan customfield objek tunggal dan nilai numerik
	jsonSingleNumericCF := `{
		"result": "success",
		"totalresults": 1,
		"startnumber": 0,
		"numreturned": 1,
		"products": {
			"product": [
				{
					"id": 77275,
					"clientid": 109524,
					"name": "VPS Pro",
					"status": "Active",
					"customfields": {
						"customfield": {"id": 155, "name": "vmid", "value": 204}
					}
				}
			]
		}
	}`

	var resp3 model.GetClientProductsResponse
	if err := json.Unmarshal([]byte(jsonSingleNumericCF), &resp3); err != nil {
		t.Fatalf("failed to unmarshal jsonSingleNumericCF: %v", err)
	}
	if len(resp3.Products.Product) != 1 {
		t.Fatalf("expected 1 product, got %d", len(resp3.Products.Product))
	}
	prod3 := resp3.Products.Product[0]
	if model.ExtractVMIDFromPtr(prod3.CustomFields) != "204" {
		t.Errorf("expected VMID 204, got %s", model.ExtractVMIDFromPtr(prod3.CustomFields))
	}
}

