package mAccounts

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Account struct {
	ID            int    `gorm:"primary_key" json:"ID"`
	IdAccount     string `json:"id_account"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	AccountNumber int    `json:"account_number"`
	Saldo         int    `json:"saldo"`
}

type Transaction struct {
	ID                     int    `gorm:"primary_key" json:"ID"`
	TransactionType        string `json:"transaction_type"`
	TransactionDescription string `json:"transaction_description"`
	Sender                 string `json:"sender"`
	Amount                 int    `json:"amount"`
	Recipient              string `json:"recipient"`
	Timestamp              string `json:"timestamp"`
}

/*func (a *Account) IsValid() error {
	if a.ID == 0 {
		return errors.New(`Id may not be empty`)
	}
	if a.IdAccount == `` {
		return errors.New(`Id Account may not be emmpty`)
	}
	if a.Name == `` {
		return errors.New(`Name may not be emmpty`)
	}
	if a.Password == `` {
		return errors.New(`Name may not be emmpty`)
	}
	if a.AccountNumber == 0 {
		return errors.New(`Account Number may not be empty`)
	}
	if a.Saldo == 0 {
		return errors.New(`Account Number may not be empty`)
	}
	return nil
}

func (a *Account) IsValidLogin() error {

	if a.IdAccount == `` {
		return errors.New(`Id Account may not be emmpty`)
	}
	if a.Password == `` {
		return errors.New(`Name may not be emmpty`)
	}
	return nil
}*/
