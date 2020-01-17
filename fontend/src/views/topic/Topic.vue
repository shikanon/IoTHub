<template>
  <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="基础通信 Topic" name="first">
            <TopicList :tableData="baseData" type="base"></TopicList>
        </el-tab-pane>
        <el-tab-pane label="物模型通信 Topic" name="second">
             <TopicList :tableData="physicsData" type="physics" lazy ></TopicList>
        </el-tab-pane>
        <el-tab-pane label="自定义 Topic" name="third">
             <TopicList :tableData="Data"  :type="type" lazy ></TopicList>
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
      },
      queryKey:{
        type:Number,
        default:0
      }
    },
    components:{TopicList},
    data() {
      return {
        activeName: 'first',
        index:'0',
        baseData:[],
        physicsData:[],
        Data:[]
      };
    },
    created(){
      this.getTopicList()
    },
    methods: {
      handleClick(tab, event) {
      
                
      },

       getTopicList(){
              this.$API_IOT.getTopicList(this.type,this.queryKey).then((res) => {
                this.baseData  = res.data.data['base']
                this.physicsData = res.data.data['physics']
                this.Data  =  res.data.data['custom']   
            })
            
      }
    }
  }
</script>