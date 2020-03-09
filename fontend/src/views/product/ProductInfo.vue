<template>
    <div name="product-info"> 
        <p class="label-title"> 产品信息
          <el-button type="text" @click="editProduct">编辑</el-button>
        </p>
        <div class="table-info">
            <div class="table-row table-row-8">
              <div class="table-row-item"> 
                <div class="table-row-label">产品名称</div> 
                <div class="table-row-info">
                    <span>{{product.name}}</span>
                </div>
              </div>
            </div>     
            <div class="table-row table-row-8">
                <div class="table-row-item">         
                  <div class="table-row-label">节点类型</div>
                  <div class="table-row-info">
                      <span>{{product.node_type }}</span>              
                  </div> 
                </div>
            </div>   
            <div class="table-row table-row-8">
              <div class="table-row-item">   
                <div class="table-row-label">创建时间</div>
                <div class="table-row-info">
                    <span>{{product.create_time}}</span>
                </div>
              </div>
            </div>      
            <div class="table-row table-row-8">
              <div class="table-row-item"> 
                <div class="table-row-label">所属品类</div> 
                <div class="table-row-info">
                    <span>{{product.object_model_name}}</span>
                </div>
              </div>
            </div>
            <div class="table-row table-row-8">
              <div class="table-row-item">             
                <div class="table-row-label">数据格式</div>
                <div class="table-row-info font-size-12">
                    <span>{{product.data_format}}</span>
                </div>
              </div>
            </div>
            <div class="table-row table-row-8">
              <div class="table-row-item"> 
                <div class="table-row-label">认证方式</div>
                <div class="table-row-info">
                    <span>{{product.auth_method}}</span>
                </div>
              </div>
            </div>
            <div class="table-row table-row-8">
                <div class="table-row-item"> 
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
                </div> 
            </div> 
            <div class="table-row table-row-8">
              <div class="table-row-item"> 
                <div class="table-row-label">状态</div>
                <div class="table-row-info">
                    <span>{{product.status}}</span>
                </div>
               </div>
             </div> 
             <div class="table-row table-row-8">
              <div class="table-row-item">     
                <div class="table-row-label">连网协议</div>
                <div class="table-row-info">
                    <span>{{product.network_way}}</span>
                </div>              
              </div>
             </div> 
            <div class="table-row table-row-24">
              <div class="table-row-item">  
                <div class="table-row-label">产品描述</div> 
                <div class="table-row-info">
                    <span>{{product.desc}}</span>
                </div>                 
             </div>
            </div> 
        </div>
        <p class="label-title">标签信息
          <el-button type="text" @click="editProductLabel">编辑</el-button>
        </p>
        <p>产品标签
         <span v-for="(item,index) in product.label" :key = "index" class="label-span">{{item['key']}}:{{item['value']}}</span>
         </p>
         <el-dialog title="编辑产品信息" :visible.sync="editProductVisible" width="26%">
            <EditProduct ref="editProduct" :product="productTemp" @close="editProductVisible = false" @success="updateSuccess"></EditProduct>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="editProductSubmit" @success="updateSuccess($event)">确 定</el-button>
                <el-button @click=" editProductCancel">取 消</el-button>
            </span>
        </el-dialog>
        <el-dialog title="添加标签" :visible.sync="addLabelVisible" width="26%">
            <AddLabel ref="addLabel" :labelArr="product.label ? JSON.parse(JSON.stringify(product.label)) :[]" type="product" @close="addLabelVisible = false"></AddLabel>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addLabelSubmit">确 定</el-button>
                <el-button @click="addLabelCancel">取 消</el-button>
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
          value:false
            
        }
      },
      computed:{
       
      },
      created(){
        //this.product.label = [{"key":"1","value":"111"},{"key":"2","value":"222"}]
      },
      methods:{
          //编辑产品信息
          editProduct(){
              this.productTemp = JSON.parse(JSON.stringify(this.product))       
              this.editProductVisible = true
              this.$refs.editProduct.init()
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
            this.product.desc = event.desc
          },

          //编辑标签
          editProductLabel(){
            this.addLabelVisible = true
            this.$refs.addLabel.init()
          },

          //添加标签提交
          addLabelSubmit(){

            let params = {pid:this.product.id, label:this.$refs.addLabel.label} 
            this.$API_IOT.editProductLabel(params).then((res) => {
               if (res.data.status === 'Y') {
                  this.$message.success('编辑标签成功')
                   this.addLabelVisible = false
                   this.product.label = this.$refs.addLabel.label
              } else {
                  this.$message.error(res.message);
              } 
               
            }) 
            
          },

          addLabelCancel(){
            this.addLabelVisible = false
            this.$refs.addLabel.init()
          }
          

      }
    }
  </script>

  <style scoped>
    .label-title{
      font-weight: bold;
    }

   
  
  </style>