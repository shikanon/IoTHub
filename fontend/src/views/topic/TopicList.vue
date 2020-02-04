<template>
    <div name="topic-list">
        <el-button  v-if="type === 'product' " type="primary" @click="addTopic " >定义Topic类</el-button>
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
             property="topic_name"
             label="Topic类"
             width="800">
            </el-table-column>

            <el-table-column          
            label="操作权限"
            width="120"
            >
                <template slot-scope="scope">
                    <span > {{scope.row.operation | topicOperationFilter}} </span>
                </template>
            </el-table-column>
            <el-table-column
            property="desc"
            label="描述">
            </el-table-column>

            <el-table-column
            v-if="type === 'product'"
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
                <el-button type="text" size="small" @click="editTopic(scope.row,scope.$index)">编辑</el-button>
                <el-divider direction="vertical"></el-divider>
                <el-button type="text" size="small" @click="deleteTopic(scope.row.id)">删除</el-button>
            </template>
            </el-table-column>
            <el-table-column
                v-if="type === 'device'"
                fixed="right"
                label="操作"
                width="200">
                <template slot-scope="scope" v-if="scope.row.operation === 2"> 
                    <el-button type="text" size="small"  @click="editTopic(scope.row)">发布消息</el-button>      
                </template>
            </el-table-column>

        </el-table>
        <el-dialog
            title="定义Topic"
            :visible.sync="dialogVisible"
            width="30%"
            >
            <addTopic ref="addTopic" :form="topicForm" @refresh="refresh"></addTopic>
            <span slot="footer" class="dialog-footer">
                <el-button @click="addTopicCancel">取 消</el-button>
                <el-button type="primary" @click="addTopicSubmit">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
  import AddTopic from './AddTopic'

  export default {
      props:{
          tableData:{
              type:Array,
              default: () => []
          } ,
          type:{
            type:String,
            default:''
          },
          queryKey:{
            type:Number,
            default:0
          }     
      },
    components: { addTopic:AddTopic },

    data() {
      return {
        dialogVisible: false,
        topicForm:{},
        group:'',
        editRowId:0
      }
    },

    created(){
    },

    methods: {
 
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

        addTopicCancel(){
            if(this.addflg){
                this.topicForm = {pid:this.queryKey,topic_name:this.tableData[0].topic_name}
            }else{
               //重置
               console.log("1111",this.tableData[this.editRowId] )
               let row  = this.tableData[this.editRowId] 
               this.topicForm = {id:row.id,operation:row.operation,desc:row.desc,topic_name:row.topic_name}
            }
            this.dialogVisible = false 
        },

        editTopic(row,index){  
            this.editRowId = index 
            this.topicForm = {id:row.id,operation:row.operation,desc:row.desc,topic_name:row.topic_name}
            this.dialogVisible = true 
            this.addflg = false

        },
        addTopic(){
            this.topicForm = {pid:this.queryKey,topic_name:this.tableData[0].topic_name}
            this.dialogVisible = true 
            this.addflg = true
        },

        refresh(){
            this.$emit('refresh')
        },

        deleteTopic(topicId){
            console.log(topicId)
             this.$confirm('此操作将永久删除该纪录, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
                this.$API_IOT.deleteTopic(topicId).then((res) => {
                    if(res.data.status === "Y"){
                        this.$message.success('删除topic成功')    
                        this.$emit('refresh')
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
        handleCurrentChange(){

        }
     
    }
  }
</script>