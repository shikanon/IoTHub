<template>
    <div name="after-addProduct">
        <el-alert
            title="您已成功完成新建产品，接下来您可以："
            type="success"
            show-icon
            center>
        </el-alert>
        <div class="flex-center">
            <div>
                <img src="https://img.alicdn.com/tfs/TB1YtStnq61gK0jSZFlXXXDKFXa-72-72.svg">
            </div>
            <div>
                <p class="p-title">添加设备</p>
                <p class="p-desc">设备归属于某个产品下的具体设备。物联网平台为设备颁发产品内唯一的证书 DeviceName。设备可以直接连接物联网平台，也可以作为子设备通过网关连接物联网平台。</p>
                <el-button @click="gotoDeviceList" size="medium">前往添加</el-button>
            </div>
        </div>
         <el-divider></el-divider>
           <div class="flex-center">
               <div>
                    <img src="https://img.alicdn.com/tfs/TB14ZWtnxD1gK0jSZFKXXcJrVXa-72-72.svg">
                </div>
                <div>
                    <p class="p-title">为产品定义物模型</p>
                    <p class="p-desc">物联网平台支持为产品定义物模型，将实际产品抽象成由属性、服务、事件所组成的数据模型，便于云端管理和数据交互。产品创建完成后，您可以为它定义物模型，产品下的设备将自动继承物模型内容。</p>
                     <el-button size="medium" @click="gotoProductDltAbility">前往定义物模型</el-button>
                </div>
        </div>
        <el-divider></el-divider>
        <el-row class="align-cneter">
            <el-button type="primary" @click="gotoProductDlt">查看产品详情</el-button>
            <el-button @click="back">返回产品列表</el-button>   
        </el-row>
    </div>
</template>

<script>
export default {
    props:{
        productId:{
            type:Number,
            default:0
        }
    },
    methods:{
         //管理设备信息
        gotoDeviceList(){
           this.$router.push({name :'device-list',params: {productId:this.productId}}) 
        },

        //查看产品详情，打开功能定义
        gotoProductDltAbility(){
             this.$router.push({name :'product-list',params:{'productId':this.productId,'activeName':'ability'}}) 
        },

        //返回产品列表
        back(){
            this.$router.push({name :'product-list'})             
        },
        //查看产品详情
        gotoProductDlt(){
            this.$router.push({name :'product-list',params:{productId:this.productId}}) 
        },
       
    }
}
</script>
<style scoped>
    .flex-center{
        display: flex;
        align-items: center;
    }

    .flex-center img{
        width: 48px;
        height: 48px;
        margin: 0 40px;
    }
    .flex-column{
        display: flex;
        flex-direction: column;
    }

    .p-title{
        color: #111;
        line-height: 22px;
        margin-bottom: 8px;
    }

    .p-desc{
        color: #555;
        line-height: 20px;
        margin-bottom: 16px;
        font-size: 14px;
    }
    .align-cneter{
        text-align: center;
    }
    
</style>