package card

import (
	"bank/pkg/bank/types"
	"strconv"
	"strings"
	// "fmt"
)

// const withdrawLimit = 20_000_00
// const depositLimit = 50_000_00
// const bonusLimit = types.Money(5_000_00)

// func IssueCard(currency types.Currency, color string, name string) types.Card {
// 	card := types.Card{
// 		ID:         1000,
// 		PAN:        "5058 xxxx xxxx 0001",
// 		Balance:    0,
// 		MinBalance: 10_000_00,
// 		Currency:   currency,
// 		Color:      color,
// 		Name:       name,
// 		Active:     true,
// 	}

// 	// fmt.Printf("%+v", card)
// 	return card
// }

// func Withdraw(card *types.Card, amount types.Money) {
// 	if amount < 0 {
// 		return
// 	}
// 	if amount > withdrawLimit {
// 		return
// 	}
// 	if !card.Active {
// 		return
// 	}
// 	if card.Balance < amount {
// 		return
// 	}

// 	(*card).Balance -= amount
// }

// func Deposit(card *types.Card, amount types.Money) {
// 	if amount < 0 {
// 		return
// 	}
// 	if amount > depositLimit {
// 		return
// 	}
// 	if !card.Active {
// 		return
// 	}

// 	(*card).Balance += amount
// }

// func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
// 	turnToPercent := float64(percent) / 100
// 	bonus := float64(card.MinBalance) * float64(turnToPercent) * float64(daysInMonth) / float64(daysInYear)

// 	if (*card).MinBalance < 0 {
// 		(*card).Balance = card.MinBalance
// 		return
// 	} else if !card.Active {
// 		return
// 	} else if bonus > float64(bonusLimit) {
// 		return
// 	} else if card.Balance < 0 {
// 		return
// 	} else {
// 		(*card).Balance += types.Money(bonus)
// 	}

// }


func Total(cards []types.Card) types.Money {
	// var totalSum
	var t types.Money
	for _, v := range cards {
		if !v.Active {
			t +=0
		} else if v.Balance < 0 {
			t+=0
		} else{
			t += v.Balance
		}
		// t += v.Balance
		// fmt.Println(v.Balance)
	}
	return t
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	PaymentSourceCards := []types.PaymentSource{}
	var identifyNumber = "01"
	var intoString string
	var putTheNymber string
i := 00
	for _, card := range cards {
		i++

		if card.Active && card.Balance > 0 {
			
			
			intoString = strconv.Itoa(int(i))
			putTheNymber = strings.Replace(identifyNumber, "1", intoString, 5)

			PaymentSourceCards = append(PaymentSourceCards, types.PaymentSource{
				Balance: card.Balance,
				Number:  putTheNymber,
				Type:    "card",
			})
		}
	}
	// var theIdNum string

	// for i := 0; i < len(paymentSourceCards); i++ {
	// 	theIdNum = paymentSourceCards[i].Number
	// 	fmt.Println(theIdNum)
	// }

	return PaymentSourceCards

}
