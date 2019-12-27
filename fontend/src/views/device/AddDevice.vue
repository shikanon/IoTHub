<template>
    <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
     <el-form-item label="产品" prop="ProductKey">
          <el-select v-model="ruleForm.ProductKey" placeholder="请选择产品">
                <el-option
                v-for="item in productArr"
                :key="item.ProductKey"
                :label="item.ProductName"
                :value="item.ProductName">
                </el-option>
            </el-select>   
    </el-form-item>
    <el-form-item label="DeviceName" prop="DeviceName">
        <el-input v-model="ruleForm.DeviceName"></el-input>
    </el-form-item>
   
    <el-form-item label="备注名称" prop="Nickname">
         <el-input v-model="ruleForm.Nickname"></el-input>

    </el-form-item>
    
    </el-form>
 
</template>
<script>
  export default {
    props:{
          productArr:{
              type:Array,
              default: () => []
          }
    },
    data() {
      return {
        ruleForm: {
          IotId: '',
          DeviceSecret: '',
          ProductKey: '',
          JoinEui: '',
          DeviceName: '',
          Nickname: '',
        },
        rules: {
          ProductKey: [
            { required: true, message: '请选择产品', trigger: 'change' },
            {  message: '请选择产品', trigger: 'blur' },

          ]
              
        }
      };
    },
    created(){
       this.init()

    },
    methods: {
        init(){
            this.ruleForm={
            IotId: '',
            DeviceSecret: '',
            ProductKey: '',
            JoinEui: '',
            DeviceName: '',
            Nickname: '',
            }
        },
        
      submitForm() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {
            alert('submit!');
            //回调父组件方法关闭弹窗
            this.$emit('close')
            let param  = {productKey:'a1ELejzj0h9',deviceName:'t4qOQEYkD0UAWiTHC08i','deviceSecret':'7wG2wOQj0pGSa51TpGXN2GyPeCblc2Uy'}
            //回调父组件方法
             this.$emit('showDeviceCertificate',param);
          } else {
            console.log('error submit!!');
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