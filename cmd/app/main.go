package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maratov-nursultan/Kubernetes/internal/config"
	"github.com/maratov-nursultan/Kubernetes/internal/database"
	"github.com/maratov-nursultan/Kubernetes/internal/handler"
	"github.com/maratov-nursultan/Kubernetes/internal/service"
	"log"
)

func main() {
	cfg := config.Get()

	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	factory := service.NewService(db)
	if err != nil {
		log.Fatalf("Kaput factory is not working %v", err)
	}

	newHandler := handler.NewHandler(factory.GetUserManager())

	e := echo.New()

	userGroup := e.Group("/user")
	userGroup.GET("get", newHandler.GetUser)
	userGroup.POST("create", newHandler.CreateUser)
	userGroup.DELETE("delete", newHandler.DeleteUser)
	userGroup.GET("update", newHandler.UpdateUser)

	e.Logger.Fatal(e.Start(":80"))
}
