package database

import (
	"database/sql"
	"github/pgabrieldeveloper/intensivo_go/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (o *OrderRepository) Save(order *entity.Order) error {
	_, err := o.Db.Exec("INSERT INTO orders (id, price, tax, final_price) values (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) GetTotal() (int, error) {
	var total int
	err := o.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
