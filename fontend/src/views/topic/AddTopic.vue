<template>
    <div name="add-topic"> 
        <p>
            Topic格式必须以“/”进行分层，区分每个类目。其中前三个类目已经规定好，
            第一个代表产品标识productKey，第二个${deviceName}通配deviceName，
            第三个user用来标识产品的自定义Topic类。简单来说，Topic类：/a15T****dhK/${deviceName}/user/update 
            是具体Topic：/a15T****dhK/mydevice1/user/update和/a15T****dhK/mydevice2/user/update等的集合。
        </p>
        <el-form label-position="top" ref="form" label-width="80px" :model="form"  :rules="rules">
            <el-form-item label="设备操作权限：" prop="operation">
                <el-select v-model="form.operation" placeholder="设备操作权限">
                    <el-option v-for="(item,key) in operationArr" :label="item.label" :value="item.value" :key="key"></el-option>                          
                </el-select>            
            </el-form-item>
            <el-form-item label="Topic类：" prop="name">
                <span>{{topic_pre}}</span>
                <el-input v-model="name" placeholder="请输入您的Topic类名"></el-input>
            </el-form-item>
            <el-form-item label="描述:">
                 <el-input type="textarea" v-model="form.desc"></el-input>
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
                if(val.id &&  val.id != oldval.id){
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
         name:"",   
         topic_pre:"", 
         operationArr:[{label:'发布',value:1},
         {label:'订阅',value:2},
         {label:'发布和订阅',value:3}],
     
         rules: {      
          operation: [
            { required: true, message: '请选择设备操作权限', trigger: 'change' }
          ],
          name: [
            { required: true, message: '请输入您的Topic类名', trigger: 'blur' },
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
            if(this.form.topic_name){
                let index = this.form.topic_name.indexOf("/user/") + 6;
                let length = this.form.topic_name.length;
                this.topic_pre = this.form.topic_name.substring(0,index);
                this.name  = this.form.topic_name.substring(index,length);
            }
            
        },
        submit(){
            this.$refs['form'].validate((valid) => {
            if (valid) {
                this.form.topic_name = `${this.topic_pre}${this.name}`
                this.form.topic_name = `${this.topic_pre}${this.name}`
            } else {
                console.log('error submit!!');
                return false;
            }
            });
        }
    }
  }
</script>