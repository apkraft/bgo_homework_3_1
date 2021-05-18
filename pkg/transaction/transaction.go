package transaction

type Transaction struct {
	Id          int64
	Value       int64
	Datetime    int64
	Mcc         string
	Status      string
	Description string
	Icon        string
}
