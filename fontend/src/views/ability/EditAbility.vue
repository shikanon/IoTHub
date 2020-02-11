<template>
    <div name="edit-ability">
        <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
          <el-form-item label="功能属性" prop="type">
              <el-input v-model="ruleForm.type"  :disabled="ruleForm.accessMode === 'r'"></el-input>
          </el-form-item>      
          <el-form-item label="功能名称" prop="name">
            <el-input v-model="ruleForm.name"  :disabled="ruleForm.accessMode === 'r'"></el-input>   
          </el-form-item>
           <el-form-item label="标识符" prop="identifier">
                <el-input v-model="ruleForm.identifier" :disabled="ruleForm.accessMode === 'r'"></el-input>   
           </el-form-item>
           <el-form-item label="数据类型" prop="data_type">
               <el-input v-model="ruleForm.data_type" :disabled="ruleForm.accessMode === 'r'"></el-input>   
          </el-form-item>  
           <el-form-item label="json对象" prop="data_condition" v-if="ruleForm.data_type === 'struct' ">
               <div v-for="(item,index) in ruleForm.data_condition" :key="index" class="condition">
                   <div>参数名称：{{item.name}}</div>
                   <div><el-button type="text" size="small">查看</el-button></div>
               </div>    
          </el-form-item>
           <el-form-item label="布尔值" prop="data_condition" v-else-if="ruleForm.data_type === 'bool'">
               <div class="flex-row" v-for="(value,key,index) in ruleForm.data_condition" :key="index">
                   <span>{{key}} - {{value}}</span>
               </div>   
          </el-form-item>
            <template v-else>
            <el-form-item label="取值范围"  >
                <div class="flex-row">
                    <el-input v-model="ruleForm.data_condition.min" :disabled="ruleForm.accessMode === 'r'"></el-input>  
                    ~  
                    <el-input  v-model="ruleForm.data_condition.max" :disabled="ruleForm.accessMode === 'r'"></el-input>  
                </div>       
            </el-form-item>
            <el-form-item label="步长" >
                <el-input v-model="ruleForm.data_condition.step" :disabled="ruleForm.accessMode === 'r'"></el-input>        
            </el-form-item>
            <el-form-item label="单位"  >
                <el-input v-model="ruleForm.data_condition.unit" :disabled="ruleForm.accessMode === 'r'"></el-input>        
            </el-form-item>
            </template> 
            <el-form-item label="读写类型" prop="accessMode">
                <el-radio v-model="ruleForm.accessMode" label="rw" disabled>读写</el-radio>
                <el-radio v-model="ruleForm.accessMode" label="r" disabled>只读</el-radio>
          </el-form-item>
          <el-form-item label="描述" prop="desc">
               <el-input type="textarea" v-model="ruleForm.desc" :disabled="ruleForm.accessMode === 'r'"></el-input>
          </el-form-item>     
      
        </el-form>
    </div>
</template>
<script>
  export default {
      props:{
          ability:{
              tyoe:Object,
              default:() => ({})
          }
      },watch:{
        //监听type,若发生变化，重新查询设备批次
        ability:{  
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
        ruleForm: {name:''},
        rules: {
          name: [
            { required: true, message: '请输入产品名称', trigger: 'blur' },
          ],
        },
           
      };
    },
    created(){
        this.init()
    },
    methods: {
      init(){   
          this.ruleForm = JSON.parse(JSON.stringify(this.ability))
          console.log(this.ruleForm)
      },

      //编辑产品提交s
      submitForm() {
        this.$refs['ruleForm'].validate((valid) => {
          if (valid) {
            
            let params ={pid:this.ruleForm.id,name:this.ruleForm.name,desc:this.ruleForm.desc}
            this.$API_IOT.updateProduct(params).then((res) => {
              if (res.data.status === 'Y') {
                  this.$message.success('更新产品成功')
                  this.$emit('success')
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
  
    .el-form-item{
      margin-bottom:0px
    }
    .el-dialog__body{
        padding-top:10px !important;
    }
    .el-form-item__label{
        padding-bottom: 0px  !important;
    }

    .condition{
        display: flex;
        padding: 0 8px;
        justify-content: space-between;
        background: #e3f2fd;
        margin: 8px 0;
    }

    .flex-row{
        display: flex;
        flex-direction: row;
    }
</style>