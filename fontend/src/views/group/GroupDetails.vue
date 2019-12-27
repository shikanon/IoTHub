<template>
    <div name="group"> 
      <el-page-header @back="goBack" :content="groupName" class="page-header">
      </el-page-header>
        <GroupInfoHead :group="group" ></GroupInfoHead>
      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="分组信息" name="group-info">
            <GroupInfo :group="group"></GroupInfo>
        </el-tab-pane>
        <el-tab-pane label="设备列表" name="device-info">
              <DeviceList type='group'></DeviceList>
        </el-tab-pane>   
        <el-tab-pane label="子分组列表" name="child-group-info">
          <GroupList></GroupList>
        </el-tab-pane>
       
      </el-tabs>
    </div>
  </template>

  <script>
   import GroupInfoHead from './GroupInfoHead'
   import GroupInfo from './GroupInfo'
   import GroupList from './GroupList'
   import DeviceList from './../device/DeviceList'

 

    export default {
       components: { GroupInfoHead,GroupInfo ,GroupList,DeviceList},
      
      data() {
        return {      
          group:{},
          activeName:'group-info'
        }
      },
      computed:{
       
      },
      created(){
        this.groupName = this.$route.params.groupName
        this.getGroupInfoByName() 
      },
      methods:{
          goBack(){
            this.$router.back(-1)
          },

         getGroupInfoByName(){
            console.log(this.groupName)
            this.group = {GroupName: "111111111111",
                            UtcCreate: "2019-12-27T02:29:45.000Z",
                            DeviceOnline: 0,
                            GroupDesc: "1111111111111111111",
                            DeviceActive: 0,
                            DeviceCount: 0,
                            GroupId: "BOfqydh6Mpb8Niu8wv5c010200"}
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