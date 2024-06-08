package stmodels

import (
	"encoding/json"
	"time"
)

const DateTimeFormat = "2006-01-02T15:04:05Z"

type DateTime struct {
	time.Time
}

func ParseTime(t time.Time) DateTime {
	return DateTime{t}
}

//goland:noinspection GoMixedReceiverTypes
func (t *DateTime) UnmarshalJSON(data []byte) error {
	var err error
	var tStr string
	err = json.Unmarshal(data, &tStr)
	if err != nil {
		return err
	}
	parsedTime, err := time.Parse(DateTimeFormat, tStr)
	if err != nil {
		return err
	}
	*t = DateTime{parsedTime}
	return nil
}

//goland:noinspection GoMixedReceiverTypes
func (t DateTime) String() string {
	return t.Format(DateTimeFormat)
}
