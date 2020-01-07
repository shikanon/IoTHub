<template>
   <div name="base">
        <div v-if="!openRouter">
                <el-tabs type="border-card" v-model="activeName" @tab-click="handleClick" >
                        <el-tab-pane label="搜索" name="serach">
                                <Search @click="clickTab($event)"></Search>
                        </el-tab-pane>
                        <el-tab-pane label="知识" name="knowledge" lazy>
                                <Knowledge :index="index" @click="gotoApiInfo($event)"></Knowledge>        
                        </el-tab-pane>
                        <el-tab-pane label="资源" name="resource" lazy>
                                <Resource :index="index" @click="gotoApiInfo($event)"></Resource>
                        </el-tab-pane>
                        <el-tab-pane label="人物" name="people"  lazy>
                                <People :index="index"></People>
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
                        index:'0',
                        openRouter:false
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
                   this.index = event.index.toString()
                },
                gotoApiInfo(groupId){
                     this.$router.push({name:'api-info',params: {groupId:groupId}})      
                }
        }
  
}
</script>

<style scoped>
     
</style>