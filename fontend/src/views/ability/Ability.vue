<template>
  <el-tabs v-model="activeName" type="card">
        <el-tab-pane label="标准功能" name="first">
            <AbilityList type="standard" :productKey="productKey"></AbilityList>
        </el-tab-pane>
        <el-tab-pane label="自定义功能" name="second">
             <AbilityList type="customize" :productKey="productKey"></AbilityList>
        </el-tab-pane>
   </el-tabs>
</template>

<script>
import AbilityList from './AbilityList'
  export default {
    components:{AbilityList},
    props:{
      productKey:{
        type:String,
        default:''
      }
    },
    data() {
      return {
        activeName: 'first'
      };
    },
    methods: {
      handleClick(tab, event) {
        console.log(tab, event);
      }
    }
  };
</script>