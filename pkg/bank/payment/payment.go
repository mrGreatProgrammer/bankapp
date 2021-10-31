package payment

import (
	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {
	var max types.Money = 0
	// var theid int
	var myId int


	for i, v := range payments {
		if v.Amount > max {
			max = v.Amount
			// theid = v.ID
			myId = i
			// fmt.Println("myId", myId)
			// fmt.Println(i)
		}
	}
	// fmt.Println(i)
	// fmt.Println("max",max)
	// fmt.Println(theid)
	return payments[myId]
}
