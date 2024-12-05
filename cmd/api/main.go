package main

import (
	"colegio/server/common/utils"
	colegioapi "colegio/server/svc/api"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load("../../environment/.env"); err != nil {
		logrus.Info("No .env file found")
	}

	bindingPort := 3000
	router := colegioapi.NewRouter()
	if utils.GetStage() == utils.Local {
		logrus.Info("Starting API server")
		logrus.Infof("Listening on port for lambda %d", bindingPort)
		server := &http.Server{
			Addr:              fmt.Sprintf(":%d", bindingPort),
			ReadHeaderTimeout: 0,
			Handler:           router,
		}
		defer server.Close()
		server.ListenAndServe()
	}
	adapter := chiadapter.New(router)
	lambda.Start(adapter.ProxyWithContext)
}
