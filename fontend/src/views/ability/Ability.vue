<template>
  <el-tabs v-model="activeName" type="card">
        <el-tab-pane label="标准功能" name="first">
            <AbilityList type="standard" :productKey="productKey" :productId="productId"></AbilityList>
        </el-tab-pane>
        <el-tab-pane label="自定义功能" name="second">
             <AbilityList type="customize" :productKey="productKey" :productId="productId"></AbilityList>
        </el-tab-pane>
   </el-tabs>
</template>

<script>
import AbilityList from './AbilityList'
  export default {
    components:{AbilityList},
    props:{
      productInf:{
        type:Object,
        default:() => ({})
      }
    },
    data() {
      return {
        activeName: 'first',
        productKey:'',
        productId:0
      };
    },

    created(){
      console.log(this.productInf)
        this.productKey = this.productInf.product_key
        this.productId = this.productInf.id
    },
    methods: {
      
      handleClick(tab, event) {
        console.log(tab, event);
      }
    }
  };
</script>