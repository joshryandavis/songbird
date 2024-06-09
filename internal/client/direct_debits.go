package client

import (
	"time"

	"github.com/joshryandavis/songbird/internal/calendar"
	"github.com/joshryandavis/songbird/starling"
)

type DirectDebit struct {
	UID              string               `json:"uid"`
	Reference        string               `json:"reference"`
	Status           string               `json:"status"`
	Source           string               `json:"source"`
	Created          starling.DateTime    `json:"created"`
	Cancelled        starling.DateTime    `json:"cancelled"`
	LastDate         time.Time            `json:"lastDate"`
	OriginatorName   string               `json:"originatorName"`
	OriginatorUID    string               `json:"originatorUid"`
	MerchantUID      string               `json:"merchantUid"`
	LastPayment      starling.LastPayment `json:"lastPayment"`
	AccountUID       string               `json:"accountUid"`
	CategoryUID      string               `json:"categoryUid"`
	ProbableNextDate time.Time            `json:"probableNextDate"`
}

func (c *Client) GetDirectDebits(api *Api, acc *starling.Account) ([]DirectDebit, error) {
	var ret []DirectDebit
	dd, err := api.Client.GetDirectDebitMandates()
	if err != nil {
		return ret, err
	}
	for _, d := range dd {
		ret = append(ret, DirectDebit{
			UID:              d.UID,
			Reference:        d.Reference,
			Status:           d.Status,
			Source:           d.Source,
			Created:          d.Created,
			Cancelled:        d.Cancelled,
			LastDate:         d.LastDate.Time,
			OriginatorName:   d.OriginatorName,
			OriginatorUID:    d.OriginatorUID,
			MerchantUID:      d.MerchantUID,
			LastPayment:      d.LastPayment,
			AccountUID:       d.AccountUID,
			CategoryUID:      d.CategoryUID,
			ProbableNextDate: getDirectDebitProbableNextDate(c, d.LastDate.Time),
		})
	}
	return ret, nil
}

func getDirectDebitProbableNextDate(c *Client, lastDate time.Time) time.Time {
	nextDate := lastDate.AddDate(0, 1, 0)
	nextDate = calendar.GetBankWorkingDay(c.Cal, nextDate)
	return nextDate
}
