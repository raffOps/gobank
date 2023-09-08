package accounts

import (
	"bank/clients"
	"bank/errors"
)

type CurrentAccount struct {
	titular               clients.Titular
	agency, numberAccount int32
	balance               float64
}

func NewCurrentAccount(titular *clients.Titular, agency, numberAccount int32, balance float64) (*CurrentAccount, error) {
	currentAccount := CurrentAccount{*titular, agency, numberAccount, balance}
	return &currentAccount, nil
}

func (account *CurrentAccount) Withdraw(value float64) error {
	if value > account.balance {
		return &errors.InsufficientBalance{Value: value}
	} else if value <= 0 {
		return &errors.InvalidValue{Value: value}
	} else {
		account.balance -= value
		return nil
	}
}

func (account *CurrentAccount) Deposit(value float64) error {
	if value <= 0 {
		return &errors.InvalidValue{Value: value}
	}
	account.balance += value
	return nil
}

func (account *CurrentAccount) Transfer(value float64, accountToTransfer *CurrentAccount) error {
	err := account.Withdraw(value)
	if err != nil {
		return err
	}

	err = accountToTransfer.Deposit(value)
	return err
}

func (account *CurrentAccount) GetBalance() float64 {
	return account.balance
}
