package main

import (
	"github.com/safayildirim/wallet-management-service/app"
	"github.com/safayildirim/wallet-management-service/pkg/logger"
)

func main() {
	err := app.New().Run()
	if err != nil {
		logger.Zap.Sugar().Fatal(err)
	}
}
