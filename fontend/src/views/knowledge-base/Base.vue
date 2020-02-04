<template>
   <div name="base">
        <div v-if="!openRouter">
                <el-tabs type="border-card" v-model="activeName" @tab-click="handleClick" >
                        <el-tab-pane label="搜索" name="serach">
                                <Search @click="clickTab($event)"></Search>
                        </el-tab-pane>
                        <el-tab-pane label="官方模型" name="knowledge" lazy>
                                <Knowledge :index="index" :arrData ="arrData" @click="gotoApiInfo($event)"></Knowledge>        
                        </el-tab-pane>
                        <el-tab-pane label="行业解决方案" name="resource" lazy>
                                <Knowledge :index="index" :arrData ="arrData" @click="gotoApiInfo($event)"></Knowledge>        
                                <!-- <Resource :index="index" @click="gotoApiInfo($event)"></Resource> -->
                        </el-tab-pane>
                        <el-tab-pane label="数据集" name="people"  lazy>
                                <!-- <People :index="index"></People> -->
                                <Knowledge :index="index" :arrData ="arrData" @click="gotoApiInfo($event)"></Knowledge>        
                        </el-tab-pane>
                </el-tabs>
        </div>
        <router-view/> 
  </div>
</template>

<script>
 import Search from './Search'
 import Knowledge from './Knowledge'
 import Resource from './Resource'
 import People from './People'

export default {
        components:{Search,Knowledge,Resource,People},
        data() {
                return {
                        activeName: 'serach',
                        index:0,
                        openRouter:false,
                        arrData:[]
                }
        },
        watch: {
                $route: function(){
                        this.openRouter = !this.openRouter  
                }
              
        },
        methods: {
                handleClick(tab, event) {
                       // console.log(tab, event)
                },
                clickTab(event){
                   this.activeName = event.name 
                   this.index = event.index
                   this.arrData = event.arrData
                },
                // gotoApiInfo(groupId){
                //      this.$router.push({name:'api-info',params: {APIGroup:groupId})      
                // }
                gotoApiInfo(APIGroup){
                     this.$router.push({name:'api-info',params: {APIGroup:APIGroup}})      
                }
        }
  
}
</script>

<style scoped>
     
</style>