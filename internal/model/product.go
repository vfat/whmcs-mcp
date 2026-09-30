package model

// --- GetProducts ---

type GetProductsRequest struct {
	PID        int    `json:"pid,omitempty" url:"pid,omitempty"`
	GID        int    `json:"gid,omitempty" url:"gid,omitempty"`
	ModuleName string `json:"module,omitempty" url:"module,omitempty"`
}

type ProductItem struct {
	PID         int    `json:"pid"`
	GID         int    `json:"gid"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Module      string `json:"module"`
	PayType     string `json:"paytype"`
	Pricing     any    `json:"pricing,omitempty"`
}

type GetProductsResponse struct {
	CommonResponse
	TotalResults int `json:"totalresults"`
	Products     struct {
		Product []ProductItem `json:"product"`
	} `json:"products"`
}

// --- GetProductGroups ---

type GetProductGroupsRequest struct{}

type ProductGroupItem struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Headline          string `json:"headline"`
	Tagline           string `json:"tagline"`
	OrderFormTemplate string `json:"orderformtemplate"`
}

type GetProductGroupsResponse struct {
	CommonResponse
	TotalResults  int `json:"totalresults"`
	ProductGroups struct {
		ProductGroup []ProductGroupItem `json:"productgroup"`
	} `json:"productgroups"`
}
