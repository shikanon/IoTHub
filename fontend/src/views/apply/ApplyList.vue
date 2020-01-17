<template>
    <div name="apply-list">
        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="ProductName"
            label="产品名称"
            width="280">
          </el-table-column>
          <el-table-column
            prop="ProductKey"
            label="ProductKey"
            width="180">
          </el-table-column>
        
          <el-table-column
            label="添加时间">
            <template slot-scope="scope">
              <span
                class="name-span"           
              >{{ scope.row.UtcCreate ?  scope.row.UtcCreate:'' }}</span>
            </template>
          </el-table-column>
         <el-table-column
            prop="Count"
            label="添加数量"
            width="180">
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            >
            <template slot-scope="scope" >
              <div class="flex-center">
                <el-button type="text" size="small" @click="showBatchDlt(scope.row.ApplyId)">详情</el-button><el-divider direction="vertical"></el-divider>
                <download-excel
                :data = "json_data"
                :fields = "json_fields"
                name = "Triad.xls">
                <el-button type="text" size="small" @click="download(scope.$index, scope.row)">下载CSV</el-button>
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
            size='40%'
            >
            <ApplyDtlList :applyId="applyId"></ApplyDtlList>
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
          applyId:'',
           json_fields: {
          "DeviceName": "DeviceName",    //常规字段
          "DeviceSecret": "DeviceSecret", //支持嵌套属性
          "ProductKey": "ProductKey"       
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
            // this.$API_IOT.getApplyList(this.currentPage,this.pageSize,this.productId).then((res) => {
            //     this.tableData = res.data.objects
            //     this.total = res.data.num_results
            // })

            this.tableData = [
                {
                    "Status": "CREATE_SUCCESS",
                    "ProductAuthType": "secret",
                    "ProductName": "123123123",
                    "UtcCreate": "2019-12-20T07:08:38.000Z",
                    "ApplyId": 1232165,
                    "Count": 10,
                    "SuccessCount": 10,
                    "ProductKey": "a1qfUCxdfqg"
                },
                {
                    "Status": "CREATE_SUCCESS",
                    "ProductAuthType": "secret",
                    "ProductName": "4444444",
                    "UtcCreate": "2019-12-19T06:00:09.000Z",
                    "ApplyId": 1221612,
                    "Count": 2,
                    "SuccessCount": 2,
                    "ProductKey": "a1ELejzj0h9"
                }
            ]
            this.total = this.tableData.length
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

        
         
           download(index,row){               
            
             this.json_data = [{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"1aTVy5xU40rPijlmzjGxrKNmYO1MdSxG","ProductKey":"a1Ibli2tqC2","DeviceName":"0iV7vxaUCOdHu57Cz61h"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"leiqy5nm5dBd0BzzsIayEv6FTXEmKhbi","ProductKey":"a1Ibli2tqC2","DeviceName":"1Cwbu6U4TJkZis7iiwZ1"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"HVi2e9hLKw9vsRkd1hoLq3rH6qzD05Cc","ProductKey":"a1Ibli2tqC2","DeviceName":"2D87VhVL3Pn0G8v9qE2K"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"cmB5fRPwbvq6Opu37KsKkqjyA9jsFOZr","ProductKey":"a1Ibli2tqC2","DeviceName":"2QDnYGpFwDNEioJYRcIw"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"LZalQwxQ2RDAUrSNDHyZUU17R3W7R5FW","ProductKey":"a1Ibli2tqC2","DeviceName":"4cUq41BU58m2QmoNN0zr"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"UvMnfHtWQ1bvI4jXuj4f3v6T6EqO5Az7","ProductKey":"a1Ibli2tqC2","DeviceName":"6Q0IphH37mIoN30lXNdf"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"FEIIbM5jzfSagEUFRqdLLF8htXYs9HyR","ProductKey":"a1Ibli2tqC2","DeviceName":"6XOocf0TTacaOG7IZ2Ww"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"JoJoHi857P8MWKrLpPom2l7GeniDnBKh","ProductKey":"a1Ibli2tqC2","DeviceName":"6c8mUv27FdrvUAeAXor6"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"5wGWMTJGfEHZSHsyOfTyuiV4SuKPnMut","ProductKey":"a1Ibli2tqC2","DeviceName":"7Qm4NdoezXiYeS5WtmwW"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"Bblk3v7NS6JQXzIzua8hlD5mt6dsOxNK","ProductKey":"a1Ibli2tqC2","DeviceName":"7elX8EQw6egJZxmeDW74"}]

          },
          showBatchDlt(applyId){
            this.drawer = true
            this.applyId =  applyId
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


 