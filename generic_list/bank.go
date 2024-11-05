package genericlist

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  sum,
	}
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	adjustBalance := func(balance float64, t Transaction) float64 {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
		return balance
	}
	return Reduce(transactions, adjustBalance, 0)
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	applyTransaction := func(target Account, t Transaction) Account {
		if t.From == target.Name {
			target.Balance -= t.Sum
		}
		if t.To == target.Name {
			target.Balance += t.Sum
		}
		return target
	}
	return Reduce(transactions, applyTransaction, account)
}
