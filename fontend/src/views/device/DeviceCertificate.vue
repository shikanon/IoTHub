<template>
    <div>
        <div class="tips">设备证书用于云端对接入的设备做鉴权认证，请妥善保管！</div>
        <div class="table-info">
            <div class="table-row">
                <div class="table-row-label">ProductKey</div> 
                <div class="table-row-info">
                    <span>{{certificateInf.productKey}}</span>
                    <CopyBtn :content="certificateInf.productKey"></CopyBtn>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">DeviceName</div>
                <div class="table-row-info">
                    <span>{{certificateInf.deviceName}}</span>
                    <CopyBtn :content="certificateInf.deviceName"></CopyBtn>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">DeviceSecret</div>
                <div v-if="!deviceSecretFlg" class="table-row-info">
                    <span>********</span>
                    <el-button  type="text" size="medium" @click="toggleFlg">显示</el-button>
                </div>
                <div v-else class="table-row-info">
                    <el-tooltip :content="certificateInf.deviceSecret" placement="bottom" effect="light">
                    <span>{{certificateInf.deviceSecret.substr(0,20)}}</span>
                    </el-tooltip>                  
                    <CopyBtn :content="certificateInf.deviceSecret"></CopyBtn>
                    <el-button  type="text" size="medium" @click="toggleFlg">隐藏</el-button>
                </div>
            </div>
        </div>
    </div> 
</template>
<script>
  export default {
    props:{
        certificateInf:{
            type:Object,
            default: () => ({})
        }
    },
    data() {
      return {
        deviceSecretFlg:false
      }
    },
    created(){

    },
    methods: {
        toggleFlg(){
            this.deviceSecretFlg = ! this.deviceSecretFlg
        }
    }
  }
</script>
<style scoped >
 .tips{
        background: #e5f3ff;
        height: 30px;
        line-height: 30px;
        margin-bottom: 10px;
        padding-left: 10px;
        color:#555
    }
  
  
   
</style>