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
	var query string
	//Database Query Execute
	err := a.sqlDb.Get(&dataTester, "select JustString from Test2 where id = $1", testerData.JustString)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return dataTester, "", err
	}

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
	//Save Log to Database
	query = fmt.Sprintf(`INSERT INTO public.table_log (request, response, log_date) VALUES('%s', '%s', now());`, string(requestMarshaler), "")
	_, err = a.sqlDb.Exec(query)
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
	//Save Log to Database
	query = fmt.Sprintf(`INSERT INTO public.table_log (request, response, log_date) VALUES('%s', '%s', now());`, "", string(resString))
	_, err = a.sqlDb.Exec(query)

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
