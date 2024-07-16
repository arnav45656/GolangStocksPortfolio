import axios from 'axios';

const API_URL = 'http://localhost:8080';

const getAuthHeaders = () => {
    const token = localStorage.getItem('jwtToken');
    return {
        headers: {
            Authorization: `Bearer ${token}`
        }
    };
};

export const getUser = (userName) => {
    return axios.get(`${API_URL}/users/username/${userName}`, getAuthHeaders())
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const createUser = (user) => {
    return axios.post(`${API_URL}/users`, user, getAuthHeaders())
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const updateUser = (id, user) => {
    return axios.put(`${API_URL}/users/${id}`, user, getAuthHeaders())
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const deleteUser = (id) => {
    return axios.delete(`${API_URL}/users/${id}`, getAuthHeaders())
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const getInvestments = (userId) => {
    return axios.get(`${API_URL}/users/${userId}/investments`, getAuthHeaders())
        .then(response => response.data)
        .catch(error => { throw error; });
};

export const createInvestment = (userId, investment) => {
    return axios.post(`${API_URL}/users/${userId}/investments`, investment, getAuthHeaders())
        .then(response => response.data)
        .catch(error => { throw error; });
};
