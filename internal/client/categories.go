package client

import (
	"strings"

	"github.com/joshryandavis/songbird/starling"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/internal/client/models"
)

func (c *Client) autoCategorise(ac *starling.Client, acc starling.Account, newT *models.Transaction) error {
	for _, cat := range c.Cfg.Categories {
		if cat.Parent != newT.SpendingCategory {
			continue
		}
		for _, name := range cat.Auto {
			if normaliseCat(name) == normaliseCat(newT.CounterParty.Name) {
				newNote := newT.Note
				newNote.Category = cat.Name
				if newT.Note.Category == newNote.Category {
					continue
				}
				log.Println("new found match", newT.CounterParty.Name, "in", cat.Name)
				_, err := UpdateNote(newT.Note, ac, acc, newT.UID, newT.CategoryUID, newNote)
				if err != nil {
					log.Fatal("error updating note", err)
					return err
				}
				break
			}
		}
	}
	return nil
}

func normaliseCat(cat string) string {
	cat = strings.ToLower(cat)
	cat = strings.Replace(cat, "-", "", -1)
	cat = strings.Replace(cat, "_", "", -1)
	cat = strings.Replace(cat, "'", "", -1)
	cat = strings.Replace(cat, ".", "", -1)
	cat = strings.Trim(cat, " ")
	return cat
}
