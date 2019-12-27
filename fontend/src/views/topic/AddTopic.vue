<template>
    <div name="add-topic"> 
        <p>
                Topic格式必须以“/”进行分层，区分每个类目。其中前三个类目已经规定好，
                第一个代表产品标识productKey，第二个${deviceName}通配deviceName，
                第三个user用来标识产品的自定义Topic类。简单来说，Topic类：/a15T****dhK/${deviceName}/user/update 
                是具体Topic：/a15T****dhK/mydevice1/user/update和/a15T****dhK/mydevice2/user/update等的集合。
        </p>
        <el-form label-position="top" ref="form" label-width="80px" :model="form"  :rules="rules">
            <el-form-item label="设备操作权限：">
                <el-select v-model="form.operation" placeholder="设备操作权限">
                    <el-option label="发布" value="0"></el-option>
                    <el-option label="订阅" value="1"></el-option>
                    <el-option label="发布和订阅" value="2"></el-option>
                </el-select>            
            </el-form-item>
            <el-form-item label="Topic类：">
                <span>123123213123213</span>
                <el-input v-model="form.topic" placeholder="请输入您的Topic类名"></el-input>
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
            id:{
                type:'number',
                default:0
            },
            operation: {
               type:'string',
                default:'0'
             },
            topic: {
                type:'string',
                default:''
            },
            desc: {
                type:'string',
                default:'0'
            },
            ProductKey:{
                type:'string',
                default:''
            }
                    
        }
    },
    data() {
      return {      
         rules: {      
          operation: [
            { required: true, message: '请选择设备操作权限', trigger: 'change' }
          ],
         topic: [
            { required: true, message: '请输入您的Topic类名', trigger: 'blur' },
            // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
        }    
      }
      
    },

    created(){
        console.log(this.form)
    },
    methods:{
        submit(){
            this.$refs['form'].validate((valid) => {
            if (valid) {
                alert('submit!');
            } else {
                console.log('error submit!!');
                return false;
            }
            });
        }
    }
  }
</script>