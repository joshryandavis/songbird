package stmodels

type Spaces struct {
	SavingGoals    []SavingGoalOrdered `json:"savingGoals"`
	SpendingSpaces []SpendingSpace     `json:"spendingSpaces"`
}
