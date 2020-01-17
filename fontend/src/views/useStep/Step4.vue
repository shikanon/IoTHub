<template>
    <div name="step4">
        <p class="step-desc">运行脚本，设备连接云端上报数据并下发指令到设备端</p>
        <P>第一步：在电脑或者嵌入式设备，解压工具包，可以执行下面的命令行：</P>
        <div class="gray-box">
             <div class="flex-center-between ">
                <p>unzip aliyun_iot_device_quickstart.zip</p>
                 <CopyBtn content="unzip aliyun_iot_device_quickstart.zip" size="small"></CopyBtn>        
            </div>
        </div>
        <P>第二步：进入文件，执行脚本，可以执行下面的命令行：</P>
        <div class="gray-box">
            <div class="flex-center-between ">
                <p>./ start.bat</p>
                <CopyBtn content="./ start.bat" size="small"></CopyBtn>        
            </div>
        </div>
        <p>第三步：运行启动脚本。下面将显示物模型的数据（注：脚本会每隔5秒上报数据，数据为：{"Status":"1","Data":"Hello, World"}）</p>
        <div class="gray-box">
            <div class="flex-center-between ">
                <p>
                    <span class="device-log-txt">设备日志</span>
                    <span class="not-active" >检测到设备 未激活</span>
                </p>             
                <div>
                    <el-button size="small">刷新</el-button>
                    <el-button size="small">清屏</el-button>
                </div>                          
            </div>   
        </div>   
        <div class="log-box">
            <div>
                <img src="https://img.alicdn.com/tfs/TB1M5KJmrvpK1RjSZFqXXcXUVXa-120-144.png" style="display: block; width: 60px;">
                <span>暂无数据</span>
            </div>
        </div>
    </div>
</template>

<script>

export default {
  
    data(){
        return{
            productName:'12312313',
            DeviceSecretShow:false ,
            abilityListShow:false  ,
            abilityData:[]         
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
        submitForm() {
            this.$refs.form.validate((valid) => {
                if (valid) {
                    alert('submit')
                    this.$emit('next')
                } else {
                    console.log('error submit!!');
                    return false;
                }
            });
        },
    }
}
</script>

<style  scoped>
    .table-row-label{
        width: 200px;
    }
    .gray-box{
        padding: 0 5px;
        background: #fbfbfc;
        border: 1px solid #ecedee;
    }
    .flex-center-between{ 
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
  

    .device-log-txt{
        font-weight: 700;
        margin-right: 16px;
        color: #373d41;
    }
    
    .not-active{
        color:#999
    }
    .not-active:before{
        content: '';
        width: 10px;
        height: 10px;
        border-radius: 50%;
        background-color: #888;
        position: absolute;
        margin-left: -15px;
        margin-top: 5px;
    }

    .log-box{
        min-height: 200px;
        border: 1px solid #ecedee;
        border-top: 0px;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .log-box div{
        display: flex;
        align-items: center;
        justify-content:center;
    }
    
  
</style>