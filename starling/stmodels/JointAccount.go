package stmodels

type JointAccount struct {
	AccountHolderUid string `json:"accountHolderUid"`
	PersonOne        struct {
		Title       string `json:"title"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		DateOfBirth string `json:"dateOfBirth"`
		Email       string `json:"email"`
		Phone       string `json:"phone"`
	} `json:"personOne"`
	PersonTwo struct {
		Title       string `json:"title"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		DateOfBirth string `json:"dateOfBirth"`
		Email       string `json:"email"`
		Phone       string `json:"phone"`
	} `json:"personTwo"`
}
