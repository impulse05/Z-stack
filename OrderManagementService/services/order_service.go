package services

import (
	"OrderManagementService/model"
	"OrderManagementService/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userID, itemID, restaurantID, deliveryAddress string) (string, error)
	AssignRider(ctx context.Context, orderID string) (string, error)
	MarkOrderDelivered(ctx context.Context, orderID string) error
	GetAllOrders(ctx context.Context) ([]*model.Order, error)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{orderRepo: orderRepo}
}

func (s *orderService) CreateOrder(ctx context.Context, userID, itemID, restaurantID, deliveryAddress string) (string, error) {
	orderID := "order_" + time.Now().Format("20060102150405")
	order := &model.Order{
		OrderID:         orderID,
		UserID:          userID,
		Timestamp:       time.Now(),
		ItemID:          itemID,
		RestaurantID:    restaurantID,
		DeliveryAddress: deliveryAddress,
		Assigned:        false,
		Delivered:       false,
	}
	err := s.orderRepo.SaveOrder(order)
	return orderID, err
}

func (s *orderService) AssignRider(ctx context.Context, orderID string) (string, error) {
	// Fetch available riders from Carthero
	resp, err := http.Get("http://localhost:8081/riders/free")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch riders from Carthero")
	}

	var riders []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Assigned bool   `json:"assigned"`
	}

	err = json.NewDecoder(resp.Body).Decode(&riders)
	if err != nil {
		return "", err
	}

	if len(riders) == 0 {
		return "", fmt.Errorf("no available riders")
	}

	idNum := riders[0].ID

	riderID := strconv.Itoa(idNum)

	// Update order with assigned rider
	order, err := s.orderRepo.GetOrder(orderID)
	if err != nil {
		return "", err
	}
	if order == nil {
		return "", fmt.Errorf("order not found")
	}
	order.RiderID = riderID
	order.Assigned = true
	err = s.orderRepo.UpdateOrder(order)
	if err != nil {
		return "", err
	}

	// Update rider status in Carthero to Busy
	_, err = http.Get(fmt.Sprintf("http://localhost:8081/riders/%s/true", riderID))
	if err != nil {
		return "", err
	}

	return riderID, nil
}

func (s *orderService) MarkOrderDelivered(ctx context.Context, orderID string) error {
	order, err := s.orderRepo.GetOrder(orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return fmt.Errorf("order not found")
	}
	order.Delivered = true
	err = s.orderRepo.UpdateOrder(order)
	if err != nil {
		return err
	}

	// Update rider status in Carthero to Free
	_, err = http.Get(fmt.Sprintf("http://localhost:8081/riders/%s/false", order.RiderID))
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) GetAllOrders(ctx context.Context) ([]*model.Order, error) {
	return s.orderRepo.GetAllOrders()
}
