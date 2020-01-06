<template>
    <div name="object-mode">
        <div class="desc-txt-gray margin-b-20">
            物模型是对设备在云端的功能描述，包括设备的属性、服务和事件。
            物联网平台通过定义一种物的描述语言来描述物模型，
            称之为 TSL (即Thing Specification Language)，
            采用JSON格式，您可以根据TSL组装上报设备的数据。您可以导出完整物模型，
            用于云端应用开发；您也可以只导出精简物模型，配合设备端SDK实现设备开发。
        </div>
        <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
            <el-tab-pane label="完整物模型" name="complete" class="tab-pane">
                <JsonView :data="completeObject"></JsonView>
            </el-tab-pane>
            <el-tab-pane label="精简物模型" name="simplify"  class="tab-pane">
                <JsonView :data="simplifyObject"></JsonView>
            </el-tab-pane>
        </el-tabs>
        
    </div> 
</template>
 

 <script>
 import FileSaver from 'file-saver'
  export default {
  
    data() {
      return {
        activeName: 'complete',
        completeObject: {},
        simplifyObject:{},
        exportData:'',
      };
    },
    created(){
       this.init()
     //this.test()
    },
    methods: {
        handleClick(tab, event){
            if(tab.index === '0'){
                this.exportData = this.completeObject
                
            }else{
                this.exportData = this.simplifyObject
              
            }
        },

        exportFile(){
            const data = JSON.stringify( this.exportData)
            const blob = new Blob([data], {type: ''})
            FileSaver.saveAs(blob, 'model.json')
        },
        init(){
             this.$axios.get('http://localhost:8080/static/complete.json').then((res) => {
                 this.completeObject  = res.data         
             }) 

            this.$axios.get('http://localhost:8080/static/simplify.json').then((res) => {
                    this.simplifyObject  = res.data       
            })
        },
      
        test(){
                this.completeObject =  {
                    "schema":"https://iotx-tsl.oss-ap-southeast-1.aliyuncs.com/schema.json",
                    "profile": {
                        "productKey": "a1L6W2BQytT"
                    },
                    "services": [
                        {
                        "outputData": [],
                        "identifier": "set",
                        "inputData": [],
                        "method": "thing.service.property.set",
                        "name": "set",
                        "required": true,
                        "callType": "async",
                        "desc": "属性设置"
                        },
                        {
                        "outputData": [
                            {
                            "identifier": "GeoLocation",
                            "dataType": {
                                "specs": [
                                {
                                    "identifier": "Longitude",
                                    "dataType": {
                                    "specs": {
                                        "unit": "°",
                                        "min": "-180",
                                        "unitName": "度",
                                        "max": "180",
                                        "step": "0.01"
                                    },
                                    "type": "double"
                                    },
                                    "name": "经度"
                                },
                                {
                                    "identifier": "Latitude",
                                    "dataType": {
                                    "specs": {
                                        "unit": "°",
                                        "min": "-90",
                                        "unitName": "度",
                                        "max": "90",
                                        "step": "0.01"
                                    },
                                    "type": "double"
                                    },
                                    "name": "纬度"
                                },
                                {
                                    "identifier": "Altitude",
                                    "dataType": {
                                    "specs": {
                                        "unit": "m",
                                        "min": "0",
                                        "unitName": "米",
                                        "max": "9999",
                                        "step": "0.01"
                                    },
                                    "type": "double"
                                    },
                                    "name": "海拔"
                                },
                                {
                                    "identifier": "CoordinateSystem",
                                    "dataType": {
                                    "specs": {
                                        "1": "WGS_84",
                                        "2": "GCJ_02"
                                    },
                                    "type": "enum"
                                    },
                                    "name": "坐标系统"
                                },
                                {
                                    "identifier": "123123",
                                    "dataType": {
                                    "specs": {
                                        "unit": "count",
                                        "min": "1",
                                        "max": "12",
                                        "step": "1"
                                    },
                                    "type": "int"
                                    },
                                    "name": "123213"
                                }
                                ],
                                "type": "struct"
                            },
                            "name": "地理位置"
                            },
                            {
                            "identifier": "BatteryLevel",
                            "dataType": {
                                "specs": {
                                "unit": "%",
                                "min": "0",
                                "unitName": "百分比",
                                "max": "100",
                                "step": "1"
                                },
                                "type": "int"
                            },
                            "name": "电池电量"
                            },
                            {
                            "identifier": "MagneticState",
                            "dataType": {
                                "specs": {
                                "0": "无车",
                                "1": "有车"
                                },
                                "type": "enum"
                            },
                            "name": "车位状态"
                            }
                        ],
                        "identifier": "get",
                        "inputData": [
                            "GeoLocation",
                            "BatteryLevel",
                            "MagneticState"
                        ],
                        "method": "thing.service.property.get",
                        "name": "get",
                        "required": true,
                        "callType": "async",
                        "desc": "属性获取"
                        }
                    ],
                    "properties": [
                        {
                        "identifier": "GeoLocation",
                        "dataType": {
                            "specs": [
                            {
                                "identifier": "Longitude",
                                "dataType": {
                                "specs": {
                                    "unit": "°",
                                    "min": "-180",
                                    "unitName": "度",
                                    "max": "180",
                                    "step": "0.01"
                                },
                                "type": "double"
                                },
                                "name": "经度"
                            },
                            {
                                "identifier": "Latitude",
                                "dataType": {
                                "specs": {
                                    "unit": "°",
                                    "min": "-90",
                                    "unitName": "度",
                                    "max": "90",
                                    "step": "0.01"
                                },
                                "type": "double"
                                },
                                "name": "纬度"
                            },
                            {
                                "identifier": "Altitude",
                                "dataType": {
                                "specs": {
                                    "unit": "m",
                                    "min": "0",
                                    "unitName": "米",
                                    "max": "9999",
                                    "step": "0.01"
                                },
                                "type": "double"
                                },
                                "name": "海拔"
                            },
                            {
                                "identifier": "CoordinateSystem",
                                "dataType": {
                                "specs": {
                                    "1": "WGS_84",
                                    "2": "GCJ_02"
                                },
                                "type": "enum"
                                },
                                "name": "坐标系统"
                            },
                            {
                                "identifier": "123123",
                                "dataType": {
                                "specs": {
                                    "unit": "count",
                                    "min": "1",
                                    "max": "12",
                                    "step": "1"
                                },
                                "type": "int"
                                },
                                "name": "123213"
                            }
                            ],
                            "type": "struct"
                        },
                        "name": "地理位置",
                        "accessMode": "r",
                        "required": true,
                        "desc": "123123213"
                        },
                        {
                        "identifier": "BatteryLevel",
                        "dataType": {
                            "specs": {
                            "unit": "%",
                            "min": "0",
                            "unitName": "百分比",
                            "max": "100",
                            "step": "1"
                            },
                            "type": "int"
                        },
                        "name": "电池电量",
                        "accessMode": "r",
                        "required": true
                        },
                        {
                        "identifier": "MagneticState",
                        "dataType": {
                            "specs": {
                            "0": "无车",
                            "1": "有车"
                            },
                            "type": "enum"
                        },
                        "name": "车位状态",
                        "accessMode": "r",
                        "required": true
                        }
                    ],
                    "events": [
                        {
                        "outputData": [
                            {
                            "identifier": "GeoLocation",
                            "dataType": {
                                "specs": [
                                {
                                    "identifier": "Longitude",
                                    "dataType": {
                                    "specs": {
                                        "unit": "°",
                                        "min": "-180",
                                        "unitName": "度",
                                        "max": "180",
                                        "step": "0.01"
                                    },
                                    "type": "double"
                                    },
                                    "name": "经度"
                                },
                                {
                                    "identifier": "Latitude",
                                    "dataType": {
                                    "specs": {
                                        "unit": "°",
                                        "min": "-90",
                                        "unitName": "度",
                                        "max": "90",
                                        "step": "0.01"
                                    },
                                    "type": "double"
                                    },
                                    "name": "纬度"
                                },
                                {
                                    "identifier": "Altitude",
                                    "dataType": {
                                    "specs": {
                                        "unit": "m",
                                        "min": "0",
                                        "unitName": "米",
                                        "max": "9999",
                                        "step": "0.01"
                                    },
                                    "type": "double"
                                    },
                                    "name": "海拔"
                                },
                                {
                                    "identifier": "CoordinateSystem",
                                    "dataType": {
                                    "specs": {
                                        "1": "WGS_84",
                                        "2": "GCJ_02"
                                    },
                                    "type": "enum"
                                    },
                                    "name": "坐标系统"
                                },
                                {
                                    "identifier": "123123",
                                    "dataType": {
                                    "specs": {
                                        "unit": "count",
                                        "min": "1",
                                        "max": "12",
                                        "step": "1"
                                    },
                                    "type": "int"
                                    },
                                    "name": "123213"
                                }
                                ],
                                "type": "struct"
                            },
                            "name": "地理位置"
                            },
                            {
                            "identifier": "BatteryLevel",
                            "dataType": {
                                "specs": {
                                "unit": "%",
                                "min": "0",
                                "unitName": "百分比",
                                "max": "100",
                                "step": "1"
                                },
                                "type": "int"
                            },
                            "name": "电池电量"
                            },
                            {
                            "identifier": "MagneticState",
                            "dataType": {
                                "specs": {
                                "0": "无车",
                                "1": "有车"
                                },
                                "type": "enum"
                            },
                            "name": "车位状态"
                            }
                        ],
                        "identifier": "post",
                        "method": "thing.event.property.post",
                        "name": "post",
                        "type": "info",
                        "required": true,
                        "desc": "属性上报"
                        },
                        {
                        "outputData": [
                            {
                            "identifier": "ErrorCode",
                            "dataType": {
                                "specs": {
                                "0": "无故障"
                                },
                                "type": "enum"
                            },
                            "name": "故障代码"
                            }
                        ],
                        "identifier": "Error",
                        "method": "thing.event.Error.post",
                        "name": "故障上报",
                        "type": "info",
                        "required": true
                        },
                        {
                        "outputData": [
                            {
                            "identifier": "AlarmType",
                            "dataType": {
                                "specs": {
                                "0": "正常",
                                "1": "强磁干扰",
                                "2": "低电量"
                                },
                                "type": "enum"
                            },
                            "name": "告警类型"
                            }
                        ],
                        "identifier": "AbnormalAlarm",
                        "method": "thing.event.AbnormalAlarm.post",
                        "name": "异常告警",
                        "type": "info",
                        "required": true
                        }
                    ]
                }
         
        

                this.simplifyObject = {
                    "properties": [
                    {
                    "identifier": "GeoLocation",
                    "dataType": {
                        "specs": [
                        {
                            "identifier": "Longitude",
                            "dataType": {
                            "type": "double"
                            }
                        },
                        {
                            "identifier": "Latitude",
                            "dataType": {
                            "type": "double"
                            }
                        },
                        {
                            "identifier": "Altitude",
                            "dataType": {
                            "type": "double"
                            }
                        },
                        {
                            "identifier": "CoordinateSystem",
                            "dataType": {
                            "type": "enum"
                            }
                        },
                        {
                            "identifier": "123123",
                            "dataType": {
                            "type": "int"
                            }
                        }
                        ],
                        "type": "struct"
                    }
                    },
                    {
                    "identifier": "BatteryLevel",
                    "dataType": {
                        "type": "int"
                    }
                    },
                    {
                    "identifier": "MagneticState",
                    "dataType": {
                        "type": "enum"
                    }
                    }
                 ],
                    "events": [
                    {
                    "outputData": [
                        {
                        "identifier": "ErrorCode",
                        "dataType": {
                            "type": "enum"
                        }
                        }
                    ],
                    "identifier": "Error",
                    "type": "info"
                    },
                    {
                    "outputData": [
                        {
                        "identifier": "AlarmType",
                        "dataType": {
                            "type": "enum"
                        }
                        }
                    ],
                    "identifier": "AbnormalAlarm",
                    "type": "info"
                    }
                 ]
                }
         }

     
        }
  }
</script>


<style scoped>
 .tab-pane{
     height: 240px !important;
     overflow-y: auto !important;
 }
</style>
