package database

import (
	"cleanarchitecture/internal/entity"
	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT count(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FetchAll() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT * FROM orders;")
	if err != nil {
		return []entity.Order{}, err
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {
		var order entity.Order

		if err := rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return orders, err
		}

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return orders, err
	}

	return orders, nil
}
