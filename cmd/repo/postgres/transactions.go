package postgres

import "gophermart/cmd/repo"

func (r *Repo) GetBalance(userID uint) (float64, error) {
	var balance float64
	dbc := r.db.Table("transactions").Joins("inner join orders on transactions.order_id = orders.id").Where("orders.user_id = ?", userID).Select("COALESCE(sum(transactions.amount), 0)").Scan(&balance)
	if dbc.Error != nil {
		return 0, dbc.Error
	}
	return balance, nil
}

func (r *Repo) GetUsed(userID uint) (float64, error) {
	var used float64
	dbc := r.db.Table("transactions").Joins("inner join orders on transactions.order_id = orders.id").Where("orders.user_id = ?", userID).Select("COALESCE(sum(transactions.amount), 0)").Where("transactions.amount < 0").Scan(&used)
	if dbc.Error != nil {
		return 0, dbc.Error
	}
	if used < 0 {
		used = -used
	}
	return used, nil
}

func (r *Repo) MakeTransaction(orderNumber string, amount float64) error {
	var order Order
	dbc := r.db.Where("number = ?", orderNumber).First(&order)
	if dbc.Error != nil {
		return dbc.Error
	}
	dbc = r.db.Create(&Transaction{
		OrderID: order.ID,
		Amount:  amount,
	})
	if dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (r *Repo) GetWithdrawals(userID uint) ([]repo.Withdrawal, error) {
	var withdrawals []repo.Withdrawal
	dbc := r.db.Table("transactions").
		Joins("inner join orders on transactions.order_id = orders.id").
		Where("orders.user_id = ? AND transactions.amount < 0", userID).
		Select("transactions.order_id, transactions.amount, transactions.created_at").
		Scan(&withdrawals)

	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return withdrawals, nil
}
