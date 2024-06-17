package main

import (
	"OrderManagementService/proto"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewOrderManagementClient(conn)

	// Test CreateOrder
	orderID := createOrder(client, "user_1", "item_456", "restaurant_789", "123 Delivery St")

	// Test AssignRider
	assignRider(client, orderID)

	// Test MarkOrderDelivered
	markOrderDelivered(client, orderID)

	// Test GetAllOrders
	getAllOrders(client)
}

func getAllOrders(client proto.OrderManagementClient) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.GetAllOrdersRequest{}

	res, err := client.GetAllOrders(ctx, req)
	if err != nil {
		log.Fatalf("GetAllOrders failed: %v", err)
	}

	for _, order := range res.Orders {
		fmt.Printf("OrderID: %v\n", order.OrderId)
		fmt.Printf("UserID: %v\n", order.UserId)
		fmt.Printf("Timestamp: %v\n", order.Timestamp)
		fmt.Printf("ItemID: %v\n", order.ItemId)
		fmt.Printf("RiderID: %v\n", order.RiderId)
		fmt.Printf("RestaurantID: %v\n", order.RestaurantId)
		fmt.Printf("DeliveryAddress: %v\n", order.DeliveryAddress)
		fmt.Printf("Assigned: %v\n", order.Assigned)
		fmt.Printf("Delivered: %v\n", order.Delivered)
		fmt.Println("----")
	}
}

func createOrder(client proto.OrderManagementClient, userID, itemID, restaurantID, deliveryAddress string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.CreateOrderRequest{
		UserId:          userID,
		ItemId:          itemID,
		RestaurantId:    restaurantID,
		DeliveryAddress: deliveryAddress,
	}

	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("CreateOrder failed: %v", err)
	}

	fmt.Printf("Order created successfully: %v\n", res.OrderId)
	return res.OrderId
}

func assignRider(client proto.OrderManagementClient, orderID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.AssignRiderRequest{OrderId: orderID}

	res, err := client.AssignRider(ctx, req)
	if err != nil {
		log.Fatalf("AssignRider failed: %v", err)
	}

	fmt.Printf("Rider assigned successfully: %v\n", res.RiderId)
}

func markOrderDelivered(client proto.OrderManagementClient, orderID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.MarkOrderDeliveredRequest{OrderId: orderID}

	_, err := client.MarkOrderDelivered(ctx, req)
	if err != nil {
		log.Fatalf("MarkOrderDelivered failed: %v", err)
	}

	fmt.Println("Order marked as delivered successfully")
}
