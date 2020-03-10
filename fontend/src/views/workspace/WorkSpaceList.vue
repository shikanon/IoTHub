<template>
    <div name="product-list"> 
      <div v-if="!openRouter">
        <el-row>
            <el-button type="primary" @click="addWorkSpaceVisible = true" size="medium">创建工作空间</el-button>
            <el-input
            v-model="search"
            placeholder='请输入空间名称查询'
            class="search-input"
            clearable
            size="medium"
            @keyup.enter.native="getWorkSpaceList(1)"
            >
              <i slot="suffix" class="el-input__icon el-icon-search" @click="getWorkSpaceList(1)"></i>
            </el-input> 
         
        </el-row>

        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column label="空间ID/名称">
              <template slot-scope="scope">
              <span
                class="name-span"           
              > {{scope.row.work_space_id}} / {{scope.row.name}}</span>
            </template>
          </el-table-column>
           <el-table-column
            label="状态"
            width="100"
            >
             <template slot-scope="scope">
              <span
                class="name-span"           
              >{{ scope.row.status_id === 0 ? '启用':'停用'}}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="接入类型"
             width="150"
            >
            <template slot-scope="scope">
              <span
                class="name-span"           
              >{{ scope.row.access_type_id |  accessTypeFilter }}</span>
            </template>
          </el-table-column>
         <el-table-column
            prop="access_address"
            label="接入地址"
            width="250"
          >
         </el-table-column>
          <el-table-column
            prop="device_count"
            label="接入设备数"
            width="150"
            >
            
          </el-table-column>          
          <el-table-column
            label="操作"
            >
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="editWorkSpace(scope.row)">查看</el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="gotoCameraList(scope.row.work_space_id)">摄像头管理</el-button>
              <el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="stopWorkSpace(scope.row.work_space_id)" v-if="scope.row.status_id === 0">停用</el-button>
              <el-divider direction="vertical" v-if="scope.row.status_id === 0"></el-divider>
              <el-button type="text" size="small" @click="deleteWorkSpace(scope.row.work_space_id)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
      </div>
      <router-view/> 
      <el-dialog title="添加工作空间" :visible.sync="addWorkSpaceVisible" width="25%">
            <AddWorkSpace ref="addWorkSpace" :workSpace="workSpace" @close="addWorkSpaceVisible = false"  @refresh="getWorkSpaceList()"></AddWorkSpace>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addWorkSpaceSubmit">确 定</el-button>
                <el-button @click="addWorkSpaceCancel">取 消</el-button>
            </span>
        </el-dialog>
    </div>
  </template>

  <script>
  import {mapState, mapMutations, mapGetters} from 'vuex'
  import AddWorkSpace from './AddWorkspace'
    export default {
      components:{AddWorkSpace},
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
          addWorkSpaceVisible:false,
          workSpaceId:null,
          workSpace:{work_space_id:0,name:'',access_type_id:1}
        }
      },
   
      watch: {
        $route: {
            handler: function(val, oldVal){ 
                if(val !== oldVal){
                    this.openRouter = !this.openRouter  
                    this.getWorkSpaceList()
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
          
        this.getWorkSpaceList(this.currentPage)  
            
       
      },
      methods:{
        //获取工作空间列表
         getWorkSpaceList(currentPage){
           if(currentPage){
              this.currentPage = currentPage
            }

             this.$API_CLOUD.getWorkSpaceList(this.currentPage,this.pageSize,this.search).then((res) => {
                this.tableData = res.data.data.data_list
                this.total = res.data.data.num_results
            })
           
         },

          //改变每页数量
         handleSizeChange(val){
            this.pageSize = val 
            this.getWorkSpaceList(1)
         },
          //改变当前页数
          handleCurrentChange(val) {
            this.currentPage = val 
            this.getWorkSpaceList()
          },

          //点击添加按钮  
          addWorkSpace(){   
              this.addWorkSpaceVisible = true
          },

          addWorkSpaceSubmit(){
            this.$refs.addWorkSpace.submitForm()
          },

          addWorkSpaceCancel(){
              this.addWorkSpaceVisible = false
              this.$refs.addWorkSpace.init()
          },
        
          //编辑工作空间
          editWorkSpace(workSpace){
            this.addWorkSpaceVisible = true
            this.workSpace = workSpace
          },

          //删除工作空间
          deleteWorkSpace(id){               
            this.$confirm('此操作将永久删除该纪录, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              //this.tableData =  this.tableData.filter(item => item.name !== row.name);

                  this.$API_CLOUD.deleteWorkSpace(id).then((res) => {
                    if(res.data.status === 'Y'){
                      this.$message.success('删除成功')    
                      this.getWorkSpaceList()              
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

          stopWorkSpace(id){               
            this.$confirm('此操作将停用该工作空间?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {

                  this.$API_CLOUD.stopWorkSpace({id:id}).then((res) => {
                    if(res.data.status === 'Y'){
                      this.$message.success('停用成功')    
                      this.getWorkSpaceList()              
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

          //跳转到摄像头列表
          gotoCameraList(workSpaceId){
            this.$router.push({name :'camera-list',params: {workSpaceId:workSpaceId}}) 

          },

          //编辑产品label
          editLabelSubmit(){
            this.searchLabel = this.$refs.addLabel.label
            this.addLabelVisible = false 
            this.getWorkSpaceList(1)

          }
      }
    }
  </script>

  <style scoped>
  .search-input{width: 200px;}
  </style>