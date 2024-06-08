package starling

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/joshryandavis/songbird/starling/stmodels"
)

func (c *Client) GetFeedItems(a *stmodels.Account, t0 time.Time) ([]stmodels.FeedItem, error) {
	var ret []stmodels.FeedItem
	timeSince := stmodels.ParseTime(t0).String()
	url := BaseEndpoint(fmt.Sprintf("feed/account/%s/category/%s?accountUid=%s&categoryUid=%s&changesSince=%s", a.AccountUID, a.DefaultCategory, a.AccountUID, a.DefaultCategory, timeSince))
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper stmodels.FeedItems
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.FeedItems
	return ret, nil
}
