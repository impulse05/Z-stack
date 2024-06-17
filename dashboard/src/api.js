
import axios from 'axios';

const getOrdersApi = () => {
    return axios.get("http://localhost:8080/orders");
}

const createNewOrderApi = (order) => {
    console.log(order);
    return axios.post("http://localhost:8080/orders", order)
}
  
const assignRiderApi = (orderId) => {
    return axios.put(`http://localhost:8080/orders/${orderId}/assign_rider`);
}

const completeOrderApi = (orderId) => {
    return axios.put(`http://localhost:8080/orders/${orderId}/deliver`);
}

const getFreeRidersApi = () => {
    return axios.get("http://localhost:8080/riders/free");
}

const createRiderApi = (rider) => {
    return axios.post("http://localhost:8080/riders", rider);
}

export {
    getOrdersApi,
    createNewOrderApi,
    assignRiderApi,
    completeOrderApi,
    getFreeRidersApi,
    createRiderApi
}