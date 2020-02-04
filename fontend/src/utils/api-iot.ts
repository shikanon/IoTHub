import { get, post ,put ,deletes,upload} from './axios'

const baseUrl = process.env.API_IOT
 // 登录
export const login = (data) => {
  return post(`/login`, data)
}

/******************* 产品管理 start *******************/

//新增
export const addProduct = (data) => {
 
  return post(`${baseUrl}/product`, data)
}

//删除
export const deleteProduct= (productId) => {
  
  return deletes(`${baseUrl}/product?`,{"pid":productId})
}

//修改
export const updateProduct= (data) => {
  return   put(`${baseUrl}/product`, data)
}

//查询列表
export const getProductList= (currentPage = 1,pageSize = 10,productName = "") => {
  //&productName=${productName}
  return get(`${baseUrl}/products?page=${currentPage}&item=${pageSize}`)
}


//查询产品详情
export const getProductById= (productKey) => {
  return get(`${baseUrl}/product?pid=${productKey}`)
}


//获取所有产品的id和名称

export const getSimpleProductList= () => {
  return get(`${baseUrl}/simpleproducts`)
}


//查询产品物模型
export const getProductMode= (productKey) => {
  return get(`${baseUrl}/tsl?pid=${productKey}`)
}

/******************* 产品管理 end *******************/

//查询节点类型
export const getNodeTypeList= () => {
  return get(`${baseUrl}/nodetypes`)
}

//查询联网方式
export const getNetTypeList= () => {
  return get(`${baseUrl}/networkways`)
}

//查询认证方式
export const getDataFormatList= () => {
  return get(`${baseUrl}/dataformats`)
}

//查询认证方式
export const getAuthTypeList= () => {
  return get(`${baseUrl}/authmethods`)
}


//查询topic列表

export const getTopicList= (type,key) => {
  if(type === 'product'){
    return get(`${baseUrl}/ptopics?pid=${key}`)
  }else{
    return get(`${baseUrl}/dtopics?did=${key}`)
  }
}


//添加topic

export const addTopic = (data) => {
 
  return post(`${baseUrl}/ptopic`, data)
}


//删除
export const deleteTopic= (topicId) => {
  
 
  return deletes(`${baseUrl}/ptopic`,{"tid":topicId})
}



export const updateTopic= (data) => {
  // "tid":   23,       // Topic的id
  // "name": "aaaaa",  // Topic类名
  // "operation": 2,   // 操作权限id
  // "desc":  ""       // 描述
  return   put(`${baseUrl}/ptopic`, data)
}

//查询标准品类
export const getCategoryList= (currentPage = 1,pageSize = 10,type = 1) => {
  return get(`${baseUrl}/models`)
}


//查询标准品类属性
export const getModelfuncs= (id ) => {
  return get(`${baseUrl}/modelfuncs?id=${id}`)
}





/******************* 设备管理 start *******************/

//新增
export const addDevice = (data) => {
  return post(`${baseUrl}/device`, data)
}

//批量自动新增
export const addDeviceArr = (data) => {
  return post(`${baseUrl}/adevice`, data)
}
//删除
export const deleteDevice= (dids = []) => {
  return deletes(`${baseUrl}/device`,{"dids":dids})
}

//修改
export const updateDevice= (productKey,data) => {
  return   put(`${baseUrl}/device?productKey=${productKey}`, data)
}

//查询列表(页码，每页数量，产品Id，设备名称，分组id)
export const getDeviceList= (currentPage = 1,pageSize = 10,productId = 0,deviceName ='',groupId ='') => {
  //&pid=${productId}$deviceName=${deviceName}$groupId=${groupId}
  return get(`${baseUrl}/devices?page=${currentPage}&item=${pageSize}&pid=${productId}`)
}

//查询设备详情
export const getDeviceDtl= ( deviceId = 0) => {
  return get(`${baseUrl}/device?&did=${deviceId}`)
}

//批次列表
export const getApplyList= (currentPage = 1,pageSize = 10) => {
  return get(`${baseUrl}/apply?page=${currentPage}&results_per_page=${pageSize}`)
}


//查询批次详情
export const getApplyListDtl= (currentPage = 1,pageSize = 10,ApplyId,ProductKey) => {
  return get(`${baseUrl}/applyDtl?page=${currentPage}&results_per_page=${pageSize}&ApplyId=${ApplyId}`)
}


//更新设备状态(禁用/启用)
export const updateDeviceStatu= (deviceKeyArr,statu) => {
  return   put(`${baseUrl}device?deviceKey=${deviceKeyArr}&statu=${statu}`, )
}

//查询设备详情
export const getRunState= ( deviceId = 0,type ='pro') => {
 if(type === 'pro'){
  return get(`${baseUrl}/prostatus?&did=${deviceId}`)
 }else{
  return get(`${baseUrl}/desstatus?&did=${deviceId}`)
 }
//   /iot/api/v1/desstatus?did=1    获取期望状态 get请求
// /iot/api/v1/prostatus?did=1   获取实时状态  get请求
 // return get(`${baseUrl}/desstatus?&did=${deviceId}`)
}



/******************* 设备管理 end *******************/


/******************* 分组管理 start *******************/

//新增
export const addGroup = (data) => {
  return post(`${baseUrl}group`, data)
}

//删除
export const deleteGroup= (DeviceKey) => {
  return deletes(`${baseUrl}/group?DeviceKey=${DeviceKey}`)
}

//修改
export const updateGroup= (productKey,data) => {
  return   put(`${baseUrl}group?productKey=${productKey}`, data)
}

//查询列表
export const getGroupList= (currentPage = 1,pageSize = 10,groupName) => {
  return get(`${baseUrl}/group?page=${currentPage}&results_per_page=${pageSize}&groupName=${groupName}`)
}


//查询分组详情
export const getGroupDtl= (currentPage = 1,pageSize = 10,groupId) => {
  return get(`${baseUrl}/group?page=${currentPage}&results_per_page=${pageSize}&groupId=${groupId}`)
}

/******************* 分组管理 end *******************/



/*********   设备分组关联管理 start  **************/


//新增设备到分组
export const connectDeviceToGroup= (productKey,data) => {
  return   put(`${baseUrl}group?productKey=${productKey}`, data)
}

//将分组在分组中删除

export const disConnectDeviceToGroup= (productKey,data) => {
  return   put(`${baseUrl}group?productKey=${productKey}`, data)
}

/*********   设备分组关联管理 end  **************/

















export const getUser = (currentPage = 1,pageSize = 10) => {
  return get(`${baseUrl}/user?page=${currentPage}&results_per_page=${pageSize}`)
}


// test获取营业厅信息
export const getHall = (currentPage, pageSize) => {
  return get(`/hall?page=${currentPage}&results_per_page=${pageSize}`)
}

