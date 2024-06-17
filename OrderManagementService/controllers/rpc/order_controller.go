package rpc

import (
	"OrderManagementService/proto"
	"OrderManagementService/services"
	"context"
)

type OrderController struct {
	proto.UnimplementedOrderManagementServer
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	orderID, err := c.orderService.CreateOrder(ctx, req.UserId, req.ItemId, req.RestaurantId, req.DeliveryAddress)
	if err != nil {
		return nil, err
	}
	return &proto.CreateOrderResponse{OrderId: orderID, Success: true}, nil
}

func (c *OrderController) AssignRider(ctx context.Context, req *proto.AssignRiderRequest) (*proto.AssignRiderResponse, error) {
	riderID, err := c.orderService.AssignRider(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &proto.AssignRiderResponse{Success: true, RiderId: riderID}, nil
}

func (c *OrderController) MarkOrderDelivered(ctx context.Context, req *proto.MarkOrderDeliveredRequest) (*proto.MarkOrderDeliveredResponse, error) {
	err := c.orderService.MarkOrderDelivered(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &proto.MarkOrderDeliveredResponse{Success: true}, nil
}

func (c *OrderController) GetAllOrders(ctx context.Context, req *proto.GetAllOrdersRequest) (*proto.GetAllOrdersResponse, error) {
	orders, err := c.orderService.GetAllOrders(ctx)
	if err != nil {
		return nil, err
	}
	var protoOrders []*proto.Order
	for _, order := range orders {
		protoOrders = append(protoOrders, &proto.Order{
			OrderId:         order.OrderID,
			UserId:          order.UserID,
			Timestamp:       order.Timestamp.String(),
			ItemId:          order.ItemID,
			RiderId:         order.RiderID,
			RestaurantId:    order.RestaurantID,
			DeliveryAddress: order.DeliveryAddress,
			Assigned:        order.Assigned,
			Delivered:       order.Delivered,
		})
	}
	return &proto.GetAllOrdersResponse{Orders: protoOrders}, nil
}
