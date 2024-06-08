package starling

type SavingGoalOrdered struct {
	SavingGoalUID   string            `json:"savingsGoalUid"`
	Name            string            `json:"name"`
	Target          CurrencyAndAmount `json:"target"`
	TotalSaved      CurrencyAndAmount `json:"totalSaved"`
	SavedPercentage int               `json:"savedPercentage"`
	SortOrder       int               `json:"sortOrder"`
	State           string            `json:"state"`
}
