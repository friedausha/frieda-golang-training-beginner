package repository

import (
	"context"
	"database/sql"
	"fmt"
	"frieda-golang-training-beginner/domain"
	"github.com/google/uuid"
	"time"
)

type paymentCodeRepository struct {
	Conn *sql.DB
}

func (m *paymentCodeRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.PaymentCode, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			fmt.Errorf(`error when fetching`, errRow)
			return
		}
	}()

	result = make([]domain.PaymentCode, 0)
	for rows.Next() {
		t := domain.PaymentCode{}
		err = rows.Scan(
			&t.ID,
			&t.PaymentCode,
			&t.Name,
			&t.Status,
			&t.ExpirationDate,
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

func (p paymentCodeRepository) GetByID(ctx context.Context, uuid uuid.UUID) (domain.PaymentCode, error) {
	query := `SELECT id, payment_code, name, status, expiration_date, created_at, updated_at
  						FROM payment_codes WHERE ID = ?`

	list, err := p.fetch(ctx, query, uuid)
	if err != nil {
		return domain.PaymentCode{}, err
	}

	if len(list) > 0 {
		res := list[0]
		return res, nil
	} else {
		return domain.PaymentCode{}, nil
	}
}

func (p paymentCodeRepository) Create(ctx context.Context, paymentCode *domain.PaymentCode) error {
	paymentCode.ID = uuid.New()
	if paymentCode.ExpirationDate.IsZero() {
		paymentCode.ExpirationDate = time.Now().AddDate(51, 0, 0).UTC()
	}
	query := `INSERT INTO payment_codes SET id=? , payment_code=? , name=?, status=?, expiration_date=?`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	_, err = stmt.ExecContext(ctx, paymentCode.ID, paymentCode.PaymentCode, paymentCode.Name, paymentCode.ExpirationDate)
	if err != nil {
		return err
	}
	return nil

}

func NewPaymentCodeRepository(Conn *sql.DB) domain.PaymentCodeRepository {
	return &paymentCodeRepository{Conn}
}
