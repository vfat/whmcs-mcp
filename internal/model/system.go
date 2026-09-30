package model

// --- GetStats ---

type GetStatsRequest struct{}

type GetStatsResponse struct {
	CommonResponse
	IncomeToday       string `json:"income_today"`
	IncomeThisMonth   string `json:"income_thismonth"`
	IncomeThisYear    string `json:"income_thisyear"`
	IncomeAllTime     string `json:"income_alltime"`
	OrdersPending     int    `json:"orders_pending"`
	OrdersActive      int    `json:"orders_active"`
	OrdersCancelled   int    `json:"orders_cancelled"`
	OrdersFraud       int    `json:"orders_fraud"`
	InvoicesPaid      int    `json:"invoices_paid"`
	InvoicesUnpaid    int    `json:"invoices_unpaid"`
	InvoicesOverdue   int    `json:"invoices_overdue"`
	InvoicesCancelled int    `json:"invoices_cancelled"`
	InvoicesRefunded  int    `json:"invoices_refunded"`
	InvoicesCollections int  `json:"invoices_collections"`
	TicketsOpen       int    `json:"tickets_open"`
	TicketsAnswered   int    `json:"tickets_answered"`
	TicketsCustomerReply int `json:"tickets_customerreply"`
	TicketsClosed     int    `json:"tickets_closed"`
	ServicesActive    int    `json:"services_active"`
	ServicesSuspended int    `json:"services_suspended"`
	ServicesTerminated int   `json:"services_terminated"`
	ServicesCancelled int    `json:"services_cancelled"`
}

// --- GetActivityLog ---

type GetActivityLogRequest struct {
	LimitStart  int    `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum    int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
	UserID      int    `json:"userid,omitempty" url:"userid,omitempty"`
	Date        string `json:"date,omitempty" url:"date,omitempty"`
	User        string `json:"user,omitempty" url:"user,omitempty"`
	Description string `json:"description,omitempty" url:"description,omitempty"`
	IPAddress   string `json:"ipaddress,omitempty" url:"ipaddress,omitempty"`
}

type ActivityLogItem struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	User        string `json:"user"`
	IPAddress   string `json:"ipaddress"`
}

type GetActivityLogResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Activity     struct {
		Entry []ActivityLogItem `json:"entry"`
	} `json:"activity"`
}

// --- GetAdminUsers ---

type GetAdminUsersRequest struct {
	RoleID int `json:"roleid,omitempty" url:"roleid,omitempty"`
}

type AdminUserItem struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	RoleID    int    `json:"roleid"`
	RoleName  string `json:"role_name"`
	Disabled  bool   `json:"disabled"`
}

type GetAdminUsersResponse struct {
	CommonResponse
	TotalResults int             `json:"totalresults"`
	AdminUsers   []AdminUserItem `json:"admin_users"`
}

// --- GetToDoItems ---

type GetToDoItemsRequest struct {
	LimitStart int    `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum   int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
	Status     string `json:"status,omitempty" url:"status,omitempty"`
}

type ToDoItem struct {
	ID          int    `json:"id"`
	Date        string `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AdminID     int    `json:"adminid"`
	Status      string `json:"status"`
	DueDate     string `json:"duedate"`
}

type GetToDoItemsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Items        struct {
		Item []ToDoItem `json:"item"`
	} `json:"items"`
}

// --- GetToDoItemStatuses ---

type ToDoStatusItem struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

type GetToDoStatusesResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	Statuses     struct {
		Status []ToDoStatusItem `json:"status"`
	} `json:"todo_item_statuses"`
}

// --- UpdateToDoItem ---

type UpdateToDoItemRequest struct {
	ID          int    `json:"itemid" url:"itemid"`
	Status      string `json:"status,omitempty" url:"status,omitempty"`
	Date        string `json:"date,omitempty" url:"date,omitempty"`
	Title       string `json:"title,omitempty" url:"title,omitempty"`
	Description string `json:"description,omitempty" url:"description,omitempty"`
	AdminID     int    `json:"adminid,omitempty" url:"adminid,omitempty"`
	DueDate     string `json:"duedate,omitempty" url:"duedate,omitempty"`
}

type UpdateToDoItemResponse struct {
	CommonResponse
	TodoID int `json:"todoid"`
}

// --- GetCurrencies ---

type CurrencyItem struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
	Format int    `json:"format"`
	Rate   string `json:"rate"`
}

type GetCurrenciesResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	Currencies   struct {
		Currency []CurrencyItem `json:"currency"`
	} `json:"currencies"`
}

// --- GetPaymentMethods ---

type PaymentMethodItem struct {
	Module      string `json:"module"`
	DisplayName string `json:"display_name"`
}

type GetPaymentMethodsResponse struct {
	CommonResponse
	TotalResults   int `json:"totalresults"`
	PaymentMethods struct {
		PaymentMethod []PaymentMethodItem `json:"paymentmethod"`
	} `json:"paymentmethods"`
}

// --- SendEmail ---

type SendEmailRequest struct {
	Messagename   string `json:"messagename,omitempty" url:"messagename,omitempty"`
	ID            int    `json:"id,omitempty" url:"id,omitempty"`
	CustomType    string `json:"customtype,omitempty" url:"customtype,omitempty"`
	CustomSubject string `json:"customsubject,omitempty" url:"customsubject,omitempty"`
	CustomMessage string `json:"custommessage,omitempty" url:"custommessage,omitempty"`
}

type SendEmailResponse struct {
	CommonResponse
}
