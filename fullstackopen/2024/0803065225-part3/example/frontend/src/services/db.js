import axios from 'axios'
const base = 'http://localhost:3001/api/notes'

const getAll = () => {
    return axios.get(base).then(response => response.data)
}

const get = id => {
    return axios.get(`${base}/${id}`).then(response => response.data)
}

const create = obj => {
    return axios.post(base, obj).then(response => response.data)
}

const update = (id, obj) => {
    return axios.put(`${base}/${id}`, obj).then(response => response.data)
}

export default {get, getAll, create, update}
