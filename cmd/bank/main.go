package main

import (
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	}

	fmt.Println(cards)
}