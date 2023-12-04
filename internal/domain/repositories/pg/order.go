package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gophermart/internal/domain/models"
	"strings"
	"time"
)

//CREATE TABLE orders
//(
//id          UUID PRIMARY KEY,
//user_id     UUID NOT NULL,
//number      VARCHAR(255) UNIQUE NOT NULL,
//status      VARCHAR(255) NOT NULL,
//accrual     BIGINT,
//uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//FOREIGN KEY (user_id) REFERENCES users (id)
//);

type orderRow struct {
	Id         string
	UserId     string
	Number     string
	Status     string
	Accrual    *int
	UploadedAt time.Time
}

func (p Pg) CreateOrder(ord models.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tr, err := p.db.BeginTx(ctx, nil)
	defer func() {
		_ = tr.Rollback()
	}()

	if err != nil {
		return fmt.Errorf("pg: CreateOrder: Begin Transaction: %w", err)
	}

	query := "INSERT INTO #table# as t (id, user_id, number, status) VALUES ($1, $2, $3, $4)"
	preparedQuery := strings.NewReplacer("#table#", p.ordersTable).Replace(query)

	if _, err = tr.ExecContext(ctx, preparedQuery, ord.Id.String(), ord.UserId.String(), ord.Number, ord.Status); err != nil {
		return fmt.Errorf("pg: CreateOrder: %w", err)
	}

	return tr.Commit()
}

func (p Pg) FindOrder(id uuid.UUID) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, user_id, number, status, accrual, uploaded_at FROM #table# WHERE id = $1"
	preparedQuery := strings.NewReplacer("#table#", p.ordersTable).Replace(query)

	rows := p.db.QueryRowContext(ctx, preparedQuery, id.String())

	row := orderRow{}
	if err := rows.Scan(&row.Id, &row.UserId, &row.Number, &row.Status, &row.Accrual, &row.UploadedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows // TODO свою ошибку
		}

		return nil, fmt.Errorf("pg: FindOrder: id=%s: %w", id.String(), err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("pg: FindOrder: Rows: id=%s: %w", id.String(), err)
	}

	return p.mapRowToOrder(row)
}

func (p Pg) FindUserOrder(user uuid.UUID, number string) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, user_id, number, status, accrual, uploaded_at FROM #table# WHERE user_id = $1 and number = $2"
	preparedQuery := strings.NewReplacer("#table#", p.ordersTable).Replace(query)

	rows := p.db.QueryRowContext(ctx, preparedQuery, user.String(), number)

	row := orderRow{}
	if err := rows.Scan(&row.Id, &row.UserId, &row.Number, &row.Status, &row.Accrual, &row.UploadedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows // TODO свою ошибку
		}

		return nil, fmt.Errorf("pg: FindUserOrder: id=%s: %w", user.String(), err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("pg: FindUserOrder: Rows: id=%s: %w", user.String(), err)
	}

	return p.mapRowToOrder(row)
}

func (p Pg) FindUserOrders(user uuid.UUID) ([]models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, user_id, number, status, accrual, uploaded_at FROM #table# WHERE user_id = $1"
	preparedQuery := strings.NewReplacer("#table#", p.ordersTable).Replace(query)

	rows, err := p.db.QueryContext(ctx, preparedQuery, user.String())
	if err != nil {
		return nil, fmt.Errorf("pg: FindUserOrders: query execution failed: %w", err)
	}
	defer rows.Close()

	var userOrders []models.Order
	for rows.Next() {
		var o models.Order
		err = rows.Scan(&o.Id, &o.UserId, &o.Number, &o.Status, &o.Accrual, &o.UploadedAt)
		if err != nil {
			return nil, fmt.Errorf("pg: FindUserOrders: failed to scan row: %w", err)
		}
		userOrders = append(userOrders, o)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("pg: FindUserOrders: error iterating rows: %w", err)
	}

	return userOrders, nil
}

func (p Pg) mapRowToOrder(row orderRow) (*models.Order, error) {
	id, err := uuid.Parse(row.Id)

	if err != nil {
		return nil, fmt.Errorf("pg: mapRowToOrder: id parse: %w", err)
	}

	uid, err := uuid.Parse(row.UserId)

	if err != nil {
		return nil, fmt.Errorf("pg: mapRowToOrder: uid parse: %w", err)
	}

	status, err := models.GetOrderStatusByValue(row.Status)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		Id:         id,
		UserId:     uid,
		Number:     row.Number,
		Status:     status,
		Accrual:    row.Accrual,
		UploadedAt: row.UploadedAt,
	}, nil
}
