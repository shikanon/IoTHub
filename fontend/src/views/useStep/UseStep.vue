<template>
    <div name="use-step">
        <div v-if="step === 0">
            <StepStart></StepStart>
        </div>
         <div v-else>
            <el-steps :active=" step - 1" align-center finish-status="success">
                <el-step title="注册设备" ></el-step>
                <el-step title="选择开发包" ></el-step>
                <el-step title="SDK下载" ></el-step>
                <el-step title="配置与测试" ></el-step>
            </el-steps>
            <div v-if="step === 1">
                <Step1 @next="next" ref="step1"></Step1>
            </div>
            <div v-if="step === 2">
                 <Step2 @next="next" ref="step2"></Step2>
            </div>
            <div v-if="step === 3">
                 <Step3 @next="next" @load="load" ref="step3"></Step3>
            </div>
            <div v-if="step === 4">
                 <Step4 @next="next" ></Step4>
            </div>
         </div>
    </div>   
</template>

<script>
import StepStart  from './StepStart'
import Step1  from './Step1'
import Step2  from './Step2'
import Step3  from './Step3'
import Step4  from './Step4'

export default {
    components:{StepStart,Step1,Step2,Step3,Step4},
    props:{
        step:{
            type:Number,
            default:0
        }
    },
    data(){
        return{
           
        }
    },
    
    methods:{
       next(){
           this.$emit('addStep',)
       },
       load(){
            this.$emit('load')
       }
    }
}
</script>

<style scoped>
.el-steps{
    background: #f8f8f8;
    padding: 28px 0;
    margin-bottom: 20px
}

</style>