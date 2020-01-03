<template>
  <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="基础通信 Topic" name="first">
            <TopicList type="base"></TopicList>
        </el-tab-pane>
        <el-tab-pane label="物模型通信 Topic" name="second">
             <TopicList type="physics" lazy ></TopicList>
        </el-tab-pane>
        <el-tab-pane label="自定义 Topic" name="third">
             <TopicList :type="type" lazy ></TopicList>
        </el-tab-pane>
   </el-tabs>
</template>

<script>
import TopicList from './TopicList'
  export default {
    props:{
      type:{
        type:String,
        default:''
      }
    },
    components:{TopicList},
    data() {
      return {
        activeName: 'first',
        index:'0'
      };
    },
    methods: {
      handleClick(tab, event) {
        // if(tab.index  === '2'){
        //   this.type = type
        // }
                
      }
    }
  };
</script>