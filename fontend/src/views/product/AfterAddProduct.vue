<template>
    <div name="after-addProduct">
        <p>您已成功完成新建产品，接下来您可以：</p>
        <div>
            <img>
            <p>添加设备</p>
            <p>设备归属于某个产品下的具体设备。物联网平台为设备颁发产品内唯一的证书 DeviceName。设备可以直接连接物联网平台，也可以作为子设备通过网关连接物联网平台。</p>
            <el-button @click="gotoDeviceList">前往添加</el-button>
        </div>
         <el-divider></el-divider>
           <div>
            <img>
            <p>为产品定义物模型</p>
            <p>物联网平台支持为产品定义物模型，将实际产品抽象成由属性、服务、事件所组成的数据模型，便于云端管理和数据交互。产品创建完成后，您可以为它定义物模型，产品下的设备将自动继承物模型内容。</p>
            <el-button>前往定义物模型</el-button>
        </div>
        <el-divider></el-divider>
        <el-row>
            <el-button type="primary" @click="gotoProductDlt">查看产品详情</el-button>
            <el-button @click="back">返回产品列表</el-button>   
        </el-row>
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
    methods:{
        //返回产品列表
        back(){
            this.$router.push({name :'product-list'})             
        },
        //查看产品详情
        gotoProductDlt(){
            this.$router.push({name :'product-list',params: {productName:this.productName}}) 
        },
        //管理设备信息
        gotoDeviceList(){
           this.$router.push({name :'device-list',params: {productName:this.productName}}) 
        }
    }
}
</script>
<style scoped>

</style>