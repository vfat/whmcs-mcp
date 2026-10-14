package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// --- GetClients ---

type GetClientsRequest struct {
	Search     string `json:"search,omitempty" url:"search,omitempty"`
	Status     string `json:"status,omitempty" url:"status,omitempty"`
	LimitStart int    `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum   int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
	Sorting    string `json:"sorting,omitempty" url:"sorting,omitempty"`
}

type ClientListItem struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	CompanyName string `json:"companyname"`
	Email       string `json:"email"`
	DateCreated string `json:"datecreated"`
	GroupId     int    `json:"groupid"`
	Status      string `json:"status"`
}

type GetClientsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Clients      struct {
		Client []ClientListItem `json:"client"`
	} `json:"clients"`
}

// --- GetClientsDetails ---

type GetClientDetailsRequest struct {
	ClientID int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	Email    string `json:"email,omitempty" url:"email,omitempty"`
	Stats    bool   `json:"stats,omitempty" url:"stats,omitempty"`
}

type ClientDetails struct {
	ID          int    `json:"id"`
	ClientID    int    `json:"clientid"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	CompanyName string `json:"companyname"`
	Email       string `json:"email"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phonenumber"`
	Status      string `json:"status"`
	Credit      string `json:"credit"`
	Currency    int    `json:"currency"`
}

type ClientStats struct {
	NumDueInvoices    int    `json:"numdueinvoices"`
	DueInvoicesBalance string `json:"dueinvoicesbalance"`
	NumOverdueInvoices int   `json:"numoverdueinvoices"`
	NumPaidInvoices   int    `json:"numpaidinvoices"`
	NumUnpaidInvoices int    `json:"numunpaidinvoices"`
	NumCancelledInvoices int `json:"numcancelledinvoices"`
	NumRefundedInvoices int  `json:"numrefundedinvoices"`
	NumCollectionsInvoices int `json:"numcollectionsinvoices"`
	NumProducts       int    `json:"numproducts"`
	NumDomains        int    `json:"numdomains"`
	NumQuotes         int    `json:"numquotes"`
	NumTickets        int    `json:"numtickets"`
	NumAffiliates     int    `json:"numaffiliates"`
}

type GetClientDetailsResponse struct {
	CommonResponse
	ClientDetails
	Stats *ClientStats `json:"stats,omitempty"`
}

// --- AddClient ---

type AddClientRequest struct {
	FirstName    string   `json:"firstname" url:"firstname"`
	LastName     string   `json:"lastname" url:"lastname"`
	CompanyName  string   `json:"companyname,omitempty" url:"companyname,omitempty"`
	Email        string   `json:"email" url:"email"`
	Address1     string   `json:"address1" url:"address1"`
	Address2     string   `json:"address2,omitempty" url:"address2,omitempty"`
	City         string   `json:"city" url:"city"`
	State        string   `json:"state" url:"state"`
	Postcode     string   `json:"postcode" url:"postcode"`
	Country      string   `json:"country" url:"country"`
	PhoneNumber  string   `json:"phonenumber,omitempty" url:"phonenumber,omitempty"`
	Password2    string   `json:"password2,omitempty" url:"password2,omitempty"`
	Currency     int      `json:"currency,omitempty" url:"currency,omitempty"`
	CustomFields []string `json:"customfields,omitempty" url:"customfields,omitempty"`
}

type AddClientResponse struct {
	CommonResponse
	ClientID int `json:"clientid"`
}

// --- UpdateClient ---

