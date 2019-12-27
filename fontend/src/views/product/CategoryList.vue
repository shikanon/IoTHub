<template>
    <div name="category-list"> 
        <div >
            <el-row>
                <el-select v-model="value" placeholder="全部领域">
                    <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
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
        </div>

        <el-table 
          :data="tableData"
          highlight-current-row
           @current-change="handleCurrentChange"
          style="width: 100%">
          <el-table-column
            label="品类名称"
            width="200">
              <template slot-scope="scope">
                  <span >{{scope.row.CategoryName}}</span>
                   <i class="el-icon-info" @click="showAttribute(scope.row.attribute)"></i>
            </template>
          </el-table-column>
           <el-table-column
            prop="type"
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
        <Pagination :currentPage="currentPage" :pageSize="pageSize" :total ="total" :pageSizes="pageSizes"
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
          pageSizes:[5,10,30,50],
            options: [{
            value: '选项1',
            label: '智能城市'
            }, {
            value: '选项2',
            label: '智能审核和'
            }, {
            value: '选项3',
            label: '智能工业'
            }, {
            value: '选项4',
            label: '边缘计算'
            }, {
            value: '选项5',
            label: '商业共享'
            }],
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
        this.getCategoryList() 
      },
      methods:{
       
         getCategoryList(){
            //  this.$API.getUser(this.currentPage,this.pageSize,this.username).then((res) => {
            //     this.tableData = res.data.objects
            //     this.total = res.data.num_results
            // })


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
             this.tableData =[
                 {
                    CategoryId: "0",
                    CategoryName: "路灯照明",
                    type: "公共服务",
                    "attribute": this.attribute,
                   
                },
                 {
                    CategoryId: "1",
                    CategoryName: "车辆定位卡",
                    type: "公共服务",
                    "attribute": this.attribute,
                   
                },
                 {
                    CategoryId: "2",
                    CategoryName: "水浸检测",
                    type: "公共服务",
                    "attribute": this.attribute,
                   
                },
                 {
                    CategoryId: "3",
                    CategoryName: "井盖移位检测",
                    type: "公共服务",
                    "attribute": this.attribute,
                   
                },
                 {
                    CategoryId: "4",
                    CategoryName: "垃圾满溢检测",
                    type: "公共服务",
                    "attribute": this.attribute,
                   
                },
                 {
                    CategoryId: "5",
                    CategoryName: "地磁检测器",
                    type: "公共服务",
                    "attribute": this.attribute,
                   
                },
                 {
                    CategoryId: "6",
                    CategoryName: "红外对射探测器",
                    type: "公共服务",
                    "attribute": this.attribute,
                   
                },           
            ]
            this.total = this.tableData.length
         },

         handleSizeChange(val) {
            console.log(`每页 ${val} 条`)
            this.pageSize = val 
            this.currentPage = 1
            this.getCategoryList()

          },
          handleCurrentChange(val) {
            console.log(`当前页: ${val}`)
            this.currentPage = val 
            this.getCategoryList()
          },

            handleClick(category){     
                this.$emit('select',category.CategoryId,category.CategoryName)            
            },
            handleCurrentChange(category){
                this.$emit('select',category.CategoryId,category.CategoryName)            
            },
            showAttribute(){
                alert(1);
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