package stmodels

type SpendingSpace struct {
	Name               string            `json:"name"`
	Balance            CurrencyAndAmount `json:"balance"`
	CardAssociationUid string            `json:"cardAssociationUid"`
	SortOrder          int               `json:"sortOrder"`
	SpendingSpaceType  string            `json:"spendingSpaceType"`
	State              string            `json:"state"`
	SpaceUid           string            `json:"spaceUid"`
}
