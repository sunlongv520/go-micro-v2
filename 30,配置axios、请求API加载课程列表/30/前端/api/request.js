import axios from 'axios'
var service = axios.create({
  baseURL: 'http://localhost:9000',
  timeout: 1000
});

export default service
