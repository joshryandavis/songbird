package starling

import (
	"encoding/json"
)

type SavingGoalOrdered struct {
	SavingGoalUID   string            `json:"savingsGoalUid"`
	Name            string            `json:"name"`
	Target          CurrencyAndAmount `json:"target"`
	TotalSaved      CurrencyAndAmount `json:"totalSaved"`
	SavedPercentage int               `json:"savedPercentage"`
	SortOrder       int               `json:"sortOrder"`
	State           string            `json:"state"`
}

type SpendingSpace struct {
	Name               string            `json:"name"`
	Balance            CurrencyAndAmount `json:"balance"`
	CardAssociationUid string            `json:"cardAssociationUid"`
	SortOrder          int               `json:"sortOrder"`
	SpendingSpaceType  string            `json:"spendingSpaceType"`
	State              string            `json:"state"`
	SpaceUid           string            `json:"spaceUid"`
}

type Spaces struct {
	SavingGoals    []SavingGoalOrdered `json:"savingGoals"`
	SpendingSpaces []SpendingSpace     `json:"spendingSpaces"`
}

func (c *Client) GetSpaces(a *Account) (Spaces, error) {
	var ret Spaces
	url := AccountEndpoint(a.AccountUID, SpacesEndpoint)
	res, err := c.Request("GET", url, "")
	if err != nil {
		return ret, err
	}
	err = json.Unmarshal(res, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
