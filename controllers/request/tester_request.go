package request

type TesterRequest struct {
	JustString string `json:"justString" binding:"required"`
}
