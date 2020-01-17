<template>
    <div name="device-info"> 
        <p class="label-title"> 设备信息</p>
  
        <div class="table-info">
          <div class="table-row">
                <div class="table-row-label">产品名称</div> 
                <div class="table-row-info">
                    <span>{{device.name}}</span>
                </div>
           
                <div class="table-row-label">ProductKey</div>
                <div class="table-row-info">
                    <span>{{device.product_key}}</span>
                    <CopyBtn :content="device.product_key"></CopyBtn>
                </div>
           
                <div class="table-row-label">区域</div>
                <div class="table-row-info">
                    <span>{{device.Region}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">节点类型</div> 
                <div class="table-row-info">
                    <span>{{device.node_type}}</span>
                </div>          
                <div class="table-row-label">DeviceName</div>
                <div class="table-row-info">
                    <span>{{device.name}}</span>
                    <CopyBtn :content="device.name"></CopyBtn>
                </div>
            
                <div class="table-row-label">认证方式</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">备注名称</div> 
                <div class="table-row-info">
                    <span>{{device.remark}}</span>
                </div>        
                <div class="table-row-label">IP地址</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                </div>
                <div class="table-row-label">固件版本</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                </div>
            </div>
             <div class="table-row">
                <div class="table-row-label">添加时间</div> 
                <div class="table-row-info">
                    <span>{{device.create_time}}</span>
                </div>         
                <div class="table-row-label">激活时间</div>
                <div class="table-row-info">
                    <span>{{device.activate_time}}</span>
                </div>
                <div class="table-row-label">最后上线时间</div>
                <div class="table-row-info">
                    <span>{{device.last_online_time}}</span>
                </div>
            </div>
             <div class="table-row">
                <div class="table-row-label">当前状态</div> 
                <div class="table-row-info">
                    <span>{{device.status_id | deviceStatusFilter}}</span>
                </div>
          
                <div class="table-row-label">实时延迟</div>
                <div class="table-row-info">
                    <el-button type="text" @click="test">测试</el-button>
                </div>
               
            </div>
        </div>
         <p class="label-title">设备扩展信息</p>
         <div class="table-info">
          <div class="table-row">
                <div class="table-row-label">SDK 语言</div> 
                <div class="table-row-info">
                    <span>Python</span>
                </div>       
                <div class="table-row-label">版本号</div>
                <div class="table-row-info">
                    <span>1.2.0</span>
                </div> 
                <div class="table-row-label">模组商</div>
                <div class="table-row-info">
                    <span>-</span>
                </div>
            </div>
             <div class="table-row">
                <div class="table-row-label">模组信息</div> 
                <div class="table-row-info">
                    <span>-</span>
                </div>                 
             </div>
         </div>
        <p class="label-title">标签信息
          <el-button type="text" @click="addLabelVisible = true ">编辑</el-button>
        </p>
        <p>设备标签
             <span v-for="(item,index) in labelArr" :key = "index" class="label-span">{{item['key']}}:{{item['value']}}</span>
        </p>
        <el-dialog title="添加标签" :visible.sync="addLabelVisible" width="26%">
            <AddLabel  ref="addLabel" :labelArr="labelArr" type="device" @close="addLabelVisible = false"></AddLabel>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="addLabelSubmit">确 定</el-button>
                <el-button @click="addLabelVisible = false">取 消</el-button>
            </span>
        </el-dialog>
    </div>
  </template>

  <script>
    export default {
      props:{
        device:{
          type:Object,
          default:() => ({})
        }      
      },
      data() {
        return {      
            addLabelVisible:false,
            labelArr:[]
        }
      },
      computed:{
       
      },
      created(){
      },
      methods:{
          
        addLabelSubmit(){
           this.labelArr =  this.$refs.addLabel.label
           this.addLabelVisible = false
           console.log(`编辑label信息提交${this.labelArr }`)

        },
        test(){
           this.$confirm('测试网络延迟时将会向设备发送一条空消息，请确认对设备端业务没有影响！', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
                }).then(() => {
                    this.$message({
                        type: 'success',
                        message: '测试成功!'
                    });
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '已取消测试'
                    });          
            });
        }
      }
    }
  </script>

  <style scoped>
    .label-title{
      font-weight: bold;
    }

  </style>

