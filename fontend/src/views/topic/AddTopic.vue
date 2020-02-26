<template>
    <div name="add-topic"> 
        <p>
            Topic格式必须以“/”进行分层，区分每个类目。其中前三个类目已经规定好，
            第一个代表产品标识productKey，第二个${deviceName}通配deviceName，
            第三个user用来标识产品的自定义Topic类。简单来说，Topic类：/a15T****dhK/${deviceName}/user/update 
            是具体Topic：/a15T****dhK/mydevice1/user/update和/a15T****dhK/mydevice2/user/update等的集合。
        </p>
        <el-form label-position="top" ref="ruleForm" label-width="80px" :model="ruleForm"  :rules="rules">
            <el-form-item label="设备操作权限：" prop="operation">
                <el-select v-model="ruleForm.operation" placeholder="设备操作权限">
                    <el-option v-for="(item,key) in operationArr" :label="item.label" :value="item.value" :key="key"></el-option>                          
                </el-select>            
            </el-form-item>
            <el-form-item label="Topic类：" prop="name">
                <span>{{topic_pre}}</span>
                <el-input v-model="ruleForm.name" placeholder="请输入您的Topic类名"></el-input>
            </el-form-item>
            <el-form-item label="描述:">
                 <el-input type="textarea" v-model="ruleForm.desc" maxlength="100" show-word-limit></el-input>
            </el-form-item>
        </el-form>        
    </div>
</template>

<script>
  export default {
     props: {
        form:{
            type:Object,
            default:() => ({})                  
        }
    },
    
    watch:{
        form:{  
            handler:function(val,oldval){           
                this.$nextTick(()=>{
                    this.init()
                })            
            },  
            immediate:true,//关键
            deep:true
          },
    },
    data() {
        var checkTopicName = (rule, value, callback) => {
           
            setTimeout(() => {
                // Topic类名用/分割，支持英文字母、数字、下划线、+和#（仅权限是订阅时支持），长度限制64
                let pattern =/[^a-z|A-Z|0-9|\_|\+|\#)]/
                let length =  Number(value.replace(/[^\x00-\xff]/g,"01").length )
                if (!( length < 65 && !pattern.test(value))) {
                    callback(new Error('Topic类名用/分割，支持英文字母、数字、下划线、+和#（仅权限是订阅时支持），长度限制64'))
                } else {     
                    callback()   
                }
                
            }, 1000)
        };
      return {  
         topic_pre:"", 
         operationArr:[{label:'发布',value:1},
         {label:'订阅',value:2},
         {label:'发布和订阅',value:3}], 
         ruleForm:{},   
         rules: {      
          operation: [
            { required: true, message: '请选择设备操作权限', trigger: 'change' }
          ],
          name: [
           { required: true, message: '请输入Topic类名', trigger: 'blur' },
           { validator: checkTopicName, trigger: 'blur' }
          ],
        }    
      }     
    },

    created(){
       
       if(this.form.topic_name){
           this.init()
       }
    
    },
    methods:{
        init(){   
            let index = this.form.topic_name.indexOf("/user/") + 6
            let length = this.form.topic_name.length
            this.topic_pre = this.form.topic_name.substring(0,index)
            //传了pid，代表新增topic 
            if(this.form.pid){
                this.ruleForm ={pid:this.form.pid,operation:1,name:"",desc:""}
            }else{
                this.ruleForm.name  = this.form.topic_name.substring(index,length)            
                this.ruleForm = {tid:this.form.id,operation:this.form.operation,name:this.ruleForm.name,desc:this.form.desc}
            }         
        },
        submit(){
            this.$refs['ruleForm'].validate((valid) => {
                if (valid) {                   
                    if(this.ruleForm.pid){ //新增
                        this.$API_IOT.addTopic(this.ruleForm).then((res) => {
                            if(res.data.status === "Y"){
                                this.$message.success('新增topic成功')    
                                this.$emit('refresh')
                            }else{
                                this.$message.error(res.message);
                            }
                        }) 
                    }else{ //修改
                        this.$API_IOT.updateTopic(this.ruleForm).then((res) => {
                            if(res.data.status === "Y"){
                                this.$message.success('修改topic成功')    
                                this.$emit('refresh')
                            }else{
                                this.$message.error(res.message);
                            }
                        }) 
                    }
                    
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        }
    }
  }
</script>