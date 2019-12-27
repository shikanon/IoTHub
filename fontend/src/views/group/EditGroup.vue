<template>
    <div name="edit-group">
        <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
          <el-form-item label="父组" prop="superGroup">
              <el-input v-model="ruleForm.superGroup" :placeholder="ruleForm.superGroup" disabled></el-input>
          </el-form-item>       
          <el-form-item label="分组名称" prop="GroupName">
            <el-input v-model="ruleForm.GroupName"  disabled></el-input>   
          </el-form-item>      
          <el-form-item label="分组描述" prop="Desc">
               <el-input type="textarea" v-model="ruleForm.desc"></el-input>
          </el-form-item>     
      
        </el-form>
    </div>
</template>
<script>
  export default {
      props:{
          product:{
              tyoe:Object,
              default:() => ({})
          }
      },
    data() {
      return {
        ruleForm: {      
          GroupName:"000000000000",
          superGroup:0,
          desc:1,
      
        },
        rules: {
          ProductName: [
            { required: true, message: '请输入产品名称', trigger: 'blur' },
            // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
        },
           
      };
    },
    created(){
       
    },
    methods: {
      
      submitForm() {
        this.$refs['ruleForm'].validate((valid) => {
          if (valid) {
            alert('submit!');
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
     .el-form {
        width:460px;
    } 
</style>