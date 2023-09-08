package accounts

type Account interface {
	Withdraw(value float64) error
}

func PayBill(account Account, value float64) error {
	err := account.Withdraw(value)
	return err
}
