<template>
    <div name="product-inf-head">
        <div class="title-inf" >  
            <div>
                <span class="text-gray title">ProductKey</span> 
                <span class="text-gray">{{product.product_key}}</span>
                <CopyBtn :content="product.product_key" size="small"></CopyBtn>        
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
                <span class="text-gray title">设备数</span> <span class="text-gray">{{product.device_count}}</span>
                <el-button  type="text" size="small" @click="gotoDeviceList">前往管理</el-button>
            </div>
        </div>
          <el-dialog title="产品证书" :visible.sync="proKeyVisible" width="30%">
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
        //跳转到设备管理列表
        gotoDeviceList(){
            this.$router.push({name :'device-list',params: {productId:this.product.id}}) 

        },
    }
}
</script>
<style  scoped>
    
</style>