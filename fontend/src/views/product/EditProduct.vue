<template>
    <div name="edit-product">
        <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
          <el-form-item label="产品名称" prop="ProductName">
              <el-input v-model="ruleForm.ProductName" :placeholder="ruleForm.ProductName"></el-input>
          </el-form-item>       
          <el-form-item label="节点类型" prop="NodeType">
            <el-input v-model="ruleForm.NodeType"  disabled></el-input>   
          </el-form-item>
           <el-form-item label="数据格式" prop="DataFormat">
                <el-input v-model="ruleForm.DataFormat" disabled></el-input>   
           </el-form-item>
           <el-form-item label="所属品类" prop="type">
               <el-input v-model="ruleForm.type" disabled></el-input>   
          </el-form-item>
          <el-form-item label="连网协议" prop="NetType">
               <el-input v-model="ruleForm.NetType" disabled></el-input>        
          </el-form-item>
          <el-form-item label="产品描述" prop="Desc">
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
          ProductName:"000000000000",
          NodeType:0,
          DataFormat:1,
          Gateway:false,
          NetType:"WIFI",
          AuthType:"secret",
          NameSpace:"ICA",
          Desc:'',
          type:''
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