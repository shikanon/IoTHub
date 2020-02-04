<template>
  <div name="ability-list">
      <el-row>
        <el-button type="primary" @click="addAbility">{{type === 'standard' ? '添加标准功能':'添加自定义功能'}}</el-button>
        <el-button @click="copyVisible = true">导入物理模式</el-button>
        <el-button @click="objectModeVisible = true">查看物流模型</el-button>
        <el-button @click="exportDeviceEndCode">生成设备端代码</el-button>
     </el-row>
    
    <el-table
      :data="tableData"
      tooltip-effect="dark"
      style="width: 100%"
      >
       <el-table-column
          label="功能类型"
          width="200"  
         >
          <template slot-scope="scope">{{ !scope.row.EventType  ? '属性':'事件' }} </template>
          
        </el-table-column>
        <el-table-column
            label="功能名称"
            width="200">
            <template slot-scope="scope">{{ scope.row.Name  }} 
                <el-tag  class="tag">{{ scope.row.AbilityType === '1' ? '必选':'可选' }}</el-tag>
            </template>


        </el-table-column>
        <el-table-column
            prop="Identifier"
            label="标识符"  
          >
        </el-table-column>
            
        <el-table-column
            label="数据类型"  >
            <template slot-scope="scope">{{ scope.row.DataType  }}</template>
        </el-table-column>
        <el-table-column
            label="数据定义"  >
            <template slot-scope="scope">{{ scope.row.LastOnlineTime  }}</template>
        </el-table-column>
         <el-table-column
           
            label="操作"
            >
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="edit(scope.row.AbilityId)">编辑</el-button>
              <el-divider v-if="scope.row.AbilityType === '2'" direction="vertical"></el-divider>
              <el-button v-if="scope.row.AbilityType === '2'" type="text" size="small" @click="delete(scope.$index, scope.row)">删除</el-button>
            </template>
          </el-table-column>
    </el-table>
    <el-dialog title="导入物模型" :visible.sync="copyVisible" width="25%">
          <CopyProductAbility ref="copy"  @close="copyVisible = false"></CopyProductAbility>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="copySubmit">确 定</el-button>
                <el-button @click="copyVisible = false">取 消</el-button>
            </span>
      </el-dialog>
      <el-dialog title="查看物模型" :visible.sync="objectModeVisible" width="35%">
          <ObjectMode ref="objectMode"  @close="objectModeVisible = false" :productId="productId"></ObjectMode>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="exportFile">导出模型文件</el-button>
            </span>
      </el-dialog>
  </div>
</template>

<script>
import CopyProductAbility from './CopyProductAbility'
import ObjectMode from './ObjectMode'
import FileSaver from 'file-saver'
import JSZip from 'jszip'

  export default {
    components:{CopyProductAbility,ObjectMode},
     props:{
       type:{
         type:String,
         default:'standard'
       },
        productKey:{
          type:String,
          default:''
        },
         productId:{
          type:Number,
          default:0
        }
     },

  
    data() {
      return {
        tableData: [],
        copyVisible:false,
        objectModeVisible:false
       
      }
    },
 

    created(){
        this.getAbilityList()
    },
    methods: {
       getAbilityList(){
            this.tableData =[
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
                {
                    "Std": true,
                    "Description": "",
                    "Identifier": "MagneticState",
                    "Required": true,
                    "AbilityType": 1,
                    "DataType": "ENUM",
                    "RwFlag": 1,
                    "CategoryType": "GeomagneticSensor",
                    "AbilityId": 6153976,
                    "Name": "车位状态",
                    "DataSpecsList": "[{\"custom\":false,\"dataType\":\"ENUM\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":16139014,\"name\":\"无车\",\"value\":0},{\"custom\":false,\"dataType\":\"ENUM\",\"gmtCreate\":1577343966000,\"gmtModified\":1577343966000,\"id\":16139015,\"name\":\"有车\",\"value\":1}]"
                },
                {
                    "Std": true,
                    "Description": "",
                    "Identifier": "Error",
                    "Required": true,
                    "AbilityType": 3,
                    "EventType": "info",
                    "RwFlag": 1,
                    "CategoryType": "GeomagneticSensor",
                    "AbilityId": 767627,
                    "Name": "故障上报"
                },
                {
                    "Std": true,
                    "Description": "",
                    "Identifier": "AbnormalAlarm",
                    "Required": true,
                    "AbilityType": 3,
                    "EventType": "info",
                    "RwFlag": 1,
                    "CategoryType": "GeomagneticSensor",
                    "AbilityId": 767628,
                    "Name": "异常告警"
                }  
            ]
       },
       addAbility(){
         if(this.type === 'standard'){

         }else{

         }
       },
       copySubmit(){
          this.$refs.copy.submit()
       },

       exportFile(){
          this.$refs.objectMode.exportFile()
       },

       exportDeviceEndCode(){
  
           this.$axios.get('http://localhost:8080/static/test.c').then((res) => {
              //用axios的方法引入地址

              const zip = new JSZip()
            
              // let txtData = ""
              // res.data.forEach((row) => {
              //   let tempStr = ''
              //   tempStr = row.toString()
              //   txtData += `${tempStr}\r\n`
              // })
              zip.file(`${this.productKey}.c`, res.data)
              zip.generateAsync({
                type: "blob"
              }).then((blob) => {
                saveAs(blob,`${this.productKey}.zip`)
              }, (err) => {
                alert('导出失败')
              })
              //console.log(res.data)
              // const blob = new Blob([res.data], {type: ''})
              // FileSaver.saveAs(blob, `${this.productKey}.zip`)
            })          
       }
      
    }
  }
</script>

<style scoped>
 .tag{
        border-radius: 20px;
        margin-left: 10px;
        height: 24px;
        line-height: 24px;
    }
</style>