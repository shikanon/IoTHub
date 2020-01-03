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
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="showBatchDlt(scope.row.ApplyId)">详情</el-button><el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="download(scope.$index, scope.row)">下载CSV</el-button>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" :pageSizes="pageSizes"
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
      components: { ApplyDtlList },
      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          pageSizes:[5,10,30,50],
          drawer:false,
          applyId:''
        }
      },
      computed:{
        ...mapState({  
          userEditFlg:'isUserEdit'                                                
        }),
          
      },
      created(){
        this.init() 
      },
      methods:{
        ...mapMutations([
            'changeUserEdit'
        ]),
         init (){
            // this.$API.getUser(this.currentPage,this.pageSize,this.username).then((res) => {
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
  </style>


 