package responses

import (
	"myGoSkeleton/models"

	"github.com/gin-gonic/gin"
)

type commonResp struct {
	g *gin.Context
}

func (cr *commonResp) SuccessResp(httpCode int, succesMessage *models.SuccessResponse) {
	succesMessage.HttpReponse = httpCode
	cr.g.JSON(httpCode, succesMessage)
	cr.g.Abort()
}

func (cr *commonResp) FailedResp(httpCode int, failedMessage *models.ErrorResponse) {
	failedMessage.HttpResponse = httpCode
	cr.g.JSON(httpCode, failedMessage)
	cr.g.Abort()
}

func NewResponse(g *gin.Context) *commonResp {
	return &commonResp{
		g,
	}
}
