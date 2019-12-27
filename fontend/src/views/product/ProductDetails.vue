<template>
    <div name="product"> 
      <el-page-header @back="goBack" :content="productName" class="page-header">
      </el-page-header>
        <productInfoHead :product="ProductInf" ></productInfoHead>
      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="产品信息" name="third2">
            <ProductInfo :product="ProductInf"></ProductInfo>
        </el-tab-pane>
        <el-tab-pane label="Topic类列表" name="third3">
              <TopicList></TopicList>
        </el-tab-pane>   
        <el-tab-pane label="功能定义" name="third4">
           
        </el-tab-pane>
        <el-tab-pane label="服务端订阅" name="third5">
           
        </el-tab-pane>
        
        <el-tab-pane label="日志服务" name="third6">
           
        </el-tab-pane>
        <el-tab-pane label="在线调试" name="fourth">
            <CopyBtn></CopyBtn>
        </el-tab-pane>
      </el-tabs>
    </div>
  </template>

  <script>
  import productInfoHead from './productInfoHead'
  import ProductInfo from './ProductInfo'
  import TopicList from '@/views/topic/TopicList'
  import {mapState, mapMutations, mapGetters} from 'vuex'

    export default {
      components: { productInfoHead,ProductInfo,TopicList },
      data() {
        return {      
          productName:'',
          ProductInf:{},
          activeName:'third2'
        }
      },
      computed:{
       
      },
      created(){
        this.productName = this.$route.params.productName
        this.getProductInf() 
      },
      methods:{
          goBack(){
            this.$router.back(-1)
          },
         getProductInf(){
            this.ProductInf ={
              CategoryId: 140,
                Owner: true,
                ProductName: "4444444",
                AliyunCommodityCode: "iothub_senior",
                ProductStatus: "DEVELOPMENT_STATUS",
                ProductSecret: "JQry1WBsBya2ZxHe",
                GmtCreate: 1576726501000,
                CategoryKey: "GarbageOverflowingDetection",
                Id2: false,
                NodeType: 0,
                DataFormat: 1,
                CategoryName: "垃圾满溢检测",
                AuthType: "secret",
                NetType: 3,
                ProductKey: "a1ELejzj0h9",
                DeviceCount: 4}
         },
           
         
          handleClick(){

          },
          //查看产品密钥
          showProKeyInfo(){

          }
      }
    }
  </script>

  <style scoped>
    
  </style>