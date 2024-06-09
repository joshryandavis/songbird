package client

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"time"
)

type Note struct {
	Updated     time.Time `json:"updated"`
	Split       bool      `json:"split"`
	Refund      bool      `json:"refund"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Comment     string    `json:"note"`
}

func UpdateNote(n Note, ac *Api, tUID string, categoryUID string, newNote Note) (Note, error) {
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
	err = ac.Client.UpdateUserNote(&ac.Account, categoryUID, tUID, string(noteJson))
	return rec, err
}
