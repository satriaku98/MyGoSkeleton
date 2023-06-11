package api

import (
	"fmt"
	"myGoSkeleton/controllers/request"
	"myGoSkeleton/controllers/responses"
	"myGoSkeleton/models"
	"myGoSkeleton/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type tester struct {
	services services.TesterService
}

// @BasePath /api/v1/example
// @Summary Tester Example
// @Schemes
// @Description Tester
// @Tags example
// @Accept json
// @Produce json
// @Param data body request.TesterRequest true "Request body"
// @Success 200 {object} responses.TesterResponse "Success"
// @Router /tester [post]
func (l *tester) Tester() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataTester request.TesterRequest
		if errBind := c.ShouldBindJSON(&dataTester); errBind != nil {
			responses.NewResponse(c).FailedResp(http.StatusInternalServerError, models.FailedMessage(errBind.Error()))
			return
		}
		responseTester, result, err := l.services.Tester(dataTester)
		if err != nil {
			responses.NewResponse(c).FailedResp(http.StatusInternalServerError, models.FailedMessage(err.Error()))
			return
		}
		if strings.Contains(result, "not") {
			responses.NewResponse(c).FailedResp(http.StatusUnauthorized, models.FailedMessage("not register"))
			return
		}
		responses.NewResponse(c).SuccessResp(http.StatusOK, models.SuccessMessage(fmt.Sprintf("%s%s", responseTester.JustString, "response"), gin.H{"result": result}))
	}
}

func NewLoginApi(routeGroup *gin.RouterGroup, testerService services.TesterService) {
	api := &tester{
		testerService,
	}
	eg := routeGroup.Group("/example")
	{
		eg.POST("/tester", api.Tester())
	}
}
