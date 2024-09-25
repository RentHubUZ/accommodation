package postgres

import (
	pb "accommodation/genproto/accommodation"
	"accommodation/pkg/logger"
	"accommodation/storage"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type HousesRepository struct {
	Db  *sql.DB
	Log *slog.Logger
}

func NewHousesRepository(db *sql.DB) storage.IUserStorage {
	return &HousesRepository{Db: db, Log: logger.NewLogger()}
}

func (h *HousesRepository) CreateHouse(ctx context.Context, req *pb.CreateHouseReq) (*pb.CreateHouseRes, error) {
	tx, err := h.Db.BeginTx(ctx, nil)
	if err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("error starting transaction: %v", err.Error()))
		return nil, err
	}

	query_property := `insert into properties (
							id, owner_id, address, price, property_type, 
							bedrooms, bathrooms, square_footage, listing_status, 
							description, roommate_count, lease_terms, lease_duration, 
							top_status, location, created_at, updated_at
			  			) values (
			  			  	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, 
			  			  	ST_SetSRID(ST_MakePoint($14, $15), 4326), $16, $17)`

	property_id := uuid.NewString()
	newtime := time.Now()

	_, err = tx.ExecContext(ctx, query_property,
		property_id, req.OwnerId, req.Address, req.Price, req.PropertyType,
		req.Bedrooms, req.Bathrooms, req.SquareFootage, req.ListingStatus,
		req.Description, req.RoommateCount, req.LeaseTerms, req.LeaseDuration,
		req.TopStatus, req.Latitude, req.Longitude,
		newtime, newtime)
	if err != nil {
		tx.Rollback()
		h.Log.ErrorContext(ctx, fmt.Sprintf("error adding property: %v", err.Error()))
		return nil, err
	}

	query_property_images := `insert into propertyimages (
									id, property_id, image_url, uploaded_at
								) values (
									$1, $2, $3, $4)`

	for _, imageUrl := range req.ImageUrl {
		property_image_url_id := uuid.NewString()
		newtimeImage := time.Now()

		_, err := tx.ExecContext(ctx, query_property_images, property_image_url_id, property_id, imageUrl, newtimeImage)
		if err != nil {
			tx.Rollback()
			h.Log.ErrorContext(ctx, fmt.Sprintf("error adding property_images: %v", err.Error()))
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("error in commenting the transaction: %v", err.Error()))
		return nil, err
	}

	response := &pb.CreateHouseRes{
		Id:            property_id,
		OwnerId:       req.OwnerId,
		Address:       req.Address,
		Price:         req.Price,
		PropertyType:  req.PropertyType,
		Bedrooms:      req.Bedrooms,
		Bathrooms:     req.Bathrooms,
		SquareFootage: req.SquareFootage,
		ListingStatus: req.ListingStatus,
		Description:   req.Description,
		RoommateCount: req.RoommateCount,
		LeaseTerms:    req.LeaseTerms,
		LeaseDuration: req.LeaseDuration,
		TopStatus:     req.TopStatus,
		Latitude:      req.Latitude,
		Longitude:     req.Longitude,
		ImageUrl:      req.ImageUrl,
		CreatedAt:     newtime.Format("2006-01-02 15:04:05"),
		UpdatedAt:     newtime.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (h *HousesRepository) UpdateHouse(ctx context.Context, req *pb.UpdateHouseReq) (*pb.UpdateHouseRes, error) {
	tx, err := h.Db.BeginTx(ctx, nil)
	if err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("error starting transaction: %v", err.Error()))
		return nil, err
	}

	query_get_property := `SELECT owner_id, address, price, property_type, bedrooms, bathrooms, 
								  square_footage, listing_status, description, roommate_count, 
								  lease_terms, lease_duration, top_status, 
								  ST_X(location) as latitude, ST_Y(location) as longitude, created_at, updated_at
						   FROM properties WHERE id = $1 AND deleted_at IS NULL`

	var oldProperty struct {
		OwnerID       string
		Address       string
		Price         float32
		PropertyType  string
		Bedrooms      int32
		Bathrooms     int32
		SquareFootage float32
		ListingStatus string
		Description   string
		RoommateCount int32
		LeaseTerms    string
		LeaseDuration int32
		TopStatus     bool
		Latitude      float32
		Longitude     float32
		CreatedAt     string
		UpdatedAt     string
	}

	err = tx.QueryRowContext(ctx, query_get_property, req.Id).Scan(
		&oldProperty.OwnerID, &oldProperty.Address, &oldProperty.Price, &oldProperty.PropertyType,
		&oldProperty.Bedrooms, &oldProperty.Bathrooms, &oldProperty.SquareFootage,
		&oldProperty.ListingStatus, &oldProperty.Description, &oldProperty.RoommateCount,
		&oldProperty.LeaseTerms, &oldProperty.LeaseDuration, &oldProperty.TopStatus,
		&oldProperty.Latitude, &oldProperty.Longitude, &oldProperty.CreatedAt, &oldProperty.UpdatedAt)
	if err != nil {
		tx.Rollback()
		h.Log.ErrorContext(ctx, fmt.Sprintf("error reading property: %v", err.Error()))
		return nil, err
	}

	ownerID := req.OwnerId
	if ownerID == "" {
		ownerID = oldProperty.OwnerID
	}
	address := req.Address
	if address == "" {
		address = oldProperty.Address
	}
	price := req.Price
	if price == 0 {
		price = oldProperty.Price
	}
	propertyType := req.PropertyType
	if propertyType == "" {
		propertyType = oldProperty.PropertyType
	}
	bedrooms := req.Bedrooms
	if bedrooms == 0 {
		bedrooms = oldProperty.Bedrooms
	}
	bathrooms := req.Bathrooms
	if bathrooms == 0 {
		bathrooms = oldProperty.Bathrooms
	}
	squareFootage := req.SquareFootage
	if squareFootage == 0 {
		squareFootage = oldProperty.SquareFootage
	}
	listingStatus := req.ListingStatus
	if listingStatus == "" {
		listingStatus = oldProperty.ListingStatus
	}
	description := req.Description
	if description == "" {
		description = oldProperty.Description
	}
	roommateCount := req.RoommateCount
	if roommateCount == 0 {
		roommateCount = oldProperty.RoommateCount
	}
	leaseTerms := req.LeaseTerms
	if leaseTerms == "" {
		leaseTerms = oldProperty.LeaseTerms
	}
	leaseDuration := req.LeaseDuration
	if leaseDuration == 0 {
		leaseDuration = oldProperty.LeaseDuration
	}
	topStatus := req.TopStatus
	latitude := req.Latitude
	if latitude == 0 {
		latitude = oldProperty.Latitude
	}
	longitude := req.Longitude
	if longitude == 0 {
		longitude = oldProperty.Longitude
	}
	createdAt := req.CreatedAt
	if createdAt == "" {
		createdAt = oldProperty.CreatedAt
	}
	updatedAt := req.UpdatedAt
	if updatedAt == "" {
		updatedAt = oldProperty.UpdatedAt
	}

	query_update_property := `UPDATE properties SET 
								owner_id = $1, address = $2, price = $3, property_type = $4, 
								bedrooms = $5, bathrooms = $6, square_footage = $7, listing_status = $8, 
								description = $9, roommate_count = $10, lease_terms = $11, lease_duration = $12, 
								top_status = $13, location = ST_SetSRID(ST_MakePoint($14, $15), 4326), 
								created_at = $16, updated_at = $17 
							  WHERE id = $18 AND deleted_at IS NULL`
	_, err = tx.ExecContext(ctx, query_update_property,
		ownerID, address, price, propertyType, bedrooms, bathrooms,
		squareFootage, listingStatus, description, roommateCount, leaseTerms,
		leaseDuration, topStatus, latitude, longitude, createdAt, updatedAt, req.Id)
	if err != nil {
		tx.Rollback()
		h.Log.ErrorContext(ctx, fmt.Sprintf("error updating property: %v", err.Error()))
		return nil, err
	}

	if len(req.ImageUrl) > 0 {
		query_delete_images := `DELETE FROM propertyimages WHERE property_id = $1`
		_, err = tx.ExecContext(ctx, query_delete_images, req.Id)
		if err != nil {
			tx.Rollback()
			h.Log.ErrorContext(ctx, fmt.Sprintf("error deleting property: %v", err.Error()))
			return nil, err
		}

		query_insert_image := `INSERT INTO propertyimages (id, property_id, image_url, uploaded_at) VALUES ($1, $2, $3, NOW())`
		for _, imageUrl := range req.ImageUrl {
			_, err := tx.ExecContext(ctx, query_insert_image, uuid.NewString(), req.Id, imageUrl)
			if err != nil {
				tx.Rollback()
				h.Log.ErrorContext(ctx, fmt.Sprintf("error new adding property_image: %v", err.Error()))
				return nil, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("error in commenting the transaction: %v", err.Error()))
		return nil, err
	}

	return &pb.UpdateHouseRes{
		Message: "Property successfully updated",
	}, nil
}

