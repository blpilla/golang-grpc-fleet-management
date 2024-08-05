package association

import (
	"context"

	"fleet_management/proto"
)

type grpcService struct {
	service Service
	proto.UnimplementedAssociationServiceServer
}

func NewGRPCService(service Service) proto.AssociationServiceServer {
	return &grpcService{service: service}
}

func (s *grpcService) AssociateDriverToVehicle(ctx context.Context, req *proto.AssociateDriverToVehicleRequest) (*proto.AssociateDriverToVehicleResponse, error) {
	if err := s.service.AssociateDriverToVehicle(int(req.GetDriverId()), int(req.GetVehicleId())); err != nil {
		return nil, err
	}
	return &proto.AssociateDriverToVehicleResponse{}, nil
}

func (s *grpcService) GetAllAssociations(ctx context.Context, req *proto.GetAllAssociationsRequest) (*proto.GetAllAssociationsResponse, error) {
	associations, err := s.service.GetAllAssociations()
	if err != nil {
		return nil, err
	}
	var grpcAssociations []*proto.Association
	for _, assoc := range associations {
		grpcAssociations = append(grpcAssociations, &proto.Association{
			DriverId:  assoc.DriverId,
			VehicleId: assoc.VehicleId,
		})
	}
	return &proto.GetAllAssociationsResponse{Associations: grpcAssociations}, nil
}

func (s *grpcService) GetDriversByVehicle(ctx context.Context, req *proto.GetDriversByVehicleRequest) (*proto.GetDriversByVehicleResponse, error) {
	driverIDs, err := s.service.GetDriversByVehicle(int(req.GetVehicleId()))
	if err != nil {
		return nil, err
	}
	var grpcDriverIDs []int32
	for _, id := range driverIDs {
		grpcDriverIDs = append(grpcDriverIDs, int32(id))
	}
	return &proto.GetDriversByVehicleResponse{DriverIds: grpcDriverIDs}, nil
}

func (s *grpcService) GetVehiclesByDriver(ctx context.Context, req *proto.GetVehiclesByDriverRequest) (*proto.GetVehiclesByDriverResponse, error) {
	vehicleIDs, err := s.service.GetVehiclesByDriver(int(req.GetDriverId()))
	if err != nil {
		return nil, err
	}
	var grpcVehicleIDs []int32
	for _, id := range vehicleIDs {
		grpcVehicleIDs = append(grpcVehicleIDs, int32(id))
	}
	return &proto.GetVehiclesByDriverResponse{VehicleIds: grpcVehicleIDs}, nil
}

func (s *grpcService) DissociateDriverFromVehicle(ctx context.Context, req *proto.DissociateDriverFromVehicleRequest) (*proto.DissociateDriverFromVehicleResponse, error) {
	if err := s.service.DissociateDriverFromVehicle(int(req.GetDriverId()), int(req.GetVehicleId())); err != nil {
		return nil, err
	}
	return &proto.DissociateDriverFromVehicleResponse{}, nil
}
