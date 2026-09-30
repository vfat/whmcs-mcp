package model

// --- RegisterDomain ---

type RegisterDomainRequest struct {
	DomainID int    `json:"domainid,omitempty" url:"domainid,omitempty"`
	Domain   string `json:"domain,omitempty" url:"domain,omitempty"`
}

type RegisterDomainResponse struct {
	CommonResponse
}

// --- TransferDomain ---

type TransferDomainRequest struct {
	DomainID int `json:"domainid" url:"domainid"`
}

type TransferDomainResponse struct {
	CommonResponse
}

// --- RenewDomain ---

type RenewDomainRequest struct {
	DomainID int `json:"domainid" url:"domainid"`
}

type RenewDomainResponse struct {
	CommonResponse
}

// --- GetDomainWhoisInfo ---

type GetDomainWhoisRequest struct {
	DomainID int `json:"domainid" url:"domainid"`
}

type GetDomainWhoisResponse struct {
	CommonResponse
	Whois string `json:"whois"`
}

// --- GetDomainNameservers ---

type GetDomainNameserversRequest struct {
	DomainID int `json:"domainid" url:"domainid"`
}

type GetDomainNameserversResponse struct {
	CommonResponse
	NS1 string `json:"ns1"`
	NS2 string `json:"ns2"`
	NS3 string `json:"ns3"`
	NS4 string `json:"ns4"`
	NS5 string `json:"ns5"`
}

// --- UpdateDomainNameservers ---

type UpdateDomainNameserversRequest struct {
	DomainID int    `json:"domainid" url:"domainid"`
	NS1      string `json:"ns1,omitempty" url:"ns1,omitempty"`
	NS2      string `json:"ns2,omitempty" url:"ns2,omitempty"`
	NS3      string `json:"ns3,omitempty" url:"ns3,omitempty"`
	NS4      string `json:"ns4,omitempty" url:"ns4,omitempty"`
	NS5      string `json:"ns5,omitempty" url:"ns5,omitempty"`
}

type UpdateDomainNameserversResponse struct {
	CommonResponse
}

// --- GetDomainLockingStatus ---

type GetDomainLockStatusRequest struct {
	DomainID int `json:"domainid" url:"domainid"`
}

type GetDomainLockStatusResponse struct {
	CommonResponse
	LockStatus string `json:"lockstatus"`
}

// --- UpdateDomainLockingStatus ---

type UpdateDomainLockStatusRequest struct {
	DomainID   int  `json:"domainid" url:"domainid"`
	LockStatus bool `json:"lockstatus" url:"lockstatus"`
}

type UpdateDomainLockStatusResponse struct {
	CommonResponse
}

// --- GetTLDPricing ---

type GetTLDPricingRequest struct {
	CurrencyID int `json:"currencyid,omitempty" url:"currencyid,omitempty"`
}

type TLDPricingItem struct {
	Register any `json:"register,omitempty"`
	Transfer any `json:"transfer,omitempty"`
	Renew    any `json:"renew,omitempty"`
}

type GetTLDPricingResponse struct {
	CommonResponse
	Currency any                       `json:"currency"`
	Pricing  map[string]TLDPricingItem `json:"pricing"`
}
