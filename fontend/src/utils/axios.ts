import axios from 'axios'
import router from '../router'
import store from '../store'
import {Message} from 'element-ui'
axios.defaults.timeout = 60 * 1000    //超时时间
axios.defaults.headers.post['Content-Type'] = 'application/json'        //配置请求头
//axios.defaults.baseURL = process.env.API_IOT   //配置接口地址
//axios.defaults.baseURL = '/api'   //配置接口地址
let cancelToken = axios.CancelToken
let retryCount = 5

//添加请求拦截器
axios.interceptors.request.use(config=>{
    // if (store.state.TOKEN) {
    //     config.headers.token = store.state.TOKEN
    //   } else {
    //     Message({
    //         type: 'warning',
    //         message: '登陆信息已过期，请重新登录',
    //         duration: 1500,
    //         offset: 80
    //       })
    //     //没有token信息，重新登陆
    //     router.push({ path: '/login'})
    //   }

     store.commit('deletePending',config)
     config.cancelToken = new cancelToken((c)=>{
        store.commit('pushPending',{ u: config.url + '&' + config.method, f: c })
    });
     return config;
   },error => {
     return Promise.reject(error);
});
 
//添加响应拦截器
axios.interceptors.response.use((response) =>{
   //在一个ajax响应后再执行一下取消操作，把已经完成的请求从pending中移除
    store.commit('deletePending',response.config)
    //没有token信息
    // if (response.data.errorCode === -2) {
    //     Message({
    //         type: 'warning',
    //         message: '登陆信息已过期，请重新登录',
    //         duration: 1500,
    //         offset: 80
    //       })
    //     store.commit('setTOKEN','')
    //     router.push({path: '/login' })   
    // } else if (response.data.errorCode === -1) {
    //     Message({
    //         type: 'warning',
    //         message: '用户无操作权限',
    //         duration: 1500,
    //         offset: 80
    //     })
        
    // }
    return response;
   },(error) =>{
    var originalRequest = error.config;
    //请求超时
    if(error.code == 'ECONNABORTED' && error.message.indexOf('timeout')!=-1 && !originalRequest._retry){
        originalRequest._retry = true
        originalRequest._retryCount = retryCount 
        //按2倍扩充超时时间
        originalRequest.timeout = originalRequest.timeout * 2
        //替换baseUrl
        originalRequest.url = originalRequest.url.replace(originalRequest.baseURL,'')
        //重发请求
        if(originalRequest._retryCount > 1){
            retryCount -= 1
            return axios.request(originalRequest);
       }else{
       
            return Promise.reject(error) 
       }
      
    }else{
        return Promise.reject(error) 

        //return { data: { } }; //返回一个空对象，主要是防止控制台报错
    }
     

   });


//返回一个Promise(发送post请求)
export function post(url, params ={}) {
    return new Promise((resolve, reject) => {
        axios.post(url, params)
            .then(response => {
                resolve(response);
            }, err => {
                reject(err);
            })
            .catch((error) => {
                reject(error)
            })
    })
}
//返回一个Promise(发送get请求)
export function get(url, params = {}) {

    return new Promise((resolve, reject) => {
        axios.get(url, {params: params})
            .then(response => {
                resolve(response)
            }, err => {
                reject(err)
            })
            .catch((error) => {
                reject(error)
            })
    })
}

//返回一个Promise(发送put请求)
export function put(url, params = {}) {
    return new Promise((resolve, reject) => {
        axios.put(url,  params)
            .then(response => {
                resolve(response)
            }, err => {
                reject(err)
            })
            .catch((error) => {
                reject(error)
            })
    })
}

//返回一个Promise(发送delete请求)
export function  deletes (url, params ={}) {
    return new Promise((resolve, reject) => {
        axios.delete(url, {data:params})
            .then(response => {
                resolve(response)
            }, err => {
                reject(err)
            })
            .catch((error) => {
                reject(error)
            })
    })
}

export function  upload (url, param ={}) {
    return new Promise((resolve, reject) => {
        axios.post(url, {params: param})
            .then(response => {
                resolve(response)
            }, err => {
                reject(err)
            })
            .catch((error) => {
                reject(error)
            })
    })
}
export default {
    post,
    get,
    put,
    deletes,
    upload,
}

