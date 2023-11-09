package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/kovalyov-valentin/orders-service/internal/repository/cache"
	"github.com/kovalyov-valentin/orders-service/internal/models"
	"github.com/sirupsen/logrus"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{
		db: db,
	}
}

var (
	shemaOrder = goqu.S("public")
	tableOrder = shemaOrder.Table("orders")
)

const (
	orderColUID  = "order_uid"
	orderColData = "data"
)

func (r *OrderPostgres) Create(ctx context.Context, orderUID string, order models.Order) error {
	record := goqu.Insert(tableOrder).
		Rows(goqu.Record{
			orderColUID:  order.OrderUID,
			orderColData: order,
		})

	query, _, err := record.ToSQL()
	if err != nil {
		logrus.Println("failed to generate insert")
		return err
	}

	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrderPostgres) GetOrderByUID(ctx context.Context, uid string) (models.Order, error) {
	var order models.Order
	record := goqu.From(tableOrder).
		Select(
			tableOrder.Col(orderColData),
		).
		Where(tableOrder.Col(orderColUID).Eq(uid))

	query, _, err := record.ToSQL()
	if err != nil {
		return models.Order{}, err
	}

	err = r.db.QueryRowContext(ctx, query).
		Scan(&order)

	if err == sql.ErrNoRows {
		logrus.Println("error: no order with whis id")
		return models.Order{}, err
	}

	if err != nil {
		logrus.Println("error select order")
		return models.Order{}, err
	}

	return order, nil
}

func (r *OrderPostgres) CacheRecovery(cache *cache.Cache) error {

	query := goqu.From(tableOrder).
		Select(tableOrder.Col(orderColData))

	rows, err := r.db.Query(query.ToSQL())
	if err != nil {
		logrus.Println("error restoring cache frompostgres")
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order); err != nil {
			logrus.Println("error scan db to go postgres")
			return err
		}

		if err = cache.Set(&order); err != nil {
			logrus.Println("error set data into cache")
			return err
		}
	}

	return nil
}
