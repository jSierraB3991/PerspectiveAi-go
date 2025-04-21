package libs

import (
	"log"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

type Enviroment struct {
	PerspectiveAPIKey string
}

func NewEnviroment() *Enviroment {
	perspectiveAPIKey, err := jsierralibs.GetDataOfEnviromentRequired("API_PERSPECTTIVE_KEY")
	if err != nil {
		log.Fatal(err)
	}

	return &Enviroment{
		PerspectiveAPIKey: perspectiveAPIKey,
	}
}
