import React, { useEffect, useState } from 'react';
import { Button, Container, Form, Table } from 'react-bootstrap';
import OrderTable from './OrderTable';
import RiderTable from './RiderTable';
import { assignRiderApi, completeOrderApi, createNewOrderApi, createRiderApi, getFreeRidersApi, getOrdersApi } from '../api';

const Dashboard = () => {
    const [orders, setOrders] = useState([]);

    const [freeRiders, setFreeRiders] = useState([]);

    const [newOrder, setNewOrder] = useState({});

    const [newRider, setNewRider] = useState({
        name: '',
    });

    async function fetchOrders() {
        getOrdersApi().then((response) => {
            setOrders(response.data);
        })

        getFreeRidersApi().then((response) => {
            setFreeRiders(response.data);
        });
    }

    const handleOrderChange = (e) => {
        setNewOrder({
            ...newOrder,
            [e.target.name]: e.target.value,
        });
    };

    const handleRiderChange = (e) => {
        setNewRider({
            ...newRider,
            [e.target.name]: e.target.value,
        });
    };

    const createOrder = () => {
        createNewOrderApi(newOrder).then((response) => {
            console.log(response);
            fetchOrders();
            setNewOrder({
                user_id: '',
                restaurant_id: '',
                item_id: '',
                delivery_address: '',
            });
        }).catch((error) => {
            console.log(error.response?.data.message);
            console.log(error);
        });
    };

    const createRider = () => {
        createRiderApi(newRider).then((response) => {
            console.log(response);
        }).then(()=>{
            fetchOrders();
            setNewRider({
                name: '',
            });
        })
    };

    const assignRider = (orderId, riderId) => {
        assignRiderApi(orderId).then((response) => {
            fetchOrders();
        });
    };

    const deliverOrder = (orderId) => {
        completeOrderApi(orderId).then((response) => {
            fetchOrders();
        });
    };

    useEffect(() => {
        fetchOrders();
    }, []);

    return (
        <Container >
          
            {OrderTable(orders, assignRider, freeRiders, deliverOrder)}
           
            {RiderTable(freeRiders)}
            <h2>Create Order</h2>
            <Form>
                <Form.Group controlId="formBasicUserId">
                    <Form.Label>User ID:</Form.Label>
                    <Form.Control type="number" name="user_id" value={newOrder.user_id} onChange={handleOrderChange} />
                </Form.Group>
                <Form.Group controlId="formBasicRestaurantId">
                    <Form.Label>Restaurant ID:</Form.Label>
                    <Form.Control type="number" name="restaurant_id" value={newOrder.restaurant_id} onChange={handleOrderChange} />
                </Form.Group>
                <Form.Group controlId="formBasicItemId">
                    <Form.Label>Item ID:</Form.Label>
                    <Form.Control type="text" name="item_id" value={newOrder.item_id} onChange={handleOrderChange} />
                </Form.Group>
                <Form.Group controlId="formBasicDeliveryAddress">
                    <Form.Label>Delivery Address:</Form.Label>
                    <Form.Control type="text" name="delivery_address" value={newOrder.delivery_address} onChange={handleOrderChange} />
                </Form.Group>
                <Button variant="success" onClick={createOrder}>
                    Create Order
                </Button>
            </Form>
            <h2>Create Rider</h2>
            <Form>
                <Form.Group controlId="formBasicRiderName">
                    <Form.Label>Name:</Form.Label>
                    <Form.Control type="text" name="name" value={newRider.name} onChange={handleRiderChange} />
                </Form.Group>
                <Button variant="success" onClick={createRider}>
                    Create Rider
                </Button>
            </Form>
        </Container>
    );
};

export default Dashboard;
