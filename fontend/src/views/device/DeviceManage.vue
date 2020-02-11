<template>
    <div name="device">
        <div v-if="!openRouter">
            <el-row :span="24" class="device-info">
                <el-select v-model="productId" placeholder="全部产品"  @change="changeProductId">
                    <el-option
                    v-for="item in productArr"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                    </el-option>
                </el-select>
                <el-divider direction="vertical"></el-divider>
                <div class="device-view">
                    <p>
                        设备总数
                        <el-tooltip content="当前指定产品的设备总数" placement="top" effect="light">
                         <i class="el-icon-question"></i>
                        </el-tooltip>                    
                    </p>
                    <p>{{device_count}}</p>
                </div>
                <el-divider direction="vertical"></el-divider>
                <div  class="device-view">           
                    <p>
                        激活设备
                        <el-tooltip content="当前已激活的设备总数" placement="top" effect="light">
                         <i class="el-icon-question"></i>
                        </el-tooltip>
                    </p>
                    <p>{{device_active_count}}</p>
                </div>
                <el-divider direction="vertical"></el-divider>
                <div  class="device-view">
                    <p>
                        当前在线
                        <el-tooltip content="当前在线的设备总数" placement="top" effect="light">
                         <i class="el-icon-question"></i>
                        </el-tooltip>
                    </p>
                    <p>{{device_online_count}}</p>
                </div>         
            </el-row>
            <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
                <el-tab-pane label="设备列表" name="device">
                    <DeviceList  ref="deviceList" :productArr="productArr" :productId="productId" @setAcount="setDeviceAcount($event)"></DeviceList>
                </el-tab-pane>
                <el-tab-pane label="批次管理" name="batch">
                    <ApplyList  ref="applyList" :productId="productId"></ApplyList>
                </el-tab-pane>     
            </el-tabs>
        </div>
         <router-view/> 
    </div>
</template>
<script>
  import DeviceList from './DeviceList'
  import ApplyList from './../apply/ApplyList'
  
export default {
    components: { DeviceList, ApplyList },
    data() {
      return {
        activeName: 'device',
        productId:0,
        productArr:[],
        openRouter:false,
        device_count:0,
        device_active_count:0,
        device_online_count:0
      }
    },
     watch: {
        $route:function(){
            this.openRouter = !this.openRouter          
        }
    },
    created(){
        
        //传了productId，需要根据productId查询设备
        if(this.$route.params.productId){
             this.productId = this.$route.params.productId
        }
        this.getSimpleProductList()
        //传了device，
        let device = this.$route.params.device
        if(device){
            this.openRouter = true
        }
        
    },
    methods: {
        //查询所有产品的id和编号
        getSimpleProductList(){
            this.$API_IOT.getSimpleProductList().then((res) => {
                this.productArr= res.data.data 
                this.productArr.unshift({id:0,name:'全部产品'})

            })
        },
      handleClick(tab, event) {
        //console.log(tab, event);
        if(tab.name === "batch"){
            this.$refs.applyList.getApplyList()
        }
      },

      changeProductId(val){
        this.productId = val
      },
    
      //设置设备的数量（总数，激活，在线）
      setDeviceAcount(event){
        this.device_count = event.device_count
        this.device_active_count =  event.device_active_count
        this.device_online_count =  event.device_online_count
      }
    }
  };
</script>

<style scoped>
    .device-info{
        display: flex;
        margin-bottom: 20px
    }
    .device-view{
        width:150px;
    }
     .device-view p{
        font-size: 24px;
        color: #373d41;
        display: flex;
        justify-content: flex-start;
        margin-bottom: 4px;
        margin-top: 0px;
        margin-left: 30px
    }
    .device-view p:nth-child(1){
        font-size: 16px;
        color: #888;   
    }
    .el-divider{
        height: 45px;
        margin-left: 50px
    }
     .el-divider:first{
        margin-left: 50px
     }
     .el-icon-question{
         line-height: 1.5;
         margin-left: 5px;
             color: #ccc;
     }
    
</style>