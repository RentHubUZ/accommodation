package postgres

import (
	"accommodation/internal/storage"
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	pt "accommodation/genproto/tariff"
)

type TariffsRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewTariffsRepository(db *sql.DB, log *slog.Logger) storage.ITariffStorage {
	return &TariffsRepository{Db: db, Log: log}
}

// Create function
func (t *TariffsRepository) Create(ctx context.Context, req *pt.CreateTariffReq) (*pt.CreateTariffRes, error) {
	query := `INSERT INTO tariffs (name, days, price, offers, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`

	var id string
	err := t.Db.QueryRow(query, req.Name, req.Days, req.Price, req.Offers).Scan(&id)
	if err != nil {
		t.Log.Error(fmt.Sprintf("Error creating tariff: %v", err.Error()))
		return nil, err
	}

	return &pt.CreateTariffRes{
		Id: id,
	}, nil
}

// Get function
func (t *TariffsRepository) Get(ctx context.Context, req *pt.GetTariffReq) (*pt.GetTariffRes, error) {
	query := `SELECT id, name, days, price, offers, created_at, updated_at FROM tariffs WHERE id = $1`

	var tariff pt.Tariff
	err := t.Db.QueryRow(query, req.Id).Scan(&tariff.Id, &tariff.Name, &tariff.Days, &tariff.Price, &tariff.Offers, &tariff.CreatedAt, &tariff.UpdatedAt)
	if err != nil {
		t.Log.Error(fmt.Sprintf("Error getting tariff: %v", err.Error()))
		return nil, err
	}

	return &pt.GetTariffRes{
		Tariff: &tariff,
	}, nil
}

// Update function
func (t *TariffsRepository) Update(ctx context.Context, req *pt.UpdateTariffReq) (*pt.UpdateTariffRes, error) {
	query := `UPDATE tariffs SET name = $1, days = $2, price = $3, offers = $4, updated_at = NOW() WHERE id = $5`

	_, err := t.Db.Exec(query, req.Name, req.Days, req.Price, req.Offers, req.Id)
	if err != nil {
		t.Log.Error(fmt.Sprintf("Error updating tariff: %v", err.Error()))
		return nil, err
	}

	return &pt.UpdateTariffRes{
		Success: true,
	}, nil
}

// Delete function
func (t *TariffsRepository) Delete(ctx context.Context, req *pt.DeleteTariffReq) (*pt.DeleteTariffRes, error) {
	query := `DELETE FROM tariffs WHERE id = $1`

	_, err := t.Db.Exec(query, req.Id)
	if err != nil {
		t.Log.Error(fmt.Sprintf("Error deleting tariff: %v", err.Error()))
		return nil, err
	}

	return &pt.DeleteTariffRes{
		Success: true,
	}, nil
}

// GetAll function
func (t *TariffsRepository) GetAll(ctx context.Context, req *pt.GetAllTariffReq) (*pt.GetAllTariffRes, error) {
	query := `SELECT id, name, days, price, offers, created_at, updated_at FROM tariffs LIMIT $1 OFFSET $2`

	rows, err := t.Db.Query(query, req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		t.Log.Error(fmt.Sprintf("Error getting all tariffs: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var tariffs []*pt.Tariff
	for rows.Next() {
		var tariff pt.Tariff
		err := rows.Scan(&tariff.Id, &tariff.Name, &tariff.Days, &tariff.Price, &tariff.Offers, &tariff.CreatedAt, &tariff.UpdatedAt)
		if err != nil {
			t.Log.Error(fmt.Sprintf("Error scanning tariffs: %v", err.Error()))
			return nil, err
		}
		tariffs = append(tariffs, &tariff)
	}

	return &pt.GetAllTariffRes{
		Tariffs: tariffs,
	}, nil
}
