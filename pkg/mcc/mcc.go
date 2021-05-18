package mcc

func TranslateMCC(code string) (string, bool) {
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5812": "Рестораны",
		"5912": "Аптеки",
	}

	mccCodeName, ok := mcc[code]

	return mccCodeName, ok
}
