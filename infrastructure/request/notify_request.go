package request

type NotifyRequest struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
