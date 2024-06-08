package client

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/client/models"
	"github.com/joshryandavis/songbird/starling"
	"github.com/joshryandavis/songbird/starling/stmodels"
)

func UpdateNote(n models.Note, c *starling.Client, acc stmodels.Account, tUID string, categoryUID string, newNote models.Note) (models.Note, error) {
	var rec = n
	if newNote == n {
		log.WithFields(log.Fields{
			"newNote": newNote,
			"oldNote": n,
		}).Printf("note is the same, skipping update")
		return rec, nil
	}
	rec = newNote
	noteJson, err := json.Marshal(rec)
	if err != nil {
		return rec, err
	}
	log.Println("updating note for client", tUID)
	err = c.PutFeedItemNote(&acc, categoryUID, tUID, string(noteJson))
	return rec, err
}

func (c *Client) GetSpendingCategories() ([]string, error) {
	return nil, nil
}
