
import axios from 'axios';

const getOrdersApi = () => {
    return axios.get("/api/orders");
}

const createNewOrderApi = (order) => {
    console.log(order);
    return axios.post("/api/orders", order)
}
  
const assignRiderApi = (orderId) => {
    return axios.patch(`/api/orders/${orderId}/assign_rider`);
}

const completeOrderApi = (orderId) => {
    return axios.patch(`/api/orders/${orderId}/deliver`);
}

const getFreeRidersApi = () => {
    return axios.get("/api/riders/free");
}

const createRiderApi = (rider) => {
    return axios.post("/api/riders", rider);
}

export {
    getOrdersApi,
    createNewOrderApi,
    assignRiderApi,
    completeOrderApi,
    getFreeRidersApi,
    createRiderApi
}