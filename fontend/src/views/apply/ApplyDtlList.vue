<template>
    <div name="apply-dtl-list">     
        <div class="apply-inf"> 
            <div >
                <div> <span>产品名称：{{productInf.name}}</span></div>
                <div class="apply-add">
                     <div><span> 添加时间：{{productInf.create_time}} </span></div>  
                     <div><span>添加数量：{{productInf.total}} </span></div>
                     <div></div>
                </div>  
            </div>
            <download-excel
              :data = "json_data"
              :fields = "json_fields"
              name = "Triad.xls">
              <el-button type="primary" @click="download">下载CSV</el-button>
            </download-excel>     
        </div>  
        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="product_key"
            label="ProductKey"
            width="300">
          </el-table-column>
          <el-table-column
            label="DeviceName"
            width="200">
              <template slot-scope="scope">
                <el-tooltip class="item" effect="dark" :content="scope.row.device_name" placement="top">
                 <span>{{ scope.row.device_name.substr(0,15) }}...</span>
             </el-tooltip>
            </template>
          </el-table-column>
         <el-table-column
            label="DeviceSecret"
            width="200">
            <template slot-scope="scope">
                <el-tooltip class="item" effect="dark" :content="scope.row.device_secret" placement="top">
                 <span>{{scope.row.device_secret.substr(0,15)}}...</span>
             </el-tooltip>
            </template>
          </el-table-column>

          <el-table-column
            label="激活状态" width="100">
            <template slot-scope="scope">
              <span          
              >{{ scope.row.status_id  }}</span>
            </template>
          </el-table-column>
            <el-table-column
            label="激活时间">
            <template slot-scope="scope">
              <span          
              >{{ scope.row.activation_time  }}</span>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
  
    </div>
  </template>

  <script>
 
    export default {
      props:{
         productInf:{
            type:Object,
            default:()=> ({})
          },
      },
      watch:{
        productInf:{
          handler(val,oldVal){
            if(val.create_time !== oldVal.create_time){
              this.currentPage = 1
              this.init()
            }
          }
        }
      },

      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          json_fields: {
          "DeviceName": "device_name",    //常规字段
          "DeviceSecret": "device_secret", //支持嵌套属性
          "ProductKey": "product_key"       
          },
          json_data:[]
        }
        
      },
      
      created(){
        this.init() 
      },
      methods:{
        
         init (){

           let productId = this.productInf.product_id
           let createTime = new Date (this.productInf.create_time).getTime().toString().substr(0,10)
           
             this.$API_IOT.getApplyListDtl(productId,createTime,this.currentPage,this.pageSize).then((res) => {
                this.tableData = res.data.data.data_list
                this.total = res.data.data.num_results
            })
         },

         handleSizeChange(val) {
            this.pageSize = val 
            this.init()

          },
          handleCurrentChange(val) {
            this.currentPage = val 
            this.init()
          },

          download(){
              this.$API_IOT.getApplyListDtl(productId,createTime,0,0).then((res) => {
                this.json_data = res.data.data.data_list
            })
          },


           deleteClick(index,row){               
              this.$confirm('此操作将永久删除该纪录, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
              }).then(() => {
                this.tableData =  this.tableData.filter(item => item.name !== row.name);
              }).catch(() => {
                this.$message({
                  type: 'info',
                  message: '已取消删除'
                });          
              });                       
          }
      }
    }
  </script>

  <style scoped>
    .apply-inf{
        display: flex;
        background-color: #fbfbfc;
        padding: 16px;
        margin-bottom: 16px;
        box-sizing: border-box;
        border: 1px solid #ecedee;
        font-size: 14px;
        align-items: center;
        justify-content: space-between;
    }
    .apply-add{
        display: flex;
    }
  </style>