package models

type AccessRequestBody struct {
	Tenant  string   `json:"tenant"`
	Sources []string `json:"sources"`
	Page    int      `json:"page"`
	Records int      `json:"records"`
}
