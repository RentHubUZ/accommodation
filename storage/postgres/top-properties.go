package postgres

import (
    "accommodation/storage"
    "context"
    "database/sql"
    "fmt"
    "log/slog"

    pt "accommodation/genproto/top-properties"
)

type TopPropertiesRepository struct {
    Db  *sql.DB
    Log *slog.Logger
}

func NewTopPropertiesRepository(db *sql.DB, log *slog.Logger) storage.ITopPropertiesStorage {
    return &TopPropertiesRepository{Db: db, Log: log}
}

func (t *TopPropertiesRepository) Create(ctx context.Context, req *pt.CreateTopPropertyReq) (*pt.CreateTopPropertyRes, error) {
    query := `INSERT INTO top_properties (property_id, user_id, start_date, end_date, tariff_name, created_at) 
              VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`

    var id string
    err := t.Db.QueryRowContext(ctx, query, req.PropertyId, req.UserId, req.StartDate, req.EndDate, req.TariffName).Scan(&id)
    if err != nil {
        t.Log.Error(fmt.Sprintf("Error creating top property: %v", err.Error()))
        return nil, err
    }

    return &pt.CreateTopPropertyRes{
        Id: id,
    }, nil
}

func (t *TopPropertiesRepository) Get(ctx context.Context, req *pt.GetTopPropertyReq) (*pt.GetTopPropertyRes, error) {
    query := `SELECT id, property_id, user_id, start_date, end_date, tariff_name, created_at FROM top_properties WHERE id = $1`

    row := t.Db.QueryRowContext(ctx, query, req.Id)
    
    var topProperty pt.TopProperty
    err := row.Scan(&topProperty.Id, &topProperty.PropertyId, &topProperty.UserId, &topProperty.StartDate, &topProperty.EndDate, &topProperty.TariffName, &topProperty.CreatedAt)
    if err != nil {
        t.Log.Error(fmt.Sprintf("Error getting top property: %v", err.Error()))
        return nil, err
    }

    return &pt.GetTopPropertyRes{
        TopProperty: &topProperty,
    }, nil
}

func (t *TopPropertiesRepository) GetAll(ctx context.Context, req *pt.GetAllTopPropertyReq) (*pt.GetAllTopPropertyRes, error) {
    offset := (req.Page - 1) * req.Limit
    query := `SELECT id, property_id, user_id, start_date, end_date, tariff_name, created_at FROM top_properties ORDER BY created_at DESC LIMIT $1 OFFSET $2`

    rows, err := t.Db.QueryContext(ctx, query, req.Limit, offset)
    if err != nil {
        t.Log.Error(fmt.Sprintf("Error getting all top properties: %v", err.Error()))
        return nil, err
    }
    defer rows.Close()

    var topProperties []*pt.TopProperty
    for rows.Next() {
        var topProperty pt.TopProperty
        err := rows.Scan(&topProperty.Id, &topProperty.PropertyId, &topProperty.UserId, &topProperty.StartDate, &topProperty.EndDate, &topProperty.TariffName, &topProperty.CreatedAt)
        if err != nil {
            t.Log.Error(fmt.Sprintf("Error scanning top property: %v", err.Error()))
            return nil, err
        }
        topProperties = append(topProperties, &topProperty)
    }

    return &pt.GetAllTopPropertyRes{
        TopProperties: topProperties,
    }, nil
}

func (t *TopPropertiesRepository) Update(ctx context.Context, req *pt.UpdateTopPropertyReq) (*pt.UpdateTopPropertyRes, error) {
    query := `UPDATE top_properties SET property_id = $1, user_id = $2, start_date = $3, end_date = $4, tariff_name = $5 WHERE id = $6`

    _, err := t.Db.ExecContext(ctx, query, req.PropertyId, req.UserId, req.StartDate, req.EndDate, req.TariffName, req.Id)
    if err != nil {
        t.Log.Error(fmt.Sprintf("Error updating top property: %v", err.Error()))
        return nil, err
    }

    return &pt.UpdateTopPropertyRes{
        Success: true,
    }, nil
}

func (t *TopPropertiesRepository) Delete(ctx context.Context, req *pt.DeleteTopPropertyReq) (*pt.DeleteTopPropertyRes, error) {
    query := `DELETE FROM top_properties WHERE id = $1`

    _, err := t.Db.ExecContext(ctx, query, req.Id)
    if err != nil {
        t.Log.Error(fmt.Sprintf("Error deleting top property: %v", err.Error()))
        return nil, err
    }

    return &pt.DeleteTopPropertyRes{
        Success: true,
    }, nil
}
