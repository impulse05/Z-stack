package server

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rider struct {
	ID       int `json:"ID"`
	Name     string
	Assigned bool
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(Cors())

	grpcClient, err := NewGRPCClient("localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/orders", func(c *gin.Context) {
		orders, err := grpcClient.GetAllOrders(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	})

	r.POST("/orders", func(c *gin.Context) {
		var req struct {
			UserId          string `json:"user_id"`
			ItemId          string `json:"item_id"`
			RestaurantId    string `json:"restaurant_id"`
			DeliveryAddress string `json:"delivery_address"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orderId, success, err := grpcClient.CreateOrder(context.Background(), req.UserId, req.ItemId, req.RestaurantId, req.DeliveryAddress)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"order_id": orderId, "success": success})
	})

	r.PATCH("/orders/:order_id/assign_rider", func(c *gin.Context) {
		orderId := c.Param("order_id")
		success, riderId, err := grpcClient.AssignRider(context.Background(), orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": success, "rider_id": riderId})
	})

	r.PATCH("/orders/:order_id/deliver", func(c *gin.Context) {
		orderId := c.Param("order_id")
		success, err := grpcClient.MarkOrderDelivered(context.Background(), orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": success})
	})

	r.POST("/riders", func(c *gin.Context) {
		var req struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// create rider
		err := createRider(req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"success": true})

	})

	r.GET("/riders/free", func(c *gin.Context) {
		riders, err := getFreeRider()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, riders)
	})

	r.GET("/", s.HelloWorldHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func createRider(name string) error {

	resp, err := http.Post("http://localhost:8081/riders", "application/json", bytes.NewBuffer([]byte(`{"name":"`+name+`"}`)))
	// this gives rider i

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	return nil

}

func getFreeRider() (riders []Rider, err error) {
	// call http end point to get free rider
	resp, err := http.Get("http://localhost:8081/riders/free")
	if err != nil {
		log.Fatal(err)
	}

	riders = make([]Rider, 0)
	// decode response
	json.NewDecoder(resp.Body).Decode(&riders)

	defer resp.Body.Close()

	return riders, nil
}

// return rider id
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, X-Requested-With, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
