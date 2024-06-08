package starling

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) GetFeedItems(a *Account, t0 time.Time) ([]FeedItem, error) {
	var ret []FeedItem
	timeSince := ParseTime(t0).String()
	url := BaseEndpoint(fmt.Sprintf("feed/account/%s/category/%s?accountUid=%s&categoryUid=%s&changesSince=%s", a.AccountUID, a.DefaultCategory, a.AccountUID, a.DefaultCategory, timeSince))
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	var wrapper FeedItems
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		return ret, err
	}
	ret = wrapper.FeedItems
	return ret, nil
}
