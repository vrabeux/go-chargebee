package chargebee

type Plan struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	InvoiceName          string `json:"invoice_name"`
	Description          string `json:"description"`
	Price                int64  `json:"price"`
	Period               int32  `json:"period"`
	PeriodUnit           string `json:"period_unit"`
	ChargeModel          string `json:"charge_model"`
	FreeQuantity         int64  `json:"free_quantity"`
	Status               string `json:"status"`
	EnabledInHostedPages bool   `json:"enabled_in_hosted_pages"`
	EnabledInPortal      bool   `json:"enabled_in_portal"`
	UpdatedAt            int64  `json:"updated_at"`
	ResourceVersion      int64  `json:"resource_version"`
	Object               string `json:"object"`
	Taxable              bool   `json:"taxable"`
	CurencyCode          string `json:"curency_code"`
}

type PlanList struct{
	Plan *Plan `json:"plan"`
}

type PlanResult struct {
	List []*PlanList `json:"list"`
}

type ListPlansParams struct {
	Limit int
}


