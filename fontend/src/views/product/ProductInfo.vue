<template>
    <div name="product-info"> 
        <p class="label-title"> 产品信息
          <el-button type="text" @click="editProduct">编辑</el-button>
        </p>
        <div class="table-info">
          <div class="table-row">
                <div class="table-row-label">产品名称</div> 
                <div class="table-row-info">
                    <span>{{product.name}}</span>
                </div>          
                <div class="table-row-label">节点类型</div>
                <div class="table-row-info">
                    <span>{{product.node_type }}</span>              
                </div> 
                <div class="table-row-label">创建时间</div>
                <div class="table-row-info">
                    <span>{{product.create_time}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">所属品类</div> 
                <div class="table-row-info">
                    <span>{{product.object_model_name}}</span>
                </div>          
                <div class="table-row-label">数据格式</div>
                <div class="table-row-info">
                    <span>{{product.data_format}}</span>
                </div>
            
                <div class="table-row-label">认证方式</div>
                <div class="table-row-info">
                    <span>{{product.auth_method}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">
                  动态注册
              
                  <el-tooltip placement="top" effect="light">
                      <i class="el-icon-question"></i>
                    <div slot="content">
                      设备动态注册无需一一烧录设备证书，每台设备烧录相同的产品证书
                      <br/>
                      即ProductKey和ProductSecret，云端鉴权通过后下发设备证书，
                      <br/>
                     云端鉴权通过后下发设备证书，您可以根据需要开启或关闭动态注册，
                     <br/>
                     保障安全性。
                      </div>
                 </el-tooltip>
                </div> 
                <div class="table-row-info">
                  <el-switch
                    v-model="value"
                     :active-text="value ? '已开启':'已关闭'"
                     >
                  </el-switch>
                </div>
          
                <div class="table-row-label">状态</div>
                <div class="table-row-info">
                    <span>{{product.status}}</span>
                </div>
                <div class="table-row-label">连网协议</div>
                <div class="table-row-info">
                    <span>{{product.network_way}}</span>
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
            <EditProduct ref="editProduct" :product="productTemp" @close="editProductVisible = false"></EditProduct>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="editProductSubmit" @success="updateSuccess($event)">确 定</el-button>
                <el-button @click=" editProductCancel">取 消</el-button>
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
          productTemp:{},
          labelArr:[],
          value:false
            
        }
      },
      computed:{
       
      },
      created(){
        
      },
      methods:{
          //编辑产品信息
          editProduct(){
              this.productTemp = JSON.parse(JSON.stringify(this.product))
              this.editProductVisible = true
          },

          //编辑产品信息提交
          editProductSubmit(){
              this.$refs.editProduct.submitForm()
          },

          //取消编辑产品信息
          editProductCancel(){
              this.editProductVisible = false
              this.productTemp = this.product
          },

          updateSuccess(event){
            this.editProductVisible = false
            this.product.name = event.name
          },

          //添加标签提交
          addLabelSubmit(){
            this.addLabelVisible = false
            this.labelArr = this.$refs.addLabel.label
          },
          

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