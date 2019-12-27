<template>
    <div name="product-key-info">    
        <p class="title">产品证书</p>
        <div class="table-info">
            <div class="table-row">
                <div class="table-row-label">ProductSecret</div> 
                <div class="table-row-info">
                    <span>{{product.ProductSecret}}</span>
                    <CopyBtn :content="product.ProductSecret"></CopyBtn>
                    <el-button type="text" @click="resetSecret" style="color:#f56c6c">重置</el-button>
                </div>
            </div>
            <div class="table-row">
                <div class="table-row-label">ProductKey</div>
                <div class="table-row-info">
                    <span>{{product.ProductKey}}</span>
                    <CopyBtn :content="product.ProductKey"></CopyBtn>
                </div>
            </div>    
        </div>
        <p>
            注：产品密钥 (ProductSecret)、设备密钥 (DeviceSecret) 为两种不同密钥。
             如您要获取的是设备密钥 (设备证书)，请前往对应的设备详情页面获取。
        </p>
        <p>烧录方式介绍</p>       
        <el-button v-if="!showMore" type = "text" size="medium"  @click="showMore = true"> 
            <i class="el-icon-arrow-down"></i>
            一机一密、一型一密介绍
         </el-button>
       <div v-if="showMore">
            <p>一型一密</p>
            <p>同一产品下所有设备可以烧录相同产品证书（即 ProductKey 和 ProductSecret ）。
            设备发送激活请求时，物联网平台进行产品身份确认，认证通过，下发该设备对应的DeviceSecret。</p>
            <p>一机一密</p>
            <p>每个设备烧录其唯一的设备证书（ProductKey、DeviceName 和 DeviceSecret）。
                当设备与物联网平台建立连接时，物联网平台对其携带的设备证书信息进行认证。
            如果您期望使用一机一密烧录方式，请前往对应的设备详情，来获取设备证书烧录，</p>
            <div class="desc-txt-blue">
                <span>如果您期望使用一机一密烧录方式，请前往对应的设备详情，来获取设备证书烧录，</span>
                <el-button  type = "text" size="medium" @click="gotoDeviceList">前往查看</el-button>
            </div>
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
            product:{
                type:Object,
                default:() =>({})
            }
        },
      data() {
        return {      
         showMore:false
        }
      },
    
      methods:{
          gotoDeviceList(){
               this.$router.push({name :'device-list',params: {product:this.product.ProductName}}) 
          },
          resetSecret(){
               this.$confirm('重置后将导致原有ProductSecret失效，您需要将新生成的ProductSecret烧录到设备中，是否确定要重置？', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
                }).then(() => {
                    // this.$message({
                    //     type: 'success',
                    //     message: '删除成功!'
                    // });
                }).catch(() => {
                    // this.$message({
                    //     type: 'info',
                    //     message: '已取消删除'
                    // });          
                });
          }
      }
    }
  </script>

  <style scoped>
    .title{
        margin-top:0px
    }
 
  </style>