package main

import (
	"github.com/jSierraB3991/PerspectiveAi-go/domain/models"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	hub := models.NewHub()
	go hub.Run()

	route(e, hub)

	e.Logger.Fatal(e.Start(":5000"))
}

func route(e *echo.Echo, hub *models.Hub) {
	ctrl := controllers.NewHubController(hub)
	e.GET("/ws", ctrl.ServeWsHanlder())

	e.POST("/notify", ctrl.NotifyHandler)
}
