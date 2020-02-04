import { get, post ,put ,deletes,upload} from './axios'
const baseUrl = process.env.API_SOTA

/******************* 产品管理 start *******************/


//查询首页栏目列表
export const getIndexList= () => {
  return get(`${baseUrl}/index`)
}

//单个栏目详情列表
export const getIndexDtlList= (index) => {
  return get(`${baseUrl}/category/${index}`)
}


export const getApiList= (apiClass) => {
  return get(`${baseUrl}/api-list?APIClass=${apiClass}`)
}


export const getApiGroup= (APIGroup='',SDKName='') => {
  return get(`${baseUrl}/api-group?APIGroup=${APIGroup}&SDKName=${SDKName}`)
}



