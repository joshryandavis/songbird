package starling

import (
	"encoding/json"
	"time"
)

const (
	DateTimeFormat = "2006-01-02T15:04:05Z"
	DateOnlyFormat = "2006-01-02"
)

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

type LastPayment struct {
	LastDate   DateOnly          `json:"lastDate"`
	LastAmount CurrencyAndAmount `json:"lastAmount"`
}

type DateOnly struct {
	time.Time
}

func (d *DateOnly) UnmarshalJSON(data []byte) error {
	var err error
	var dStr string
	err = json.Unmarshal(data, &dStr)
	if err != nil {
		return err
	}
	parsedTime, err := time.Parse(DateOnlyFormat, dStr)
	if err != nil {
		return err
	}
	*d = DateOnly{parsedTime}
	return nil
}
