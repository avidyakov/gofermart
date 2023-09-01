package postgres

func (r *Repo) CreateOrder(number string, userID uint) (orderID uint, err error) {
	order := Order{
		Number: number,
		UserID: userID,
	}
	dbc := r.db.Create(&order)
	if dbc.Error != nil {
		return 0, dbc.Error
	}
	return order.ID, nil
}
