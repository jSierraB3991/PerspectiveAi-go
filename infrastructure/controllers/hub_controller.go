package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jSierraB3991/PerspectiveAi-go/domain/models"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/request"
	"github.com/labstack/echo/v4"
)

type HubController struct {
	hub      *models.Hub
	upgrader *websocket.Upgrader
}

func NewHubController(hub *models.Hub) *HubController {

	// Actualizador de websocket
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &HubController{hub: hub, upgrader: &upgrader}
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

	ctrl.hub.SendTo <- models.TargetedMessage{
		Name:    req.Name,
		Message: []byte(req.Message),
	}
	ctrl.hub.SendTo <- models.TargetedMessage{
		Name:    c.QueryParam("name"),
		Message: []byte(req.Message),
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "notificación enviada"})

}
