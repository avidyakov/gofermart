package postgres

import "gophermart/cmd/repo"

func (r *Repo) CreateOrder(number string, userID uint) (orderID uint, err error) {
	var order Order
	_ = r.db.Where("number = ?", number).First(&order)
	if order.UserID == userID {
		return 0, repo.ErrOrderAlreadyUploaded
	} else if order.ID != 0 {
		return 0, repo.ErrOrderExists
	}

	order = Order{
		Number: number,
		UserID: userID,
		Status: New,
	}
	dbc := r.db.Create(&order)
	if dbc.Error != nil {
		return 0, dbc.Error
	}
	return order.ID, nil
}

func (r *Repo) GetOrders(userID uint) ([]repo.Order, error) {
	var orders []Order
	dbc := r.db.Where("user_id = ?", userID).Order("id").Find(&orders)
	if dbc.Error != nil {
		return nil, dbc.Error
	}

	var result []repo.Order
	for _, order := range orders {
		result = append(result, repo.Order{
			Number:    order.Number,
			CreatedAt: order.CreatedAt,
			Status:    order.Status.String(),
			Accrual:   order.Accrual,
		})
	}
	return result, nil
}
