<template>
    <div name="user-list">     
      <div v-if="!isRouter">
        <div >
            <el-row>
              <el-button type="primary" @click="editOrAddClick()">添加用户</el-button>
              <el-input
                v-model="username"
                placeholder='输入用户名称查询'
                class="search-input"
                clearable
                @keyup.enter.native="init(1)"
              >
                <i slot="suffix" class="el-input__icon el-icon-search" @click="init(1)"></i>
              </el-input> 
            </el-row>
        </div>
        <el-table 
          :data="tableData"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="user_name"
            label="账号"
            width="280">
          </el-table-column>
          <el-table-column
            prop="email"
            label="邮箱"
            width="480">
          </el-table-column>
        
          <el-table-column
            label="角色">
            <template slot-scope="scope">
              <span
                class="name-span"           
              >{{ scope.row.role ?  scope.row.role.role_name:'' }}</span>
            </template>
          </el-table-column>
        
          <el-table-column
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="editOrAddClick(scope.row.user_name)">编辑</el-button><el-divider direction="vertical"></el-divider>
              <el-button type="text" size="small" @click="deleteClick(scope.$index, scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" :pageSizes="pageSizes"
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
      </div>
       <router-view /> 
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
          username:'',
          isRouter:false
        }
      },
      computed:{
        
          
      },
      created(){
        this.init() 
      },
      methods:{
        
         init (){
            this.$API.getUser(this.currentPage,this.pageSize,this.username).then((res) => {
                this.tableData = res.data.objects
                this.total = res.data.num_results
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
    .search-input{
      width:200px;
    }
  </style>