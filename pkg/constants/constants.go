package constants

const (
	DeviceTopicPrefix        = "sys"
	DeviceTopicPrefixIndex   = 1
	ProductKeyIndex          = 2
	DeviceNameIndex          = 3
	EventIndex               = 6
	ServiceIndex             = 6
	Success                  = 200
	SuccessMsg               = "Success"
	RequestError             = 400
	RequestErrorMsg          = "服务器繁忙"
	RequestParameterError    = 460
	RequestParameterErrorMsg = "请求参数错误"
	TooManyRequestsError     = 429
	TooManyRequestsErrorMsg  = "请求过于频繁"
	DeviceNotFoundError      = 6100
	DeviceNotFoundErrorMsg   = "设备不存在"
	Version                  = "1"
	SetProperty              = "thing.service.property.set"
	CallService              = "thing.service.%s"
	ShadowTopic              = "/shadow/get/%s/%s"
	ShadowTopicPrefix        = "shadow"
	ShadowTopicPrefixIndex   = 1
	ShadowProductKeyIndex    = 3
	ShadowDeviceNameIndex    = 4
	JsonFormatError          = 400
	JsonFormatErrorMsg       = "不正确的JSON格式。"
	MissingMethodError       = 401
	MissingMethodErrorMsg    = "影子数据缺少method信息。"
	MissingStateError        = 402
	MissingStateErrorMsg     = "影子数据缺少state字段。"
	VersionIsNotIntError     = 403
	VersionIsNotIntErrorMsg  = "影子数据中version值不是数字。"
	MissingReportedError     = 404
	MissingReportedErrorMsg  = "影子数据缺少reported字段。"
	ReportedIsNullError      = 405
	ReportedIsNullErrorMsg   = "影子数据中 reported属性字段为空。"
	ValidMethodError         = 406
	ValidMethodErrorMsg      = "影子数据中 method是无效的方法。"
	ShadowIsNullError        = 407
	ShadowIsNullErrorMsg     = "影子内容为空。"
	TooManyReportedError     = 408
	TooManyReportedErrorMsg  = "影子数据中reported属性个数超过128个。"
	VersionConflictError     = 409
	VersionConflictErrorMsg  = "影子版本冲突。"
	ServerExceptionError     = 500
	ServerExceptionErrorMsg  = "服务端处理异常。"
	ModelJson                = `{
  "schema": "https://iotx-tsl.oss-ap-southeast-1.aliyuncs.com/schema.json",
  "profile": {
    "productKey": "a1hDUAnwB19"
  },
  "services": [
    {
      "outputData": [],
      "identifier": "set",
      "inputData": [
        {
          "identifier": "LightStatus",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "工作状态"
        },
        {
          "identifier": "LightAdjustLevel",
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
          "name": "调光等级"
        },
        {
          "identifier": "ErrorPowerThreshold",
          "dataType": {
            "specs": {
              "unit": "W",
              "min": "0",
              "unitName": "瓦特",
              "max": "1000",
              "step": "1"
            },
            "type": "int"
          },
          "name": "故障功率门限"
        },
        {
          "identifier": "ErrorCurrentThreshold",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0.1",
              "unitName": "安培",
              "max": "9",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "故障电流门限"
        },
        {
          "identifier": "TiltThreshold",
          "dataType": {
            "specs": {
              "unit": "°",
              "min": "0",
              "unitName": "度",
              "max": "90",
              "step": "1"
            },
            "type": "int"
          },
          "name": "倾斜阈值"
        },
        {
          "identifier": "UnderVoltThreshold",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "1"
            },
            "type": "int"
          },
          "name": "欠压阈值"
        },
        {
          "identifier": "OverCurrentThreshold",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0",
              "unitName": "安培",
              "max": "9",
              "step": "1"
            },
            "type": "int"
          },
          "name": "过流阈值"
        },
        {
          "identifier": "OverVoltThreshold",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "1"
            },
            "type": "int"
          },
          "name": "过压阈值"
        },
        {
          "identifier": "LightErrorEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "灯具故障使能"
        },
        {
          "identifier": "OverCurrentEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "过流告警使能"
        },
        {
          "identifier": "OverVoltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "过压告警使能"
        },
        {
          "identifier": "UnderVoltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "欠压告警使能"
        },
        {
          "identifier": "LeakageEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "漏电告警使能"
        },
        {
          "identifier": "OverTiltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "倾斜告警使能"
        },
        {
          "identifier": "GeoLocation",
          "dataType": {
            "specs": [
              {
                "identifier": "longitude",
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
                "identifier": "latitude",
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
                "name": "纬度"
              },
              {
                "identifier": "altitude",
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
                "name": "高度"
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
              }
            ],
            "type": "struct"
          },
          "name": "地理位置"
        },
        {
          "identifier": "DeviceTime",
          "dataType": {
            "specs": {},
            "type": "date"
          },
          "name": "设备时间"
        },
        {
          "identifier": "CurrentVoltage",
          "dataType": {
            "specs": {
              "item": {
                "type": "int"
              },
              "size": "10"
            },
            "type": "array"
          },
          "name": "wsp"
        },
        {
          "identifier": "Nono",
          "dataType": {
            "specs": {
              "item": {
                "specs": [
                  {
                    "identifier": "a",
                    "dataType": {
                      "specs": {
                        "min": "0",
                        "max": "10",
                        "step": "1"
                      },
                      "type": "int"
                    },
                    "name": "a"
                  },
                  {
                    "identifier": "b",
                    "dataType": {
                      "specs": {
                        "min": "0",
                        "max": "10",
                        "step": "1"
                      },
                      "type": "float"
                    },
                    "name": "b"
                  }
                ],
                "type": "struct"
              },
              "size": "10"
            },
            "type": "array"
          },
          "name": "uuss"
        }
      ],
      "method": "thing.service.property.set",
      "name": "set",
      "required": true,
      "callType": "async",
      "desc": "属性设置"
    },
    {
      "outputData": [
        {
          "identifier": "LightStatus",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "工作状态"
        },
        {
          "identifier": "LightAdjustLevel",
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
          "name": "调光等级"
        },
        {
          "identifier": "LightVolt",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "4",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "工作电压"
        },
        {
          "identifier": "LightCurrent",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0",
              "unitName": "安培",
              "max": "9",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "工作电流"
        },
        {
          "identifier": "ActivePower",
          "dataType": {
            "specs": {
              "unit": "W",
              "min": "0",
              "unitName": "瓦特",
              "max": "1000",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "有功功率值"
        },
        {
          "identifier": "PowerRatio",
          "dataType": {
            "specs": {
              "unit": "pF",
              "min": "0.01",
              "unitName": "皮法",
              "max": "1",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "功率因数"
        },
        {
          "identifier": "PowerConsumption",
          "dataType": {
            "specs": {
              "unit": "kW·h",
              "min": "0",
              "unitName": "千瓦·时",
              "max": "2147483647",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "用电量"
        },
        {
          "identifier": "DrainVoltage",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "漏电压"
        },
        {
          "identifier": "TiltValue",
          "dataType": {
            "specs": {
              "unit": "°",
              "min": "0",
              "unitName": "度",
              "max": "90",
              "step": "1"
            },
            "type": "int"
          },
          "name": "倾斜角度值"
        },
        {
          "identifier": "ErrorPowerThreshold",
          "dataType": {
            "specs": {
              "unit": "W",
              "min": "0",
              "unitName": "瓦特",
              "max": "1000",
              "step": "1"
            },
            "type": "int"
          },
          "name": "故障功率门限"
        },
        {
          "identifier": "ErrorCurrentThreshold",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0.1",
              "unitName": "安培",
              "max": "9",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "故障电流门限"
        },
        {
          "identifier": "TiltThreshold",
          "dataType": {
            "specs": {
              "unit": "°",
              "min": "0",
              "unitName": "度",
              "max": "90",
              "step": "1"
            },
            "type": "int"
          },
          "name": "倾斜阈值"
        },
        {
          "identifier": "UnderVoltThreshold",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "1"
            },
            "type": "int"
          },
          "name": "欠压阈值"
        },
        {
          "identifier": "OverCurrentThreshold",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0",
              "unitName": "安培",
              "max": "9",
              "step": "1"
            },
            "type": "int"
          },
          "name": "过流阈值"
        },
        {
          "identifier": "OverVoltThreshold",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "1"
            },
            "type": "int"
          },
          "name": "过压阈值"
        },
        {
          "identifier": "LightErrorEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "灯具故障使能"
        },
        {
          "identifier": "OverCurrentEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "过流告警使能"
        },
        {
          "identifier": "OverVoltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "过压告警使能"
        },
        {
          "identifier": "UnderVoltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "欠压告警使能"
        },
        {
          "identifier": "LeakageEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "漏电告警使能"
        },
        {
          "identifier": "OverTiltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "倾斜告警使能"
        },
        {
          "identifier": "LampError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "灯具故障告警"
        },
        {
          "identifier": "OverCurrentError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "过流告警"
        },
        {
          "identifier": "OverVoltError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "过压告警"
        },
        {
          "identifier": "UnderVoltError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "欠压告警"
        },
        {
          "identifier": "OverTiltError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "倾斜告警"
        },
        {
          "identifier": "LeakageError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "漏电告警"
        },
        {
          "identifier": "GeoLocation",
          "dataType": {
            "specs": [
              {
                "identifier": "longitude",
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
                "identifier": "latitude",
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
                "name": "纬度"
              },
              {
                "identifier": "altitude",
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
                "name": "高度"
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
              }
            ],
            "type": "struct"
          },
          "name": "地理位置"
        },
        {
          "identifier": "DeviceTime",
          "dataType": {
            "specs": {},
            "type": "date"
          },
          "name": "设备时间"
        },
        {
          "identifier": "ManufacturerName",
          "dataType": {
            "specs": {
              "length": "256"
            },
            "type": "text"
          },
          "name": "制造商名称"
        },
        {
          "identifier": "CurrentVoltage",
          "dataType": {
            "specs": {
              "item": {
                "type": "int"
              },
              "size": "10"
            },
            "type": "array"
          },
          "name": "wsp"
        },
        {
          "identifier": "Nono",
          "dataType": {
            "specs": {
              "item": {
                "specs": [
                  {
                    "identifier": "a",
                    "dataType": {
                      "specs": {
                        "min": "0",
                        "max": "10",
                        "step": "1"
                      },
                      "type": "int"
                    },
                    "name": "a"
                  },
                  {
                    "identifier": "b",
                    "dataType": {
                      "specs": {
                        "min": "0",
                        "max": "10",
                        "step": "1"
                      },
                      "type": "float"
                    },
                    "name": "b"
                  }
                ],
                "type": "struct"
              },
              "size": "10"
            },
            "type": "array"
          },
          "name": "uuss"
        }
      ],
      "identifier": "get",
      "inputData": [
        "LightStatus",
        "LightAdjustLevel",
        "LightVolt",
        "LightCurrent",
        "ActivePower",
        "PowerRatio",
        "PowerConsumption",
        "DrainVoltage",
        "TiltValue",
        "ErrorPowerThreshold",
        "ErrorCurrentThreshold",
        "TiltThreshold",
        "UnderVoltThreshold",
        "OverCurrentThreshold",
        "OverVoltThreshold",
        "LightErrorEnable",
        "OverCurrentEnable",
        "OverVoltEnable",
        "UnderVoltEnable",
        "LeakageEnable",
        "OverTiltEnable",
        "LampError",
        "OverCurrentError",
        "OverVoltError",
        "UnderVoltError",
        "OverTiltError",
        "LeakageError",
        "GeoLocation",
        "DeviceTime",
        "ManufacturerName",
        "CurrentVoltage",
        "Nono"
      ],
      "method": "thing.service.property.get",
      "name": "get",
      "required": true,
      "callType": "async",
      "desc": "属性获取"
    },
    {
      "outputData": [],
      "identifier": "TimeReset",
      "inputData": [
        {
          "identifier": "TimeReset",
          "dataType": {
            "specs": {
              "length": "255"
            },
            "type": "text"
          },
          "name": "TimeReset"
        }
      ],
      "method": "thing.service.TimeReset",
      "name": "设备校时服务",
      "required": false,
      "callType": "async"
    }
  ],
  "properties": [
    {
      "identifier": "LightStatus",
      "dataType": {
        "specs": {
          "0": "关闭",
          "1": "打开"
        },
        "type": "bool"
      },
      "name": "工作状态",
      "accessMode": "rw",
      "required": true
    },
    {
      "identifier": "LightAdjustLevel",
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
      "name": "调光等级",
      "accessMode": "rw",
      "required": true,
      "desc": "调光等级采用百分比表示"
    },
    {
      "identifier": "LightVolt",
      "dataType": {
        "specs": {
          "unit": "V",
          "min": "0",
          "unitName": "伏特",
          "max": "4",
          "step": "0.1"
        },
        "type": "float"
      },
      "name": "工作电压",
      "accessMode": "r",
      "required": true,
      "desc": "显示设备电压；电参数采用4个字节浮点型数据"
    },
    {
      "identifier": "LightCurrent",
      "dataType": {
        "specs": {
          "unit": "A",
          "min": "0",
          "unitName": "安培",
          "max": "9",
          "step": "0.1"
        },
        "type": "float"
      },
      "name": "工作电流",
      "accessMode": "r",
      "required": true,
      "desc": "电参数采用4个字节浮点型数据"
    },
    {
      "identifier": "ActivePower",
      "dataType": {
        "specs": {
          "unit": "W",
          "min": "0",
          "unitName": "瓦特",
          "max": "1000",
          "step": "0.1"
        },
        "type": "float"
      },
      "name": "有功功率值",
      "accessMode": "r",
      "required": true,
      "desc": "电参数采用4个字节浮点型数据"
    },
    {
      "identifier": "PowerRatio",
      "dataType": {
        "specs": {
          "unit": "pF",
          "min": "0.01",
          "unitName": "皮法",
          "max": "1",
          "step": "0.1"
        },
        "type": "float"
      },
      "name": "功率因数",
      "accessMode": "r",
      "required": true,
      "desc": "电参数采用4个字节浮点型数据"
    },
    {
      "identifier": "PowerConsumption",
      "dataType": {
        "specs": {
          "unit": "kW·h",
          "min": "0",
          "unitName": "千瓦·时",
          "max": "2147483647",
          "step": "0.1"
        },
        "type": "float"
      },
      "name": "用电量",
      "accessMode": "r",
      "required": true,
      "desc": "耗电量；电参数采用4个字节浮点型数据"
    },
    {
      "identifier": "DrainVoltage",
      "dataType": {
        "specs": {
          "unit": "V",
          "min": "0",
          "unitName": "伏特",
          "max": "400",
          "step": "0.1"
        },
        "type": "float"
      },
      "name": "漏电压",
      "accessMode": "r",
      "required": true,
      "desc": "电参数采用4个字节浮点型数据"
    },
    {
      "identifier": "TiltValue",
      "dataType": {
        "specs": {
          "unit": "°",
          "min": "0",
          "unitName": "度",
          "max": "90",
          "step": "1"
        },
        "type": "int"
      },
      "name": "倾斜角度值",
      "accessMode": "r",
      "required": true,
      "desc": "路灯的倾斜角度；采用1个字节16进制数"
    },
    {
      "identifier": "ErrorPowerThreshold",
      "dataType": {
        "specs": {
          "unit": "W",
          "min": "0",
          "unitName": "瓦特",
          "max": "1000",
          "step": "1"
        },
        "type": "int"
      },
      "name": "故障功率门限",
      "accessMode": "rw",
      "required": true,
      "desc": "故障功率门限"
    },
    {
      "identifier": "ErrorCurrentThreshold",
      "dataType": {
        "specs": {
          "unit": "A",
          "min": "0.1",
          "unitName": "安培",
          "max": "9",
          "step": "0.1"
        },
        "type": "float"
      },
      "name": "故障电流门限",
      "accessMode": "rw",
      "required": true
    },
    {
      "identifier": "TiltThreshold",
      "dataType": {
        "specs": {
          "unit": "°",
          "min": "0",
          "unitName": "度",
          "max": "90",
          "step": "1"
        },
        "type": "int"
      },
      "name": "倾斜阈值",
      "accessMode": "rw",
      "required": true
    },
    {
      "identifier": "UnderVoltThreshold",
      "dataType": {
        "specs": {
          "unit": "V",
          "min": "0",
          "unitName": "伏特",
          "max": "400",
          "step": "1"
        },
        "type": "int"
      },
      "name": "欠压阈值",
      "accessMode": "rw",
      "required": true
    },
    {
      "identifier": "OverCurrentThreshold",
      "dataType": {
        "specs": {
          "unit": "A",
          "min": "0",
          "unitName": "安培",
          "max": "9",
          "step": "1"
        },
        "type": "int"
      },
      "name": "过流阈值",
      "accessMode": "rw",
      "required": true
    },
    {
      "identifier": "OverVoltThreshold",
      "dataType": {
        "specs": {
          "unit": "V",
          "min": "0",
          "unitName": "伏特",
          "max": "400",
          "step": "1"
        },
        "type": "int"
      },
      "name": "过压阈值",
      "accessMode": "rw",
      "required": true
    },
    {
      "identifier": "LightErrorEnable",
      "dataType": {
        "specs": {
          "0": "关闭",
          "1": "打开"
        },
        "type": "bool"
      },
      "name": "灯具故障使能",
      "accessMode": "rw",
      "required": true,
      "desc": "1：打开，0：关闭"
    },
    {
      "identifier": "OverCurrentEnable",
      "dataType": {
        "specs": {
          "0": "关闭",
          "1": "打开"
        },
        "type": "bool"
      },
      "name": "过流告警使能",
      "accessMode": "rw",
      "required": true,
      "desc": "1：打开，0：关闭"
    },
    {
      "identifier": "OverVoltEnable",
      "dataType": {
        "specs": {
          "0": "关闭",
          "1": "打开"
        },
        "type": "bool"
      },
      "name": "过压告警使能",
      "accessMode": "rw",
      "required": true,
      "desc": "1：打开，0：关闭"
    },
    {
      "identifier": "UnderVoltEnable",
      "dataType": {
        "specs": {
          "0": "关闭",
          "1": "打开"
        },
        "type": "bool"
      },
      "name": "欠压告警使能",
      "accessMode": "rw",
      "required": true,
      "desc": "1：打开，0：关闭"
    },
    {
      "identifier": "LeakageEnable",
      "dataType": {
        "specs": {
          "0": "关闭",
          "1": "打开"
        },
        "type": "bool"
      },
      "name": "漏电告警使能",
      "accessMode": "rw",
      "required": true,
      "desc": "1：打开，0：关闭"
    },
    {
      "identifier": "OverTiltEnable",
      "dataType": {
        "specs": {
          "0": "关闭",
          "1": "打开"
        },
        "type": "bool"
      },
      "name": "倾斜告警使能",
      "accessMode": "rw",
      "required": true,
      "desc": "1：打开，0：关闭"
    },
    {
      "identifier": "LampError",
      "dataType": {
        "specs": {
          "0": "正常",
          "1": "告警"
        },
        "type": "bool"
      },
      "name": "灯具故障告警",
      "accessMode": "r",
      "required": true,
      "desc": "1-告警；0-正常"
    },
    {
      "identifier": "OverCurrentError",
      "dataType": {
        "specs": {
          "0": "正常",
          "1": "告警"
        },
        "type": "bool"
      },
      "name": "过流告警",
      "accessMode": "r",
      "required": true,
      "desc": "1-告警；0-正常"
    },
    {
      "identifier": "OverVoltError",
      "dataType": {
        "specs": {
          "0": "正常",
          "1": "告警"
        },
        "type": "bool"
      },
      "name": "过压告警",
      "accessMode": "r",
      "required": true,
      "desc": "1-告警；0-正常"
    },
    {
      "identifier": "UnderVoltError",
      "dataType": {
        "specs": {
          "0": "正常",
          "1": "告警"
        },
        "type": "bool"
      },
      "name": "欠压告警",
      "accessMode": "r",
      "required": true,
      "desc": "1-告警；0-正常"
    },
    {
      "identifier": "OverTiltError",
      "dataType": {
        "specs": {
          "0": "正常",
          "1": "告警"
        },
        "type": "bool"
      },
      "name": "倾斜告警",
      "accessMode": "r",
      "required": true,
      "desc": "1-告警；0-正常"
    },
    {
      "identifier": "LeakageError",
      "dataType": {
        "specs": {
          "0": "正常",
          "1": "告警"
        },
        "type": "bool"
      },
      "name": "漏电告警",
      "accessMode": "r",
      "required": true,
      "desc": "1-告警；0-正常"
    },
    {
      "identifier": "GeoLocation",
      "dataType": {
        "specs": [
          {
            "identifier": "longitude",
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
            "identifier": "latitude",
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
            "name": "纬度"
          },
          {
            "identifier": "altitude",
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
            "name": "高度"
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
          }
        ],
        "type": "struct"
      },
      "name": "地理位置",
      "accessMode": "rw",
      "required": true
    },
    {
      "identifier": "DeviceTime",
      "dataType": {
        "specs": {},
        "type": "date"
      },
      "name": "设备时间",
      "accessMode": "rw",
      "required": false,
      "desc": "设备时间"
    },
    {
      "identifier": "ManufacturerName",
      "dataType": {
        "specs": {
          "length": "256"
        },
        "type": "text"
      },
      "name": "制造商名称",
      "accessMode": "r",
      "required": false,
      "desc": "制造商名称"
    },
    {
      "identifier": "CurrentVoltage",
      "dataType": {
        "specs": {
          "item": {
            "type": "int"
          },
          "size": "10"
        },
        "type": "array"
      },
      "name": "wsp",
      "accessMode": "rw",
      "required": false
    },
    {
      "identifier": "Nono",
      "dataType": {
        "specs": {
          "item": {
            "specs": [
              {
                "identifier": "a",
                "dataType": {
                  "specs": {
                    "min": "0",
                    "max": "10",
                    "step": "1"
                  },
                  "type": "int"
                },
                "name": "a"
              },
              {
                "identifier": "b",
                "dataType": {
                  "specs": {
                    "min": "0",
                    "max": "10",
                    "step": "1"
                  },
                  "type": "float"
                },
                "name": "b"
              }
            ],
            "type": "struct"
          },
          "size": "10"
        },
        "type": "array"
      },
      "name": "uuss",
      "accessMode": "rw",
      "required": false
    }
  ],
  "events": [
    {
      "outputData": [
        {
          "identifier": "LightStatus",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "工作状态"
        },
        {
          "identifier": "LightAdjustLevel",
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
          "name": "调光等级"
        },
        {
          "identifier": "LightVolt",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "4",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "工作电压"
        },
        {
          "identifier": "LightCurrent",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0",
              "unitName": "安培",
              "max": "9",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "工作电流"
        },
        {
          "identifier": "ActivePower",
          "dataType": {
            "specs": {
              "unit": "W",
              "min": "0",
              "unitName": "瓦特",
              "max": "1000",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "有功功率值"
        },
        {
          "identifier": "PowerRatio",
          "dataType": {
            "specs": {
              "unit": "pF",
              "min": "0.01",
              "unitName": "皮法",
              "max": "1",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "功率因数"
        },
        {
          "identifier": "PowerConsumption",
          "dataType": {
            "specs": {
              "unit": "kW·h",
              "min": "0",
              "unitName": "千瓦·时",
              "max": "2147483647",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "用电量"
        },
        {
          "identifier": "DrainVoltage",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "漏电压"
        },
        {
          "identifier": "TiltValue",
          "dataType": {
            "specs": {
              "unit": "°",
              "min": "0",
              "unitName": "度",
              "max": "90",
              "step": "1"
            },
            "type": "int"
          },
          "name": "倾斜角度值"
        },
        {
          "identifier": "ErrorPowerThreshold",
          "dataType": {
            "specs": {
              "unit": "W",
              "min": "0",
              "unitName": "瓦特",
              "max": "1000",
              "step": "1"
            },
            "type": "int"
          },
          "name": "故障功率门限"
        },
        {
          "identifier": "ErrorCurrentThreshold",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0.1",
              "unitName": "安培",
              "max": "9",
              "step": "0.1"
            },
            "type": "float"
          },
          "name": "故障电流门限"
        },
        {
          "identifier": "TiltThreshold",
          "dataType": {
            "specs": {
              "unit": "°",
              "min": "0",
              "unitName": "度",
              "max": "90",
              "step": "1"
            },
            "type": "int"
          },
          "name": "倾斜阈值"
        },
        {
          "identifier": "UnderVoltThreshold",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "1"
            },
            "type": "int"
          },
          "name": "欠压阈值"
        },
        {
          "identifier": "OverCurrentThreshold",
          "dataType": {
            "specs": {
              "unit": "A",
              "min": "0",
              "unitName": "安培",
              "max": "9",
              "step": "1"
            },
            "type": "int"
          },
          "name": "过流阈值"
        },
        {
          "identifier": "OverVoltThreshold",
          "dataType": {
            "specs": {
              "unit": "V",
              "min": "0",
              "unitName": "伏特",
              "max": "400",
              "step": "1"
            },
            "type": "int"
          },
          "name": "过压阈值"
        },
        {
          "identifier": "LightErrorEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "灯具故障使能"
        },
        {
          "identifier": "OverCurrentEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "过流告警使能"
        },
        {
          "identifier": "OverVoltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "过压告警使能"
        },
        {
          "identifier": "UnderVoltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "欠压告警使能"
        },
        {
          "identifier": "LeakageEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "漏电告警使能"
        },
        {
          "identifier": "OverTiltEnable",
          "dataType": {
            "specs": {
              "0": "关闭",
              "1": "打开"
            },
            "type": "bool"
          },
          "name": "倾斜告警使能"
        },
        {
          "identifier": "LampError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "灯具故障告警"
        },
        {
          "identifier": "OverCurrentError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "过流告警"
        },
        {
          "identifier": "OverVoltError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "过压告警"
        },
        {
          "identifier": "UnderVoltError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "欠压告警"
        },
        {
          "identifier": "OverTiltError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "倾斜告警"
        },
        {
          "identifier": "LeakageError",
          "dataType": {
            "specs": {
              "0": "正常",
              "1": "告警"
            },
            "type": "bool"
          },
          "name": "漏电告警"
        },
        {
          "identifier": "GeoLocation",
          "dataType": {
            "specs": [
              {
                "identifier": "longitude",
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
                "identifier": "latitude",
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
                "name": "纬度"
              },
              {
                "identifier": "altitude",
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
                "name": "高度"
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
              }
            ],
            "type": "struct"
          },
          "name": "地理位置"
        },
        {
          "identifier": "DeviceTime",
          "dataType": {
            "specs": {},
            "type": "date"
          },
          "name": "设备时间"
        },
        {
          "identifier": "ManufacturerName",
          "dataType": {
            "specs": {
              "length": "256"
            },
            "type": "text"
          },
          "name": "制造商名称"
        },
        {
          "identifier": "CurrentVoltage",
          "dataType": {
            "specs": {
              "item": {
                "type": "int"
              },
              "size": "10"
            },
            "type": "array"
          },
          "name": "wsp"
        },
        {
          "identifier": "Nono",
          "dataType": {
            "specs": {
              "item": {
                "specs": [
                  {
                    "identifier": "a",
                    "dataType": {
                      "specs": {
                        "min": "0",
                        "max": "10",
                        "step": "1"
                      },
                      "type": "int"
                    },
                    "name": "a"
                  },
                  {
                    "identifier": "b",
                    "dataType": {
                      "specs": {
                        "min": "0",
                        "max": "10",
                        "step": "1"
                      },
                      "type": "float"
                    },
                    "name": "b"
                  }
                ],
                "type": "struct"
              },
              "size": "10"
            },
            "type": "array"
          },
          "name": "uuss"
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
          "identifier": "alarmType",
          "dataType": {
            "specs": {
              "0": "防拆报警",
              "1": "防拆报警解除"
            },
            "type": "enum"
          },
          "name": "报警类型"
        }
      ],
      "identifier": "alarmEvent",
      "method": "thing.event.alarmEvent.post",
      "name": "报警事件",
      "type": "alert",
      "required": false
    }
  ]
}`
)

