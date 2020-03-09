<template>
    <div name="edit-product">
        <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
          <el-form-item label="产品名称" prop="name">
              <el-input v-model="ruleForm.name" :placeholder="ruleForm.name"></el-input>
          </el-form-item>      
          <el-form-item label="节点类型" prop="NodeType">
            <el-input v-model="ruleForm.node_type"  disabled></el-input>   
          </el-form-item>
           <el-form-item label="数据格式" prop="DataFormat">
                <el-input v-model="ruleForm.data_format" disabled></el-input>   
           </el-form-item>
           <el-form-item label="所属品类" prop="type">
               <el-input v-model="ruleForm.object_model_name" disabled></el-input>   
          </el-form-item>
          <el-form-item label="连网协议" prop="NetType">
               <el-input v-model="ruleForm.network_way" disabled></el-input>        
          </el-form-item>
          <el-form-item label="产品描述" prop="desc">
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
       var checkProductName = (rule, value, callback) => {    
        setTimeout(() => {
          // 支持中文、英文字母、数字、和特殊字符_-@()，长度限制4~30，中文算2位
          let pattern =/[^a-z|A-Z|\u4E00-\u9FA5|0-9|\-|_|\@|\(|\))]/
          let length =  Number(value.replace(/[^\x00-\xff]/g,"01").length )
          if (!(length > 3 &&  length < 33 && !pattern.test(value))) {
            callback(new Error('支持中文、英文字母、数字、和特殊字符_-@()，长度限制4~30，中文算2位'))
          } else {     
            callback()    
          }
        },1000)
      }
      return {
        ruleForm: {name:''},
        rules: {
          name: [
            { required: true, message: '请输入产品名称', trigger: 'blur' },
            { validator: checkProductName, trigger: 'blur' }
          ],
        },
           
      };
    },
    created(){
        this.init()
    },
    methods: {
      init(){   
          this.ruleForm = JSON.parse(JSON.stringify(this.product))
      },

      //编辑产品提交
      submitForm() {
        this.$refs['ruleForm'].validate((valid) => {
          if (valid) {
            let params ={pid:this.ruleForm.id,name:this.ruleForm.name,desc:this.ruleForm.desc}
            this.$API_IOT.updateProduct(params).then((res) => {
              if (res.data.status === 'Y') {
                  this.$message.success('更新产品成功')
                  this.$emit('success',params)
              } else {
                  this.$message.error(res.message);
              }       
            })
      
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
    .el-form-item{
      margin-bottom:0px
    }
</style>