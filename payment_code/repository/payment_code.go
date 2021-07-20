package repository

import (
	"context"
	"database/sql"
	"fmt"
	"frieda-golang-training-beginner/domain"
	"github.com/google/uuid"
	"time"
)

type PaymentCodeRepository struct {
	Conn *sql.DB
}

func (m *PaymentCodeRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.PaymentCode, err error) {
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

func (p PaymentCodeRepository) GetByID(ctx context.Context, uuid string) (domain.PaymentCode, error) {
	query := `SELECT id, payment_code, name, status, expiration_date, created_at, updated_at
  						FROM payment_codes WHERE id = $1`

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

func (p PaymentCodeRepository) Create(ctx context.Context, paymentCode *domain.PaymentCode) error {
	paymentCode.ID = uuid.New()
	query := `INSERT INTO payment_codes (id, payment_code, name, status, expiration_date) 
				VALUES ($1 , $2 , $3, $4, $5)`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, paymentCode.ID, paymentCode.PaymentCode, paymentCode.Name, paymentCode.Status, paymentCode.ExpirationDate)
	if err != nil {
		return err
	}

	return nil
}

func (p PaymentCodeRepository) Expire(ctx context.Context) error {
	now := time.Now()
	query := `UPDATE payment_codes SET status=$1, updated_at=$2 WHERE status='ACTIVE' and expiration_date <= $3`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, "EXPIRED", now, now)
	if err != nil {
		return err
	}
	return nil
}

func NewPaymentCodeRepository(Conn *sql.DB) *PaymentCodeRepository {
	return &PaymentCodeRepository{Conn}
}
