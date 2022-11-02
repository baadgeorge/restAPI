package main

import (
	"github.com/sirupsen/logrus"
	"test"
	"test/internal/app"
	"test/internal/repository"
	"test/internal/service"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(new(logrus.TextFormatter))
	logrus.SetReportCaller(true)

	DB, err := repository.NewDB()
	if err != nil {
		logrus.Fatal(err)
		return
	}
	repo := repository.NewRepository(DB)
	signingKey := "Zq4t7w!z%C*F-JaNdRgUkXp2s5u8x/A?"
	tokenService := service.NewTokenService(signingKey)
	dataService := service.NewService(repo, tokenService)
	handler := app.NewHandler(dataService)

	server := new(test.Server)
	if err := server.Run("9000", handler.Router()); err != nil {
		logrus.Fatalf("error while running server: %s", err)
	}

	logrus.Print("server started")

}
