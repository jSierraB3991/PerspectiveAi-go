package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jSierraB3991/PerspectiveAi-go/domain/models"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/request"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/service"
	"github.com/labstack/echo/v4"
)

type HubController struct {
	hub                *models.Hub
	upgrader           *websocket.Upgrader
	perspectiveService *service.PerspectiveService
}

func NewHubController(hub *models.Hub, perspectiveService *service.PerspectiveService) *HubController {

	// Actualizador de websocket
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &HubController{hub: hub, upgrader: &upgrader, perspectiveService: perspectiveService}
}

func (ctrl *HubController) ServeWsHanlder() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "name query param is required")
		}

		conn, err := ctrl.upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		client := &models.Client{
			Name: name,
			Conn: conn,
			Send: make(chan []byte, 256),
		}

		ctrl.hub.Register <- client

		go client.WritePump()

		return nil
	}
}

func (ctrl *HubController) NotifyHandler(c echo.Context) error {

	var req request.NotifyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	message, err := ctrl.perspectiveService.Analyze(message)
	if err != nil {
		log.Println(err)
		return err
	}

	ctrl.hub.SendTo <- models.TargetedMessage{
		Name:    req.Name,
		Message: []byte(message),
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "notificación enviada"})

}
