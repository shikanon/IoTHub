<template>
    <div name="group-info"> 
        <p class="label-title"> 分组信息
          <el-button type="text" @click="editProductVisible = true">编辑</el-button>
        </p>
        <div class="table-info">
          <div class="table-row">
                <div class="table-row-label">分组名称</div> 
                <div class="table-row-info">
                    <span>{{group.GroupName}}</span>
                </div>
           
                <div class="table-row-label">分组层级</div>
                <div class="table-row-info">
                    <span>分组/{{group.GroupName}}</span>
                </div>
           
                <div class="table-row-label">分组ID</div>
                <div class="table-row-info">
                    <span>{{group.GroupId}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">设备总数</div> 
                <div class="table-row-info">
                    <span>{{group.DeviceCount}}</span>
                </div>          
                <div class="table-row-label">激活设备</div>
                <div class="table-row-info">
                    <span>{{group.DeviceActive}}</span>
                </div>
            
                <div class="table-row-label">当前在线</div>
                <div class="table-row-info">
                    <span>{{group.DeviceOnline}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">创建时间</div> 
                <div class="table-row-info">
                    <span>{{group.UtcCreate}}</span>
                </div>
            </div>
              <div class="table-row">
                <div class="table-row-label">分组描述</div> 
                <div class="table-row-info">
                    <span>{{group.GroupDesc}}</span>
                </div>
            </div>
        </div>
        <p class="label-title">标签信息
          <el-button type="text" @click="addLabelVisible = true">编辑</el-button>
        </p>
        <p>产品标签
         <span v-for="(item,index) in labelArr" :key = "index" class="label-span">{{item['key']}}:{{item['value']}}</span>
         </p>
         <el-dialog title="编辑分组信息" :visible.sync="editProductVisible" width="26%">
            <EditGroup ref="editGroup" :group="group" @close="editProductVisible = false"></EditGroup>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="editProductSubmit">确 定</el-button>
                <el-button @click="editProductVisible = false">取 消</el-button>
            </span>
        </el-dialog>
        <el-dialog title="添加标签" :visible.sync="addLabelVisible" width="26%">
            <AddLabel ref="editProduct" :labelArr="labelArr" type="product" @close="addLabelVisible = false"></AddLabel>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addLabelSubmit">确 定</el-button>
                <el-button @click="addLabelVisible = false">取 消</el-button>
            </span>
        </el-dialog>
    </div>
  </template>

  <script>
  import EditGroup  from './EditGroup'
    export default {
      components:{EditGroup},
      props:{
        group:{
          type:Object,
          default:() =>({})
        }      
      },
      data() {
        return { 
          editProductVisible:false, 
          addLabelVisible:false,
          labelArr:[] 
            
        }
      },
      computed:{
       
      },
      created(){
        this.productName = this.$route.params.product
      },
      methods:{
          editProductSubmit(){
              this.$refs.editProduct.submitForm()
          },
          addLabelSubmit(){
            this.addLabelVisible = false
            this.labelArr = this.$refs.editProduct.label
          }

      }
    }
  </script>

  <style scoped>
    .label-title{
      font-weight: bold;
    }

   
    .label-span{
      display: inline-block;
      padding: 0 8px;
      background: #ebedef;
      color: #333;
      margin-right: 8px;
      margin-bottom: 8px;
      cursor: pointer;
      position: relative;
      border-radius: 4px;
    }
  </style>