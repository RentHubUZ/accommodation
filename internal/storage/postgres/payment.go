package postgres

import (
	"accommodation/internal/storage"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	pt "accommodation/genproto/payment"
)

type PaymentRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewPaymentRepository(db *sql.DB, log *slog.Logger) storage.IPaymentStorage {
	return &PaymentRepository{Db: db, Log: log}
}
// Get payment by ID
func (r *PaymentRepository) Get(ctx context.Context, req *pt.GetPaymentReq) (*pt.GetPaymentRes, error) {
	query := "SELECT id, amount, status, transaction_date FROM Payments WHERE id = $1"
	row := r.Db.QueryRowContext(ctx, query, req.Id)

	var payment pt.Payment
	err := row.Scan(&payment.Id, &payment.Amount, &payment.Status, &payment.TransactionDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found")
		}
		r.Log.Error("error fetching payment", slog.Any("error", err))
		return nil, err
	}

	return &pt.GetPaymentRes{Payment: &payment}, nil
}

// Get all payments with pagination
func (r *PaymentRepository) GetAll(ctx context.Context, req *pt.GetAllPaymentReq) (*pt.GetAllPaymentRes, error) {
	offset := (req.Page - 1) * req.Limit
	query := "SELECT id, amount, status, transaction_date FROM Payments LIMIT $1 OFFSET $2"

	rows, err := r.Db.QueryContext(ctx, query, req.Limit, offset)
	if err != nil {
		r.Log.Error("error fetching payments", slog.Any("error", err))
		return nil, err
	}
	defer rows.Close()

	var payments []*pt.Payment
	for rows.Next() {
		var payment pt.Payment
		if err := rows.Scan(&payment.Id, &payment.Amount, &payment.Status, &payment.TransactionDate); err != nil {
			r.Log.Error("error scanning payment row", slog.Any("error", err))
			return nil, err
		}
		payments = append(payments, &payment)
	}

	if err := rows.Err(); err != nil {
		r.Log.Error("error in payment rows", slog.Any("error", err))
		return nil, err
	}

	return &pt.GetAllPaymentRes{Payments: payments}, nil
}

// Create a new payment
func (r *PaymentRepository) Create(ctx context.Context, req *pt.CreatePaymentReq) (*pt.CreatePaymentRes, error) {
	query := "INSERT INTO Payments (amount, status, transaction_date) VALUES ($1, $2, $3) RETURNING id"
	var id string
	err := r.Db.QueryRowContext(ctx, query, req.Amount, req.Status, time.Now().Format("2006-01-02")).Scan(&id)
	if err != nil {
		r.Log.Error("error creating payment", slog.Any("error", err))
		return nil, err
	}

	return &pt.CreatePaymentRes{Id: id}, nil
}

// Delete a payment by ID
func (r *PaymentRepository) Delete(ctx context.Context, req *pt.DeletePaymentReq) (*pt.DeletePaymentRes, error) {
	query := "DELETE FROM Payments WHERE id = $1"
	result, err := r.Db.ExecContext(ctx, query, req.Id)
	if err != nil {
		r.Log.Error("error deleting payment", slog.Any("error", err))
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.Log.Error("error getting rows affected", slog.Any("error", err))
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("payment not found")
	}

	return &pt.DeletePaymentRes{Success: true}, nil
}
