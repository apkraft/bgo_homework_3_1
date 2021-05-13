package card

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Transactions []Transaction
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func SumByMcc(transactions []Transaction, mccs []string) int64 {

	var result int64 = 0

	for _, transaction := range transactions {
		for _, mcc := range mccs {
			if mcc == transaction.Mcc {
				result += transaction.Value
			}
		}
	}

	return result
}
