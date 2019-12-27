<template>
    <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
        <el-form-item label="产品" prop="product">
            <el-select v-model="ruleForm.product" placeholder="请选择产品">
                    <el-option
                    v-for="item in productArr"
                    :key="item.ProductKey"
                    :label="item.ProductName"
                    :value="item.ProductName">
                    </el-option>
                </el-select>   
        </el-form-item>
        <el-form-item label="添加方式" prop="type">
              <el-radio v-model="ruleForm.type" label="1"  border size="medium">自动生成</el-radio>
              <el-radio v-model="ruleForm.type" label="2"  border size="medium">批量上传</el-radio>
        </el-form-item>
    
        <el-form-item label="设备数量" prop="count">
            <el-input-number v-model="ruleForm.count" controls-position="right"  :min="1"  label="请输入设备数量"></el-input-number>
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
          product: '',
          type: '1',
          count: 1 ,
        },
        rules: {
          product: [
            { required: true, message: '请选择产品', trigger: 'change' },
            {  message: '请选择产品', trigger: 'blur' },
          ],
            type: [
                { required: true, message: '请选择添加方式', trigger: 'change' },
            ],
          
           count: [
            { required: true,type:'number',message: '请输入设备数量', trigger: 'blur' },

          ]
              
        }
      };
    },
    created(){
       this.init()

    },
    methods: {
        init(){
            this.ruleForm= {    
            product: '',
            type: '1',
            count: 1 ,
            }
        },    
    
      submitForm() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {
            alert('submit!');
            //回调父组件方法
            this.$emit('close')
            //回调父组件方法显示结果
             this.$emit('showAddResult');
            
            setTimeout(()=>{
                this.$emit('addSuccess',this.ruleForm.count)
             },3*1000)
             
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