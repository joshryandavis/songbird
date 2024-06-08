package client

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/client/models"
)

func WriteToTmp(output []models.Transaction) error {
	tmp, err := os.CreateTemp("tmp", "transactions-*.json")
	if err != nil {
		log.Fatal("error creating temp file", err)
		return err
	}
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		log.Fatal("error marshalling output to json", err)
		return err
	}
	_, err = tmp.WriteString(string(jsonOutput))
	if err != nil {
		log.Fatal("error writing json to temp file", err)
		return err
	}
	return nil
}
