package main

import "fmt"

type account struct {
	titular       string
	agency        int32
	accountNumber int32
	balance       float64
}

func (account *account) withdraw(value float64) bool {
	if value >= account.balance || value <= 0 {
		return false
	} else {
		account.balance -= value
		return true
	}
}

func main() {
	accountRafael := account{"Rafael", 765, 767867, 100.0}
	valueWithdraw := -2.0
	fmt.Printf("%+v\n", accountRafael)
	accountRafael.withdraw(valueWithdraw)
	fmt.Printf("%+v\n", accountRafael)
}
