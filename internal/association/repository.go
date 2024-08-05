package association

import (
	"database/sql"
	"fleet_management/proto"
)

type Repository interface {
	AssociateDriverToVehicle(driverID int, vehicleID int) error
	GetAllAssociations() ([]*proto.Association, error)
	GetDriversByVehicle(vehicleID int) ([]int, error)
	GetVehiclesByDriver(driverID int) ([]int, error)
	DissociateDriverFromVehicle(driverID int, vehicleID int) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) AssociateDriverToVehicle(driverID int, vehicleID int) error {
	query := "INSERT INTO associations (driver_id, vehicle_id) VALUES ($1, $2)"
	_, err := r.db.Exec(query, driverID, vehicleID)
	return err
}

func (r *postgresRepository) GetAllAssociations() ([]*proto.Association, error) {
	query := "SELECT driver_id, vehicle_id FROM associations"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	associations := []*proto.Association{}
	for rows.Next() {
		assoc := &proto.Association{}
		if err := rows.Scan(&assoc.DriverId, &assoc.VehicleId); err != nil {
			return nil, err
		}
		associations = append(associations, assoc)
	}
	return associations, nil
}

func (r *postgresRepository) GetDriversByVehicle(vehicleID int) ([]int, error) {
	query := "SELECT driver_id FROM associations WHERE vehicle_id = $1"
	rows, err := r.db.Query(query, vehicleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var driverIDs []int
	for rows.Next() {
		var driverID int
		if err := rows.Scan(&driverID); err != nil {
			return nil, err
		}
		driverIDs = append(driverIDs, driverID)
	}
	return driverIDs, nil
}

func (r *postgresRepository) GetVehiclesByDriver(driverID int) ([]int, error) {
	query := "SELECT vehicle_id FROM associations WHERE driver_id = $1"
	rows, err := r.db.Query(query, driverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicleIDs []int
	for rows.Next() {
		var vehicleID int
		if err := rows.Scan(&vehicleID); err != nil {
			return nil, err
		}
		vehicleIDs = append(vehicleIDs, vehicleID)
	}
	return vehicleIDs, nil
}

func (r *postgresRepository) DissociateDriverFromVehicle(driverID int, vehicleID int) error {
	query := "DELETE FROM associations WHERE driver_id = $1 AND vehicle_id = $2"
	_, err := r.db.Exec(query, driverID, vehicleID)
	return err
}
