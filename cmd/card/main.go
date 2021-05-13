package main

import (
	"fmt"
	"time"

	"github.com/apkraft/bgo_homework_1_4/pkg/card"
)

func main() {

	master := &card.Card{
		Id:       12345,
		Issuer:   "The Bank",
		Balance:  100_000_00,
		Currency: "RUB",
		Number:   "1234 5678 9012 3456",
		Transactions: []card.Transaction{
			{
				Id:          1,
				Value:       -120191,
				Datetime:    1610627520,
				Mcc:         "5812",
				Status:      "done",
				Description: "Gastrobar",
				Icon:        "link",
			},
			{
				Id:          2,
				Value:       200000,
				Datetime:    1610645640,
				Mcc:         "",
				Status:      "done",
				Description: "Пополнение через Сбербанк",
				Icon:        "link",
			},
			{
				Id:          3,
				Value:       -73555,
				Datetime:    1610710260,
				Mcc:         "5411",
				Status:      "done",
				Description: "Gipermarket",
				Icon:        "link",
			},
		},
	}

	fmt.Printf("%v\n\n", master)

	newTransaction := &card.Transaction{
		Id:          4,
		Value:       100,
		Datetime:    time.Now().Unix(),
		Mcc:         "",
		Status:      "done",
		Description: "Новая транзакция",
		Icon:        "link",
	}

	fmt.Printf("%v\n\n", newTransaction)
	card.AddTransaction(master, newTransaction)
	fmt.Printf("%v\n\n", master)

	mccs := []string{"5812", "5411"}
	sumByMcc := card.SumByMcc(master.Transactions, mccs)

	fmt.Printf("Сумма транзакций с кодами MCC %v составляет %v\n\n", mccs, sumByMcc)

	for i := range master.Transactions {
		category, codeDeclared := card.TranslateMCC(master.Transactions[i].Mcc)
		if codeDeclared {
			fmt.Printf("Коду '%s' соответствует категория '%s'\n", master.Transactions[i].Mcc, category)
		} else {
			fmt.Printf("Для кода '%s' категория не указана\n", master.Transactions[i].Mcc)
		}
	}
}
