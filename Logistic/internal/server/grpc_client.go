package server

import (
	"context"

	pb "Logistic/internal/proto"

	"google.golang.org/grpc"
)

type GRPCClient struct {
	orderManagementClient pb.OrderManagementClient
}

func NewGRPCClient(addr string) (*GRPCClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderManagementClient(conn)
	return &GRPCClient{orderManagementClient: client}, nil
}

func (c *GRPCClient) GetAllOrders(ctx context.Context) ([]*pb.Order, error) {
	req := &pb.GetAllOrdersRequest{}
	resp, err := c.orderManagementClient.GetAllOrders(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Orders, nil
}

func (c *GRPCClient) CreateOrder(ctx context.Context, userId, itemId, restaurantId, deliveryAddress string) (string, bool, error) {
	req := &pb.CreateOrderRequest{
		UserId:          userId,
		ItemId:          itemId,
		RestaurantId:    restaurantId,
		DeliveryAddress: deliveryAddress,
	}
	resp, err := c.orderManagementClient.CreateOrder(ctx, req)
	if err != nil {
		return "", false, err
	}
	return resp.OrderId, resp.Success, nil
}

func (c *GRPCClient) AssignRider(ctx context.Context, orderId string) (bool, string, error) {
	req := &pb.AssignRiderRequest{OrderId: orderId}
	resp, err := c.orderManagementClient.AssignRider(ctx, req)
	if err != nil {
		return false, "", err
	}
	return resp.Success, resp.RiderId, nil
}

func (c *GRPCClient) MarkOrderDelivered(ctx context.Context, orderId string) (bool, error) {
	req := &pb.MarkOrderDeliveredRequest{OrderId: orderId}
	resp, err := c.orderManagementClient.MarkOrderDelivered(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}
