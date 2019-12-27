<template>
    <div name="group-list"> 
        <div v-if="!openRouter">
          <el-row>
            <el-button type="primary" @click="addGroupVisible = true" size="medium">新建分组</el-button>
            <el-input
              v-model="search"
              placeholder='请输入分组名称查询'
              class="search-input"
              clearable
              size="medium"
              @keyup.enter.native="getGroupList(1)"
            >
              <i slot="suffix" class="el-input__icon el-icon-search" @click="getGroupList(1)"></i>
            </el-input> 
             <el-select  v-model="value" placeholder="请选择标签分组" @visible-change="addLabelVisible = true " > </el-select>
          </el-row>    
          <el-table 
            :data="tableData"
            highlight-current-row
            style="width: 100%">
            <el-table-column
              prop="GroupName"
              label="分组名称"
              width="300">
            </el-table-column>
            
            <el-table-column         
            prop="GroupId"
              label="分组ID"
              width="300">
          
            </el-table-column>
          <el-table-column
              prop="UtcCreate"
              label="添加时间"
              >
            </el-table-column>
            
            <el-table-column
              label="操作"
              width="100">
              <template slot-scope="scope">
                <el-button type="text" size="small" @click="gotoGroupDtl(scope.row.GroupName)">查看</el-button>
                <el-divider direction="vertical"></el-divider>
                <el-button type="text" size="small" @click="deleteClick(scope.$index, scope.row.GroupId)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
          <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" :pageSizes="pageSizes"
            @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
          ></Pagination>
          <el-dialog title="新建分组" :visible.sync="addGroupVisible" width="25%">
              <AddGroup ref="addGroup"  @close="addGroupVisible = false"></AddGroup>
              <span slot="footer" class="dialog-footer">
                  <el-button type="primary" @click="addGroupSubmit">确 定</el-button>
                  <el-button @click="addGroupVisible = false">取 消</el-button>
              </span>
          </el-dialog>
            <el-dialog title="添加标签" :visible.sync="addLabelVisible" width="26%">
            <AddLabel ref="addLabel" :labelArr="labelArr"  @close="addLabelVisible = false"></AddLabel>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addLabelSubmit">确 定</el-button>
                <el-button @click="addLabelVisible = false">取 消</el-button>
            </span>
        </el-dialog>
        </div>   
        <router-view/> 
    </div>
  </template>

  <script>
  import AddGroup from './AddGroup'

    export default {
      components:{AddGroup},
      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          pageSizes:[5,10,30,50],
          search:'',
          addGroupVisible:false,
          GroupName:'',
          addLabelVisible:false,
          labelArr:[],
          value:''
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
          this.GroupName = this.$route.params.GroupName
          if(this.GroupName ){
            this.gotoGroupDtl(this.GroupId)
          }
        
        this.getGroupList() 
      },
      methods:{
         getGroupList(){
            //  this.$API.getUser(this.currentPage,this.pageSize,this.username).then((res) => {
            //     this.tableData = res.data.objects
            //     this.total = res.data.num_results
            // })
             this.tableData =[
               {
                "GroupName": "111111111111",
                "UtcCreate": "2019-12-27T02:29:45.000Z",
                "GroupDesc": "1111111111111111111",
                "GroupId": "BOfqydh6Mpb8Niu8wv5c010200"
                },
                {
                "GroupName": "3423424",
                "UtcCreate": "2019-12-27T01:31:20.000Z",
                "GroupDesc": "24234",
                "GroupId": "Xx4Ebu2hNwlfsuCfZlP2010200"
                }
             ]
            this.total = this.tableData.length
         },

         handleSizeChange(val) {
            this.pageSize = val 
            this.currentPage = 1
            this.getGroupList()

          },
          handleCurrentChange(val) {
            this.currentPage = val 
            this.getGroupList()
          },

            addGroupSubmit(){    
                this.$refs.addGroup.submitForm();
            },
        
            gotoGroupDtl(groupName){
                this.$router.push({name :'group-details',params: {groupName:groupName}})             
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
         
          addLabelSubmit(){

          }
        

      }
    }
  </script>

  <style scoped>
    .search-input{
      width:200px;
    }
  </style>