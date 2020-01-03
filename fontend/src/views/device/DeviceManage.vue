<template>
    <div name="device">
        <div v-if="!openRouter">
            <el-row :span="24" class="device-info">
                <el-select v-model="product" placeholder="全部产品">
                    <el-option
                    v-for="item in products"
                    :key="item.ProductKey"
                    :label="item.ProductName"
                    :value="item.ProductName">
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
                    <p>4</p>
                </div>
                <el-divider direction="vertical"></el-divider>
                <div  class="device-view">
                    <!-- <el-button 
                    type="primary"  
                    circle 
                    size="mini"
                  ></el-button> -->
                    <p>
                        激活设备
                        <el-tooltip content="当前已激活的设备总数" placement="top" effect="light">
                         <i class="el-icon-question"></i>
                        </el-tooltip>
                    </p>
                    <p>1</p>
                </div>
                <el-divider direction="vertical"></el-divider>
                <div  class="device-view">
                    <p>
                        <!-- <el-button
                        type="success" 
                        circle 
                        size="mini"></el-button> -->
                        当前在线
                        <el-tooltip content="当前在线的设备总数" placement="top" effect="light">
                         <i class="el-icon-question"></i>
                        </el-tooltip>
                    </p>
                    <p>0</p>
                </div>         
            </el-row>
            <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
                <el-tab-pane label="设备列表" name="device">
                    <DeviceList :productArr="products"></DeviceList>
                </el-tab-pane>
                <el-tab-pane label="批次管理" name="batch">
                    <ApplyList></ApplyList>
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
        product:'',
        products:[],
        openRouter:false
      }
    },
     watch: {
        $route:function(){
            this.openRouter = !this.openRouter          
        }
    },
    created(){
        this.product = this.$route.params.product
        this.getProductList()
        let device = this.$route.params.device
        if(device){
            this.openRouter = true
        }
        
    },
    methods: {
        getProductList(){
            this.products = [{
                    "Type": "owned",
                    "ProductName": "4444444",
                    "ProductKey": "a1ELejzj0h9",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "123123123",
                    "ProductKey": "a1qfUCxdfqg",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "3231323",
                    "ProductKey": "a1Ibli2tqC2",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "22222",
                    "ProductKey": "a1RIivcyWf0",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "1111",
                    "ProductKey": "a1Vr6NBfa0i",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "123213",
                    "ProductKey": "a10TD9i7WFi",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "213213",
                    "ProductKey": "a1qJfSBZss0",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "3123",
                    "ProductKey": "a1VT5Q8pKbu",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                },
                {
                    "Type": "owned",
                    "ProductName": "test",
                    "ProductKey": "a1zJA7k9sjd",
                    "PublishStatus": "DEVELOPMENT_STATUS"
                }
            ]
        },
      handleClick(tab, event) {
        //console.log(tab, event);
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