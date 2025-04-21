package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jSierraB3991/PerspectiveAi-go/domain/libs"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/request"
	"github.com/jSierraB3991/PerspectiveAi-go/infrastructure/response"
)

type PerspectiveService struct {
	PerspectiveAPIKey string
}

func NewPerspectiveService(env *libs.Enviroment) *PerspectiveService {
	return &PerspectiveService{PerspectiveAPIKey: env.PerspectiveAPIKey}
}

func (s *PerspectiveService) Analyze(text string) (string, error) {
	url := "https://commentanalyzer.googleapis.com/v1alpha1/comments:analyze"
	apiKey := s.PerspectiveAPIKey
	requestBody := request.PerspectiveRequest{
		Comment: request.Comment{
			Text: text,
		},
		Languages: []string{"es"},
		RequestedAttributes: map[string]interface{}{
			"TOXICITY": map[string]interface{}{},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url+"?key="+apiKey, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var result response.PerspectiveResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	score := result.AttributeScores.Toxicity.SummaryScore.Value

	if score > 0.7 {
		return "*****", nil
	}

	return text, nil

}
