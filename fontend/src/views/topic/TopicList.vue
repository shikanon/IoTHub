<template>
    <div name="topic-list">
        <el-button type="primary" @click="addTopicShow()" v-if="addBtnVisible">定义Topic类</el-button>
        <el-table
            ref="singleTable"
            :data="tableData"
            highlight-current-row
            @current-change="handleCurrentChange"
            style="width: 100%">
            
            <el-table-column
            property="TopicShortName"
            label="Topic类"
            width="500">
            </el-table-column>
            <el-table-column
            label="Operation"
            width="120">
                <template slot-scope="scope">
                        <span>{{scope.Operation === 0 ? '发布': scope.Operation === 1 ? '订阅':'发布和订阅'}}</span>
                </template>
            </el-table-column>
            <el-table-column
            property="Desc"
            label="描述">
            </el-table-column>
            <el-table-column
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
                <el-button type="text" size="small" @click="editTopicClick(scope.row)">编辑</el-button>
                <el-divider direction="vertical"></el-divider>
                <el-button type="text" size="small" @click="deleteClick(scope.$index, scope.row.Id)">删除</el-button>
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
          addBtnVisible:{
              type:Boolean,
              default:true
          }
      },
    components: { addTopic:AddTopic },

    data() {
      return {
        tableData: [],
        dialogVisible: false,
        topicForm:{}
      }
    },
    created(){
        this.getTopicList()
    },

    methods: {

     getTopicList(){
        this.tableData = [ 
            {
            "Desc": "13123",
            "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/123123",
            "Operation": "2",
            "Id": "7069078",
            "ProductKey": "a1ELejzj0h9"
            },
            {
            "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update",
            "Operation": "0",
            "Id": "7040770",
            "ProductKey": "a1ELejzj0h9"
            },
            {
            "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/update/error",
            "Operation": "0",
            "Id": "7040771",
            "ProductKey": "a1ELejzj0h9"
            },
            {
            "TopicShortName": "/a1ELejzj0h9/${deviceName}/user/get",
            "Operation": "1",
            "Id": "7040772",
            "ProductKey": "a1ELejzj0h9"
          }]
     },
     addTopicShow(){
         this.dialogVisible = true 
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