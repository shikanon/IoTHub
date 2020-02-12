<template>
    <div id="category-list"> 
        <el-row>
            <el-select v-model="value" placeholder="全部领域" size="medium" @change="getCategoryList(1)">
              <el-option label="全部领域" value=""></el-option>
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
           @current-change="selectThisRow"
          style="width: 100%">
          <el-table-column
            label="品类名称"
            width="200">
              <template slot-scope="scope">
                  <span >{{scope.row.name}}</span>
                   <i class="el-icon-info" @click="showAttribute(scope.row.id)"></i>
            </template>
          </el-table-column>
           <el-table-column
            prop="scene"
            label="所属场景"
            width="150">
          </el-table-column>
          <el-table-column
            label="操作"
            >
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="selectCategory(scope.row)" :disabled="categoryId === scope.row.id">
                {{categoryId === scope.row.id ? '已选择':'选择'}} 
                </el-button>             
            </template>
          </el-table-column>
        </el-table>
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" 
          @handleSizeChange="handleSizeChange" @handleCurrentChange="handleCurrentChange"
        ></Pagination>     
    </div>
  </template>

  <script>
    import AttributeList from './AttributeList'
    export default {
      props:{
        categoryId:{
          type:Number,
          default:0
        }
      },
      components:{AttributeList},
      data() {
        return {
          tableData:[],
          currentPage:1,
          pageSize:10,
          total:0,
          options: [],
          value: '',
          attributes:[],
          query:'',
        }
      },
   
      watch: {
        $route:function(){
            this.openRouter = !this.openRouter          
        }
    },
      created(){
        this.openRouter = false
        this.getModelTypeList()
        this.getCategoryList() 
                       this.tableData =  [{'id':1,'name':"1",'scene':"123",'territory':'12312312'}]

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

        

         handleSizeChange(val) {
            this.pageSize = val 
            this.getCategoryList(1)

          },
          handleCurrentChange(val) {
            this.currentPage = val 
            this.getCategoryList()
          },

          selectThisRow(value){

          },

            selectCategory(category){     
                this.$emit('select',category)            
            },
            // handleCurrentChange(category){
            //     this.$emit('select',category)            
            // },
            showAttribute(id){        
                this.$emit('showAttributeList',id)
                
            },

         
        
      }
    }
  </script>

  <style scoped>
    .search-input{
      width:200px;
    }
    #category-list{
      margin: 0 20px;
    }
  </style>