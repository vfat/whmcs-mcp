package model

// --- Orders ---

type GetOrdersRequest struct {
	LimitStart int    `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum   int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
	ID         int    `json:"id,omitempty" url:"id,omitempty"`
	UserID     int    `json:"userid,omitempty" url:"userid,omitempty"`
	Status     string `json:"status,omitempty" url:"status,omitempty"`
}

type OrderItem struct {
	ID             int    `json:"id"`
	OrderNum       string `json:"ordernum"`
	UserID         int    `json:"userid"`
	ContactID      int    `json:"contactid"`
	Date           string `json:"date"`
	Nameservers    string `json:"nameservers"`
	TransferSecret string `json:"transfersecret"`
	Renewals       string `json:"renewals"`
	Promocode      string `json:"promocode"`
	PromoType      string `json:"promotype"`
	PromoValue     string `json:"promovalue"`
	OrderData      string `json:"orderdata"`
	Amount         string `json:"amount"`
	PaymentMethod  string `json:"paymentmethod"`
	InvoiceID      int    `json:"invoiceid"`
	Status         string `json:"status"`
}

type GetOrdersResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Orders       struct {
		Order []OrderItem `json:"order"`
	} `json:"orders"`
}

type AcceptOrderRequest struct {
	OrderID         int    `json:"orderid" url:"orderid"`
	ServerID        int    `json:"serverid,omitempty" url:"serverid,omitempty"`
	ServiceUsername string `json:"serviceusername,omitempty" url:"serviceusername,omitempty"`
	ServicePassword string `json:"servicepassword,omitempty" url:"servicepassword,omitempty"`
	Registrar       string `json:"registrar,omitempty" url:"registrar,omitempty"`
	AutoSetup       bool   `json:"autosetup,omitempty" url:"autosetup,omitempty"`
	SendEmail       bool   `json:"sendemail,omitempty" url:"sendemail,omitempty"`
}

type AcceptOrderResponse struct {
	CommonResponse
}

type CancelOrderRequest struct {
	OrderID   int  `json:"orderid" url:"orderid"`
	CancelSub bool `json:"cancelsub,omitempty" url:"cancelsub,omitempty"`
	NoEmail   bool `json:"noemail,omitempty" url:"noemail,omitempty"`
}

type CancelOrderResponse struct {
	CommonResponse
}

type DeleteOrderRequest struct {
	OrderID int `json:"orderid" url:"orderid"`
}

type DeleteOrderResponse struct {
	CommonResponse
}

type FraudOrderRequest struct {
	OrderID   int  `json:"orderid" url:"orderid"`
	CancelSub bool `json:"cancelsub,omitempty" url:"cancelsub,omitempty"`
}

type FraudOrderResponse struct {
	CommonResponse
}

type PendingOrderRequest struct {
	OrderID int `json:"orderid" url:"orderid"`
}

type PendingOrderResponse struct {
	CommonResponse
}

// --- Quotes ---

type GetQuotesRequest struct {
	QuoteID  int    `json:"quoteid,omitempty" url:"quoteid,omitempty"`
	ClientID int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	Subject  string `json:"subject,omitempty" url:"subject,omitempty"`
	Stage    string `json:"stage,omitempty" url:"stage,omitempty"`
}

type QuoteItem struct {
	ID          int    `json:"id"`
	Subject     string `json:"subject"`
	Stage       string `json:"stage"`
	Date        string `json:"date"`
	DateCreated string `json:"datecreated"`
	ValidUntil  string `json:"validuntil"`
	Total       string `json:"total"`
	Subtotal    string `json:"subtotal"`
	UserID      int    `json:"userid"`
}

type GetQuotesResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Quotes       struct {
		Quote []QuoteItem `json:"quote"`
	} `json:"quotes"`
}

type CreateQuoteRequest struct {
	Subject    string `json:"subject" url:"subject"`
	Stage      string `json:"stage" url:"stage"`
	Date       string `json:"date,omitempty" url:"date,omitempty"`
	ValidUntil string `json:"validuntil,omitempty" url:"validuntil,omitempty"`
	UserID     int    `json:"userid,omitempty" url:"userid,omitempty"`
}

type CreateQuoteResponse struct {
	CommonResponse
	QuoteID int `json:"quoteid"`
}

type AcceptQuoteRequest struct {
	QuoteID int `json:"quoteid" url:"quoteid"`
}

type AcceptQuoteResponse struct {
	CommonResponse
	InvoiceID int `json:"invoiceid"`
}

type DeleteQuoteRequest struct {
	QuoteID int `json:"quoteid" url:"quoteid"`
}

type DeleteQuoteResponse struct {
	CommonResponse
}
