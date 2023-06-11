package models

type OtherApiRequest struct {
	Data      OtherApiRequestData `json:"data"`
	Message   []string            `json:"message"`
	IsSuccess bool                `json:"isSuccess"`
}
type OtherApiRequestData struct {
	AccountFrom    string `json:"accountFrom"`
	InstitutionID  string `json:"institutionId"`
	NdcProductCode string `json:"ndcProductCode"`
}
