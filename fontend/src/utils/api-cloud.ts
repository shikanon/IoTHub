import { get, post ,put ,deletes} from './axios'


const cloudUrl = process.env.API_CLOUD

/******************** 视频云 start********************/


//新增
export const addWorkSpace = (data) => {
 
    return post(`${cloudUrl}/workspace`, data)
  }
  
  //删除
  export const deleteWorkSpace= (WorkSpaceId) => {
    
    return deletes(`${cloudUrl}/workspace?`,{"id":WorkSpaceId})
  }
  
  //修改
  export const updateWorkSpace= (data) => {
    return   put(`${cloudUrl}/workspace`, data)
  }
  
  //查询列表
  export const getWorkSpaceList= (currentPage = 1,pageSize = 10,workSpaceName = "") => {
    return get(`${cloudUrl}/workspaces?page=${currentPage}&item=${pageSize}&name=${workSpaceName}`)
  }
  

  export const getSimpleWorkSpaceList = () => {
    return get(`${cloudUrl}/wss`)
  }
  

  
  //查看
  export const getWorkSpaceDtl= ( WorkSpaceId = 0) => {
    return get(`${cloudUrl}/workspace?&id=${WorkSpaceId}`)
  }
  
  //停用
  export const stopWorkSpace= ( data) => {
    return put(`${cloudUrl}/wsstatus`,data)
  }
  
  
  
//新增
export const addCamera = (data) => {
 
    return post(`${cloudUrl}/camera`, data)
  }
  
  //删除
  export const deleteCamera= (cameraIds = []) => {
    
    return deletes(`${cloudUrl}/camera?`,{"id":cameraIds})
  }
  
  //修改
  export const updateCamera= (data) => {
    return   put(`${cloudUrl}/camera`, data)
  }
  
  //查询列表
  export const getCameraList= (currentPage = 1,pageSize = 10,sin = "",workSpaceId= 0) => {
    return get(`${cloudUrl}/cameras?page=${currentPage}&item=${pageSize}&sin=${sin}&ws=${workSpaceId}`)
  }
  
  //查看
  export const getCameraDtl= ( cameraId = 0) => {
    return get(`${cloudUrl}/camera?&id=${cameraId}`)
  }
  
  //停用/开启
  export const changeCameraStatu= ( cameraIds = []) => {
      alert(1)
    return put(`${cloudUrl}/cstatus`,{"id":cameraIds})
  }


  export const getSimpleCameraList = () => {
    return get(`${cloudUrl}/cas`)
  }
  
  /******************** 视频云 end********************/