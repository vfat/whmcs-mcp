package model

// CommonResponse adalah struktur dasar yang diterima dari seluruh respon API WHMCS.
type CommonResponse struct {
	Result  string `json:"result"`
	Message string `json:"message,omitempty"`
}

// PaginationParams mendefinisikan parameter paginasi dan pengurutan standar WHMCS.
type PaginationParams struct {
	LimitStart int    `json:"limitstart,omitempty" url:"limitstart,omitempty"`
	LimitNum   int    `json:"limitnum,omitempty" url:"limitnum,omitempty"`
	Sorting    string `json:"sorting,omitempty" url:"sorting,omitempty"`
}