type UpdateClientRequest struct {
	ClientID    int    `json:"clientid" url:"clientid"`
	FirstName   string `json:"firstname,omitempty" url:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty" url:"lastname,omitempty"`
	CompanyName string `json:"companyname,omitempty" url:"companyname,omitempty"`
	Email       string `json:"email,omitempty" url:"email,omitempty"`
	Address1    string `json:"address1,omitempty" url:"address1,omitempty"`
	Address2    string `json:"address2,omitempty" url:"address2,omitempty"`
	City        string `json:"city,omitempty" url:"city,omitempty"`
	State       string `json:"state,omitempty" url:"state,omitempty"`
	Postcode    string `json:"postcode,omitempty" url:"postcode,omitempty"`
	Country     string `json:"country,omitempty" url:"country,omitempty"`
	PhoneNumber string `json:"phonenumber,omitempty" url:"phonenumber,omitempty"`
	Status      string `json:"status,omitempty" url:"status,omitempty"`
}

type UpdateClientResponse struct {
	CommonResponse
	ClientID int `json:"clientid"`
}

// --- CloseClient ---

type CloseClientRequest struct {
	ClientID int `json:"clientid" url:"clientid"`
}

type CloseClientResponse struct {
	CommonResponse
	ClientID int `json:"clientid"`
}

// --- GetClientsProducts ---

type GetClientProductsRequest struct {
	ClientID  int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	ServiceID int    `json:"serviceid,omitempty" url:"serviceid,omitempty"`
	PID       int    `json:"pid,omitempty" url:"pid,omitempty"`
	Domain    string `json:"domain,omitempty" url:"domain,omitempty"`
	Username  string `json:"username,omitempty" url:"username,omitempty"`
}

// CustomFieldItem merepresentasikan sebuah item customfield pada produk WHMCS.
type CustomFieldItem struct {
	ID    interface{} `json:"id"`
	Name  string      `json:"name"`
	Value string      `json:"value"`
}

func (c *CustomFieldItem) UnmarshalJSON(data []byte) error {
	type Alias CustomFieldItem
	aux := struct {
		RawValue interface{} `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.RawValue == nil {
		c.Value = ""
	} else {
		switch v := aux.RawValue.(type) {
		case string:
			c.Value = v
		case float64:
			if v == float64(int64(v)) {
				c.Value = strconv.FormatInt(int64(v), 10)
			} else {
				c.Value = fmt.Sprintf("%v", v)
			}
		default:
			c.Value = fmt.Sprintf("%v", v)
		}
	}
	return nil
}

// RawCustomFields merepresentasikan pembungkus customfields dari WHMCS API.
type RawCustomFields struct {
	CustomField []CustomFieldItem `json:"customfield"`
}

func (r *RawCustomFields) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" || trimmed == `""` || trimmed == "null" || trimmed == "{}" || trimmed == "[]" {
		r.CustomField = nil
		return nil
	}

	var wrapper struct {
		CustomField json.RawMessage `json:"customfield"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		r.CustomField = nil
		return nil
	}

	cfTrimmed := strings.TrimSpace(string(wrapper.CustomField))
	if cfTrimmed == "" || cfTrimmed == `""` || cfTrimmed == "null" || cfTrimmed == "{}" || cfTrimmed == "[]" {
		r.CustomField = nil
		return nil
	}

	var items []CustomFieldItem
	if err := json.Unmarshal(wrapper.CustomField, &items); err == nil {
		r.CustomField = items
		return nil
	}

	var single CustomFieldItem
	if err := json.Unmarshal(wrapper.CustomField, &single); err == nil {
		r.CustomField = []CustomFieldItem{single}
		return nil
	}

	r.CustomField = nil
	return nil
}

// ExtractVMID mengekstrak nilai VMID dari RawCustomFields (case-insensitive) dengan default "000".
func ExtractVMID(customFields RawCustomFields) string {
	for _, cf := range customFields.CustomField {
		val := strings.TrimSpace(cf.Value)
		if strings.EqualFold(cf.Name, "vmid") && val != "" {
			return val
		}
	}
	return "000" // Default jika tidak ditemukan atau kosong
}

// ExtractVMIDFromPtr mengekstrak nilai VMID dari pointer *RawCustomFields.
func ExtractVMIDFromPtr(customFields *RawCustomFields) string {
	if customFields == nil {
		return "000"
	}
	return ExtractVMID(*customFields)
}

