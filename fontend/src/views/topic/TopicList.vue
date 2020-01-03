<template>
    <div name="topic-list">
        <el-button  v-if="type === 'product'" type="primary" @click="dialogVisible = true " >定义Topic类</el-button>
        <p v-if="type === 'base' || type === 'physics'">{{type === 'base' ? '基础通信 Topic 列表': '物模型通信 Topic 列表'}}</p>
        <el-table
            :data="tableData"
            :span-method="objectSpanMethod"
            :border="type === 'base' || type === 'physics'"
            highlight-current-row
            @current-change="handleCurrentChange"
            style="width: 100%">
            
             <el-table-column
                v-if="type === 'base' || type === 'physics'"
                property="name"
                label="功能"
                width="100">
            </el-table-column>

            <el-table-column
             property="TopicShortName"
             label="Topic类"
             width="400">
            </el-table-column>

            <el-table-column          
            label="操作权限"
            width="120"
            >
                <template slot-scope="scope">
                    <span > {{scope.row.Operation | topicOperationFilter}} </span>
                </template>
            </el-table-column>
            <el-table-column
            property="Desc"
            label="描述">
            </el-table-column>

            <el-table-column
            v-if="type === 'product'"
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
                <el-button type="text" size="small" @click="editTopicClick(scope.row)">编辑</el-button>
                <el-divider direction="vertical"></el-divider>
                <el-button type="text" size="small" @click="deleteClick(scope.$index, scope.row.Id)">删除</el-button>
            </template>
            </el-table-column>
              <el-table-column
            v-if="type === 'device' "
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
                <el-button type="text" size="small" v-if="scope.row.Operation === '1'" @click="editTopicClick(scope.row)">发布消息</el-button>
                
            </template>
            </el-table-column>

        </el-table>
        <el-dialog
            title="定义Topic"
            :visible.sync="dialogVisible"
            width="30%"
            >
            <addTopic ref="addTopic" :form="topicForm" ></addTopic>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="addTopicSubmit">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
  import AddTopic from './AddTopic'

  export default {
      props:{
          type:{
              type:String,
              default:''
          } ,
           watch:{
            type:{  
                handler:function(val,oldval){ 
                    if(val!=oldval){
                        this.$nextTick(()=>{
                            this.getTopicList()
                        })
                    }
                },  
                immediate:true,//关键
                deep:true
             },
            }
      },
    components: { addTopic:AddTopic },

    data() {
      return {
        tableData: [],
        dialogVisible: false,
        topicForm:{},
        group:''
      }
    },

    created(){
        this.getTopicList()
    },

    methods: {

     getTopicList(){
         console.log(this.type)
         if(this.type === 'base'){
             this.tableData = [               
                 {
                    "name":"固件升级",
                    "Desc": "1",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/123123",
                    "Operation": "2",
                    "Id": "7069078",
                    "ProductKey": "a1ELejzj0h9"
                    },
                    {
                    "name":"固件升级",
                    "Desc": "2",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update",
                    "Operation": "0",
                    "Id": "7040770",
                    "ProductKey": "a1ELejzj0h9"
                    },
                    {
                    "name":"设备影子",
                     "Desc": "3",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update/error",
                    "Operation": "0",
                    "Id": "7040771",
                    "ProductKey": "a1ELejzj0h9"
                    },
                     {
                    "name":"设备影子",
                    "Desc": "4",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update/error",
                    "Operation": "0",
                    "Id": "7040771",
                    "ProductKey": "a1ELejzj0h9"
                    },
                    {
                    "name":"设备影子",
                    "Desc": "5",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/get",
                    "Operation": "1",
                    "Id": "7040772",
                    "ProductKey": "a1ELejzj0h9"
                }]
            
            
         }else if(this.type === 'physics'){
             this.tableData = [               
                 {
                    "name":"属性设置",
                    "Desc": "1",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/123123",
                    "Operation": "2",
                    "Id": "7069078",
                    "ProductKey": "a1ELejzj0h9"
                    },
                    {
                    "name":"属性设置",
                    "Desc": "2",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update",
                    "Operation": "0",
                    "Id": "7040770",
                    "ProductKey": "a1ELejzj0h9"
                    },
                    {
                    "name":"服务调用",
                     "Desc": "3",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update/error",
                    "Operation": "0",
                    "Id": "7040771",
                    "ProductKey": "a1ELejzj0h9"
                    },
                     {
                    "name":"服务调用",
                    "Desc": "4",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update/error",
                    "Operation": "0",
                    "Id": "7040771",
                    "ProductKey": "a1ELejzj0h9"
                    },
                    {
                    "name":"服务调用",
                    "Desc": "5",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/get",
                    "Operation": "1",
                    "Id": "7040772",
                    "ProductKey": "a1ELejzj0h9"
                     },
                    {
                    "name":"属性上报",
                    "Desc": "6",
                    "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/get",
                    "Operation": "1",
                    "Id": "7040772",
                    "ProductKey": "a1ELejzj0h9"
                     }
                ]
            
            
         }else{
              this.tableData = [ 
                {
                "Desc": "1",
                "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/123123",
                "Operation": "2",
                "Id": "7069078",
                "ProductKey": "a1ELejzj0h9"
                },
                {
                "Desc": "2",
                "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update",
                "Operation": "0",
                "Id": "7040770",
                "ProductKey": "a1ELejzj0h9"
                },
                {
                 "Desc": "3",
                "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update/error",
                "Operation": "0",
                "Id": "7040771",
                "ProductKey": "a1ELejzj0h9"
                },
                {
                "Desc": "4",
                "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/get",
                "Operation": "1",
                "Id": "7040772",
                "ProductKey": "a1ELejzj0h9"
            }]
         }
       
     },
     //当前行，当前列,第几行,第几列
     objectSpanMethod({ row, column, rowIndex, columnIndex }){
        if (this.type === 'base' || this.type === 'physics'){
            if (columnIndex === 0) {              
                let length = this.tableData.filter((item,index) => item.name === row.name).length //获取name相同的数据长度
                let index = this.tableData.findIndex((item,index) => item.name === row.name) //获取第一个name相同的索引
                if(rowIndex === index ){
                    return  [length,1]
                }else{
                    return  [0,1]
                }                             
            }
        }else{
             return  [1,1]    
        }     
     },
    

     addTopicSubmit(){
         this.dialogVisible = false 
         this.$refs.addTopic.submit()
     },

     editTopicClick(row){
         this.topicForm = {Id:row.Id,Operation:row.Operation,Desc:row.Desc,TopicShortName:row.TopicShortName,ProductKey:row.ProductKey}
         console.log(this.topicForm)
         this.dialogVisible = true 
     },
     handleCurrentChange(){

     }
     
    }
  }
</script>