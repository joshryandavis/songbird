package stmodels

type AssociatedFeedRoundUp struct {
	GoalCategoryUid string            `json:"goalCategoryUid"`
	Amount          CurrencyAndAmount `json:"amount"`
}
