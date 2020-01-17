<template>
    <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
     <el-form-item v-if="productId" label="产品" >        
         <el-input v-model="productName"  disabled></el-input>
    </el-form-item>
     <el-form-item v-else  label="产品" prop="pid">   
         <el-select v-model="ruleForm.pid" placeholder="请选择产品">
          <el-option
              v-for="item in productArrTemp"
              :key="item.id"
              :label="item.name"
              :value="item.id">
          </el-option>
        </el-select> 
    </el-form-item>
    <el-form-item label="设备名称" prop="name">
        <el-input v-model="ruleForm.name"></el-input>
    </el-form-item>
   
    <el-form-item label="备注名称" prop="remark">
         <el-input v-model="ruleForm.remark"></el-input>
    </el-form-item>
    
    </el-form>
 
</template>
<script>
  export default {
    props:{
      productArr:{
          type:Array,
          default: () => []
      },
      productId:{
        type:Number,
        default:0
      }
    },
     watch:{
        //监听productId,若发生变化，重新查询设备列表
        productId:{  
            handler:function(val,oldval){ 
                if(val!=oldval){
                    this.$nextTick(()=>{
                        this.init()
                    })
                }
            },  
            immediate:true,//关键
            deep:true
          },
      },
    
    data() {
      return {   
        ruleForm: { },
        productName:'',
        productArrTemp:[],
        rules: {
          pid: [
            { required: true, message: '请选择产品', trigger: 'blur' },
          ],
         name: [
            { required: true, message: '请输入设备名称', trigger: 'blur' },

          ],
          remark: [
            { required: true, message: '请输入备注', trigger: 'blur' },

          ],
              
        }
      };
    },
    created(){
       
       this.init()

    },
    methods: {
        init(){
          this.productArrTemp = JSON.parse(JSON.stringify(this.productArr))
          this.productArrTemp.shift();
          if(this.productId){
             this.ruleForm.pid = this.productId
             this.productName = this.productArrTemp.filter(item => item.id === this.productId)[0].name
             this.ruleForm =  {
              pid: this.productId,
              name: '',
              remark: '',       
            }
          }else{
              this.ruleForm =  {
              pid: null,
              name: '',
              remark: '',       
            }
          }
         
            
        },
        
      submitForm() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {
              this.$API_IOT.addDevice(this.ruleForm).then((res) => {
                if(res.data.status  === 'Y'){
                  let data  = res.data.data
                  //回调父组件方法关闭弹窗
                  this.$emit('close')
                  let param  = {'productKey':data.product_key,'deviceName':data.device_name,'deviceSecret':data.device_secret}  
                  //回调父组件方法,显示设备证书信息
                  this.$emit('showDeviceCertificate',param);
              
                }else{
                    this.$message.error(res.message);
                }
               
            })          
          } else {
            return false;
          }
        });
      },


     
    }
  }
</script>
<style scoped >
  .el-select{
      width: 440px;
  }
</style>