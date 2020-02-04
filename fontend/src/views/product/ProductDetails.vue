<template>
    <div name="product"> 
      <div class="page-header">
        <el-page-header @back="goBack" :content="ProductInf.name" class="page-header">
        </el-page-header>
        <el-button type="primary" size="medium" @click="dialogVisible = true" >发布</el-button>
      </div>
      <productInfoHead :product="ProductInf" ></productInfoHead>
      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="产品信息" name="product">
            <ProductInfo :product="ProductInf"></ProductInfo>
        </el-tab-pane>
        <el-tab-pane label="Topic类列表" name="topic" lazy>
              <Topic type="product" :queryKey="productId"></Topic>
        </el-tab-pane>   
        <el-tab-pane label="功能定义" name="ability" lazy>
            <Ability :productInf="ProductInf"></Ability>
        </el-tab-pane>
        <el-tab-pane label="服务端订阅" name="fourth" lazy>
           
        </el-tab-pane>
        
        <!-- <el-tab-pane label="日志服务" name="third6">
           
        </el-tab-pane>
        <el-tab-pane label="在线调试" name="fourth">
            <CopyBtn></CopyBtn>
        </el-tab-pane> -->
      </el-tabs>

       <el-dialog title="确认发布产品" :visible.sync="dialogVisible" width="35%">
           <PublishProduct ref="pulish" :productName="ProductInf.name" @setBtnStatus="setBtnStatus($event)" ></PublishProduct>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="publishSubmit" :disabled="btnDisabled">确 定</el-button>
                <el-button @click="dialogVisible = false">取 消</el-button>
            </span>
        </el-dialog>
      
    </div>
  </template>

  <script>
  import productInfoHead from './productInfoHead'
  import ProductInfo from './ProductInfo'
  import Ability from '@/views/ability/Ability'
  import PublishProduct from './PublishProduct'

  import {mapState, mapMutations, mapGetters} from 'vuex'

    export default {
      components: { productInfoHead,ProductInfo,Ability,PublishProduct },
    
      data() {
        return {      
          productId:'',
          activeName:'product',
          ProductInf:{},
          dialogVisible:false,
          btnDisabled:true
        }
      },
      computed:{
       
      },
      created(){
        this.productId = this.$route.params.productId
        this.activeName = this.$route.params.activeName
        this.getProductInf() 
        //  this.btnDisabled = !this.$refs.pulish.publishBtn
      },
      methods:{
          goBack(){
            this.$router.back(-1)
          },
         getProductInf(){
              this.$API_IOT.getProductById(this.productId).then((res) => {
                this.ProductInf = res.data.data
            })
            
         },
           
         
          handleClick(){

          },

          setBtnStatus(event){
              this.btnDisabled = event
          },
          //发布产品
          publishSubmit(){

          }
      }
    }
  </script>

  <style scoped>
    
  </style>