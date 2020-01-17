<template>
    <div name="step3">
        <p class="step-desc">下载设备连接工具包</p>
        <P>已创建的资源：</P>
        <div class="gray-box">
            <p>{{productName}}</p>
            <div class="title-inf">
                <div>
                    <span class="text-gray title">ProductKey:</span>
                    <span class="text-gray">{{ProductKey}}</span>
                </div>
                <div>
                    <span class="text-gray title">DeviceName:</span>
                    <span class="text-gray">{{DeviceName}}</span>
                    <CopyBtn :content="DeviceName" size="small"></CopyBtn>        
                </div>
                <div>
                    <span class="text-gray title">DeviceSecret:</span>
                    <el-tooltip class="item" effect="dark" :content="DeviceSecret" placement="top">
                           <span class="text-gray">{{ DeviceSecretShow ? `${DeviceSecret.substr (0,10)}...` :'*********'}}</span>
                    </el-tooltip>
                    <el-button type="text" @click="DeviceSecretShow = !DeviceSecretShow">{{ DeviceSecretShow ? '隐藏':'显示'}}</el-button>
                </div>
            </div>
        </div>
        <P>设备功能：</P>
        <div class="gray-box">
            <div class="ability-title">
                <p>功能列表</p>
                <el-button type="text" @click="abilityListShow = !abilityListShow">{{ abilityListShow ? '展开':'收起'}}</el-button>
            </div>
             <el-table v-if="abilityListShow"
                :data="abilityData"
                style="width: 100%">
                <el-table-column
                    label="功能类型"
                    >
                    <template slot-scope="scope">{{ !scope.row.EventType  ? '属性':'事件' }} </template>
                </el-table-column>
                <el-table-column
                    prop="Name"
                    label="功能名称"
                    >
                </el-table-column>
                <el-table-column
                    prop="Identifier"
                    label="标识符">
                </el-table-column>
                <el-table-column
                    prop="DataType"
                    label="数据类型">
                </el-table-column>
                <el-table-column
                    prop="address"
                    label="功能定义">
                </el-table-column>
                <el-table-column
                    prop="address"
                    label="读写类型">
                </el-table-column>
                </el-table>
        </div>
        <p>设备开发工具包:</p>
        <div class="table-info">
            <div  class="table-row">
                <div  class="table-row-label">设备三元组</div>
                <div class="table-row-info">
                    {{`ProductKey：${ProductKey}，DeviceName：${DeviceName}，DeviceSecret：`}}
                     <el-tooltip class="item" effect="dark" :content="DeviceSecret" placement="top">
                        <span class="text-gray"> {{ DeviceSecretShowFlg ? `${DeviceSecret.substr(0,7)}...` :'*********'}}</span>
                    </el-tooltip>
                    <el-button type="text" @click="DeviceSecretShowFlg = !DeviceSecretShowFlg">{{ DeviceSecretShowFlg ? '隐藏':'显示'}}</el-button>
                </div> 
            </div>   
            <div  class="table-row">
                <div  class="table-row-label">设备软件开发工具包</div>
                <div  class="table-row-info"><span >MQTT 协议的 Java SDK</span></div> 
            </div>  
            <div  class="table-row">
                <div  class="table-row-label">用于接收和发送消息的脚本</div>
                <div  class="table-row-info"><span >start.sh</span></div> 
            </div>  
            <div  class="table-row">
                <div  class="table-row-label">下载连接工具包</div>
                <div  class="table-row-info">
                    <el-button type="primary" @click="load">下载windows工具包</el-button>    
                </div> 
            </div>             
        </div>
        
    </div>
</template>

<script>

