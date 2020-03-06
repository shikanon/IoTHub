<template>
    <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
     <el-form-item v-if="workSpaceId" label="工作空间" >        
         <el-input v-model="workSpaceName"  disabled></el-input>
    </el-form-item>
     <el-form-item v-else  label="工作空间" prop="ws">   
         <el-select v-model="ruleForm.ws" placeholder="请选择工作空间">
          <el-option
              v-for="item in workSpaceArr"
              :key="item.id"
              :label="item.name"
              :value="item.id">
          </el-option>
        </el-select> 
    </el-form-item>
    <el-form-item label="摄像头SIN码" prop="sin">
        <el-input v-model="ruleForm.sin"></el-input>
    </el-form-item>
   
    <!-- <el-form-item label="备注" prop="remark">
         <el-input v-model="ruleForm.remark"></el-input>
    </el-form-item> -->
    
    </el-form>
 
</template>
<script>
  export default {
    props:{
      workSpaceArr:{
          type:Array,
          default: () => []
      },
      workSpaceId:{
        type:Number,
        default:0
      }
    },
     watch:{
        //监听workSpaceId,若发生变化，重新查询设备列表
        workSpaceId:{  
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
      var checkDeviceName = (rule, value, callback) => {
        if (!value) {
          return callback(new Error('请输入设备名称'));
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
      var checkRemark = (rule, value, callback) => {
        if (!value) {
          return callback(new Error('请输入设备名称'));
        }
        setTimeout(() => {
          // 长度4~64；中文算2个；支持特殊字符“_”
          let pattern =/[^a-z|A-Z|0-9|\_)]/
          let length =  Number(value.replace(/[^\x00-\xff]/g,"01").length )
          if (!(length > 3 &&  length < 65 && !pattern.test(value))) {
            callback(new Error('支持中文、英文字母、数字和下划线，4~64个字符，中文算两个'))
          } else {     
            callback();      
          }
          
        }, 1000)
      };

      return {   
        ruleForm: { },
        workSpaceName:'',
        rules: {
          ws: [
            { required: true, message: '请选择工作空间', trigger: 'blur' },
          ],
         sin: [
            { required: true, message: '请输入摄像头SIN码', trigger: 'blur' },
            { validator: checkDeviceName, trigger: 'blur' }
          ],
          // remark: [
          //   { required: true, message: '请输入备注', trigger: 'blur' },
          //   { validator: checkRemark, trigger: 'blur' }

          // ],
              
        }
      };
    },
    created(){
       console.log(123213,this.workSpaceArr)
       this.init()

    },
    methods: {
        init(){
         
          if(this.workSpaceId){
             this.ruleForm.ws = this.workSpaceId
          
             this.workSpaceName = this.workSpaceArr.filter(item => item.id === this.workSpaceId)[0].name
             this.ruleForm =  {
              ws: this.workSpaceId,
              sin: '',
            }
          }else{
              this.ruleForm =  {
              ws: null,
              sin: '',
            }
          }
         
            
        },
        
      submitForm() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {
              this.$API_CLOUD.addCamera(this.ruleForm).then((res) => {
                if(res.data.status  === 'Y'){
                  let data  = res.data.data
                  //回调父组件方法关闭弹窗
                  this.$emit('close')
                  this.$message.success("添加摄像头成功");
              
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