package accounts

import (
	"bank/clients"
	"bank/errors"
)

type SavingsAccount struct {
	titular                          clients.Titular
	agency, numberAccount, operation int32
	balance                          float64
}

func NewSavingsAccount(titular *clients.Titular, agency, numberAccount, operation int32, balance float64) (*SavingsAccount, error) {
	savingsAccount := SavingsAccount{titular: *titular, agency: agency, numberAccount: numberAccount, operation: operation, balance: balance}
	return &savingsAccount, nil
}

func (account *SavingsAccount) Withdraw(value float64) error {
	if value > account.balance {
		return &errors.InsufficientBalance{Value: value}
	} else if value <= 0 {
		return &errors.InvalidValue{Value: value}
	} else {
		account.balance -= value
		return nil
	}
}

func (account *SavingsAccount) Deposit(value float64) error {
	if value <= 0 {
		return &errors.InvalidValue{Value: value}
	}
	account.balance += value
	return nil
}

func (account *SavingsAccount) Transfer(value float64, accountToTransfer *CurrentAccount) error {
	err := account.Withdraw(value)
	if err != nil {
		return err
	}

	err = accountToTransfer.Deposit(value)
	return err
}

func (account *SavingsAccount) GetBalance() float64 {
	return account.balance
}
