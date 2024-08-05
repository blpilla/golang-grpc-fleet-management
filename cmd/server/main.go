package main

import (
	"net"

	"fleet_management/internal/association"
	"fleet_management/internal/db"
	"fleet_management/internal/driver"
	"fleet_management/internal/grpc"
	"fleet_management/internal/util"
	"fleet_management/internal/vehicle"
	"fleet_management/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Inicializa o logger
	logger := util.NewLogger()

	// Conecta ao banco de dados
	database, err := db.NewPostgreSQL()
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the database
	database.AutoMigrate(&driver.Driver{}, &vehicle.Vehicle{}, &association.Association{})

	// Inicializa os repositórios
	driverRepo := driver.NewRepository(database)
	vehicleRepo := vehicle.NewRepository(database)
	associationRepo := association.NewRepository(database)

	// Inicializa os serviços
	driverService := driver.NewService(driverRepo)
	vehicleService := vehicle.NewService(vehicleRepo)
	associationService := association.NewService(associationRepo)

	// Cria o servidor gRPC
	grpcServer := grpc.NewServer()

	// Registra os serviços gRPC
	proto.RegisterDriverServiceServer(grpcServer, driver.NewGRPCService(driverService))
	proto.RegisterVehicleServiceServer(grpcServer, vehicle.NewGRPCService(vehicleService))
	proto.RegisterAssociationServiceServer(grpcServer, association.NewGRPCService(associationService))

	// Habilita a reflexão gRPC para facilitar o desenvolvimento
	reflection.Register(grpcServer)

	// Inicializa o listener na porta 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal("Failed to listen on port 50051:", err)
	}

	// Inicia o servidor gRPC
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Failed to serve gRPC server:", err)
	}
}
