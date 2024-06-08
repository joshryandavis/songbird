package starling

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Account struct {
	AccountUID      string   `json:"accountUid"`
	AccountType     string   `json:"accountType"`
	DefaultCategory string   `json:"defaultCategory"`
	Currency        string   `json:"currency"`
	CreatedAt       DateTime `json:"createdAt"`
	Name            string   `json:"name"`
}

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

func (c *Client) GetAccounts() ([]Account, error) {
	var ret []Account
	url := fmt.Sprintf("%s/%s", BaseUrlProd, AccountsEndpoint)
	res, err := c.Request("GET", url, "")
	log.Printf("request: %s", res)
	if err != nil {
		log.Panic("request error:", err)
		return ret, err
	}
	var wrapper Accounts
	err = json.Unmarshal(res, &wrapper)
	if err != nil {
		log.Panic("json unmarshal error:", err)
		return ret, err
	}
	ret = wrapper.Accounts
	return ret, nil
}

type AccountIdentifier struct {
	IdentifierType    string `json:"identifierType"`
	BankIdentifier    string `json:"bankIdentifier"`
	AccountIdentifier string `json:"accountIdentifier"`
}

func (c *Client) GetAccountIdentifiers(a *Account) (AccountIdentifiers, error) {
	var ret AccountIdentifiers
	url := AccountEndpoint(a.AccountUID, IdentifiersEndpoint)
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

type AccountIdentifiers struct {
	AccountIdentifier  string              `json:"accountIdentifier"`
	BankIdentifier     string              `json:"bankIdentifier"`
	IBAN               string              `json:"iban"`
	BIC                string              `json:"bic"`
	AccountIdentifiers []AccountIdentifier `json:"accountIdentifiers"`
}

type AccountHolder struct {
	AccountHolderUid  string `json:"accountHolderUid"`
	AccountHolderType string `json:"accountHolderType"`
}

type AccountHolderName struct {
	AccountHolderName string `json:"accountHolderName"`
}

func (c *Client) GetAccountHolder() (AccountHolder, error) {
	var ret AccountHolder
	url := BaseEndpoint(AccountHolderEndpoint)
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

func (c *Client) GetAccountHolderName() (AccountHolderName, error) {
	var ret AccountHolderName
	url := BaseEndpoint(AccountHolderEndpoint)
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
