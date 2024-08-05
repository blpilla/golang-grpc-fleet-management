package driver

import (
	"database/sql"
	"fleet_management/proto"
)

type Repository interface {
	Create(driver *proto.CreateDriverRequest) (*proto.DriverResponse, error)
	GetByID(id int32) (*proto.DriverResponse, error)
	GetAll() ([]*proto.DriverResponse, error)
	Update(driver *proto.UpdateDriverRequest) (*proto.DriverResponse, error)
	Delete(id int32) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(driver *proto.CreateDriverRequest) (*proto.DriverResponse, error) {
	query := "INSERT INTO drivers (first_name, last_name, license) VALUES ($1, $2, $3) RETURNING id"
	var id int32
	err := r.db.QueryRow(query, driver.FirstName, driver.LastName, driver.License).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &proto.DriverResponse{
		Id:        id,
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		License:   driver.License,
	}, nil
}

func (r *postgresRepository) GetByID(id int32) (*proto.DriverResponse, error) {
	driver := &proto.DriverResponse{}
	query := "SELECT id, first_name, last_name, license FROM drivers WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&driver.Id, &driver.FirstName, &driver.LastName, &driver.License)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return driver, err
}

func (r *postgresRepository) GetAll() ([]*proto.DriverResponse, error) {
	query := "SELECT id, first_name, last_name, license FROM drivers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []*proto.DriverResponse
	for rows.Next() {
		driver := &proto.DriverResponse{}
		if err := rows.Scan(&driver.Id, &driver.FirstName, &driver.LastName, &driver.License); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *postgresRepository) Update(driver *proto.UpdateDriverRequest) (*proto.DriverResponse, error) {
	query := "UPDATE drivers SET first_name = $1, last_name = $2, license = $3 WHERE id = $4"
	_, err := r.db.Exec(query, driver.FirstName, driver.LastName, driver.License, driver.Id)
	if err != nil {
		return nil, err
	}
	return &proto.DriverResponse{
		Id:        driver.Id,
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		License:   driver.License,
	}, nil
}

func (r *postgresRepository) Delete(id int32) error {
	query := "DELETE FROM drivers WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
