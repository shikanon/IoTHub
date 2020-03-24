<template>
  <div name="device-list">
        <el-row >
           <el-select v-model="workSpaceId" placeholder="全部空间"  @change="changeWorkSpaceId">
              <el-option label="全部空间" value=""></el-option>
              <el-option
              v-for="item in workSpaceArr"
              :key="item.id"
              :label="item.name"
              :value="item.id">
              </el-option>
            </el-select>
            <el-button type="primary" @click="addCameraVisible = true"    size="medium">添加摄像头</el-button>
           
             <el-input
                v-model="searchVal"
                placeholder="请输入摄像头SIN码查询"
                class="search-input"
                clearable
                size="medium"
                @keyup.enter.native="getCameraList(1)"
                >
              <i slot="suffix" class="el-input__icon el-icon-search" @click="getCameraList(1)"></i>
            </el-input>    
        </el-row>  
   
        <el-table
            ref="multipleTable"
            :data="tableData"
            tooltip-effect="dark"
            style="width: 100%"
            @selection-change="handleSelectionChange">
        <el-table-column
            type="selection"
            width="55">
        </el-table-column>     

        <el-table-column
            label="摄像头sin码"    
            >
            <template slot-scope="scope">{{ scope.row.sin }}</template>
        </el-table-column>
        <el-table-column
            prop="space_name"
            label="工作空间"
            >
        </el-table-column>
        <el-table-column 
            label="摄像头接入地址" 
            show-overflow-tooltip>
                <template slot-scope="scope">{{ scope.row.access_address }}</template>
        </el-table-column>
        <el-table-column
            label="状态"
            width="100"
           >
            <template slot-scope="scope">
            
                 <span> {{ scope.row.status_id === 0 ? '启用':'停用'}}</span>

            </template>
        </el-table-column>
        <el-table-column
            label="最后连接时间"  width="150">
            <template slot-scope="scope">{{ scope.row.last_connect_time  }}</template>
        </el-table-column>
        <el-table-column 
                fixed="right"
                label="操作"
                width="200">
                <template slot-scope="scope">
                <el-button @click="goToLaboratoryPage(scope.row.camera_id)" type="text" size="small">实验室 </el-button>
                <el-divider direction="vertical"></el-divider>                
                <el-button   type="text" size="small" @click="clickHandle( scope.row.camera_id,1)">删除</el-button>
                <el-divider direction="vertical" ></el-divider >                
                <el-button   v-if="scope.row.status_id === 0" type="text" size="small" @click="clickHandle( scope.row.camera_id,2)">停用</el-button>
                <el-button   v-if="scope.row.status_id === 1" type="text" size="small" @click="clickHandle( scope.row.camera_id,3)">开启</el-button>

                </template>
            </el-table-column>
        </el-table>
         <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
        <div   style="margin-top: 20px;margin-left: 15px;" >
            <el-checkbox v-model="selectFlg" ></el-checkbox>
            <el-button @click="handleSelection('1')" :disabled="!selectFlg" type="primary" size="medium">删除</el-button>
            <el-button @click="handleSelection('2')" :disabled="!selectFlg" type="primary" size="medium">停用</el-button>
            <el-button @click="handleSelection('3')" :disabled="!selectFlg" type="primary" size="medium">启用</el-button>
        </div>
      
 
        
        <el-dialog title="添加摄像头" :visible.sync="addCameraVisible" width="25%">
            <AddCamera ref="addCamera" :workSpaceArr="workSpaceArr" :workSpaceId="workSpaceId" @close="addCameraSuccess"></AddCamera>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addCameraSubmit">确 定</el-button>
                <el-button @click="addCameraCancel">取 消</el-button>
            </span>
        </el-dialog>
    
        
      
        
  </div>
</template>

