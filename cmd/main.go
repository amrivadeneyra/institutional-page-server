package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load("../configs/config.dev.yaml"); err != nil {
		logrus.Info("No .env file found")
	}


	bindingPort := 3000
}
