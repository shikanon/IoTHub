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
            <el-input
              placeholder="请选择设备标签"
              suffix-icon="el-icon-arrow-down"
              @focus="addLabelVisible = true"
              class="search-input"
              size="medium"
              value="">
            </el-input>
        </el-row>
        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="name"
            label="产品名称"
            width="300">
          </el-table-column>
           <el-table-column
            prop="product_key"
            label="ProductKey"
            width="280">
          </el-table-column>
          <el-table-column
            label="节点类型"
            width="200">
            <template slot-scope="scope">
              <span
                class="name-span"           
              >{{ scope.row.node_type  }}</span>
            </template>
          </el-table-column>
         <el-table-column
            prop="create_time"
            label="添加时间"
            >
          </el-table-column>          
          <el-table-column
            label="操作"
            width="400">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="gotoProductDtl(scope.row.id)">查看</el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="gotoDeviceList(scope.row.id)">管理设备</el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="deleteProduct(scope.row.id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
      </div>
      <router-view/> 
        <el-dialog title="标签筛选" :visible.sync="addLabelVisible" width="26%">
            <AddLabel ref="addLabel" type="product" @close="addLabelVisible = false"></AddLabel>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="editLabelSubmit">确 定</el-button>
                <el-button @click="addLabelVisible = false">取 消</el-button>
            </span>
        </el-dialog>
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
          productId:0,
          search:'',
          searchLabel:[],
          openRouter:false,
          addLabelVisible:false,
          
        }
      },
   
      watch: {
        $route: {
            handler: function(val, oldVal){ 
                if(val !== oldVal){
                    this.openRouter = !this.openRouter  
                    this.getProductList()
                }
                
                this.productId = this.$route.params.productId
                let activeName = this.$route.params.activeName

                //传了productId，需要直接跳转到详情页
                if(this.productId ){
                  this.gotoProductDtl(this.productId,activeName)
                }

            },
            // 深度观察监听
            deep: true
        }  
      },
      created(){
        this.openRouter = false   
          
        this.getProductList(this.currentPage)  
            
       
      },
      methods:{
        //获取产品列表
         getProductList(currentPage){
           if(currentPage){
              this.currentPage = currentPage
            }

             this.$API_IOT.getProductList(this.currentPage,this.pageSize,this.search).then((res) => {
                this.tableData = res.data.data.data_list
                this.total = res.data.data.num_results
            })
           
         },

          //改变每页数量
         handleSizeChange(val){
            this.pageSize = val 
            this.getProductList(1)
         },
          //改变当前页数
          handleCurrentChange(val) {
            this.currentPage = val 
            this.getProductList()
          },

          //点击添加产品按钮  
          addProduct(){    
              this.$router.push({name :'product-add'})
          },
        
          //跳转到产品详情页面
          gotoProductDtl(productId,activeName = 'product'){
              this.$router.push({name :'product-details',params: {"productId":productId,"activeName":activeName}})             
          },

          //删除产品
          deleteProduct(id){               
            this.$confirm('此操作将永久删除该纪录, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              //this.tableData =  this.tableData.filter(item => item.name !== row.name);

                  this.$API_IOT.deleteProduct(id).then((res) => {
                    if(res.data.status === 'Y'){
                      this.$message.success('删除成功')    
                      this.getProductList()              
                    }else{
                      this.$message.error(res.message);
                    } 
                  })
            }).catch(() => {
              this.$message({
                type: 'info',
                message: '已取消删除'
              });          
            });                       
          },

          //跳转到设备列表
          gotoDeviceList(productId){
            this.$router.push({name :'device-list',params: {productId:productId}}) 

          },

          //编辑产品label
          editLabelSubmit(){
            this.searchLabel = this.$refs.addLabel.label
            this.addLabelVisible = false 
            this.getProductList(1)

          }
      }
    }
  </script>

  <style scoped>
  .search-input{width: 200px;}
  </style>