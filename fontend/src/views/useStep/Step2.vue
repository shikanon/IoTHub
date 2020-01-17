<template>
    <div name="step2">
        <p class="step-desc">选择最有利于您连接到阿里云物联网平台的设备平台和开发工具包</p>
        <el-form label-position="top"  ref="form" label-width="100px" >
            <el-form-item label="选择设备平台：" >
                <el-radio-group v-model="radio1" @change="changeRadio1">
                    <el-radio label="1" border >Linux</el-radio>
                    <el-radio label="2" border :disabled="radio3 === '3'">OSX</el-radio>
                    <el-radio label="3" border :disabled="radio3 === '3'">Windows</el-radio>
                    <el-radio label="4" border :disabled="radio3 === '3'">Android</el-radio>
                </el-radio-group>
                
            </el-form-item>
            <el-form-item label="选择设备连接协议：" >
                <el-radio v-model="radio2" label="1" border>MQTT</el-radio>
            </el-form-item>
                <el-form-item label="选择设备开发工具包："  >
                    <el-radio-group v-model="radio3" @change="changeRadio3">
                        <el-radio label="1" border :disabled="radio1 === '4'">Node.js</el-radio>
                        <el-radio label="2" border :disabled="radio1 === '4'">Java</el-radio>
                        <el-radio label="3" border :disabled="radio1 !== '1' || radio1 === '4' ">嵌入式C</el-radio>
                        <el-radio label="4" border :disabled="radio1 !== '4'">Android</el-radio>
                    </el-radio-group>              
            </el-form-item>
        </el-form> 
        <div class="tips">
            <img  class="tips-image" src="https://img.alicdn.com/tfs/TB1LhWDmzTpK1RjSZKPXXa3UpXa-66-66.png">
            <div>
                <div class="tips-title">
                {{ radio3  === '1' ? 'Node.js' : radio3  === '2' ? 'Java':  radio3  === '3' ? '嵌入式C' :'Android' }} 工具包，开发使用条件：
                </div>
                <div v-if="radio3 === '1'" class="tips-content">
                    请确保已安装Node.js LTS，并且版本不低于8.10
                </div>
                <div v-else-if="radio3 === '2'" class="tips-content">
                    请确保已安装下面的软件，并且版本不低于指定的版本：Ubuntu 16.04、JDK 1.8.0.144
                </div>
                <div v-else-if="radio3 === '3'" class="tips-content">
                    请确保已安装以下软件，并且版本不低于指定的版本：
                    Ubuntu 16.04 LTS、make-4.1、git-2.7.4、gcc-5.4.0、gcov-5.4.0、lcov-1.12、bash-4.3.48、tar-1.28、cmake3.51、awk4.1.3、sed 4.2.2
                </div>
                <div v-else class="tips-content">
                    如果您是Windows开发环境，请确保已安装下面的软件，并且版本不低于指定的版本：JDK 1.8.0.144、gradle 4.1、Android SDK version 27；
                    如果您是Linux/OSX开发环境，请确保已安装下面的软件，并且版本不低于指定的版本：Ubuntu 16.04 / OSX 10.13.6、JDK 1.8.0.144、gradle 4.1、
                    Android SDK version 27
                </div>
             </div>
            
        </div>
    </div>
</template>

<script>


export default {
   
    data(){
        return{
            radio1:'1',
            radio2:'1',
            radio3:'1',
        }
    },
  
    methods:{
        changeRadio1(value){
            if(value === '4'){
                this.radio3 = '4'
            }else{
                this.radio3 = '1'
            }
            
        },
        changeRadio3(event,value2){
            console.log(value2)
        },
         nextStep() {
               alert('submit')
               this.$emit('next')
            // this.$refs.form.validate((valid) => {
            //     if (valid) {
            //         alert('submit')
            //         this.$emit('next')
            //     } else {
            //         console.log('error submit!!');
            //         return false;
            //     }
            // });
        },
    }
}
</script>

<style  scoped>
    .tips{
        display: flex;
        align-items: center;
        background: #e3f2fd;
        border: 1px solid #a0cbed;
        border-radius: 3px;
        padding: 10px 16px;
        }
    .tips-image{
        width: 24px;
        height: 24px;
        margin-right: 20px;
    }
    .tips-title{
        color: #666;
        font-size: 14px;
        font-weight: bolder;
    }
    .tips-content{
        margin-top: 4px;
            color: #666;
            font-size: 14px;
    }
</style>