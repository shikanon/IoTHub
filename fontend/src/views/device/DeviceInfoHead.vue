<template>
    <div name="device-info-head">
        <div class="title-inf">  
            <div >
                <span class="text-gray title">产品</span> 
                <span class="text-gray">{{device.ProductName}}</span>
                <CopyBtn :content="device.ProductName"></CopyBtn>        
            </div>
            <div >
                <span class="text-gray title">DeviceSecret</span> 
                <span class="text-gray">********</span>
                <el-button  type="text" size="small" @click="devKeyVisible = true">查看</el-button>   
            </div>  
            <div></div>
        </div>
        <div class="title-inf">
            <div>
                <span class="text-gray title">ProductKey</span> <span class="text-gray">{{device.ProductKey}}</span>
                <CopyBtn :content="device.ProductKey"></CopyBtn>    
            </div>    
        </div>
        <el-dialog title="设备证书" :visible.sync="devKeyVisible" width="26%">
            <DeviceKeyInfo ref="devKeyInfo" :device="device" @close="devKeyVisible = false"></DeviceKeyInfo>
            <span slot="footer" class="dialog-footer">
                <el-button @click="devKeyVisible = false">关 闭</el-button>
            </span>
        </el-dialog>
    </div>    
</template>
<script>
import DeviceKeyInfo from './DeviceKeyInfo'
export default {
    components:{DeviceKeyInfo},
    props:{
         device:{
            type:Object,
            default: () => ({}),
        }
    },
     data(){
         return {
             devKeyVisible:false
         }
     },
   
    methods:{
  
    }
}
</script>
<style  scoped>
   
</style>