func (h *HousesRepository) GetAllHouse(ctx context.Context, req *pb.GetallHouseReq) ([]*pb.GetAllHouseRes, error) {
	limit := req.Limit
	page := req.Page
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	query := `SELECT p.id, p.owner_id, p.address, p.price, p.property_type, p.bedrooms, 
					 p.bathrooms, p.square_footage, p.listing_status, p.description, p.roommate_count, 
					 p.lease_terms, p.lease_duration, p.top_status, ST_X(p.location) as latitude, 
					 ST_Y(p.location) as longitude, p.created_at, p.updated_at,
					 ARRAY(SELECT image_url FROM propertyimages WHERE property_id = p.id) as image_url
			  FROM properties p
			  WHERE p.deleted_at IS NULL
			  ORDER BY p.created_at DESC
			  LIMIT $1 OFFSET $2`

	rows, err := h.Db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("error reading property and property_images: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var houses []*pb.GetAllHouseRes
	for rows.Next() {
		var house pb.GetAllHouseRes
		var imageUrls []string

		err := rows.Scan(
			&house.Id, &house.OwnerId, &house.Address, &house.Price, &house.PropertyType,
			&house.Bedrooms, &house.Bathrooms, &house.SquareFootage, &house.ListingStatus,
			&house.Description, &house.RoommateCount, &house.LeaseTerms, &house.LeaseDuration,
			&house.TopStatus, &house.Latitude, &house.Longitude, &house.CreatedAt, &house.UpdatedAt,
			pq.Array(&imageUrls),
		)
		if err != nil {
			h.Log.ErrorContext(ctx, fmt.Sprintf("error scanning property and property_images: %v", err.Error()))
			return nil, err
		}

		house.ImageUrl = imageUrls
		houses = append(houses, &house)
	}

	if err = rows.Err(); err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("error rows: %v", err.Error()))
		return nil, err
	}

	return houses, nil
}

