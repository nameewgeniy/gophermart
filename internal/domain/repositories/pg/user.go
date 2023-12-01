package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"gophermart/internal/domain/models"
	"strings"
	"time"
)

func (p Pg) CreateUser(user models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tr, err := p.db.BeginTx(ctx, nil)
	defer func() {
		_ = tr.Rollback()
	}()

	if err != nil {
		return fmt.Errorf("pg: CreateUser: Begin Transaction: %w", err)
	}

	//CREATE TABLE users
	//(
	//	id            UUID UNIQUE PRIMARY KEY,
	//	login         VARCHAR(100) UNIQUE NOT NULL,
	//	password_hash VARCHAR(255)        NOT NULL,
	//	balance       BIGINT,
	//	created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	//	deleted_at    TIMESTAMP
	//);

	query := "INSERT INTO #table# as t (id, login, password_hash, balance) VALUES ($1, $2, $3, $4)"
	preparedQuery := strings.NewReplacer("#table#", p.usersTable).Replace(query)

	if _, err = tr.ExecContext(ctx, preparedQuery, user.Id.String(), user.Login, user.PasswordHash, user.Balance.Amount()); err != nil {
		return fmt.Errorf("pg: CreateUser: %w", err)
	}

	return tr.Commit()
}

func (p Pg) FindUser(id uuid.UUID) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res := models.User{}
	query := "SELECT * FROM #table# WHERE id = $1"
	preparedQuery := strings.NewReplacer("#table#", p.usersTable).Replace(query)

	rows := p.db.QueryRowContext(ctx, preparedQuery, id)

	//CREATE TABLE users
	//(
	//	id            UUID UNIQUE PRIMARY KEY,
	//	login         VARCHAR(100) UNIQUE NOT NULL,
	//	password_hash VARCHAR(255)        NOT NULL,
	//	balance       BIGINT,
	//	created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	//	deleted_at    TIMESTAMP
	//);

	if err := rows.Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &res, sql.ErrNoRows // TODO свою ошибку
		}

		return &res, fmt.Errorf("pg: FindUser: id=%s: %w", id.String(), err)
	}

	if err := rows.Err(); err != nil {
		return &res, fmt.Errorf("pg: FindUser: Rows: id=%s: %w", id.String(), err)
	}

	return &res, nil
}

func (p Pg) FindUserByLogin(login string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM #table# WHERE login = $1"
	preparedQuery := strings.NewReplacer("#table#", p.usersTable).Replace(query)

	rows := p.db.QueryRowContext(ctx, preparedQuery, login)

	row := struct {
		Id           string
		Login        string
		PasswordHash string
		Balance      uint64
		CreatedAt    time.Time
		DeletedAt    *time.Time
	}{}

	if err := rows.Scan(&row.Id, &row.Login, &row.PasswordHash, &row.Balance, &row.CreatedAt, &row.DeletedAt); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows // TODO свою ошибку
		}

		return nil, fmt.Errorf("pg: FindUserByLogin: login=%s: %w", login, err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("pg: FindUserByLogin: Rows: login=%s: %w", login, err)
	}

	uid, err := uuid.Parse(row.Id)

	if err != nil {
		return nil, fmt.Errorf("pg: FindUserByLogin: uuid parse: %w", err)
	}

	return &models.User{
		Id:           uid,
		Login:        row.Login,
		PasswordHash: row.PasswordHash,
		Balance:      *money.New(int64(row.Balance), money.RUB),
		CreatedAt:    row.CreatedAt,
		DeletedAt:    row.DeletedAt,
	}, nil
}
