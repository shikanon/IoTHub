<template>
  <div name="device-list">
        <el-row v-if="type === 'device'">
            <el-button type="primary" @click="addDeviceVisible = true"    size="medium">添加设备</el-button>
            <el-button @click="addDeviceGroupVisible = true"    size="medium">批量添加</el-button>
             <el-select v-model="searchKey" placeholder="DeviceName" @change="queryParamChange"  size="medium">
                <el-option label="DeviceName" value="DeviceName"></el-option> 
                <el-option label="备注名称" value="Nickname" ></el-option>
                </el-select>
             <el-input
                v-model="searchVal"
                :placeholder="searchKey === 'DeviceName' ? '请输入Device查询' : '请输入备注名称查询' "
                class="search-input"
                clearable
                size="medium"
                @keyup.enter.native="getDeviceList(1)"
                >
              <i slot="suffix" class="el-input__icon el-icon-search" @click="getDeviceList(1)"></i>
            </el-input> 
            <el-input
              placeholder="请选择设备标签"
              suffix-icon="el-icon-arrow-down"
              @focus="addLabelVisible = true"
              class="search-input"
              size="medium"
              v-module="label">
            </el-input>
        </el-row>  
        <el-row v-if="type === 'group'">
            <el-button type="primary" @click="drawer = true">添加设备到分组</el-button>
            <el-select v-model="product" placeholder="全部产品">
                    <el-option
                    v-for="item in products"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                    </el-option>
             </el-select>
              <el-input
                v-model="search"
                placeholder='请输入Device查询'
                class="search-input"
                clearable
                size="medium"
                @keyup.enter.native="getDeviceList(1)"
                >
              <i slot="suffix" class="el-input__icon el-icon-search" @click="getDeviceList(1)"></i>
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
            label="DeviceName"    
            >
            <template slot-scope="scope">{{ scope.row.name }}</template>
        </el-table-column>
        <el-table-column
            prop="their_product_name"
            label="设备所属产品"
            >
        </el-table-column>
        <el-table-column v-if="!drawer"
            prop="node_type"
            label="节点类型"
            width="120"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            label="状态/启用状态"
            width="200"
            show-overflow-tooltip
            :filters="[{text: '已禁用', value: 2}, {text: '未激活', value: 1},
            {text: '在线', value: 3}, {text: '离线', value: 4}]"
            :filter-method="filterHandler">
            <template slot-scope="scope">
                <el-button 
                    :type="scope.row.status_id === 3 ? 'success': scope.row.status_id === 1 ? 'info': scope.row.status_id === 2 ? 'danger':'warning'"  
                    circle 
                    size="mini"
                    class="statu-circle">
                </el-button>
                 <span> {{ scope.row.status_id | deviceStatusFilter}}</span>
                 <el-switch  v-model="scope.row.status"  @change="changeDeviceStatus(scope.row)"></el-switch>  

            </template>
        </el-table-column>
        <el-table-column
            label="最后上线时间"  width="200">
            <template slot-scope="scope">{{ scope.row.last_on_line_time  }}</template>
        </el-table-column>
        <el-table-column 
                fixed="right"
                label="操作"
                width="100">
                <template slot-scope="scope">
                <el-button v-if="type === 'device'" @click="goToDeviceDetails(scope.row.id)" type="text" size="small">查看 </el-button>
                <el-button  v-else  @click="goToDeviceList(scope.row.id)" type="text" size="small">查看</el-button>
                <el-divider direction="vertical"></el-divider>                
                <el-button  v-if="type === 'device'" type="text" size="small" @click="deleteClick('1', scope.row.id)">删除</el-button>
                <el-button  v-else type="text" size="small" @click="deleteClick('2',scope.row.id)">从分组中删除</el-button>
                </template>
            </el-table-column>
        </el-table>
         <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
        <div v-if="type === 'device'"  style="margin-top: 20px;margin-left: 15px;" >
            <el-checkbox v-model="selectFlg" ></el-checkbox>
            <el-button @click="handleSelection('1')" :disabled="!selectFlg" type="primary" size="medium">删除</el-button>
            <el-button @click="handleSelection('2')" :disabled="!selectFlg" type="primary" size="medium">禁用</el-button>
            <el-button @click="handleSelection('3')" :disabled="!selectFlg" type="primary" size="medium">启用</el-button>
        </div>
         <div v-if="type === 'gruop' && !drawer" style="margin-top: 20px;margin-left: 15px;" >
            <el-checkbox v-model="selectFlg" ></el-checkbox>
            <el-button @click="handleSelection('4')" :disabled="!selectFlg" type="primary" size="medium">从分组中删除</el-button>     
        </div>
        <el-dialog title="标签筛选" :visible.sync="addLabelVisible" width="26%">
            <AddLabel ref="addLabel"  type="device" @close="addLabelVisible = false"></AddLabel>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addLabelSubmit">确 定</el-button>
                <el-button @click="addLabelVisible = false">取 消</el-button>
            </span>
        </el-dialog>
        
        <el-dialog title="添加设备" :visible.sync="addDeviceVisible" width="25%">
            <AddDevice ref="addDevice" :productArr="productArr" :productId="productId" @close="addDeviceVisible = false"
            @showDeviceCertificate="showDeviceCertificate"></AddDevice>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addDeviceSubmit">确 定</el-button>
                <el-button @click="addDeviceCancel">取 消</el-button>
            </span>
        </el-dialog>
        <el-dialog title="查看设备证书" :visible.sync="deviceCertificateVisible" width="25%">
            <DeviceCertificate ref="deviceCertificate" :certificateInf="certificateInf" ></DeviceCertificate>
            <span slot="footer" class="dialog-footer">
                <CopyBtn :type ="'primary'"  :content="JSON.stringify(certificateInf) " :btnTxt="'一键复制'"></CopyBtn>
                <el-button  @click="deviceCertificateVisible = false">关闭</el-button>
            </span>
        </el-dialog>
         <el-dialog title="批量添加设备" :visible.sync="addDeviceGroupVisible" width="25%">
            <AddDeviceGroup ref="addDeviceGroup"  :productArr="productArr"   :productId="productId"
              @close="addDeviceGroupVisible = false" 
              @showAddResult="showAddResult"
              @addSuccess="addSuccess">
            </AddDeviceGroup>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addDeviceGroupSubmit">确 定</el-button>
                <el-button @click="addDeviceGroupCancel">取 消</el-button>
            </span>
        </el-dialog>
        <el-dialog :title="addResultTitle" :visible.sync="addResultVisible" width="25%">
            <AddResult :activities="activities"></AddResult>
            <el-button type="primary" v-if="loadBtn" class="loadBtn">下载设备证书</el-button>
            <span slot="footer" class="dialog-footer">
                <el-button @click="addResultVisible = false">关 闭</el-button>
            </span>
        </el-dialog>
        <el-drawer
          title="添加设备到分组"
          :visible.sync="drawer"
           size='35%'>
            <DeviceListCheck @submit="addDeviceToGroup"  @close="drawer = false" :products="products"></DeviceListCheck>        
        </el-drawer>
  </div>
