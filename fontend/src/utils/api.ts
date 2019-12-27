import { get, post ,put ,deletes,upload} from './axios'



 // 登录
export const login = (data) => {
  return post(`/login`, data)
}


export const getUser = (currentPage = 1,pageSize = 10) => {
  return get(`/user?page=${currentPage}&results_per_page=${pageSize}`)
}


// 获取营业厅信息
export const getHall = (currentPage, pageSize) => {
  return get(`/hall?page=${currentPage}&results_per_page=${pageSize}`)
}