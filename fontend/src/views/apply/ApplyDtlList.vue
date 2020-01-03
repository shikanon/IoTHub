<template>
    <div name="apply-dtl-list">     
        <div class="apply-inf"> 
            <div >
                <div> <span>产品名称：{{123}}</span></div>
                <div class="apply-add">
                     <div><span> 添加时间：{{123}} </span></div>  
                     <div><span>添加数量：{{123}} </span></div>
                     <div></div>
                </div>  
            </div>
            <el-button type="primary">下载CSV</el-button>
        </div>
        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="ProductKey"
            label="ProductKey"
            width="100">
          </el-table-column>
          <el-table-column
            label="DeviceName"
            width="200">
              <template slot-scope="scope">
                <el-tooltip class="item" effect="dark" content="scope.row.DeviceName" placement="top">
                 <span>{{ scope.row.DeviceName.substr(0,15) }}...</span>
             </el-tooltip>
            </template>
          </el-table-column>
         <el-table-column
            label="DeviceSecret"
            width="200">
            <template slot-scope="scope">
                <el-tooltip class="item" effect="dark" content="scope.row.DeviceSecret" placement="top">
                 <span>{{scope.row.DeviceSecret.substr(0,15)}}...</span>
             </el-tooltip>
            </template>
          </el-table-column>

          <el-table-column
            label="激活状态">
            <template slot-scope="scope">
              <span          
              >{{ scope.row.Status  }}</span>
            </template>
          </el-table-column>
            <el-table-column
            label="激活时间">
            <template slot-scope="scope">
              <span          
              >{{ scope.row.role  }}</span>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" :pageSizes="pageSizes"
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
  
    </div>
  </template>

  <script>
 
    export default {
      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          pageSizes:[5,10,30,50],
        }
      },
      computed:{
        
          
      },
      created(){
        this.init() 
      },
      methods:{
        
         init (){
            // this.$API.getApplyList(this.currentPage,this.pageSize,this.username).then((res) => {
            //     this.tableData = res.data.objects
            //     this.total = res.data.num_results
            // })

             this.tableData  = [{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"1aTVy5xU40rPijlmzjGxrKNmYO1MdSxG","ProductKey":"a1Ibli2tqC2","DeviceName":"0iV7vxaUCOdHu57Cz61h"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"leiqy5nm5dBd0BzzsIayEv6FTXEmKhbi","ProductKey":"a1Ibli2tqC2","DeviceName":"1Cwbu6U4TJkZis7iiwZ1"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"HVi2e9hLKw9vsRkd1hoLq3rH6qzD05Cc","ProductKey":"a1Ibli2tqC2","DeviceName":"2D87VhVL3Pn0G8v9qE2K"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"cmB5fRPwbvq6Opu37KsKkqjyA9jsFOZr","ProductKey":"a1Ibli2tqC2","DeviceName":"2QDnYGpFwDNEioJYRcIw"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"LZalQwxQ2RDAUrSNDHyZUU17R3W7R5FW","ProductKey":"a1Ibli2tqC2","DeviceName":"4cUq41BU58m2QmoNN0zr"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"UvMnfHtWQ1bvI4jXuj4f3v6T6EqO5Az7","ProductKey":"a1Ibli2tqC2","DeviceName":"6Q0IphH37mIoN30lXNdf"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"FEIIbM5jzfSagEUFRqdLLF8htXYs9HyR","ProductKey":"a1Ibli2tqC2","DeviceName":"6XOocf0TTacaOG7IZ2Ww"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"JoJoHi857P8MWKrLpPom2l7GeniDnBKh","ProductKey":"a1Ibli2tqC2","DeviceName":"6c8mUv27FdrvUAeAXor6"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"5wGWMTJGfEHZSHsyOfTyuiV4SuKPnMut","ProductKey":"a1Ibli2tqC2","DeviceName":"7Qm4NdoezXiYeS5WtmwW"},{"Status":"UNACTIVE","UtcActive":"","DeviceSecret":"Bblk3v7NS6JQXzIzua8hlD5mt6dsOxNK","ProductKey":"a1Ibli2tqC2","DeviceName":"7elX8EQw6egJZxmeDW74"}]
             this.total = 10
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

        
          editOrAddClick(username){
             this.isRouter = false
             if(username){
                this.$router.push({name :'user-update',params: {username:username}}) 
             }else{
               this.$router.push({ path:`/home/system/user/add`})                     
             }
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