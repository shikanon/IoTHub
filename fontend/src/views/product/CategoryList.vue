<template>
    <div name="category-list"> 
        <el-row>
            <el-select v-model="value" placeholder="全部领域">
                <el-option
                v-for="item in options"
                :key="item.id"
                :label="item.name"
                :value="item.id">
                </el-option>
            </el-select>
            <el-input
                v-model="query"
                placeholder='请输入品类名称或者所属场景'
                class="search-input"
                clearable
                size="medium"
                @keyup.enter.native="getCategoryList(1)"
            >
                <i slot="suffix" class="el-input__icon el-icon-search" @click="getCategoryList(1)"></i>
            </el-input> 
        </el-row>

        <el-table 
          :data="tableData"
          highlight-current-row
           @current-change="handleCurrentChange"
          style="width: 100%">
          <el-table-column
            label="品类名称"
            width="200">
              <template slot-scope="scope">
                  <span >{{scope.row.name}}</span>
                   <i class="el-icon-info" @click="showAttribute(scope.row.attribute)"></i>
            </template>
          </el-table-column>
           <el-table-column
            prop="scene"
            label="所属场景"
            width="150">
          </el-table-column>
          <el-table-column
            label="操作"
            width="100">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="handleClick(scope.row)">选择</el-button>
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>
        <el-drawer
          title="标准功能定义"
          :visible.sync="drawer"
          size='30%'>
            <AttributeList ></AttributeList>
        </el-drawer>
      </div>
  </template>

  <script>
    import AttributeList from './AttributeList'
    export default {
        components:{AttributeList},
      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          options: [],
          value: '',
          attribute:[],
          query:'',
          drawer:false
        }
      },
   
      watch: {
        $route:function(){
            this.openRouter = !this.openRouter          
        }
    },
      created(){
        this.openRouter = false
        this.get
        this.getCategoryList() 
      },
      methods:{
       
         getCategoryList(num){
           if(num){
             this.currentPage = num 
           }
             this.$API_IOT.getCategoryList(this.currentPage,this.pageSize,this.value).then((res) => {
                this.tableData = res.data.data
                this.total = res.data.num_results
            })           
         },
         
         getModelTypeList(){
            this.options = [{id:1,name:"智能城市"},
            {id:2,name:"智能生活"},
            {id:3,name:"	智能工业"},
            {id:4,name:"	边缘计算"},
            {id:5,name:"	商业共享"},
            {id:6,name:"	智能模板"},
            {id:7,name:"	智能电力"},
           {id: 8,name:"	智能农业"},
            {id:9,name:"	智能建筑"},
           {id: 10,name:"	智能园区"}]

            //  this.$API_IOT.getCategoryList(this.currentPage,this.pageSize,this.username).then((res) => {
            //     this.tableData = res.data.data
            //     this.total = res.data.num_results
            // })           
         },

         getAttributeList(){
            this.attribute =[
                {type:'属性',
                 name:'工作状态',
                 symbal:'LightStatus',
                 dataType:'BOOL',
                },
                 {type:'属性',
                 name:'工作电压',
                 symbal:'LightVolt',
                 dataType:'FLOAT',
                },
                  {type:'属性',
                 name:'地理位置',
                 symbal:'GeoLocation',
                 dataType:'STRUCT',
                },           
            ]
         },

         handleSizeChange(val) {
            this.pageSize = val 
            this.getCategoryList(1)

          },
          handleCurrentChange(val) {
            this.currentPage = val 
            this.getCategoryList()
          },

            handleClick(category){     
                this.$emit('select',category)            
            },
            handleCurrentChange(category){
                this.$emit('select',category)            
            },
            showAttribute(){        
                this.drawer = true 
            }
        
      }
    }
  </script>

  <style scoped>
    .search-input{
      width:200px;
    }
  </style>