package model

// GetServersRequest mendefinisikan parameter request untuk aksi GetServers.
type GetServersRequest struct {
	FetchStatus bool `json:"fetchStatus,omitempty" url:"fetchStatus,omitempty"`
}

// ServerItem merepresentasikan rincian satu server hosting.
type ServerItem struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Hostname      string `json:"hostname"`
	IPAddress     string `json:"ipaddress"`
	AssignedIPs   string `json:"assignedips,omitempty"`
	StatusAddress string `json:"statusaddress,omitempty"`
	MaxAccounts   int    `json:"maxaccounts"`
	Type          string `json:"type"`
	Active        int    `json:"active"`
	ActiveServer  bool   `json:"activeServer"`
	Disabled      int    `json:"disabled"`
}

// ServerListContainer membungkus array server dari WHMCS.
type ServerListContainer struct {
	Server []ServerItem `json:"server"`
}

// GetServersResponse mendefinisikan struktur balasan GetServers WHMCS.
type GetServersResponse struct {
	CommonResponse
	FetchStatus bool                `json:"fetchStatus,omitempty"`
	Servers     ServerListContainer `json:"servers"`
}

// ModuleCreateRequest mendefinisikan parameter untuk ModuleCreate.
type ModuleCreateRequest struct {
	AccountID int `json:"accountid" url:"accountid"`
}

// ModuleSuspendRequest mendefinisikan parameter untuk ModuleSuspend.
type ModuleSuspendRequest struct {
	AccountID     int    `json:"accountid" url:"accountid"`
	SuspendReason string `json:"suspendreason,omitempty" url:"suspendreason,omitempty"`
}

// ModuleUnsuspendRequest mendefinisikan parameter untuk ModuleUnsuspend.
type ModuleUnsuspendRequest struct {
	AccountID int `json:"accountid" url:"accountid"`
}

// ModuleTerminateRequest mendefinisikan parameter untuk ModuleTerminate.
type ModuleTerminateRequest struct {
	AccountID int `json:"accountid" url:"accountid"`
}

// ModuleChangePasswordRequest mendefinisikan parameter untuk ModuleChangePassword.
type ModuleChangePasswordRequest struct {
	AccountID       int    `json:"accountid" url:"accountid"`
	ServicePassword string `json:"servicepassword,omitempty" url:"servicepassword,omitempty"`
}

// ModuleResponse adalah respon umum dari aksi perintah modul.
type ModuleResponse struct {
	CommonResponse
	Message string `json:"message,omitempty"`
}
