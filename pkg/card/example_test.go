package card

import (
	// "bank/pkg/bank/card"
	"bank/pkg/bank/types"
	// "bank/pkg/bank/card"
	"fmt"
)

// func ExampleWithdraw() {

// 	card := types.Card{Balance: 20_000_00, Active: true}
// 	Withdraw(&card, 10_000_00)

// 	fmt.Println(card.Balance)

// 	// result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
// 	// fmt.Println(result.Balance)

// 	// Output:
// 	// 1000000

// }

// func ExampleWithdraw_noMoney(){
// 	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
// 	fmt.Println(result.Balance)

// 	// Output:
// 	// 0

// }

// func ExampleWithdraw_inactive(){
// 	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
// 	fmt.Println(result.Active)

// 	// Output:
// 	// false

// }

// func ExampleWithdraw_limit(){
// 	result := Withdraw(types.Card{Balance: 40_000_00, Active: true}, 40_000_00)
// 	fmt.Println(result.Balance)

// 	// Output:
// 	// 4000000

// }

// func ExampleDeposit() {
// 	card := types.Card{Balance: 20_000_00, Active: true}
// 	Deposit(&card, 10_000_00)

// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 3000000
// }

// func ExampleDeposit_inactive() {
// 	card := types.Card{Balance: 20_000_00, Active: false}
// 	Deposit(&card, 10_000_00)

// 	fmt.Println(card.Active)

// 	// Output:
// 	// false
// }

// func ExampleDeposit_limit() {
// 	card := types.Card{Balance: 20_000_00, Active: true}
// 	Deposit(&card, 51_000_00)

// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 2000000
// }

// func ExampleAddBonus_positive() {
// 	card := types.Card{MinBalance: 10_000_00, Active: true}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.Balance)

// 	// Output:
// 	// 2465
// }

// func ExampleAddBonus_negativeBalance() {
// 	card := types.Card{MinBalance: -1, Active: true}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.Balance)

// 	// Output:
// 	// -1
// }

// func ExampleAddBonus_inactive() {
// 	card := types.Card{MinBalance: 10_000_00, Active: false}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.MinBalance)

// 	// Output:
// 	// 1000000
// }

// func ExampleAddBonus_limitEqual(){
// 	card := types.Card{MinBalance: 5_000_00, Active: true}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.MinBalance)

// 	// Output:
// 	// 500000
// }

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Balance: -10_000_00,
			Active:  false,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}

	fmt.Println(Total(cards))
	//Output: 2000000
}

func ExamplePaymentSources() {

	cards := []types.Card{
		{
			Balance: 999_99,
			Active:  true,
		},
		{
			Balance: 888_88,
			Active:  true,
		},
		{
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
	}

	PaymentSources(cards)

	//Output:
	// 	5058 xxxx xxxx 99999
	// 5058 xxxx xxxx 88888

}