<script>
  import AddCamera from './AddCamera'
  import DeviceCertificate from '../device/DeviceCertificate'
  import AddDeviceGroup from '../device/AddDeviceGroup'
  import AddResult from '../device/AddResult'
  import DeviceListCheck from '../device/DeviceListCheck'


  export default {
      name:"device-list",
      components: { AddCamera,DeviceCertificate,AddDeviceGroup,AddResult,DeviceListCheck},  
    
      data() {
        return {
          tableData: [],
          multipleSelection: [],
          workSpaceId:"",
          workSpaceArr:[],
          addCameraVisible:false,   
          search:'',
          product:'',       
          currentPage:1,
          pageSize:10,
          total:0,
          searchVal:'',     
        }
    },
    computed:{
      selectFlg:function(){
          return this.multipleSelection.length > 0
      }
    },

    watch:{
        //监听productId,若发生变化，重新查询设备列表
        workSpaceId:{  
            handler:function(val,oldval){ 
                if(val!=oldval){
                    this.$nextTick(()=>{
                        this.currentPage = 1
                        this.getCameraList()
                    })
                }
            },  
            immediate:true,//关键
            deep:true
          },
      },
  
    created(){
        if(this.$route.params.workSpaceId){
          this.workSpaceId = this.$route.params.workSpaceId
        }

         this.getSimpleWorkSpaceList()
     
         this.getCameraList() 
          
    },

    methods: {

        //cha
        getSimpleWorkSpaceList(){
             this.$API_CLOUD.getSimpleWorkSpaceList().then((res) => {
                this.workSpaceArr = res.data.data
            })
        },
        //查询摄像头列表
        getCameraList(currentPage){
          if(currentPage){
            this.currentPage = currentPage
          }
    
          this.$API_CLOUD.getCameraList(this.currentPage,this.pageSize,this.searchVal,this.workSpaceId).then((res) => {
            if(res.data.status === "Y"){
                this.tableData = res.data.data.data_list
                this.total = res.data.data.num_results
                let params = {}
                params.device_count = res.data.data.num_results
                params.device_active_count =  res.data.data.device_active_count
                params.device_online_count =  res.data.data.device_online_count
                this.$emit('setAcount',params)
            }
                //this.tableData
            })

        },

        //改变每页数量
        handleSizeChange(val){
          this.pageSize = val
          this.getCameraList(1)
        },

        //改变当前页数
        handleCurrentChange(val){
            this.currentPage = val
            this.getCameraList()
        },

        //改变多选框的回调函数
        handleSelectionChange(val) {
          this.multipleSelection = val;
        },

        //筛选操作
        filterHandler(value, row, column) {
          const property = column['property'];
          return row[property] === value;
        },

        changeWorkSpaceId(val){
          this.workSpaceId = val
        },
     
        //确认添加设备
        addCameraSubmit(){  
            this.$refs['addCamera'].submitForm()
        },

        //取消添加设备
        addCameraCancel(){
          this.addCameraVisible = false
          this.$refs['addCamera'].init()
        },
        addCameraSuccess(){
          this.addCameraVisible = false
          this.getCameraList()
        },

         
        //跳转到实验室页面
        goToLaboratoryPage(cameraId){
          //this.openRouter = true 
          this.$router.push({name :'laboratory-list',params: {cameraId:cameraId}}) 
        },

        //操作单条纪录，,type 1-删除，2-停用
        clickHandle(cameraId,type = 1){
          let msg = '此操作将永久删除该纪录, 是否继续?'
          if(type === 2){
                msg ='此操作将停用该摄像头, 是否继续?'
          }else if(type === 3){
             msg ='此操作将开启该摄像头, 是否继续?'
          }
            
           
            this.$confirm(msg, '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
                  let arr = new Array();
                  arr.push(cameraId)
                  if(type === 1){
                       this.deleteCameras(arr)
                  }else if(type === 2 || type === 3){
                      this.changeCameraStatu(arr)
                  }
                
                  //this.tableData =  this.tableData.filter(item => item.name !== row.name);
            }).catch(() => {
              // this.$message({
              //   type: 'info',
              //   message: '已取消删除'
              // });          
            });  
        },

     
       

        //多选操作，1-删除，2-禁用，3-启用,4-从分组中删除
        handleSelection(type){
          
          let arr = this.multipleSelection.map((item) => {
            return item.id
          })

          //.join(",")
       
          if(type === '1'){//删除
            this.deleteCameras(arr)
          }else if(type === '2'){//禁用
            this.changeCameraStatu(arr)
          }else if(type === '3'){//启用
            this.changeCameraStatu(arr)
          }else if(type === '4'){
              this.deleteCamerasFromGroup(arr)
          }
          
        },

        //删除设备
        deleteCameras(arr){
          console.log(arr)
          this.$API_CLOUD.deleteCamera(arr).then((res) => {
                if(res.data.status  === 'Y'){
                  this.$message.success('删除成功')
                  this.getCameraList()
                }else{
                  this.$message.error(res.message);
                }              
            })
        },

         //停用/开启
        changeCameraStatu(arr){
          this.$API_CLOUD.changeCameraStatu(arr).then((res) => {
                if(res.data.status  === 'Y'){
                  this.$message.success('操作成功')
                  this.getCameraList()                             
                }else{
                    this.$message.error(res.message);
                }              
            })
        },

         //变更单个设备状态
        changeDeviceStatus(deviceInf){
          console.log(deviceInf.id)
          let array = new Array();
          array.push(deviceInf.id)
          if(deviceInf.status_id === 1){ //未激活 -> 激活
            this.activeDevices(array)
          }else { //禁用
            this.disabledDevices(array)
          }
        },
       
        //启用设备
        activeDevices(arr){
          this.$API_CLOUD.activeDevices(arr).then((res) => {
                if(res.data.status  === 'Y'){
                  this.$message.sucess('启用成功',function(){
                    this.getCameraList()
                  })
                }else{
                    this.$message.error(res.message);
                }              
            })
        },
                    

        //改变查询条件
        queryParamChange(value){
            this.searchVal = ""
        },

     

        

    }
  }
</script>
<style scoped>
 .loadBtn{
     margin-left: 65px;
 }
 .statu-circle{
     padding: 5px !important;
     margin-left: 10px;
 }
 .statuCircle:before{
     content:'';
     width: 10px;
    height: 10px;
    border-radius: 50%;
 }
 .success{background-color: #b3e19d}
 .info{background-color: #c8c9cc}
 .danger{background-color: #fab6b6}
 .warning{background-color: #f3d19e}
 .search-input{width: 200px;}
</style>