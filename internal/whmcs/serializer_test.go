package whmcs_test

import (
	"testing"

	"github.com/vfat/whmcs-mcp/internal/whmcs"
)

func TestSerializeParams_Flat(t *testing.T) {
	input := map[string]interface{}{
		"action":     "GetClients",
		"limitnum":   25,
		"limitstart": 0,
		"search":     "john",
	}

	values, err := whmcs.SerializeParams(input)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if values.Get("action") != "GetClients" {
		t.Errorf("expected action GetClients, got %s", values.Get("action"))
	}
	if values.Get("limitnum") != "25" {
		t.Errorf("expected limitnum 25, got %s", values.Get("limitnum"))
	}
	if values.Get("limitstart") != "0" {
		t.Errorf("expected limitstart 0, got %s", values.Get("limitstart"))
	}
	if values.Get("search") != "john" {
		t.Errorf("expected search john, got %s", values.Get("search"))
	}
}

func TestSerializeParams_NestedMap(t *testing.T) {
	input := map[string]interface{}{
		"action": "DomainUpdateNameservers",
		"nameservers": map[string]string{
			"ns1": "ns1.example.com",
			"ns2": "ns2.example.com",
		},
	}

	values, err := whmcs.SerializeParams(input)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if values.Get("action") != "DomainUpdateNameservers" {
		t.Errorf("expected action DomainUpdateNameservers, got %s", values.Get("action"))
	}
	if values.Get("nameservers[ns1]") != "ns1.example.com" {
		t.Errorf("expected nameservers[ns1]=ns1.example.com, got %s", values.Get("nameservers[ns1]"))
	}
	if values.Get("nameservers[ns2]") != "ns2.example.com" {
		t.Errorf("expected nameservers[ns2]=ns2.example.com, got %s", values.Get("nameservers[ns2]"))
	}
}

func TestSerializeParams_Slice(t *testing.T) {
	input := map[string]interface{}{
		"action":       "AddClient",
		"customfields": []string{"fieldA", "fieldB"},
	}

	values, err := whmcs.SerializeParams(input)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if values.Get("customfields[0]") != "fieldA" {
		t.Errorf("expected customfields[0]=fieldA, got %s", values.Get("customfields[0]"))
	}
	if values.Get("customfields[1]") != "fieldB" {
		t.Errorf("expected customfields[1]=fieldB, got %s", values.Get("customfields[1]"))
	}
}

type SampleRequest struct {
	Action   string            `url:"action"`
	ClientID int               `url:"clientid"`
	Domain   string            `url:"domain,omitempty"`
	EmptyOpt string            `url:"emptyopt,omitempty"`
	Config   map[string]string `url:"config,omitempty"`
}

func TestSerializeParams_Struct(t *testing.T) {
	req := SampleRequest{
		Action:   "GetClientsProducts",
		ClientID: 123,
		Config: map[string]string{
			"pkg": "cpanel-pro",
		},
	}

	values, err := whmcs.SerializeParams(req)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if values.Get("action") != "GetClientsProducts" {
		t.Errorf("expected action GetClientsProducts, got %s", values.Get("action"))
	}
	if values.Get("clientid") != "123" {
		t.Errorf("expected clientid 123, got %s", values.Get("clientid"))
	}
	if values.Get("config[pkg]") != "cpanel-pro" {
		t.Errorf("expected config[pkg]=cpanel-pro, got %s", values.Get("config[pkg]"))
	}
	if values.Has("emptyopt") {
		t.Errorf("expected emptyopt omitted, but got present")
	}
}
