package vehicle

import (
	"database/sql"
	"fleet_management/proto"
)

type Repository interface {
	Create(vehicle *proto.CreateVehicleRequest) (*proto.VehicleResponse, error)
	GetByID(id int32) (*proto.VehicleResponse, error)
	GetAll() ([]*proto.VehicleResponse, error)
	Update(vehicle *proto.UpdateVehicleRequest) (*proto.VehicleResponse, error)
	Delete(id int32) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(vehicle *proto.CreateVehicleRequest) (*proto.VehicleResponse, error) {
	query := "INSERT INTO vehicles (make, model) VALUES ($1, $2) RETURNING id"
	var id int32
	err := r.db.QueryRow(query, vehicle.Make, vehicle.Model).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &proto.VehicleResponse{
		Id:    id,
		Make:  vehicle.Make,
		Model: vehicle.Model,
	}, nil
}

func (r *postgresRepository) GetByID(id int32) (*proto.VehicleResponse, error) {
	vehicle := &proto.VehicleResponse{}
	query := "SELECT id, make, model FROM vehicles WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&vehicle.Id, &vehicle.Make, &vehicle.Model)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return vehicle, err
}

func (r *postgresRepository) GetAll() ([]*proto.VehicleResponse, error) {
	query := "SELECT id, make, model FROM vehicles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []*proto.VehicleResponse
	for rows.Next() {
		vehicle := &proto.VehicleResponse{}
		if err := rows.Scan(&vehicle.Id, &vehicle.Make, &vehicle.Model); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (r *postgresRepository) Update(vehicle *proto.UpdateVehicleRequest) (*proto.VehicleResponse, error) {
	query := "UPDATE vehicles SET make = $1, model = $2 WHERE id = $3"
	_, err := r.db.Exec(query, vehicle.Make, vehicle.Model, vehicle.Id)
	if err != nil {
		return nil, err
	}
	return &proto.VehicleResponse{
		Id:    vehicle.Id,
		Make:  vehicle.Make,
		Model: vehicle.Model,
	}, nil
}

func (r *postgresRepository) Delete(id int32) error {
	query := "DELETE FROM vehicles WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
