package model

// --- GetInvoices ---

type GetInvoicesRequest struct {
	LimitStart int    `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum   int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
	UserID     int    `json:"userid,omitempty" url:"userid,omitempty"`
	Status     string `json:"status,omitempty" url:"status,omitempty"`
	OrderBy    string `json:"orderby,omitempty" url:"orderby,omitempty"`
	Order      string `json:"order,omitempty" url:"order,omitempty"`
}

type InvoiceListItem struct {
	ID            int    `json:"id"`
	UserID        int    `json:"userid"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	CompanyName   string `json:"companyname"`
	InvoiceNum    string `json:"invoicenum"`
	Date          string `json:"date"`
	DueDate       string `json:"duedate"`
	DatePaid      string `json:"datepaid"`
	SubTotal      string `json:"subtotal"`
	Credit        string `json:"credit"`
	Tax           string `json:"tax"`
	Tax2          string `json:"tax2"`
	Total         string `json:"total"`
	TaxRate       string `json:"taxrate"`
	TaxRate2      string `json:"taxrate2"`
	Status        string `json:"status"`
	PaymentMethod string `json:"paymentmethod"`
	Notes         string `json:"notes"`
}

type GetInvoicesResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Invoices     struct {
		Invoice []InvoiceListItem `json:"invoice"`
	} `json:"invoices"`
}

// --- GetInvoice ---

type GetInvoiceRequest struct {
	InvoiceID int `json:"invoiceid" url:"invoiceid"`
}

type InvoiceLineItem struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	RelID       int    `json:"relid"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Taxed       int    `json:"taxed"`
}

type InvoiceTransactionItem struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Gateway   string `json:"gateway"`
	TransID   string `json:"transid"`
	AmountIn  string `json:"amountin"`
	Fees      string `json:"fees"`
	AmountOut string `json:"amountout"`
}

type GetInvoiceResponse struct {
	CommonResponse
	InvoiceID     int    `json:"invoiceid"`
	InvoiceNum    string `json:"invoicenum"`
	UserID        int    `json:"userid"`
	Date          string `json:"date"`
	DueDate       string `json:"duedate"`
	DatePaid      string `json:"datepaid"`
	SubTotal      string `json:"subtotal"`
	Credit        string `json:"credit"`
	Tax           string `json:"tax"`
	Tax2          string `json:"tax2"`
	Total         string `json:"total"`
	Balance       string `json:"balance"`
	TaxRate       string `json:"taxrate"`
	TaxRate2      string `json:"taxrate2"`
	Status        string `json:"status"`
	PaymentMethod string `json:"paymentmethod"`
	Notes         string `json:"notes"`
	Items         struct {
		Item []InvoiceLineItem `json:"item"`
	} `json:"items"`
	Transactions struct {
		Transaction []InvoiceTransactionItem `json:"transaction"`
	} `json:"transactions"`
}

// --- CreateInvoice ---

type CreateInvoiceRequest struct {
	UserID          int       `json:"userid" url:"userid"`
	Status          string    `json:"status,omitempty" url:"status,omitempty"`
	SendInvoice     bool      `json:"sendinvoice,omitempty" url:"sendinvoice,omitempty"`
	PaymentMethod   string    `json:"paymentmethod,omitempty" url:"paymentmethod,omitempty"`
	TaxRate         float64   `json:"taxrate,omitempty" url:"taxrate,omitempty"`
	TaxRate2        float64   `json:"taxrate2,omitempty" url:"taxrate2,omitempty"`
	Date            string    `json:"date,omitempty" url:"date,omitempty"`
	DueDate         string    `json:"duedate,omitempty" url:"duedate,omitempty"`
	Notes           string    `json:"notes,omitempty" url:"notes,omitempty"`
	ItemDescription []string  `json:"itemdescription,omitempty" url:"itemdescription,omitempty"`
	ItemAmount      []float64 `json:"itemamount,omitempty" url:"itemamount,omitempty"`
	ItemTaxed       []bool    `json:"itemtaxed,omitempty" url:"itemtaxed,omitempty"`
}

type CreateInvoiceResponse struct {
	CommonResponse
	InvoiceID int `json:"invoiceid"`
}

// --- UpdateInvoice ---

type UpdateInvoiceRequest struct {
	InvoiceID           int    `json:"invoiceid" url:"invoiceid"`
	Status              string `json:"status,omitempty" url:"status,omitempty"`
	PaymentMethod       string `json:"paymentmethod,omitempty" url:"paymentmethod,omitempty"`
	Date                string `json:"date,omitempty" url:"date,omitempty"`
	DueDate             string `json:"duedate,omitempty" url:"duedate,omitempty"`
	Notes               string `json:"notes,omitempty" url:"notes,omitempty"`
	Publish             bool   `json:"publish,omitempty" url:"publish,omitempty"`
	PublishAndSendEmail bool   `json:"publishandsendemail,omitempty" url:"publishandsendemail,omitempty"`
}

type UpdateInvoiceResponse struct {
	CommonResponse
	InvoiceID int `json:"invoiceid"`
}

// --- AddInvoicePayment ---

type AddInvoicePaymentRequest struct {
	InvoiceID int     `json:"invoiceid" url:"invoiceid"`
	TransID   string  `json:"transid" url:"transid"`
	Gateway   string  `json:"gateway" url:"gateway"`
	Date      string  `json:"date,omitempty" url:"date,omitempty"`
	Amount    float64 `json:"amount,omitempty" url:"amount,omitempty"`
	Fees      float64 `json:"fees,omitempty" url:"fees,omitempty"`
	NoEmail   bool    `json:"noemail,omitempty" url:"noemail,omitempty"`
}

type AddInvoicePaymentResponse struct {
	CommonResponse
}

// --- ApplyCredit ---

type ApplyCreditRequest struct {
	InvoiceID int     `json:"invoiceid" url:"invoiceid"`
	Amount    float64 `json:"amount" url:"amount"`
	NoEmail   bool    `json:"noemail,omitempty" url:"noemail,omitempty"`
}

type ApplyCreditResponse struct {
	CommonResponse
	InvoiceID int    `json:"invoiceid"`
	Amount    string `json:"amount"`
}

// --- GetTransactions ---

type GetTransactionsRequest struct {
	InvoiceID int    `json:"invoiceid,omitempty" url:"invoiceid,omitempty"`
	ClientID  int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	TransID   string `json:"transid,omitempty" url:"transid,omitempty"`
}

type TransactionItem struct {
	ID          int    `json:"id"`
	RefundID    int    `json:"refundid"`
	UserID      int    `json:"userid"`
	Currency    int    `json:"currency"`
	Gateway     string `json:"gateway"`
	Date        string `json:"date"`
	Description string `json:"description"`
	AmountIn    string `json:"amountin"`
	Fees        string `json:"fees"`
	AmountOut   string `json:"amountout"`
	Rate        string `json:"rate"`
	TransID     string `json:"transid"`
	InvoiceID   int    `json:"invoiceid"`
}

type GetTransactionsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Transactions struct {
		Transaction []TransactionItem `json:"transaction"`
	} `json:"transactions"`
}
