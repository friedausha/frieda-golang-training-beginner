package repository

import (
	"context"
	"database/sql"
	"fmt"
	"frieda-golang-training-beginner/domain"
	"github.com/google/uuid"
)

type PaymentRepository struct {
	Conn *sql.DB
}

func (m PaymentRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Payment,
	err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			fmt.Errorf("error when fetching %d", errRow)
			return
		}
	}()

	result = make([]domain.Payment, 0)
	for rows.Next() {
		t := domain.Payment{}
		err = rows.Scan(
			&t.ID,
			&t.PaymentCode,
			&t.TransactionID,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m PaymentRepository) GetByID(ctx context.Context, id string) (domain.Payment, error) {
	query := `SELECT id, payment_code, transaction_id, created_at, updated_at
  						FROM payment WHERE id = $1`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Payment{}, err
	}

	if len(list) > 0 {
		res := list[0]
		return res, nil
	} else {
		return domain.Payment{}, nil
	}
}

func (m PaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	payment.ID = uuid.New()
	query := `INSERT INTO payment (id, name, amount, payment_code, transaction_id) 
				VALUES ($1 , $2 , $3, $4, $5)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, payment.ID, payment.Name, payment.Amount, payment.PaymentCode, payment.TransactionID)
	if err != nil {
		return err
	}
	return nil
}

func NewPaymentRepository(Conn *sql.DB) *PaymentRepository {
	return &PaymentRepository{Conn}
}
