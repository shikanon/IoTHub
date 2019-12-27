<template>
  <div name="device-list">
    <div v-if="!openRouter">
        <el-row v-if="type === 'device'">
            <el-button type="primary" @click="addDeviceVisible = true">添加设备</el-button>
            <el-button @click="addDeviceGroupVisible = true">批量添加</el-button>
        </el-row>  
        <el-row v-else>
            <el-button type="primary" >添加设备到分组</el-button>
            <el-select v-model="product" placeholder="全部产品">
                    <el-option
                    v-for="item in products"
                    :key="item.ProductKey"
                    :label="item.ProductName"
                    :value="item.ProductName">
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
            label="DeviceName/备注名称"
            width="300"  
            >
            <template slot-scope="scope">{{ scope.row.DeviceName }}</template>
        </el-table-column>
        <el-table-column
            prop="ProductName"
            label="设备所属产品"
            width="300">
        </el-table-column>
        <el-table-column
            prop="NodeType"
            label="节点类型"
            width="120"
            show-overflow-tooltip>
        </el-table-column>
        <el-table-column
            label="状态/启用状态"
            show-overflow-tooltip
            :filters="[{text: '已禁用', value: 'DISABLE'}, {text: '未激活', value: 'UNACTIVE'},
            {text: '在线', value: 'ONLINE'}, {text: '离线', value: 'OFFLINE'}]"
            :filter-method="filterHandler">
            <template slot-scope="scope">
                <el-button 
                    :type="scope.row.Status === 'ONLINE' ? 'success': scope.row.Status === 'UNACTIVE' ? 'info': scope.row.Status === 'DISABLE' ? 'danger':'warning'"  
                    circle 
                    size="mini"
                    class="statu-circle">
                </el-button>
                 <span> {{ scope.row.Status | deviceStatusFilter}}</span>

                    <!-- <span :class="['statuCircle',scope.row.Status === 'ONLINE' ? 'success': scope.row.Status === 'UNACTIVE' ? 'info': scope.row.Status === 'DISABLE' ? 'danger':'warning']">{{ scope.row.Status | deviceStatusFilter}}</span> -->
                <!-- <el-switch 
                ></el-switch>  -->
            </template>
        </el-table-column>
        <el-table-column
            label="最后上线时间"  >
            <template slot-scope="scope">{{ scope.row.LastOnlineTime  }}</template>
        </el-table-column>
        <el-table-column 
                fixed="right"
                label="操作"
                width="200">
                <template slot-scope="scope">
                <el-button @click="handleClick(scope.row)" type="text" size="small">查看 </el-button>
                <el-divider direction="vertical"></el-divider>                
                <el-button  v-if="type === 'device'" type="text" size="small" @click="deleteClick(scope.$index, scope.row.DeviceName)">删除1</el-button>
                <el-button  v-else type="text" size="small" @click="deleteClick(scope.$index, scope.row.DeviceName)">从分组中删除</el-button>

                <el-divider v-if="type === 'device'"  direction="vertical"></el-divider>     
                <el-popconfirm v-if="type === 'device'" 
                    confirmButtonText='确定'
                    cancelButtonText='取消'
                    icon="el-icon-info"
                    iconColor="red"
                    title="确定删除吗"
                >
                    <el-button slot="reference" type="text" size="small" >删除2</el-button>
                </el-popconfirm>
                </template>
            </el-table-column>
        </el-table>
        <div v-if="type === 'device'"  style="margin-top: 20px;margin-left: 15px;" >
            <el-checkbox v-model="selectFlg" ></el-checkbox>
            <el-button @click="handleSelection('1')" :disabled="!selectFlg" type="primary">删除</el-button>
            <el-button @click="handledSelection('2')" :disabled="!selectFlg" type="primary">禁用</el-button>
            <el-button @click="handleSelection('3')" :disabled="!selectFlg" type="primary">启用</el-button>
        </div>
         <div v-else  style="margin-top: 20px;margin-left: 15px;" >
            <el-checkbox v-model="selectFlg" ></el-checkbox>
            <el-button @click="handleSelection('1')" :disabled="!selectFlg" type="primary">从分组中删除</el-button>
         
        </div>
        <el-dialog title="添加设备" :visible.sync="addDeviceVisible" width="25%">
            <AddDevice ref="addDevice" :productArr="productArr" @close="addDeviceVisible = false"
            @showDeviceCertificate="showDeviceCertificate"></AddDevice>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addDeviceSubmit">确 定</el-button>
                <el-button @click="addDeviceVisible = false">取 消</el-button>
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
            <AddDeviceGroup ref="addDeviceGroup"  :productArr="productArr" 
            @close="addDeviceGroupVisible = false" 
            @showAddResult="showAddResult"
            @addSuccess="addSuccess"></AddDeviceGroup>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addDeviceGroupSubmit">确 定</el-button>
                <el-button @click="addDeviceGroupVisible = false">取 消</el-button>
            </span>
        </el-dialog>
        <el-dialog :title="addResultTitle" :visible.sync="addResultVisible" width="25%">
            <AddResult :activities="activities"></AddResult>
            <el-button type="primary" v-if="loadBtn" class="loadBtn">下载设备证书</el-button>
            <span slot="footer" class="dialog-footer">
                <el-button @click="addResultVisible = false">关 闭</el-button>
            </span>
        </el-dialog>
    </div>
  </div>
</template>

