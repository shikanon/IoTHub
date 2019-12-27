<template>
    <div name="product-key-info"> 
        <p class="title">设备证书<CopyBtn btnTxt="一键复制" :content="JSON.stringify(copyTxt)"></CopyBtn></p>
        <div class="table-info">
            <div class="table-row">
                <div class="table-row-label">ProductKey</div> 
                <div class="table-row-info">
                    <span>{{device.ProductKey}}</span>
                    <CopyBtn :content="device.ProductKey"></CopyBtn>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">DeviceName</div>
                <div class="table-row-info">
                    <span>{{device.DeviceName}}</span>
                    <CopyBtn :content="device.DeviceName"></CopyBtn>
                </div>
            </div>    
            <div class="table-row">
                <div class="table-row-label">DeviceSecret</div>
                <div class="table-row-info">
                    <span>{{device.DeviceSecret}}</span>
                    <CopyBtn :content="device.DeviceSecret"></CopyBtn>
                </div>
            </div> 
        </div>
  
        <p>烧录方式介绍</p>       
        <el-button v-if="!showMore" type = "text" size="medium"  @click="showMore = true"> 
            <i class="el-icon-arrow-down"></i>
            一机一密、一型一密介绍
         </el-button>
       <div v-if="showMore">
            <p>一型一密</p>
            <p>同一产品下所有设备可以烧录相同产品证书（即 ProductKey 和 ProductSecret ）。
                设备发送激活请求时，物联网平台进行产品身份确认，认证通过，下发该设备对应的DeviceSecret。</p>
             <div class="desc-txt-blue">
                <span>如果您期望使用一型一密烧录方式，请前往对应的产品详情，来获取产品证书烧录，</span>
                <el-button  type = "text" size="medium" @click="gotoProductDtl">前往查看</el-button>
            </div>
            <p>一机一密</p>
            <p>每个设备烧录其唯一的设备证书（ProductKey、DeviceName 和 DeviceSecret）。
                当设备与物联网平台建立连接时，物联网平台对其携带的设备证书信息进行认证。</p>
           
        </div>
         <el-button v-if="showMore" type = "text" size="medium"  @click="showMore = false"> 
            <i class="el-icon-arrow-up"></i>
            收起
         </el-button>
    </div>
  </template>

  <script>

    export default {
        props:{
            device:{
                type:Object,
                default:() =>({})
            }
        },
      data() {
        return {     
         copyTxt:{
            ProductKey: '',
            DeviceName: '',
            DeviceSecret: ''
         },
         showMore:false
        }
      },
      created(){        
         this.copyTxt.ProductKey =  this.device.ProductKey
         this.copyTxt.DeviceName =  this.device.DeviceName
         this.copyTxt.DeviceSecret =  this.device.DeviceSecret  
     },
      methods:{
          gotoProductDtl(){
               this.$router.push({name :'product-list',params: {productName:this.device.ProductName}}) 
          }
      }
    }
  </script>

  <style scoped>
 
  .title{
        margin-top:0px
    }
  </style>