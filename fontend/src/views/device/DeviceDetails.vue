
<template>
  <div name="device-info">
    <div class="page-header">
        <el-page-header @back="goBack" :content="device.DeviceName" >       
        </el-page-header>
        <el-tag :type="device.Status === 'ONLINE' ? 'success': device.Status === 'UNACTIVE' ? 'info': device.Status === 'DISABLE' ? 'danger':'warning' " 
        class="tag">{{ device.Status | deviceStatusFilter }}</el-tag>
    </div>
       <DeviceInfoHead :device="device" ></DeviceInfoHead>

      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="设备信息" name="first">
              <DeviceInfo :deviceName="device.DeviceName"></DeviceInfo>
        </el-tab-pane>
        <el-tab-pane label="Topic列表" name="first1">
              <TopicList :addBtnVisible="addBtnVisible"></TopicList>
        </el-tab-pane>
        <el-tab-pane label="运行状态" name="second">     
            <runState></runState>
        </el-tab-pane>
        <el-tab-pane label="事件管理" name="third2">
           
        </el-tab-pane>
        <el-tab-pane label="服务调用" name="third3">
            
        </el-tab-pane>   
        <el-tab-pane label="设备影子" name="third4">
           
        </el-tab-pane>
        <el-tab-pane label="文件管理" name="third5">
           
        </el-tab-pane>
        
        <el-tab-pane label="日志服务" name="third6">
           
        </el-tab-pane>
        <el-tab-pane label="在线调试" name="fourth">
            
        </el-tab-pane>
      </el-tabs>
    
  </div>
</template>
  <script>
  import DeviceInfoHead from './DeviceInfoHead'
  import DeviceInfo from './DeviceInfo'
  import RunState from '@/views/runState/RunState'
  
    export default {
      components: { DeviceInfoHead,DeviceInfo,runState:RunState },
      data() {
        return {
          activeName:'first',
          device:{},
          addBtnVisible:false
        }
      },
     
      created(){     
        this.device = JSON.parse(this.$route.params.device)
      },
     
      methods:{
       
        goBack(){
            this.$emit('changeIsRouter')
            this.$router.back(-1)
        },
       
        handleClick(tab, event) {
          //console.log(tab, event);
        },
       
      
        queryDataByName(){
          this.dialogTableVisible = true 
        }
         
         
      }
    }
  </script>

  <style scoped>
    .page-header{
      margin-bottom: 20px;
      display: flex;
     line-height: 32px;  
    }
    .tag{
        border-radius: 20px;
        margin-left: 20px;
        height: 24px;
        line-height: 24px;
    }
    
  </style>