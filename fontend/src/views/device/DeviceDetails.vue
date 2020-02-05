
<template>
  <div name="device-info">
    <div class="header">
        <el-page-header @back="goBack" :content="deviceInfo.name" >       
        </el-page-header>
        <el-tag :type="deviceInfo.status_id === 5 ? 'success': deviceInfo.status_id === 1 ? 'info': deviceInfo.status_id === 4 ? 'danger':'warning' " 
        class="tag">{{ deviceInfo.status_id | deviceStatusFilter }}</el-tag>
    </div>
       <DeviceInfoHead :device="deviceInfo" ></DeviceInfoHead>

      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="设备信息" name="deviceInfo">
              <DeviceInfo :device="deviceInfo"></DeviceInfo>
        </el-tab-pane>
        <el-tab-pane label="Topic列表" name="topic" lazy>
              <Topic type="device" :queryKey="deviceId"></Topic>
        </el-tab-pane>
        <el-tab-pane label="运行状态" name="runState" lazy>     
            <runState :deviceId="deviceId"></runState>
        </el-tab-pane>
        <el-tab-pane label="事件管理" name="EventList" lazy>
          <EventList :deviceId="deviceId"  type="event"></EventList>
        </el-tab-pane>
        <el-tab-pane label="服务调用" name="serviceList" lazy>
            <EventList :deviceId="deviceId" type="service"></EventList>
        </el-tab-pane>   
        <el-tab-pane label="设备影子" name="DeviceShadow" lazy> 
             <DeviceShadow :status="deviceInfo.status_id"></DeviceShadow>
        </el-tab-pane>
        <el-tab-pane label="文件管理" name="fileList" lazy>
           
        </el-tab-pane>
        
        <el-tab-pane label="日志服务" name="logService" lazy>
           
        </el-tab-pane>
        <el-tab-pane label="在线调试" name="debug" lazy>
            
        </el-tab-pane>
      </el-tabs>
    
  </div>
</template>
  <script>
  import DeviceInfoHead from './DeviceInfoHead'
  import DeviceInfo from './DeviceInfo'
  import RunState from '@/views/runState/RunState'
  import EventList from '@/views/event/EventList'
  import DeviceShadow from './DeviceShadow'

  
    export default {
      components: { DeviceInfoHead,DeviceInfo,runState:RunState, EventList,DeviceShadow},
      data() {
        return {
          activeName:'deviceInfo',
          deviceId:'',
          deviceInfo:{},
          addBtnVisible:false
        }
      },
     
      created(){     
        this.deviceId = this.$route.params.deviceId
        this.getDeviceDtl()
      },
     
      methods:{
       
        goBack(){
            this.$emit('changeIsRouter')
            this.$router.back(-1)
        },
       
        handleClick(tab, event) {
          //console.log(tab, event);
        },

        getDeviceDtl(){
             this.$API_IOT.getDeviceDtl(this.deviceId).then((res) => {
                 this.deviceInfo = res.data.data         
            })
        },
       
      
        queryDataByName(){
          this.dialogTableVisible = true 
        }
         
         
      }
    }
  </script>

  <style scoped>
    .header{
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