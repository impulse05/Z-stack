import React from 'react'
import { Table } from 'react-bootstrap';
export default function RiderTable(freeRiders) {
    return <>
        <h2>Free Riders</h2>
    <Table striped bordered hover>
        <thead>
            <tr>
                <th>Rider ID</th>
                <th>Name</th>
                <th>Assigned</th>
            </tr>
        </thead>
        <tbody>
            {freeRiders.map((rider) => (
                <tr key={rider.ID}>
                    <td>{rider.ID}</td>
                    <td>{rider.Name}</td>
                    <td>{rider.Assigned ? 'Yes' : 'No'}</td>
                </tr>
            ))}
        </tbody>
    </Table >
    </>;
}