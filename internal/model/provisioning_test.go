package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
)

func TestProvisioningModels_Serialization(t *testing.T) {
	// 1. GetServersRequest & Response
	srvReq := model.GetServersRequest{
		FetchStatus: true,
	}
	if !srvReq.FetchStatus {
		t.Errorf("expected FetchStatus to be true")
	}

	srvRespJSON := `{
		"result": "success",
		"fetchStatus": true,
		"servers": {
			"server": [
				{
					"id": 1,
					"name": "cPanel Prod 01",
					"hostname": "srv1.example.com",
					"ipaddress": "192.168.1.10",
					"maxaccounts": 200,
					"type": "cpanel",
					"active": 1,
					"activeServer": true
				}
			]
		}
	}`
	var srvResp model.GetServersResponse
	if err := json.Unmarshal([]byte(srvRespJSON), &srvResp); err != nil {
		t.Fatalf("failed to unmarshal GetServersResponse: %v", err)
	}
	if srvResp.Result != "success" || len(srvResp.Servers.Server) != 1 {
		t.Errorf("unexpected server list result: %+v", srvResp)
	}
	if srvResp.Servers.Server[0].Hostname != "srv1.example.com" {
		t.Errorf("expected hostname srv1.example.com, got %s", srvResp.Servers.Server[0].Hostname)
	}

	// 2. Module Command Requests
	modCreate := model.ModuleCreateRequest{AccountID: 101}
	if modCreate.AccountID != 101 {
		t.Errorf("expected AccountID 101, got %d", modCreate.AccountID)
	}

	modSuspend := model.ModuleSuspendRequest{AccountID: 101, SuspendReason: "Overdue payment"}
	if modSuspend.SuspendReason != "Overdue payment" {
		t.Errorf("expected SuspendReason 'Overdue payment', got %s", modSuspend.SuspendReason)
	}

	modUnsuspend := model.ModuleUnsuspendRequest{AccountID: 101}
	if modUnsuspend.AccountID != 101 {
		t.Errorf("expected AccountID 101, got %d", modUnsuspend.AccountID)
	}

	modTerminate := model.ModuleTerminateRequest{AccountID: 101}
	if modTerminate.AccountID != 101 {
		t.Errorf("expected AccountID 101, got %d", modTerminate.AccountID)
	}

	modPass := model.ModuleChangePasswordRequest{AccountID: 101, ServicePassword: "NewSecretPassword123!"}
	if modPass.ServicePassword != "NewSecretPassword123!" {
		t.Errorf("expected password 'NewSecretPassword123!', got %s", modPass.ServicePassword)
	}
}
