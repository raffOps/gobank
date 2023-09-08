package main

import (
	"bank/accounts"
	"bank/clients"
	"fmt"
	"os"
)

func main() {
	titularRafael, err := clients.NewTitular("Rafael", "12345667778", "Lawyer")
	accountRafael, _ := accounts.NewCurrentAccount(titularRafael, 765, 767867, 100.0)

	titularMaria, err := clients.NewTitular("Maria", "1234566734", "Doctor")
	accountMaria, _ := accounts.NewCurrentAccount(titularMaria, 435, 766, 300)

	valueWithdraw := 2.0
	fmt.Printf("%+v\n", accountRafael)
	fmt.Printf("Sacando %.2f da conta %+v\n", valueWithdraw, accountRafael)
	err = accountRafael.Withdraw(valueWithdraw)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", accountRafael)

	valueToTransfer := 2.0
	fmt.Printf("\nTransferindo %.2f da conta %+v para conta %+v\n", valueToTransfer, accountRafael, accountMaria)
	err = accountRafael.Transfer(valueToTransfer, accountMaria)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("\n%+v\n", accountRafael)
	fmt.Printf("%+v\n", accountMaria)
	err = accountMaria.Deposit(300)
	if err != nil {
		return
	}
	fmt.Printf("Saldo da Maria: %.2f\n", accountRafael.GetBalance())

	savingsAccountRafael, _ := accounts.NewSavingsAccount(titularRafael, 765, 7676567, 5, 200.0)
	err = savingsAccountRafael.Deposit(200)
	if err != nil {
		os.Exit(1)
		return
	}
	fmt.Printf("Saldo da poupança Rafael: %.2f\n", savingsAccountRafael.GetBalance())

	err = accounts.PayBill(savingsAccountRafael, 30999999)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Saldo da poupança Rafael: %.2f\n", savingsAccountRafael.GetBalance())
}