</template>

<script>
  import AddDevice from './AddDevice'
  import DeviceCertificate from './DeviceCertificate'
  import AddDeviceGroup from './AddDeviceGroup'
  import AddResult from './AddResult'
  import DeviceListCheck from './DeviceListCheck'


  export default {
      name:"device-list",
      components: { AddDevice,DeviceCertificate,AddDeviceGroup,AddResult,DeviceListCheck},
      props:{
        productArr:{
            type:Array,
            default:() => []       
        },
        productId:{
             type:Number,
             default:0
        },
        type:{
            type:String,
            default:'device'
        },
      
      },
      watch:{
        //监听productId,若发生变化，重新查询设备列表
        productId:{  
            handler:function(val,oldval){ 
                if(val!=oldval){
                    this.$nextTick(()=>{
                        this.currentPage = 1
                        this.getDeviceList()
                    })
                }
            },  
            immediate:true,//关键
            deep:true
          },
      },
      data() {
      return {
        tableData: [],
        multipleSelection: [],
        addDeviceVisible:false,
        deviceCertificateVisible:false,
        certificateInf:{productKey:'',deviceName:'',deviceSecret:''},
        addDeviceGroupVisible:false,
        addResultVisible:false,
        addResultTitle:'正在添加设备...',
        loadBtn:false,
        activities:[], 
        products:[],
        search:'',
        product:'',
        drawer:false,
        currentPage:1,
        pageSize:10,
        total:0,
        searchKey:"DeviceName",
        searchVal:'',
        searchLabel:[],
        label:'',
        addLabelVisible:false

      }
    },
    computed:{
      selectFlg:function(){
          return this.multipleSelection.length > 0
      }
    },
  
    created(){
        this.getDeviceList() 
        if(this.type === 'group'){
            this.getProducts()         
        }        
    },

    methods: {
        //查询设备列表
        getDeviceList(currentPage){
          if(currentPage){
            this.currentPage = currentPage
          }
         // let params = {};
        //  params[this.searchKey] = this.searchVal
          this.$API_IOT.getDeviceList(this.currentPage,this.pageSize,this.productId,this.searchVal).then((res) => {
                this.tableData = res.data.data.data_list
                this.total = res.data.data.num_results
                let params = {}
                params.device_count = res.data.data.num_results
                params.device_active_count =  res.data.data.device_active_count
                params.device_online_count =  res.data.data.device_online_count
                this.$emit('setAcount',params)

                //this.tableData
            })

        },

        //改变每页数量
        handleSizeChange(val){
          this.pageSize = val
          this.getDeviceList(1)
        },

        //改变当前页数
        handleCurrentChange(val){
            this.currentPage = val
            this.getDeviceList()
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
     
        //确认添加设备
        addDeviceSubmit(){  
            this.$refs['addDevice'].submitForm()
        },

        //取消添加设备
        addDeviceCancel(){
          this.addDeviceVisible = false
          this.$refs['addDevice'].init()
        },

        //跳转到设备详情页面
        goToDeviceDetails(deviceId){  
          this.$router.push({name :'device-details',params: {deviceId:deviceId}}) 
        },

        //跳转到设备列表
        goToDeviceList(obj){
          this.openRouter = true 
          this.$router.push({name :'device-details',params: {device:JSON.stringify(obj)}}) 
        },

        //删除设备,type 1-删除设备，2-从分组中删除设备
        deleteClick(type,deviceId){
            let msg = '此操作将永久删除该纪录, 是否继续?'
            if(type === '2'){
                msg = '此操作将永久从分组中删除该纪录, 是否继续?'
            }
            this.$confirm(msg, '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
                if(type === '1'){  //删除设备
                  let arr = new Array();
                  arr.push(deviceId)
                  this.deleteDevices(arr)
                }else if(type === '2'){ //从分组中删除设备

                }
              this.tableData =  this.tableData.filter(item => item.name !== row.name);
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
            this.deleteDevices(arr)
          }else if(type === '2'){//禁用
            this.disabledDevices(arr)
          }else if(type === '3'){//启用
            this.activeDevices(arr)
          }else if(type === '4'){
              this.deleteDevicesFromGroup(arr)
          }
          
        },

        //删除设备
        deleteDevices(arr){
          this.$API_IOT.deleteDevice(arr).then((res) => {
                if(res.data.status  === 'Y'){
                  this.$message.success('删除成功')
                  this.getDeviceList()
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
        //禁用设备
        disabledDevices(arr){
          this.$API_IOT.disabledDevice(arr).then((res) => {
                if(res.data.status  === 'Y'){
                  this.$message.success('禁用成功',function(){
                    this.getDeviceList()
                  })
                  
                }else{
                    this.$message.error(res.message);
                }              
            })
        },
        //启用设备
        activeDevices(arr){
          this.$API_IOT.activeDevices(arr).then((res) => {
                if(res.data.status  === 'Y'){
                  this.$message.sucess('启用成功',function(){
                    this.getDeviceList()
                  })
                }else{
                    this.$message.error(res.message);
                }              
            })
        },
    
        //展示设备证书信息
        showDeviceCertificate(obj){
          this.deviceCertificateVisible = true
          this.certificateInf = obj
          this.getDeviceList()

        },

        //批量添加提交
        addDeviceGroupSubmit(){
              this.$refs['addDeviceGroup'].submit()
        },

        //批量添加取消
        addDeviceGroupCancel(){
           this.addDeviceGroupVisible = false
           this.$refs['addDeviceGroup'].init()
        },      

        //批量添加设备-结果过渡
        showAddResult(){
          this.addResultVisible = true
          this.addResultTitle = '正在添加设备...'
          this.loadBtn = false
          this.activities =  [
          {
            content: '正在添加设备，请等待',
            timestamp: '添加设备可能需要一些时间，请您耐心等待',
            size: 'large',
            type: 'primary',
          }, {
            content: '下载设备证书',
            timestamp: '添加完成后可下载设备证书,您可以在“批次管理”中随时下载本文件。',
            color: '#0bbd87'
          }
          ]
        },

        //批量添加设备成功
        addSuccess(num){
            this.addResultTitle = '添加完成'
            this.activities = [
              {
              content: '添加完成',
              timestamp: `数量:${num}`,
              color: '#0bbd87'
              }, {
              content: '下载设备证书',
              timestamp: '添加完成后可下载设备证书,您可以在“批次管理”中随时下载本文件。',
              color: '#0bbd87'
              }
          ]
          this.loadBtn = true
          this.getDeviceList()
          this.$refs['addDeviceGroup'].init()

        },

        //查询产品信息
        getProducts(){
            this.$API_IOT.getSimpleProductList().then((res) => {
                this.products = res.data.data.data
            })
        },

        //从分组中删除设备
        deleteDevicesFromGroup(){

        },

        //添加设备到分组
        addDeviceToGroup(deviceArr){
            console.log(deviceArr)
            this.drawer = false
            console.log(this.drawer )
        },

        //改变查询条件
        queryParamChange(value){
            this.searchVal = ""
        },

        //添加标签提交
        addLabelSubmit(){
            this.searchLabel = this.$refs.addLabel.label
            this.addLabelVisible = true 
            this.getDeviceList(1)
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