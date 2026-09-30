package model

// --- GetTickets ---

type GetTicketsRequest struct {
	LimitStart int    `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum   int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
	DeptID     int    `json:"deptid,omitempty" url:"deptid,omitempty"`
	ClientID   int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	Email      string `json:"email,omitempty" url:"email,omitempty"`
	Status     string `json:"status,omitempty" url:"status,omitempty"`
	Subject    string `json:"subject,omitempty" url:"subject,omitempty"`
}

type TicketListItem struct {
	ID                 int    `json:"id"`
	TID                string `json:"tid"`
	DeptID             int    `json:"deptid"`
	DeptName           string `json:"deptname"`
	UserID             int    `json:"userid"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	CC                 string `json:"cc"`
	CName              string `json:"cname"`
	Date               string `json:"date"`
	Subject            string `json:"subject"`
	Status             string `json:"status"`
	Priority           string `json:"priority"`
	Admin              string `json:"admin"`
	LastReply          string `json:"lastreply"`
	Flag               int    `json:"flag"`
	Service            string `json:"service"`
	Unread             bool   `json:"unread"`
}

type GetTicketsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	StartNumber  int `json:"startnumber"`
	NumReturned  int `json:"numreturned"`
	Tickets      struct {
		Ticket []TicketListItem `json:"ticket"`
	} `json:"tickets"`
}

// --- GetTicket ---

type GetTicketRequest struct {
	TicketID int    `json:"ticketid,omitempty" url:"ticketid,omitempty"`
	TicketNum string `json:"ticketnum,omitempty" url:"ticketnum,omitempty"`
}

type TicketReplyItem struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userid"`
	ContactID int   `json:"contactid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Date     string `json:"date"`
	Message  string `json:"message"`
	Attachment string `json:"attachment"`
	Admin    string `json:"admin"`
}

type TicketNoteItem struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	Admin   string `json:"admin"`
	Message string `json:"message"`
}

type GetTicketResponse struct {
	CommonResponse
	TicketID   int    `json:"ticketid"`
	TID        string `json:"tid"`
	C          string `json:"c"`
	DeptID     int    `json:"deptid"`
	DeptName   string `json:"deptname"`
	UserID     int    `json:"userid"`
	ContactID  int    `json:"contactid"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	CC         string `json:"cc"`
	Date       string `json:"date"`
	Subject    string `json:"subject"`
	Status     string `json:"status"`
	Priority   string `json:"priority"`
	Admin      string `json:"admin"`
	LastReply  string `json:"lastreply"`
	Flag       int    `json:"flag"`
	Replies    struct {
		Reply []TicketReplyItem `json:"reply"`
	} `json:"replies"`
	Notes struct {
		Note []TicketNoteItem `json:"note"`
	} `json:"notes"`
}

// --- OpenTicket ---

type OpenTicketRequest struct {
	DeptID    int    `json:"deptid" url:"deptid"`
	Subject   string `json:"subject" url:"subject"`
	Message   string `json:"message" url:"message"`
	ClientID  int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	ContactID int    `json:"contactid,omitempty" url:"contactid,omitempty"`
	Name      string `json:"name,omitempty" url:"name,omitempty"`
	Email     string `json:"email,omitempty" url:"email,omitempty"`
	Priority  string `json:"priority,omitempty" url:"priority,omitempty"`
	ServiceID int    `json:"serviceid,omitempty" url:"serviceid,omitempty"`
	DomainID  int    `json:"domainid,omitempty" url:"domainid,omitempty"`
	Admin     bool   `json:"admin,omitempty" url:"admin,omitempty"`
	Markdown  bool   `json:"markdown,omitempty" url:"markdown,omitempty"`
}

type OpenTicketResponse struct {
	CommonResponse
	ID     int    `json:"id"`
	TID    string `json:"tid"`
	C      string `json:"c"`
}

// --- AddTicketReply ---

type AddTicketReplyRequest struct {
	TicketID      int    `json:"ticketid" url:"ticketid"`
	Message       string `json:"message" url:"message"`
	ClientID      int    `json:"clientid,omitempty" url:"clientid,omitempty"`
	ContactID     int    `json:"contactid,omitempty" url:"contactid,omitempty"`
	Name          string `json:"name,omitempty" url:"name,omitempty"`
	Email         string `json:"email,omitempty" url:"email,omitempty"`
	AdminUsername string `json:"adminusername,omitempty" url:"adminusername,omitempty"`
	Status        string `json:"status,omitempty" url:"status,omitempty"`
	NoEmail       bool   `json:"noemail,omitempty" url:"noemail,omitempty"`
	Markdown      bool   `json:"markdown,omitempty" url:"markdown,omitempty"`
}

type AddTicketReplyResponse struct {
	CommonResponse
}

// --- AddTicketNote ---

type AddTicketNoteRequest struct {
	TicketID int    `json:"ticketid" url:"ticketid"`
	Message  string `json:"message" url:"message"`
	Markdown bool   `json:"markdown,omitempty" url:"markdown,omitempty"`
}

type AddTicketNoteResponse struct {
	CommonResponse
}

// --- UpdateTicket ---

type UpdateTicketRequest struct {
	TicketID int    `json:"ticketid" url:"ticketid"`
	DeptID   int    `json:"deptid,omitempty" url:"deptid,omitempty"`
	Subject  string `json:"subject,omitempty" url:"subject,omitempty"`
	UserID   int    `json:"userid,omitempty" url:"userid,omitempty"`
	Name     string `json:"name,omitempty" url:"name,omitempty"`
	Email    string `json:"email,omitempty" url:"email,omitempty"`
	Priority string `json:"priority,omitempty" url:"priority,omitempty"`
	Status   string `json:"status,omitempty" url:"status,omitempty"`
	Flag     int    `json:"flag,omitempty" url:"flag,omitempty"`
}

type UpdateTicketResponse struct {
	CommonResponse
	TicketID int `json:"ticketid"`
}

// --- DeleteTicket ---

type DeleteTicketRequest struct {
	TicketID int `json:"ticketid" url:"ticketid"`
}

type DeleteTicketResponse struct {
	CommonResponse
	TicketID int `json:"ticketid"`
}

// --- GetSupportDepartments ---

type GetSupportDepartmentsRequest struct {
	IgnoreDeptAssignments bool `json:"ignore_dept_assignments,omitempty" url:"ignore_dept_assignments,omitempty"`
}

type SupportDepartmentItem struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AwaitingReply int   `json:"awaitingreply"`
	OpenTickets  int    `json:"opentickets"`
}

type GetSupportDepartmentsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	Departments  struct {
		Department []SupportDepartmentItem `json:"department"`
	} `json:"departments"`
}

// --- GetSupportStatuses ---

type GetSupportStatusesRequest struct {
	DeptID int `json:"deptid,omitempty" url:"deptid,omitempty"`
}

type SupportStatusItem struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

type GetSupportStatusesResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	Statuses     struct {
		Status []SupportStatusItem `json:"status"`
	} `json:"statuses"`
}
