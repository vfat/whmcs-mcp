package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestTicketModels_Serialization(t *testing.T) {
	req := model.OpenTicketRequest{
		DeptID:   1,
		Subject:  "Server unreachable",
		Message:  "My VPS instance is not responding on port 22.",
		ClientID: 42,
		Priority: "High",
		Markdown: true,
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize OpenTicketRequest: %v", err)
	}

	if values.Get("deptid") != "1" {
		t.Errorf("expected deptid=1, got %s", values.Get("deptid"))
	}
	if values.Get("subject") != "Server unreachable" {
		t.Errorf("expected subject='Server unreachable', got %s", values.Get("subject"))
	}
	if values.Get("clientid") != "42" {
		t.Errorf("expected clientid=42, got %s", values.Get("clientid"))
	}
	if values.Get("priority") != "High" {
		t.Errorf("expected priority=High, got %s", values.Get("priority"))
	}

	jsonBlob := `{
		"result": "success",
		"totalresults": 1,
		"startnumber": 0,
		"numreturned": 1,
		"tickets": {
			"ticket": [
				{
					"id": 55,
					"tid": "ABC-12345",
					"deptid": 1,
					"deptname": "Technical Support",
					"userid": 42,
					"name": "Jane Doe",
					"email": "jane@example.com",
					"subject": "Server unreachable",
					"status": "Open",
					"priority": "High"
				}
			]
		}
	}`

	var resp model.GetTicketsResponse
	if err := json.Unmarshal([]byte(jsonBlob), &resp); err != nil {
		t.Fatalf("failed to unmarshal GetTicketsResponse: %v", err)
	}

	if resp.Result != "success" {
		t.Errorf("expected success result, got %s", resp.Result)
	}
	if len(resp.Tickets.Ticket) != 1 || resp.Tickets.Ticket[0].ID != 55 {
		t.Errorf("unexpected ticket items: %+v", resp.Tickets)
	}
}
