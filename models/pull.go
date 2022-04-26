package models

type PullRequestBody struct {
	Source string `json:"source"`
	Tenant string `json:"tenant"`
	Params Params `json:"params"`
}
