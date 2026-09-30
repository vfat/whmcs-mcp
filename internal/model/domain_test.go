package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestDomainModels_Serialization(t *testing.T) {
	req := model.UpdateDomainNameserversRequest{
		DomainID: 88,
		NS1:      "ns1.example.com",
		NS2:      "ns2.example.com",
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("failed to serialize UpdateDomainNameserversRequest: %v", err)
	}

	if values.Get("domainid") != "88" {
		t.Errorf("expected domainid=88, got %s", values.Get("domainid"))
	}
	if values.Get("ns1") != "ns1.example.com" {
		t.Errorf("expected ns1=ns1.example.com, got %s", values.Get("ns1"))
	}

	jsonBlob := `{
		"result": "success",
		"ns1": "ns1.example.com",
		"ns2": "ns2.example.com"
	}`

	var resp model.GetDomainNameserversResponse
	if err := json.Unmarshal([]byte(jsonBlob), &resp); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if resp.Result != "success" || resp.NS1 != "ns1.example.com" {
		t.Errorf("unexpected response: %+v", resp)
	}
}