export default {
  
    data(){
        return{
            productName:'12312313',
            ProductKey:'a1VBSA20hnl',
            DeviceName:'111111',
            DeviceSecret:'D7EPs7FgtjGSJWglKn9AdlOENUIDrtx5',
            DeviceSecretShow:false ,
            abilityListShow:false  ,
            abilityData:[],
            DeviceSecretShowFlg:false         
        }
    },
    created(){
        this.getAbilityData()
    },
    

    methods:{
        getAbilityData(){
             this.abilityData =[
                {
                    "Std": true,
                    "Description": "",
                    "Identifier": "GeoLocation",
                    "Required": true,
                    "AbilityType": 1,
                    "DataType": "STRUCT",
                    "RwFlag": 1,
                    "CategoryType": "GeomagneticSensor",
                    "AbilityId": 6153974,
                    "Name": "地理位置",
                    "DataSpecsList": "[{\"childDataType\":\"DOUBLE\",\"childName\":\"经度\",\"childSpecsDTO\":{\"custom\":false,\"dataType\":\"DOUBLE\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":8044122,\"max\":\"180\",\"min\":\"-180\",\"precise\":7,\"step\":\"0.01\",\"unit\":\"°\",\"unitName\":\"度\"},\"dataSpecs\":{\"$ref\":\"$[0].childSpecsDTO\"},\"dataType\":\"STRUCT\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":7676616,\"identifier\":\"Longitude\",\"isStd\":1,\"name\":\"地理位置\"},{\"childDataType\":\"DOUBLE\",\"childName\":\"纬度\",\"childSpecsDTO\":{\"custom\":false,\"dataType\":\"DOUBLE\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":8044123,\"max\":\"90\",\"min\":\"-90\",\"precise\":7,\"step\":\"0.01\",\"unit\":\"°\",\"unitName\":\"度\"},\"dataSpecs\":{\"$ref\":\"$[1].childSpecsDTO\"},\"dataType\":\"STRUCT\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":7676617,\"identifier\":\"Latitude\",\"isStd\":1,\"name\":\"地理位置\"},{\"childDataType\":\"DOUBLE\",\"childName\":\"海拔\",\"childSpecsDTO\":{\"custom\":false,\"dataType\":\"DOUBLE\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":8044124,\"max\":\"9999\",\"min\":\"0\",\"precise\":7,\"step\":\"0.01\",\"unit\":\"m\",\"unitName\":\"米\"},\"dataSpecs\":{\"$ref\":\"$[2].childSpecsDTO\"},\"dataType\":\"STRUCT\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":7676618,\"identifier\":\"Altitude\",\"isStd\":1,\"name\":\"地理位置\"},{\"childDataType\":\"ENUM\",\"childEnumSpecsDTO\":[{\"custom\":false,\"dataType\":\"ENUM\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":16139012,\"name\":\"WGS_84\",\"value\":1},{\"custom\":false,\"dataType\":\"ENUM\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":16139013,\"name\":\"GCJ_02\",\"value\":2}],\"childName\":\"坐标系统\",\"dataSpecsList\":[{\"custom\":false,\"dataType\":\"ENUM\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":16139012,\"name\":\"WGS_84\",\"value\":1},{\"custom\":false,\"dataType\":\"ENUM\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":16139013,\"name\":\"GCJ_02\",\"value\":2}],\"dataType\":\"STRUCT\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":7676619,\"identifier\":\"CoordinateSystem\",\"isStd\":1,\"name\":\"地理位置\"}]"
                },
                {
                    "DataSpecs": "{\"custom\":false,\"dataType\":\"INT\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":8044125,\"max\":\"100\",\"min\":\"0\",\"precise\":0,\"step\":\"1\",\"unit\":\"%\",\"unitName\":\"百分比\"}",
                    "Std": true,
                    "Description": "",
                    "Identifier": "BatteryLevel",
                    "Required": true,
                    "AbilityType": 1,
                    "DataType": "INT",
                    "RwFlag": 1,
                    "CategoryType": "GeomagneticSensor",
                    "AbilityId": 6153975,
                    "Name": "电池电量"
                },
                
            ]
        },

        load(){
            this.$emit('load')
        },
        nextStep() {
            this.$emit('next')        
        },
    }
}
</script>

<style  scoped>
    .table-row-label{
        width: 170px;
    }
    .gray-box{
        padding: 20px;
        background: #fbfbfc;
        border: 1px solid #ecedee;
    }
    .ability-title{ 
        display: flex;
        justify-content: space-between;
    }
    .table-row-info{
        width: 100%;
    }
    .title-inf .title{
       min-width:80px;
    }
</style>