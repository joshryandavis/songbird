package models

import "time"

type Note struct {
	Updated     time.Time `json:"updated"`
	Split       bool      `json:"split"`
	Refund      bool      `json:"refund"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Comment     string    `json:"note"`
}
