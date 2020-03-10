<template>
    <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
    
    <el-form-item label="工作空间名称" prop="name">
        <el-input v-model="ruleForm.name"></el-input>
    </el-form-item>
     <el-form-item   label="推流类型" prop="access_type_id">   
         <el-select v-model="ruleForm.access_type_id" placeholder="请选择推流类型">
          <el-option
              v-for="item in typeArr"
              :key="item.id"
              :label="item.name"
              :value="item.id">
          </el-option>
        </el-select> 
    </el-form-item>  
    </el-form>
 
</template>
<script>
  export default {
    props:{
     
      workSpace:{
        type:Object,
        default:() => ({})
      }
    },
     watch:{
        //监听productId,若发生变化，重新查询设备列表
        workSpace:{  
            handler:function(val,oldval){   
                if (val && oldval && val.work_space_id != oldval.work_space_id ){
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
      var checkDeviceName = (rule, value, callback) => {
        if (!value) {
          return callback(new Error('请输入工作空间名称'));
        }
        setTimeout(() => {
          // 支持英文字母、数字和特殊字符-_@.:，长度限制4~32
          let pattern =/[^a-z|A-Z|0-9|\-|\_|\@|\.|\:)]/
          let length =  Number(value.replace(/[^\x00-\xff]/g,"01").length )
          if (!(length > 3 &&  length < 33 && !pattern.test(value))) {
            callback(new Error('支持英文字母、数字和特殊字符-_@.:，长度限制4~32'))
          } else {     
            callback();      
          }
          
        }, 1000)
      };
     

      return {   
        ruleForm: {work_space_id:0,name:'',access_type_id:1, },
        productName:'',
        typeArr:[{id:1,name:'RTMP'},{id:2,name:'HLS'},{id:3,name:'HTTP-FLV'}],
        rules: {
          access_type_id: [
            { required: true, message: '请选择推流类型', trigger: 'blur' },
          ],
         name: [
            { required: true, message: '请输入工作空间名称', trigger: 'blur' },
            { validator: checkDeviceName, trigger: 'blur' }
          ],          
        }
      };
    },
    created(){
       
       this.init()

    },
    methods: {
        init(){
          if(this.workSpace){
           this.ruleForm = {id:this.workSpace.work_space_id,name:this.workSpace.name,access_type_id:this.workSpace.access_type_id}
          }else{
             this.ruleForm = {name:'',access_type_id:1 }
          }                 
        },
        
      submitForm() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {

            if(this.workSpace){//修改
                 this.$API_CLOUD.updateWorkSpace(this.ruleForm).then((res) => {
                  if(res.data.status  === 'Y'){
                    this.$message.success('修改成功')    

                    //回调父组件方法关闭弹窗
                    this.$emit('close')  
                    //刷新列表
                     this.$emit('refresh')                  
                  }else{
                      this.$message.error(res.message);
                  }
                
              }) 
            }else{
                this.$API_CLOUD.addWorkSpace(this.ruleForm).then((res) => {
                  if(res.data.status  === 'Y'){
                     this.$message.success('添加成功')    
                    //回调父组件方法关闭弹窗
                    this.$emit('close')                   
                  }else{
                      this.$message.error(res.message);
                  }
                
              }) 
            }
                      
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