type ClientProductItem struct {
	ID              int              `json:"id"`
	ClientID        int              `json:"clientid"`
	OrderID         int              `json:"orderid"`
	PID             int              `json:"pid"`
	RegDate         string           `json:"regdate"`
	Name            string           `json:"name"`
	Group           string           `json:"translated_group_name"`
	Domain          string           `json:"domain"`
	DedicatedIP     string           `json:"dedicatedip"`
	ServerID        int              `json:"serverid"`
	ServerName      string           `json:"servername"`
	ServerIP        string           `json:"serverip"`
	ServerHostname  string           `json:"serverhostname"`
	FirstPayment    string           `json:"firstpaymentamount"`
	RecurringAmount string           `json:"recurringamount"`
	BillingCycle    string           `json:"billingcycle"`
	NextDueDate     string           `json:"nextduedate"`
	Status          string           `json:"status"`
	Username        string           `json:"username"`
	VMID            string           `json:"vmid"` // Field flattened untuk AI Agent
	CustomFields    *RawCustomFields `json:"customfields,omitempty"`
}

// Product adalah alias dari ClientProductItem
type Product = ClientProductItem

type ClientProductsContainer struct {
	Product []ClientProductItem `json:"product"`
}

func (p *ClientProductsContainer) UnmarshalJSON(data []byte) error {
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" || trimmed == `""` || trimmed == "null" || trimmed == "{}" || trimmed == "[]" {
		p.Product = nil
		return nil
	}

	var wrapper struct {
		RawProduct json.RawMessage `json:"product"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil
	}

	pTrimmed := strings.TrimSpace(string(wrapper.RawProduct))
	if pTrimmed == "" || pTrimmed == `""` || pTrimmed == "null" || pTrimmed == "{}" || pTrimmed == "[]" {
		p.Product = nil
		return nil
	}

	var items []ClientProductItem
	if err := json.Unmarshal(wrapper.RawProduct, &items); err == nil {
		p.Product = items
		return nil
	}

	var single ClientProductItem
	if err := json.Unmarshal(wrapper.RawProduct, &single); err == nil {
		p.Product = []ClientProductItem{single}
		return nil
	}

	return nil
}

type GetClientProductsResponse struct {
	CommonResponse
	TotalResults int                     `json:"totalresults"`
	StartNumber  int                     `json:"startnumber"`
	NumReturned  int                     `json:"numreturned"`
	Products     ClientProductsContainer `json:"products"`
}

// --- GetClientsDomains ---

type GetClientDomainsRequest struct {
	ClientID int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	DomainID int    `json:"domainid,omitempty" url:"domainid,omitempty"`
	Domain   string `json:"domain,omitempty" url:"domain,omitempty"`
}

type ClientDomainItem struct {
	ID              int    `json:"id"`
	UserID          int    `json:"userid"`
	Order       int    `json:"orderid"`
	RegType         string `json:"regtype"`
	DomainName      string `json:"domainname"`
	Registrar       string `json:"registrar"`
	RegDate         string `json:"regdate"`
	ExpiryDate      string `json:"expirydate"`
	NextDueDate     string `json:"nextduedate"`
	Status          string `json:"status"`
	SubscriptionID  string `json:"subscriptionid"`
	DNSManagement   bool   `json:"dnsmanagement"`
	EmailForwarding bool   `json:"emailforwarding"`
	IDProtection    bool   `json:"idprotection"`
	DoNotRenew      bool   `json:"donotrenew"`
}

type GetClientDomainsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Domains      struct {
		Domain []ClientDomainItem `json:"domain"`
	} `json:"domains"`
}

// --- GetClientInvoices ---

type GetClientInvoicesRequest struct {
	ClientID int    `json:"clientid" url:"userid"`
	Status   string `json:"status,omitempty" url:"status,omitempty"`
	LimitNum int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
}
