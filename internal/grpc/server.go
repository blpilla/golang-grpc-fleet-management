package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"fleet_management/internal/association"
	"fleet_management/internal/driver"
	"fleet_management/internal/vehicle"
	"fleet_management/proto"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer(driverService driver.Service, vehicleService vehicle.Service, associationService association.Service) *Server {
	grpcServer := grpc.NewServer()

	proto.RegisterDriverServiceServer(grpcServer, driver.NewGRPCService(driverService))
	proto.RegisterVehicleServiceServer(grpcServer, vehicle.NewGRPCService(vehicleService))
	proto.RegisterAssociationServiceServer(grpcServer, association.NewGRPCService(associationService))

	return &Server{grpcServer: grpcServer}
}

func (s *Server) Serve(lis net.Listener) error {
	log.Println("Starting gRPC server...")
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
