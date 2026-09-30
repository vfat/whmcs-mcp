package model

// GetAffiliatesRequest mendefinisikan parameter request untuk aksi GetAffiliates.
type GetAffiliatesRequest struct {
	LimitStart int `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum   int `json:"limitnum,omitempty" url:"limitnum,omitempty"`
}

// AffiliateItem merepresentasikan akun afiliasi.
type AffiliateItem struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userid"`
	Date      string `json:"date,omitempty"`
	Visitors  int    `json:"visitors"`
	PayType   string `json:"paytype,omitempty"`
	PayAmount string `json:"payamount,omitempty"`
	OneTime   string `json:"onetime,omitempty"`
	Balance   string `json:"balance,omitempty"`
	Withdrawn string `json:"withdrawn,omitempty"`
}

// AffiliateListContainer membungkus array afiliasi.
type AffiliateListContainer struct {
	Affiliate []AffiliateItem `json:"affiliate"`
}

// GetAffiliatesResponse mendefinisikan balasan GetAffiliates dari WHMCS.
type GetAffiliatesResponse struct {
	CommonResponse
	TotalResults int                    `json:"totalresults"`
	StartNumber  int                    `json:"startnumber"`
	NumReturned  int                    `json:"numreturned"`
	Affiliates   AffiliateListContainer `json:"affiliates"`
}

// AffiliateActivateRequest mendefinisikan parameter untuk AffiliateActivate.
type AffiliateActivateRequest struct {
	UserID int `json:"userid" url:"userid"`
}

// AffiliateActivateResponse mendefinisikan balasan AffiliateActivate.
type AffiliateActivateResponse struct {
	CommonResponse
	AffiliateID int `json:"affiliateid"`
}

// GetPromotionsRequest mendefinisikan parameter pencarian promosi.
type GetPromotionsRequest struct {
	Code string `json:"code,omitempty" url:"code,omitempty"`
}

// PromotionItem merepresentasikan kupon diskon/promosi.
type PromotionItem struct {
	ID               int    `json:"id"`
	Code             string `json:"code"`
	Type             string `json:"type"`
	Recurring        int    `json:"recurring"`
	Value            string `json:"value"`
	Cycles           string `json:"cycles,omitempty"`
	AppliesTo        string `json:"appliesto,omitempty"`
	Requires         string `json:"requires,omitempty"`
	RequiresExisting int    `json:"requiresexisting,omitempty"`
	StartDate        string `json:"startdate,omitempty"`
	ExpirationDate   string `json:"expirationdate,omitempty"`
	MaxUses          int    `json:"maxuses,omitempty"`
	Uses             int    `json:"uses,omitempty"`
	LifetimePromo    int    `json:"lifetimepromo,omitempty"`
	ApplyOnce        int    `json:"applyonce,omitempty"`
	NewSignups       int    `json:"newsignups,omitempty"`
	ExistingClient   int    `json:"existingclient,omitempty"`
	OncePerClient    int    `json:"onceperclient,omitempty"`
	RecurFor         int    `json:"recurfor,omitempty"`
	Upgrades         int    `json:"upgrades,omitempty"`
	UpgradeConfig    string `json:"upgradeconfig,omitempty"`
	Notes            string `json:"notes,omitempty"`
}

// PromotionListContainer membungkus array promosi.
type PromotionListContainer struct {
	Promotion []PromotionItem `json:"promotion"`
}

// GetPromotionsResponse mendefinisikan balasan GetPromotions WHMCS.
type GetPromotionsResponse struct {
	CommonResponse
	TotalResults int                    `json:"totalresults"`
	Promotions   PromotionListContainer `json:"promotions"`
}

// LogActivityRequest mendefinisikan parameter untuk LogActivity.
type LogActivityRequest struct {
	Description string `json:"description" url:"description"`
	UserID      int    `json:"userid,omitempty" url:"userid,omitempty"`
}

// GetEmailTemplatesRequest mendefinisikan parameter filter GetEmailTemplates.
type GetEmailTemplatesRequest struct {
	Type     string `json:"type,omitempty" url:"type,omitempty"`
	Language string `json:"language,omitempty" url:"language,omitempty"`
}

// EmailTemplateItem merepresentasikan templat email.
type EmailTemplateItem struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Subject string `json:"subject"`
	Custom bool   `json:"custom"`
}

// EmailTemplateContainer membungkus array templat email.
type EmailTemplateContainer struct {
	EmailTemplate []EmailTemplateItem `json:"emailtemplate"`
}

// GetEmailTemplatesResponse mendefinisikan balasan GetEmailTemplates.
type GetEmailTemplatesResponse struct {
	CommonResponse
	TotalResults   int                    `json:"totalresults"`
	EmailTemplates EmailTemplateContainer `json:"emailtemplates"`
}

// DeleteClientRequest mendefinisikan parameter untuk DeleteClient.
type DeleteClientRequest struct {
	ClientID int `json:"clientid" url:"clientid"`
}
