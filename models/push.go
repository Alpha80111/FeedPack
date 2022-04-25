package models

type PushRequestBody struct {
	Source string `json:"source"`
	Tenant string `json:"tenant"`
}
