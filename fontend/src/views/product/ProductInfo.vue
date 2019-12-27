<template>
    <div name="product-info"> 
        <p class="label-title"> 产品信息
          <el-button type="text" @click="editProductVisible = true">编辑</el-button>
        </p>
        <div class="table-info">
          <div class="table-row">
                <div class="table-row-label">产品名称</div> 
                <div class="table-row-info">
                    <span>{{product.ProductName}}</span>
                </div>
           
                <div class="table-row-label">节点类型</div>
                <div class="table-row-info">
                    <span>{{product.AuthType}}</span>
                </div>
           
                <div class="table-row-label">创建时间</div>
                <div class="table-row-info">
                    <span>{{product.GmtCreate}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">所属品类</div> 
                <div class="table-row-info">
                    <span>{{product.CategoryName}}</span>
                </div>          
                <div class="table-row-label">数据格式</div>
                <div class="table-row-info">
                    <span>{{product.deviceName}}</span>
                </div>
            
                <div class="table-row-label">认证方式</div>
                <div class="table-row-info">
                    <span>{{product.deviceName}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">动态注册</div> 
                <div class="table-row-info">
                    <span>{{product.productKey}}</span>
                </div>
          
                <div class="table-row-label">状态</div>
                <div class="table-row-info">
                    <span>{{product.deviceName}}</span>
                </div>
                <div class="table-row-label">连网协议</div>
                <div class="table-row-info">
                    <span>{{product.deviceName}}</span>
                </div>
            </div>
        </div>
        <p class="label-title">标签信息
          <el-button type="text" @click="addLabelVisible = true">编辑</el-button>
        </p>
        <p>产品标签
         <span v-for="(item,index) in labelArr" :key = "index" class="label-span">{{item['key']}}:{{item['value']}}</span>
         </p>
         <el-dialog title="编辑产品信息" :visible.sync="editProductVisible" width="26%">
            <EditProduct ref="editProduct" :product="product" @close="editProductVisible = false"></EditProduct>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="editProductSubmit">确 定</el-button>
                <el-button @click="editProductVisible = false">取 消</el-button>
            </span>
        </el-dialog>
        <el-dialog title="添加标签" :visible.sync="addLabelVisible" width="26%">
            <AddLabel ref="addLabel" :labelArr="labelArr" type="product" @close="addLabelVisible = false"></AddLabel>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addLabelSubmit">确 定</el-button>
                <el-button @click="addLabelVisible = false">取 消</el-button>
            </span>
        </el-dialog>
    </div>
  </template>

  <script>
  import EditProduct from './EditProduct'

    export default {
      components:{EditProduct},
      props:{
        product:{
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