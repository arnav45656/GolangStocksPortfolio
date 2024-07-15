import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const getUser = (userName) => {
    return axios.get(`${API_URL}/users/username/${userName}`)
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const createUser = (user) => {
    return axios.post(`${API_URL}/users`, user) // Send user directly
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const updateUser = (id, user) => {
    return axios.put(`${API_URL}/users/${id}`, user)
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const deleteUser = (id) => {
    return axios.delete(`${API_URL}/users/${id}`)
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const getInvestments = (userId) => {
    return axios.get(`${API_URL}/users/${userId}/investments`)
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const createInvestment = (userId, investment) => {
    return axios.post(`${API_URL}/users/${userId}/investments`, investment)
        .then(response => response.data)
        .catch(error => { throw error; });
};
