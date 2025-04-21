package main

import (
	"github.com/jSierraB3991/PerspectiveAi-go/domain/libs"
	"github.com/jSierraB3991/PerspectiveAi-go/domain/models"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/controllers"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()
	e := echo.New()
	hub := models.NewHub()
	go hub.Run()

	route(e, hub)

	e.Logger.Fatal(e.Start(":5000"))
}

func route(e *echo.Echo, hub *models.Hub) {
	perspectiveService := service.NewPerspectiveService(libs.NewEnviroment())
	ctrl := controllers.NewHubController(hub, perspectiveService)
	e.GET("/ws", ctrl.ServeWsHanlder())

	e.POST("/notify", ctrl.NotifyHandler)
}