<script>
  import AddDevice from './AddDevice'
  import DeviceCertificate from './DeviceCertificate'
  import AddDeviceGroup from './AddDeviceGroup'
  import AddResult from './AddResult'

  export default {
      props:{
        productArr:{
            type:Array,
            default:() => []       
        },
        type:{
            type:String,
            default:'device'
        }
      },
    components: { AddDevice,DeviceCertificate,AddDeviceGroup,AddResult},

    data() {
      return {
        tableData: [],
        multipleSelection: [],
        addDeviceVisible:false,
        openRouter:false,
        deviceCertificateVisible:false,
        certificateInf:{productKey:'',deviceName:'',deviceSecret:''},
        addDeviceGroupVisible:false,
        addResultVisible:false,
        addResultTitle:'正在添加设备...',
        loadBtn:false,
        activities:[],
        products:[]
      }
    },
    computed:{
      selectFlg:function(){
          return this.multipleSelection.length > 0
      }
    },
    watch: {
        $route:function(){
            console.log(123)
            this.openRouter = !this.openRouter          
        }
    },
    created(){
           this.getDiviceList() 
           if(type === 'group'){
               this.getProducts()
           }
           
    },

    methods: {
        getDiviceList(){
            this.tableData =[
                {
                    "UtcActiveTime": "",
                    "LastOnlineTime": "",
                    "Status": "DISABLE",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 14:00:09",
                    "IotId": "DFTwwlvDhW8DmKJP2EuK000100",
                    "UtcCreateTime": "2019-12-19T06:00:09.000Z",
                    "CategoryKey": "GarbageOverflowingDetection",
                    "ActiveTime": "",
                    "UtcLastOnlineTime": "",
                    "NodeType": 0,
                    "ProductKey": "a1ELejzj0h9",
                    "DeviceName": "DFTwwlvDhW8DmKJP2EuK"
                },
                {
                    "UtcActiveTime": "",
                    "LastOnlineTime": "",
                    "Status": "UNACTIVE",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 14:00:09",
                    "IotId": "l8s2H6EOs46qlMW7jQKd000100",
                    "UtcCreateTime": "2019-12-19T06:00:09.000Z",
                    "CategoryKey": "GarbageOverflowingDetection",
                    "ActiveTime": "",
                    "UtcLastOnlineTime": "",
                    "NodeType": 0,
                    "ProductKey": "a1ELejzj0h9",
                    "DeviceName": "l8s2H6EOs46qlMW7jQKd"
                },
                {
                    "UtcActiveTime": "",
                    "LastOnlineTime": "",
                    "Status": "ONLINE",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 13:59:24",
                    "IotId": "kw57odVwBBxmesgfpBGF000100",
                    "UtcCreateTime": "2019-12-19T05:59:24.000Z",
                    "CategoryKey": "GarbageOverflowingDetection",
                    "ActiveTime": "",
                    "UtcLastOnlineTime": "",
                    "NodeType": 0,
                    "ProductKey": "a1ELejzj0h9",
                    "DeviceName": "q123",
                    "Nickname": "123123"
                },
                {
                    "UtcActiveTime": "2019-12-19T02:30:48.731Z",
                    "LastOnlineTime": "2019-12-19 10:30:48",
                    "Status": "OFFLINE",
                    "ProductName": "test",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 10:30:10",
                    "IotId": "gQj9TNrlkqvU1UgtcLI9000100",
                    "UtcCreateTime": "2019-12-19T02:30:10.000Z",
                    "CategoryKey": "Lighting",
                    "ActiveTime": "2019-12-19 10:30:48",
                    "UtcLastOnlineTime": "2019-12-19T02:30:48.731Z",
                    "NodeType": 0,
                    "ProductKey": "a1zJA7k9sjd",
                    "DeviceName": "test_device",
                    "Nickname": "test_device"
                }
            ]
        },
      toggleSelection(rows) {
        if (rows) {
          rows.forEach(row => {
            this.$refs.multipleTable.toggleRowSelection(row);
          });
        } else {
          this.$refs.multipleTable.clearSelection();
        }
      },
      handleSelectionChange(val) {
        this.multipleSelection = val;
      },
      filterHandler(value, row, column) {
        const property = column['property'];
        return row[property] === value;
      },
      deleteSelection(type){
          console.log(this.multipleSelection)
      },
      
    
      addDeviceSubmit(){  
          this.$refs['addDevice'].submitForm()
      },
      handleClick(obj){
       // let device =  {DeviceName:obj.DeviceName,DeviceSecret:obj.CategoryKey,ProductName:obj.ProductName,ProductKey:obj.ProductKey}
        this.$router.push({name :'device-details',params: {device:JSON.stringify(obj)}}) 
      },
   
      showDeviceCertificate(obj){
        this.deviceCertificateVisible = true
        this.certificateInf =obj
      },
     
      addDeviceGroupSubmit(){
            this.$refs['addDeviceGroup'].submitForm()
      },
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
      },
       getProducts(){
            this.products = [{
                    "Type": "owned",
                    "ProductName": "4444444",
                    "ProductKey": "a1ELejzj0h9",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "123123123",
                    "ProductKey": "a1qfUCxdfqg",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "3231323",
                    "ProductKey": "a1Ibli2tqC2",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "22222",
                    "ProductKey": "a1RIivcyWf0",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "1111",
                    "ProductKey": "a1Vr6NBfa0i",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "123213",
                    "ProductKey": "a10TD9i7WFi",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "213213",
                    "ProductKey": "a1qJfSBZss0",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "3123",
                    "ProductKey": "a1VT5Q8pKbu",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "test",
                    "ProductKey": "a1zJA7k9sjd",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                }
            ]
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
 .search-input{width: 300px;}
</style>