<template>
  <div>
     <el-row >
        <el-input
            v-model="search"
            placeholder='请输入事件标识符'
            class="search-input"
            clearable
            size="medium"
            @keyup.enter.native="getDeviceList(1)"
            >
          <i slot="suffix" class="el-input__icon el-icon-search" @click="getDeviceList(1)"></i>
        </el-input> 
        <el-select v-model="eventType" placeholder="全部类型"    size="medium">
                <el-option
                v-for="item in eventTypes"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
          </el-select>
          <el-select v-model="time" :placeholder="times[0].label" multiple :multiple-limit="1"  size="medium">
                <el-option
                v-for="item in times"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
          </el-select>
    </el-row> 
    
    <el-table
      :data="tableData"
      tooltip-effect="dark"
      style="width: 100%"
      >
     <el-table-column
            label="时间"
            width="200"  
            >
            <template slot-scope="scope">{{ scope.row.DeviceName }}</template>
        </el-table-column>
        <el-table-column
            prop="ProductName"
            label="标识符"
            width="200">
        </el-table-column>
        <el-table-column
            label="事件名称"  >
            <template slot-scope="scope">{{ scope.row.LastOnlineTime  }}</template>
        </el-table-column>
            
        <el-table-column
            label="事件类型"  >
            <template slot-scope="scope">{{ scope.row.LastOnlineTime  }}</template>
        </el-table-column>
        <el-table-column
            label="输出参数"  >
            <template slot-scope="scope">{{ scope.row.LastOnlineTime  }}</template>
        </el-table-column>
    </el-table>
  
  </div>
</template>

<script>
  export default {
  
    data() {
      return {
        tableData: [],
        eventTypes: [{label:'信息',value:'info'},{label:'告警',value:'alert'},{label:'故障',value:'error'}],
        search:'',
        eventType:'',
        times:[{label:'1小时',value:'1'},{label:'24小时',value:'2'},{label:'7天',value:'3'},{label:'自定义',value:'4'}],
        time:'1'
      }
    },
 

    created(){
        this.getDiviceList()
    },
    methods: {
       getDiviceList(){
            this.tableData =[
                {
                    "UtcActiveTime": "",
                    "LastOnlineTime": "",
                    "Status": "DISABLE",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 14:00:09",
                    "IotId": "DFTwwlvDhW8DmKJP2EuK000100",
                    "UtcCreateTime": "2019-12-19T06:00:09.000Z",
                    "CategoryKey": "GarbageOverflowingDetection",
                    "ActiveTime": "",
                    "UtcLastOnlineTime": "",
                    "NodeType": 0,
                    "ProductKey": "a1ELejzj0h9",
                    "DeviceName": "DFTwwlvDhW8DmKJP2EuK"
                },
                {
                    "UtcActiveTime": "",
                    "LastOnlineTime": "",
                    "Status": "UNACTIVE",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 14:00:09",
                    "IotId": "l8s2H6EOs46qlMW7jQKd000100",
                    "UtcCreateTime": "2019-12-19T06:00:09.000Z",
                    "CategoryKey": "GarbageOverflowingDetection",
                    "ActiveTime": "",
                    "UtcLastOnlineTime": "",
                    "NodeType": 0,
                    "ProductKey": "a1ELejzj0h9",
                    "DeviceName": "l8s2H6EOs46qlMW7jQKd"
                },
                {
                    "UtcActiveTime": "",
                    "LastOnlineTime": "",
                    "Status": "ONLINE",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 13:59:24",
                    "IotId": "kw57odVwBBxmesgfpBGF000100",
                    "UtcCreateTime": "2019-12-19T05:59:24.000Z",
                    "CategoryKey": "GarbageOverflowingDetection",
                    "ActiveTime": "",
                    "UtcLastOnlineTime": "",
                    "NodeType": 0,
                    "ProductKey": "a1ELejzj0h9",
                    "DeviceName": "q123",
                    "Nickname": "123123"
                },
                {
                    "UtcActiveTime": "2019-12-19T02:30:48.731Z",
                    "LastOnlineTime": "2019-12-19 10:30:48",
                    "Status": "OFFLINE",
                    "ProductName": "test",
                    "AliyunCommodityCode": "iothub_senior",
                    "CreateTime": "2019-12-19 10:30:10",
                    "IotId": "gQj9TNrlkqvU1UgtcLI9000100",
                    "UtcCreateTime": "2019-12-19T02:30:10.000Z",
                    "CategoryKey": "Lighting",
                    "ActiveTime": "2019-12-19 10:30:48",
                    "UtcLastOnlineTime": "2019-12-19T02:30:48.731Z",
                    "NodeType": 0,
                    "ProductKey": "a1zJA7k9sjd",
                    "DeviceName": "test_device",
                    "Nickname": "test_device"
                }
            ]
       }
      
    }
  }
</script>

<style scoped>
 .search-input{width: 300px}

</style>