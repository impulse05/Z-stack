import React from 'react'
import { Table , Button } from 'react-bootstrap'

export default function OrderTable(orders, assignRider, freeRiders, deliverOrder) {
    return <>
        <h2>Orders</h2>
        <Table striped bordered hover>
            <thead>
                <tr>
                    <th>Order ID</th>
                    <th>User ID</th>
                    <th>Timestamp</th>
                    <th>Rider ID</th>
                    <th>Restaurant ID</th>
                    <th>Delivery Address</th>
                    <th>Status</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                {orders?.map((order) => (
                    <tr key={order.order_id}>
                        <td>{order.order_id}</td>
                        <td>{order.user_id}</td>
                        <td>{order.timestamp}</td>
                        <td>{order.rider_id}</td>
                        <td>{order.restaurant_id}</td>
                        <td>{order.delivery_address}</td>
                        {/* if assigned not delived then show assigned */}
                        <td>{order?.assigned ? (order?.delivered ? "Delivered" : "Assigned") : "Not Assigned"}</td>
                        <td>
                            {!order.assigned  && (
                                <Button variant="success" onClick={() => assignRider(order.order_id)}>
                                    Assign
                                </Button>
                            )}
                            {order.assigned && !order.delivered  && (
                                <Button variant="success" onClick={() => deliverOrder(order.order_id)}>
                                    Deliver
                                </Button>
                            )}
                        </td>
                    </tr>
                ))}
            </tbody>
        </Table>;
    </>
}
