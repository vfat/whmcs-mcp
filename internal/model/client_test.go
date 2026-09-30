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
