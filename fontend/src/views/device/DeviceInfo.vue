<template>
    <div name="device-info"> 
        <p class="label-title"> 设备信息
        </p>
        <div class="table-info">
          <div class="table-row">
                <div class="table-row-label">产品名称</div> 
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                </div>
           
                <div class="table-row-label">ProductKey</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                    <CopyBtn :content="device.ProductName"></CopyBtn>
                </div>
           
                <div class="table-row-label">区域</div>
                <div class="table-row-info">
                    <span>{{device.Region}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">节点类型</div> 
                <div class="table-row-info">
                    <span>{{device.CategoryName}}</span>
                </div>          
                <div class="table-row-label">DeviceName</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                    <CopyBtn :content="device.ProductName"></CopyBtn>
                </div>
            
                <div class="table-row-label">认证方式</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">备注名称</div> 
                <div class="table-row-info">
                    <span>{{device.deviceKey}}</span>
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
                    <span>{{device.ProductName}}</span>
                </div>
          
                <div class="table-row-label">激活时间</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                </div>
                <div class="table-row-label">最后上线时间</div>
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
                </div>
            </div>
             <div class="table-row">
                <div class="table-row-label">当前状态</div> 
                <div class="table-row-info">
                    <span>{{device.ProductName}}</span>
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
        <p>设备标签</p>
        <el-dialog title="添加标签" :visible.sync="addLabelVisible" width="26%">
            <AddLabel  :labelArr="labelArr" type="device" @close="addLabelVisible = false"></AddLabel>
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
        ProductName:{
          type:String,
          default:''
        }      
      },
      data() {
        return {      
            device:{},
            addLabelVisible:false,
            labelArr:[]
        }
      },
      computed:{
       
      },
      created(){
          this.getDeviceInf()
      },
      methods:{
          
        getDeviceInf(){
                this.device ={Status: "UNACTIVE",
                Owner: true,
                ProductName: "3231323",
                DeviceSecret: "HVi2e9hLKw9vsRkd1hoLq3rH6qzD05Cc",
                UtcOnline: "",
                IotId: "2D87VhVL3Pn0G8v9qE2K000100",
                GmtCreate: "2019-12-23 15:43:25",
                UtcCreate: "2019-12-23T07:43:25.000Z",
                UtcActive: "",
                GmtActive: "",
                NodeType: 0,
                Region: "cn-shanghai",
                GmtOnline: "",
                ProductKey: "a1Ibli2tqC2",
                DeviceName: "2D87VhVL3Pn0G8v9qE2K",
            }
        }
        ,addLabelSubmit(){

        }
        ,test(){
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