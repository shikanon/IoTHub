<template>
    <div name="step1">
        <p class="step-desc">注册连接的设备信息</p>
        <el-form label-position="top" :model="form" ref="form" :rules="rules" label-width="100px" >
            <el-form-item label="产品名称" prop="ProductName">
                <el-input v-model="form.ProductName" placeholder="请输入产品名称"></el-input>
            </el-form-item>
            <el-form-item label="设备名称（可选）" prop="DeviceName">
                <el-input v-model="form.DeviceName" placeholder="请输入设备名称"></el-input>
            </el-form-item>
        </el-form>  
    </div>
</template>

<script>

export default {
  
    data(){
        return{
            form:{ProductName:'',DeviceName:''},
            rules:{
                 ProductName:[
                     { required: true, message: '请输入产品名称', trigger: 'blur' }
                     ],
                 DeviceName: [
                    { required: true, message: '请输入设备名称', trigger: 'blur' }
                ]
            }                
        }
    },
    

    methods:{
        nextStep() {
            this.$refs.form.validate((valid) => {
                if (valid) {
                    alert('submit')
                    this.$emit('next')
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
    }
}
</script>