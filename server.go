package main

import (
	"fmt"
	docs "myGoSkeleton/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log"
	"myGoSkeleton/config"
	"myGoSkeleton/controllers/api"
	"myGoSkeleton/manager"
	"os/exec"
)

type AppServer interface {
	Run()
}
type serverConfig struct {
	gin            *gin.Engine
	Name           string
	Port           string
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	ServiceManager manager.ServiceManager
	Config         *config.Config
}

func (s *serverConfig) initHeader() {
	s.gin.Use()
	s.routeGroupApi()
	s.ServiceManager.TesterService().Scheduler()
}

func (s *serverConfig) routeGroupApi() {
	apiLogin := s.gin.Group("/api/v1")
	api.NewLoginApi(apiLogin, s.ServiceManager.TesterService())
}

func (s *serverConfig) Run() {
	s.initHeader()
	s.gin.Run(fmt.Sprintf("%s:%s", s.Name, s.Port))
}

func Server() AppServer {
	ginStart := gin.Default()
	//swagger starts

	docs.SwaggerInfo.BasePath = "/api/v1/example"

	ginStart.GET("api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//swagger end
	config := config.NewConfig()
	infra := manager.NewInfraManager(config.ConfigDatabase)
	repo := manager.NewRepoManager(infra.PostgreConn())
	service := manager.NewServiceManager(repo)
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:8000/api/v1/swagger/index.html").Start()
	if err != nil {
		log.Fatal(err)
	}
	return &serverConfig{
		ginStart,
		config.ConfigServer.Url,
		config.ConfigServer.Port,
		infra,
		repo,
		service,
		config,
	}
}
