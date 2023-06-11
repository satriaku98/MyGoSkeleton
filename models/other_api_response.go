package models

type OtherApiResponse struct {
	Data      OtherApiResponseData `json:"data"`
	Message   []string             `json:"message"`
	IsSuccess bool                 `json:"isSuccess"`
}
type OtherApiResponseData struct {
	ResponseCode string `json:"responseCode"`
	Details      string `json:"details"`
}
