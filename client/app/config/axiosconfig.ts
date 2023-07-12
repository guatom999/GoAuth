// import axios, { AxiosError, InternalAxiosRequestConfig } from 'axios'

// const onRequest = ( config : InternalAxiosRequestConfig) => {
//     const token = localStorage.getItem("token");
//     config.headers.Authorization = 'Bearer ' + token;
//     return config
// }

// const onRequestError = ( err : AxiosError) => {

//     console.log("error is ===>" , err)

//     return err

// }

// axios.interceptors.request.use(onRequest,onRequestError)