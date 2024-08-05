package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"fleet_management/proto"
)

type Client struct {
	driverClient      proto.DriverServiceClient
	vehicleClient     proto.VehicleServiceClient
	associationClient proto.AssociationServiceClient
}

func NewClient(address string) (*Client, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	return &Client{
		driverClient:      proto.NewDriverServiceClient(conn),
		vehicleClient:     proto.NewVehicleServiceClient(conn),
		associationClient: proto.NewAssociationServiceClient(conn),
	}, nil
}

func (c *Client) Close() {
	// Closing the connection is handled by the gRPC library when the client is garbage collected.
	// But we can add any additional cleanup if needed.
}

// Exemplo de uso do cliente
func (c *Client) CreateDriver(ctx context.Context, driver *proto.CreateDriverRequest) (*proto.DriverResponse, error) {
	return c.driverClient.CreateDriver(ctx, driver)
}

func (c *Client) CreateVehicle(ctx context.Context, vehicle *proto.CreateVehicleRequest) (*proto.VehicleResponse, error) {
	return c.vehicleClient.CreateVehicle(ctx, vehicle)
}

func (c *Client) AssociateDriverToVehicle(ctx context.Context, association *proto.AssociateDriverToVehicleRequest) (*proto.AssociateDriverToVehicleResponse, error) {
	return c.associationClient.AssociateDriverToVehicle(ctx, association)
}
