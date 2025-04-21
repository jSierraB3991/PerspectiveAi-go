package request

type Comment struct {
	Text string `json:"text"`
}

type PerspectiveRequest struct {
	Comment             Comment                `json:"comment"`
	Languages           []string               `json:"languages"`
	RequestedAttributes map[string]interface{} `json:"requestedAttributes"`
}
