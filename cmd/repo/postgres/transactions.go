package postgres

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
