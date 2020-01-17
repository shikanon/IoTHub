<template>
    <div name="publish-product">
        <div class="desc-txt-gray">
              
            <p>您即将发布的产品为：<el-link type="primary" class="product-name"  :underline="false">{{productName}}</el-link></p>
            <div>产品发布后将从开发阶段进入到正式投产或使用阶段。</div>
        </div>
        <p>请勾选并确认该产品的信息和各项功能已具备发布条件：</p>
        <div>
            <div class="step-box">
                <div class="left">
                    <div class="step-title">第一步</div> 
                    <el-divider direction="vertical"></el-divider>
                    <div class="step-desc">  
                        请确认产品的各项基本信息准确无误，产品发布后将无法再做修改和删除
                    </div>
                </div>
                <div><el-checkbox v-model="step1">{{step1 === false ? '请确认':'已确认'}}</el-checkbox></div>
            </div>
             <div  class="step-box">
                <div class="left">
                    <div class="step-title">第二步 </div><el-divider direction="vertical"></el-divider>
                    <div class="step-desc">  
                       请确认设备的各项功能已经完成开发和调试，发布后产品的功能改动请通过OTA升级。
                    </div>
                </div>
                <div><el-checkbox v-model="step2">{{step2 === false ? '请确认':'已确认'}}</el-checkbox></div>
            </div>
             <div  class="step-box">
                <div class="left">
                    <div class="step-title">第三步</div><el-divider direction="vertical"></el-divider>
                    <div class="step-desc">
                         
                        请确认产品已经具备上线发布条件，开始进入规模化接入和部署。
                    </div>
                </div>
                <div><el-checkbox v-model="step3">{{step3 === false ? '请确认':'已确认'}}</el-checkbox></div>
            </div>
             <div  class="step-box">
                <div class="left">
                    <div class="step-title">可选</div><el-divider direction="vertical"></el-divider>
                    <div class="step-desc">
                        您可以选择将现在的产品模型，升级为新的行业模型标准。
                    </div>
                 </div>
                <div><el-checkbox v-model="step4">推荐</el-checkbox></div>
            </div>
        </div>
    </div>
</template>


<script>
export default {
    props:{
        productName:{
            type:String,
             default:''
        }
    },
    data(){
        return{
            step1:false,
            step2:false,
            step3:false,
            step4:false,
        }
    },
    computed:{
        //确定按钮
        publishBtn:function(){        
            return this.step1 && this.step2  && this.step3  
        }
    },
    watch:{
        publishBtn:function(){
            this.$emit('setBtnStatus',!this.publishBtn)
        }    
    }
}
</script>

<style scoped>

    .product-name{
        font-size: 20px;
        color: #0070cc;
    }
    .step-box{
        display: flex;
        border: 1px solid #ccc;
        align-items: center;
        height: 80px;
        border-bottom:none; 
        justify-content: space-around;
    }
    

    .left{
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .step-box:last-child{
        border-bottom: 1px solid #ccc;
    }


    .step-title{
        width: 60px;
        display: flex;
        align-items: center;
    }
    .step-desc{
        font-size: 14px;
        color: #666;
        width: 410px;
        margin-left: 20px;
    }
</style>