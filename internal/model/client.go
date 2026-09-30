package model

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

type ClientProductItem struct {
	ID             int    `json:"id"`
	ClientID       int    `json:"clientid"`
	Order       int    `json:"orderid"`
	PID            int    `json:"pid"`
	RegDate        string `json:"regdate"`
	Name           string `json:"name"`
	Group          string `json:"translated_group_name"`
	Domain         string `json:"domain"`
	DedicatedIP    string `json:"dedicatedip"`
	ServerID       int    `json:"serverid"`
	ServerName     string `json:"servername"`
	ServerIP       string `json:"serverip"`
	ServerHostname string `json:"serverhostname"`
	FirstPayment   string `json:"firstpaymentamount"`
	RecurringAmount string `json:"recurringamount"`
	BillingCycle   string `json:"billingcycle"`
	NextDueDate    string `json:"nextduedate"`
	Status         string `json:"status"`
	Username       string `json:"username"`
}

type GetClientProductsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Products     struct {
		Product []ClientProductItem `json:"product"`
	} `json:"products"`
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
