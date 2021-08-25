package repository

import (
	"context"
	"database/sql"
	"fmt"
	"frieda-golang-training-beginner/domain"
	"github.com/google/uuid"
)

type InquiryRepository struct {
	Conn *sql.DB
}

func (m InquiryRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Inquiry, err error) {
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

	result = make([]domain.Inquiry, 0)
	for rows.Next() {
		t := domain.Inquiry{}
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

func (m InquiryRepository) GetByTransactionID(ctx context.Context, transactionID string) (domain.Inquiry, error) {
	query := `SELECT id, payment_code, transaction_id, created_at, updated_at
  						FROM inquiry WHERE transaction_id = $1`

	list, err := m.fetch(ctx, query, transactionID)
	if err != nil {
		return domain.Inquiry{}, err
	}

	if len(list) > 0 {
		res := list[0]
		return res, nil
	} else {
		return domain.Inquiry{}, nil
	}
}

func (m InquiryRepository) Create(ctx context.Context, inquiry *domain.Inquiry) error {
	inquiry.ID = uuid.New()
	query := `INSERT INTO inquiry (id, payment_code, transaction_id) 
				VALUES ($1 , $2 , $3)`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, inquiry.ID, inquiry.PaymentCode, inquiry.TransactionID,
		inquiry.CreatedAt, inquiry.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func NewInquiryRepository(Conn *sql.DB) *InquiryRepository {
	return &InquiryRepository{Conn}
}
