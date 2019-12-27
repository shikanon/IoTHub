<template>
    <div name="product-list"> 
      <div v-if="!openRouter">
        <el-row>
            <el-button type="primary" @click="addProduct()" size="medium">创建产品</el-button>
            <el-input
            v-model="search"
            placeholder='请输入产品名称查询'
            class="search-input"
            clearable
            size="medium"
            @keyup.enter.native="getProductList(1)"
            >
            <i slot="suffix" class="el-input__icon el-icon-search" @click="getProductList(1)"></i>
            </el-input> 
        </el-row>
     
        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="ProductName"
            label="产品名称"
            width="300">
          </el-table-column>
           <el-table-column
            prop="ProductKey"
            label="ProductKey"
            width="280">
          </el-table-column>
          <el-table-column
            label="节点类型"
            width="200">
            <template slot-scope="scope">
              <span
                class="name-span"           
              >{{ scope.row.NodeType === 0 ? '设备':'' }}</span>
            </template>
          </el-table-column>
         <el-table-column
            prop="CreateTime"
            label="添加时间"
            >
          </el-table-column>          
          <el-table-column
            label="操作"
            width="400">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="gotoProductDtl(scope.row.ProductName)">查看</el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="gotoDeviceList(scope.row.ProductName)">管理设备</el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="deleteClick(scope.$index, scope.row.ProductName)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" :pageSizes="pageSizes"
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
      </div>
      <router-view/> 
    </div>
  </template>

  <script>
  import {mapState, mapMutations, mapGetters} from 'vuex'

    export default {
      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          pageSizes:[5,10,30,50],
          productname:'',
          search:'',
          openRouter:false
        }
      },
   
      watch: {
        $route: {
            handler: function(val, oldVal){ 
                this.openRouter = !this.openRouter  
            },
            // 深度观察监听
            deep: true
        }
       
    },
      created(){
          this.openRouter = false
          this.productname = this.$route.params.productName
          if(this.productname ){
            this.gotoProductDtl(this.productname)
          }
        
        this.getProductList() 
      },
      methods:{
         getProductList(){
            //  this.$API.getUser(this.currentPage,this.pageSize,this.username).then((res) => {
            //     this.tableData = res.data.objects
            //     this.total = res.data.num_results
            // })
             this.tableData =[
                {
                    "GmtCreate": "2019-12-19 11:35:01",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726501000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1ELejzj0h9"
                },
                {
                    "GmtCreate": "2019-12-19 11:34:50",
                    "ProductName": "123123123",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726490000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1qfUCxdfqg"
                },
                {
                    "GmtCreate": "2019-12-19 11:34:39",
                    "ProductName": "3231323",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726479000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1Ibli2tqC2"
                },
                {
                    "GmtCreate": "2019-12-19 11:34:29",
                    "ProductName": "22222",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726469000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1RIivcyWf0"
                },
                {
                    "GmtCreate": "2019-12-19 11:34:18",
                    "ProductName": "1111",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726458000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1Vr6NBfa0i"
                },
                {
                    "GmtCreate": "2019-12-19 11:34:04",
                    "ProductName": "123213",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726444000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a10TD9i7WFi"
                },
                {
                    "GmtCreate": "2019-12-19 11:33:55",
                    "ProductName": "213213",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726435000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1qJfSBZss0"
                },
                {
                    "GmtCreate": "2019-12-19 11:33:45",
                    "ProductName": "3123",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726425000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1VT5Q8pKbu"
                },
                {
                    "GmtCreate": "2019-12-19 10:29:38",
                    "Description": "test",
                    "ProductName": "test",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576722578000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1zJA7k9sjd"
                }
            ]
            this.total = this.tableData.length
         },

         handleSizeChange(val) {
            this.pageSize = val 
            this.currentPage = 1
            this.getProductList()

          },
          handleCurrentChange(val) {
            this.currentPage = val 
            this.getProductList()
          },

            addProduct(){    
                this.$router.push({name :'product-add'})
            },
        
            gotoProductDtl(productName){
                this.$router.push({name :'product-details',params: {productName:productName}})             
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
          },
          gotoDeviceList(productname){
            this.$router.push({name :'device-list',params: {product:productname}}) 

          }
      }
    }
  </script>

  <style scoped>
    .search-input{
      width:200px;
    }
  </style>