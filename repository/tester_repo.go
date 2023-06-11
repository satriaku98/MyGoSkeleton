package repository

import (
	"encoding/json"
	"fmt"
	configs "myGoSkeleton/config"
	"myGoSkeleton/controllers/request"
	"myGoSkeleton/controllers/responses"
	"myGoSkeleton/models"
	"myGoSkeleton/utility"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
)

type TesterRepo interface {
	Tester(testerData request.TesterRequest) (responses.TesterResponse, string, error)
	Scheduler()
}

type testerRepo struct {
	sqlDb *sqlx.DB
}

func (a *testerRepo) Tester(testerData request.TesterRequest) (responses.TesterResponse, string, error) {
	var dataTester responses.TesterResponse

	//Database Query Execute
	err := a.sqlDb.Get(&dataTester, "select JustString from Test2 where id = $1", testerData.JustString)

	//HTTP Request
	//Create Request Message
	request := &models.OtherApiRequest{
		Data: models.OtherApiRequestData{
			AccountFrom:    testerData.JustString,
			InstitutionID:  "777000028",
			NdcProductCode: "210000",
		},
		Message:   nil,
		IsSuccess: true,
	}
	requestMarshaler, _ := json.Marshal(request)
	getConfig := configs.NewConfig().ConfigApp
	// params := map[string]string{"key1": "val1", "key2": "val2"}
	//var user1, errPost = HttpPost(requestMarshaler, baseURL, params)
	var user1, errPost = utility.HttpPost(requestMarshaler, getConfig.OTHER_API_URL, nil)
	if errPost != nil {
		fmt.Println("Error!", err.Error())
		reqq := string(requestMarshaler)
		return dataTester, reqq, err
	}
	resString, _ := json.Marshal(user1)
	fmt.Println("Response", string(resString))

	//Logic
	result := utility.CheckStrings(testerData.JustString, dataTester.JustString)
	//Error handling
	if err != nil {
		return dataTester, result, err
	}
	//success return
	return dataTester, result, nil
}

func (a *testerRepo) Scheduler() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Tag("tag").Do(utility.Task)
	s.StartAsync()
}

func NewTesterRepo(sqlDb *sqlx.DB) TesterRepo {
	return &testerRepo{
		sqlDb,
	}
}
