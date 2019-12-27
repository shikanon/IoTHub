<template>
    <div name="product-inf-head">
        <div class="title-inf" >  
            <div>
                <span class="text-gray title">ProductKey</span> 
                <span class="text-gray">{{product.ProductKey}}</span>
                <CopyBtn :content="product.ProductKey" size="small"></CopyBtn>        
            </div> 
            <div>
                <span class="text-gray title">ProductSecret</span> 
                <span class="text-gray">********</span>
                <el-button  type="text" size="small" @click="proKeyVisible = true">查看</el-button>   
            </div>  
            <div>
            </div>
        </div>
        <div  class="title-inf" >
            <div>
                <span class="text-gray title">设备数</span> <span class="text-gray">{{product.DeviceCount}}</span>
                <el-button  type="text" size="small" @click="gotoDeviceList">前往管理</el-button>
            </div>
        </div>
          <el-dialog title="产品证书" :visible.sync="proKeyVisible" width="26%">
            <ProductKeyInfo ref="proKeyInfo" :product="product" @close="proKeyVisible = false"></ProductKeyInfo>
            <span slot="footer" class="dialog-footer">
                <el-button @click="proKeyVisible = false">关 闭</el-button>
            </span>
        </el-dialog>
    </div>    
</template>
<script>
import ProductKeyInfo from './ProductKeyInfo'
export default {
    components:{ProductKeyInfo},
    props:{
        product:{
            type:Object,
            default: () => ({}),
        },
      
    },
    data(){
        return {
            proKeyVisible:false
        }
    },

    methods:{
        showProKeyInfo(){

        },
      
         //跳转到设备列表
        gotoDeviceList(){
            this.$router.push({name :'device-list',params: {product:this.product.ProductName}}) 
        },
    }
}
</script>
<style  scoped>
    
</style>