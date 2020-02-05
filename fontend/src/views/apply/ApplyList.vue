<template>
    <div name="apply-list">
        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="name"
            label="产品名称"
            width="280">
          </el-table-column>
          <el-table-column
            prop="product_key"
            label="ProductKey"
            width="380">
          </el-table-column>    
          <el-table-column
            label="添加时间"
            width="180">
            <template slot-scope="scope">
              <span>{{ scope.row.create_time ?  scope.row.create_time:'' }}</span>
            </template>
          </el-table-column>
         <el-table-column
            prop="total"
            label="添加数量"
            width="180">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            >
            <template slot-scope="scope" >
              <div class="flex-center">
                <el-button type="text" size="small" @click="showBatchDlt(scope.row)">详情</el-button><el-divider direction="vertical"></el-divider>
                <download-excel
                :data = "json_data"
                :fields = "json_fields"
                name = "Triad.xls">
                <el-button type="text" size="small" @click="download(scope.row)">下载CSV</el-button>
              </download-excel> 
              </div>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>

        <el-drawer
            title="批次详情"
            :visible.sync="drawer"
            size='50%'
            >
            <ApplyDtlList :productInf="productInf" ></ApplyDtlList>
        </el-drawer>
    </div>
  </template>

  <script>
  import ApplyDtlList from './ApplyDtlList'
  import {mapState, mapMutations, mapGetters} from 'vuex'

    export default {
      props:{
        productId:{
            type:Number,
            default:0
        }
      },
      components: { ApplyDtlList },
      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          drawer:false,
          productInf:{},
          json_fields: {
          "DeviceName": "device_name",    //常规字段
          "DeviceSecret": "device_secret", //支持嵌套属性
          "ProductKey": "product_key"       
          },
          json_data:[]
        }
      },
      watch:{
        //监听productId,若发生变化，重新查询设备批次
        productId:{  
            handler:function(val,oldval){ 
                if(val!=oldval){
                    this.$nextTick(()=>{
                        this.getApplyList()
                    })
                }
            },  
            immediate:true,//关键
            deep:true
          },
      },
     
      created(){
        this.getApplyList() 
      },
      methods:{
     
         getApplyList(){
            this.$API_IOT.getApplyList(this.currentPage,this.pageSize,this.productId).then((res) => {
                this.tableData = res.data.data.data_list
                this.total = res.data.data.num_results
            })

            
         },

         handleSizeChange(val) {
            console.log(`每页 ${val} 条`)
            this.pageSize = val 
            this.currentPage = 1
            this.init()

          },
          handleCurrentChange(val) {
            console.log(`当前页: ${val}`)
            this.currentPage = val 
            this.init()
          },

        
         
           download(row){  
              let createTime = new Date (row.create_time).getTime().toString().substr(0,10)
              this.$API_IOT.getApplyListDtl(row.product_id,createTime,0,0).then((res) => {
                  console.log( res.data.data.data_list)
                  this.json_data = res.data.data.data_list
              })
          },

          showBatchDlt(row){
            this.drawer = true 
            this.productInf = row
          
          }
        

      }
    }
  </script>

  <style scoped>
    .search-input{
      width:200px;
    }
    .flex-center{
      display: flex;
      align-items: center;
    }
  </style>


 