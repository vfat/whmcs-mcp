package model_test

import (
	"encoding/json"
	"testing"

	"github.com/vfat/whmcs-mcp/internal/model"
)

func TestMarketingModels_Serialization(t *testing.T) {
	// 1. GetAffiliates
	affReq := model.GetAffiliatesRequest{LimitStart: 0, LimitNum: 10}
	if affReq.LimitNum != 10 {
		t.Errorf("expected LimitNum 10, got %d", affReq.LimitNum)
	}

	affJSON := `{
		"result": "success",
		"totalresults": 1,
		"affiliates": {
			"affiliate": [
				{
					"id": 5,
					"userid": 12,
					"visitors": 150,
					"balance": "45.00"
				}
			]
		}
	}`
	var affResp model.GetAffiliatesResponse
	if err := json.Unmarshal([]byte(affJSON), &affResp); err != nil {
		t.Fatalf("failed to unmarshal GetAffiliatesResponse: %v", err)
	}
	if affResp.TotalResults != 1 || len(affResp.Affiliates.Affiliate) != 1 {
		t.Errorf("unexpected affiliate count: %+v", affResp)
	}

	// 2. AffiliateActivate
	actReq := model.AffiliateActivateRequest{UserID: 12}
	if actReq.UserID != 12 {
		t.Errorf("expected UserID 12, got %d", actReq.UserID)
	}

	// 3. GetPromotions
	promReq := model.GetPromotionsRequest{Code: "SUMMER50"}
	if promReq.Code != "SUMMER50" {
		t.Errorf("expected Code SUMMER50, got %s", promReq.Code)
	}

	promJSON := `{
		"result": "success",
		"totalresults": 1,
		"promotions": {
			"promotion": [
				{
					"id": 2,
					"code": "SUMMER50",
					"type": "percentage",
					"value": "50.00"
				}
			]
		}
	}`
	var promResp model.GetPromotionsResponse
	if err := json.Unmarshal([]byte(promJSON), &promResp); err != nil {
		t.Fatalf("failed to unmarshal GetPromotionsResponse: %v", err)
	}
	if promResp.TotalResults != 1 || promResp.Promotions.Promotion[0].Code != "SUMMER50" {
		t.Errorf("unexpected promotions response: %+v", promResp)
	}

	// 4. LogActivity
	logReq := model.LogActivityRequest{Description: "Manual maintenance performed", UserID: 1}
	if logReq.Description != "Manual maintenance performed" {
		t.Errorf("expected description match, got %s", logReq.Description)
	}

	// 5. GetEmailTemplates
	emailReq := model.GetEmailTemplatesRequest{Type: "invoice", Language: "english"}
	if emailReq.Type != "invoice" {
		t.Errorf("expected Type invoice, got %s", emailReq.Type)
	}

	// 6. DeleteClient
	delReq := model.DeleteClientRequest{ClientID: 99}
	if delReq.ClientID != 99 {
		t.Errorf("expected ClientID 99, got %d", delReq.ClientID)
	}
}