func (h *HousesRepository) GetByIdHouse(ctx context.Context, req *pb.GetByIdHouseReq) (*pb.GetByIdHouseRes, error) {
	query := `SELECT p.id, p.owner_id, p.address, p.price, p.property_type, p.bedrooms, 
					 p.bathrooms, p.square_footage, p.listing_status, p.description, 
					 p.roommate_count, p.lease_terms, p.lease_duration, p.top_status, 
					 ST_X(p.location) as latitude, ST_Y(p.location) as longitude, 
					 p.created_at, p.updated_at,
					 ARRAY(SELECT image_url FROM propertyimages WHERE property_id = p.id) as image_url
			  FROM properties p
			  WHERE p.id = $1 AND deleted_at IS NULL`

	var house pb.GetByIdHouseRes
	var imageUrls []string

	err := h.Db.QueryRowContext(ctx, query, req.Id).Scan(
		&house.Id, &house.OwnerId, &house.Address, &house.Price, &house.PropertyType,
		&house.Bedrooms, &house.Bathrooms, &house.SquareFootage, &house.ListingStatus,
		&house.Description, &house.RoommateCount, &house.LeaseTerms, &house.LeaseDuration,
		&house.TopStatus, &house.Latitude, &house.Longitude, &house.CreatedAt,
		&house.UpdatedAt, pq.Array(&imageUrls),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			h.Log.ErrorContext(ctx, fmt.Sprintf("no such house was found: %v", err.Error()))
			return nil, fmt.Errorf("house with id %s not found", req.Id)
		}
		h.Log.ErrorContext(ctx, fmt.Sprintf("%s --> Error retrieving information on this id: %v", req.Id, err.Error()))
		return nil, err
	}

	house.ImageUrl = imageUrls

	return &house, nil
}

func (h *HousesRepository) DeleteHouse(ctx context.Context, req *pb.DeleteHouseReq) (*pb.DeleteHouseRes, error) {
	query_property_image_url := `update 
								propertyimages
							set 
								deleted_at = $1
							where 
								property_id = $2`

	query_property_delete := `update 
									properties
								set
									deleted_at = $1
								where
									id = $2`

	NewTimeDelete := time.Now()

	_, err := h.Db.ExecContext(ctx, query_property_image_url,NewTimeDelete,req.Id)
	if err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("failed to soft delete images for house with id %s: %v", req.Id, err.Error()))
		return nil,err
	}

	_,err = h.Db.ExecContext(ctx,query_property_delete,NewTimeDelete,req.Id)
	if err != nil {
		h.Log.ErrorContext(ctx, fmt.Sprintf("failed to soft delete house with id %s: %v", req.Id, err))
		return nil,err
	}

	return &pb.DeleteHouseRes{
		Message: "Your home has been successfully launched",
	},nil
}