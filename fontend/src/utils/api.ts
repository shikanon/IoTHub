import { get, post ,put ,deletes,upload} from './axios'



 // 登录
export const login = (data) => {
  return post(`/login`, data)
}

/******************* 产品管理 start *******************/

//新增
export const addProduct = (data) => {
  return post(`product`, data)
}

//删除
export const deleteProduct= (DeviceKey) => {
  return deletes(`/product?DeviceKey=${DeviceKey}`)
}

//修改
export const updateProduct= (productKey,data) => {
  return   put(`product?productKey=${productKey}`, data)
}

//查询列表
export const getProductList= (currentPage = 1,pageSize = 10,productName) => {
  return get(`/product?page=${currentPage}&results_per_page=${pageSize}&productName=${productName}`)
}


//查询产品详情
export const getProductDtlByKey= (currentPage = 1,pageSize = 10,productKey) => {
  return get(`/product?page=${currentPage}&results_per_page=${pageSize}&productKey=${productKey}`)
}

/******************* 产品管理 end *******************/

/******************* 设备管理 start *******************/

//新增
export const addDevice = (data) => {
  return post(`product`, data)
}

//批量自动新增(productName,count)

export const addDeviceArr = (data) => {
  return post(`product`, data)
}
//删除
export const deleteDevice= (DeviceKey) => {
  return deletes(`/device?DeviceKey=${DeviceKey}`)
}

//修改
export const updateDevice= (productKey,data) => {
  return   put(`device?productKey=${productKey}`, data)
}

//查询列表(页码，每页数量，产品名称，设备名称，分组id)
export const getDeviceList= (currentPage = 1,pageSize = 10,productName = '',deviceName ='',groupId ='') => {
  return get(`/device?page=${currentPage}&results_per_page=${pageSize}&productName=${productName}$deviceName=${deviceName}$groupId=${groupId}`)
}

//查询设备详情
export const getDeviceDtl= (currentPage = 1,pageSize = 10, deviceName) => {
  return get(`/device?page=${currentPage}&results_per_page=${pageSize}&productName=${deviceName}`)
}

//批次列表
export const getApplyList= (currentPage = 1,pageSize = 10) => {
  return get(`/apply?page=${currentPage}&results_per_page=${pageSize}`)
}


//查询批次详情
export const getApplyListDtl= (currentPage = 1,pageSize = 10,ApplyId,ProductKey) => {
  return get(`/applyDtl?page=${currentPage}&results_per_page=${pageSize}&ApplyId=${ApplyId}`)
}


//更新设备状态(禁用/启用)
export const updateDeviceStatu= (deviceKeyArr,statu) => {
  return   put(`device?deviceKey=${deviceKeyArr}&statu=${statu}`, )
}



/******************* 设备管理 end *******************/


/******************* 分组管理 start *******************/

//新增
export const addGroup = (data) => {
  return post(`group`, data)
}

//删除
export const deleteGroup= (DeviceKey) => {
  return deletes(`/group?DeviceKey=${DeviceKey}`)
}

//修改
export const updateGroup= (productKey,data) => {
  return   put(`group?productKey=${productKey}`, data)
}

//查询列表
export const getGroupList= (currentPage = 1,pageSize = 10,groupName) => {
  return get(`/group?page=${currentPage}&results_per_page=${pageSize}&groupName=${groupName}`)
}


//查询分组详情
export const getGroupDtl= (currentPage = 1,pageSize = 10,groupId) => {
  return get(`/group?page=${currentPage}&results_per_page=${pageSize}&groupId=${groupId}`)
}

/******************* 分组管理 end *******************/



/*********   设备分组关联管理 start  **************/


//新增设备到分组
export const connectDeviceToGroup= (productKey,data) => {
  return   put(`group?productKey=${productKey}`, data)
}

//将分组在分组中删除

export const disConnectDeviceToGroup= (productKey,data) => {
  return   put(`group?productKey=${productKey}`, data)
}

/*********   设备分组关联管理 end  **************/

















export const getUser = (currentPage = 1,pageSize = 10) => {
  return get(`/user?page=${currentPage}&results_per_page=${pageSize}`)
}


// test获取营业厅信息
export const getHall = (currentPage, pageSize) => {
  return get(`/hall?page=${currentPage}&results_per_page=${pageSize}`)
}

