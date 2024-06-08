package starling

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) PutFeedItemNote(a *stmodels.Account, categoryUid string, itemUid string, note string) error {
	url := BaseEndpoint(fmt.Sprintf("feed/account/%s/category/%s/%s/user-note", a.AccountUID, categoryUid, itemUid))
	note = strings.Replace(note, "\"", "\\\"", -1)
	body := fmt.Sprintf("{\"userNote\":\"%s\"}", note)
	log.Println("Updating note for", itemUid, "to", note)
	_, err := c.Request("PUT", url, body)
	if err != nil {
		return err
	}
	return nil
}
