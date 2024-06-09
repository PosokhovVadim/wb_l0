package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"order/internal/model"
	"order/internal/storage"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(postgresPath string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", postgresPath)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) CreateOrder(ctx context.Context, order_uid uuid.UUID, order model.Order) error {
	query := `
		INSERT INTO orders
		(order_uid, order_data)
		VALUES
		($1, $2)
		ON CONFLICT (order_uid) DO NOTHING
	`

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	res, err := s.db.ExecContext(ctx, query, order_uid, orderJSON)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return storage.ErrOrderAlreadyExists
	}

	return nil
}

func (s *PostgresStorage) GetOrder(ctx context.Context, uuid uuid.UUID) (*model.Order, error) {
	query := `
		SELECT order_data
		FROM orders
		WHERE order_uid = $1
	`
	row := s.db.QueryRowContext(ctx, query, uuid)

	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return nil, storage.ErrOrderNotFound
		}
		return nil, err
	}

	var orderData []byte
	err := row.Scan(&orderData)
	if err != nil {
		return nil, err
	}

	var order model.Order
	err = json.Unmarshal(orderData, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil

}

func (s *PostgresStorage) DeleteOrder(ctx context.Context, uuid uuid.UUID) error {
	query := `
		DELETE FROM orders
		WHERE order_uid = $1
	`
	res, err := s.db.ExecContext(ctx, query, uuid)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return storage.ErrOrderNotFound
	}

	return nil
}
