<template>
  <div>
     <el-row >
        <el-select v-model="product" placeholder="请选择产品"  size="medium">
                <el-option
                v-for="item in products"
                :key="item.ProductKey"
                :label="item.ProductName"
                :value="item.ProductName">
                </el-option>
          </el-select>
          <el-input
            v-model="search"
            placeholder='请输入Device查询'
            class="search-input"
            clearable
            size="medium"
            @keyup.enter.native="getDeviceList(1)"
            >
          <i slot="suffix" class="el-input__icon el-icon-search" @click="getDeviceList(1)"></i>
        </el-input> 
    </el-row> 
    
    <el-table
      ref="multipleTable"
      :data="tableData"
      tooltip-effect="dark"
      style="width: 100%"
      @selection-change="handleSelectionChange">
    
        <el-table-column
            type="selection"
            width="55">
        </el-table-column>     
        <el-table-column
            label="DeviceName/备注名称"
            width="200"  
            >
            <template slot-scope="scope">{{ scope.row.DeviceName }}</template>
        </el-table-column>
        <el-table-column
            prop="ProductName"
            label="设备所属产品"
            width="200">
        </el-table-column>
       
        <el-table-column
            label="状态/启用状态"
            show-overflow-tooltip
            :filters="[{text: '已禁用', value: 'DISABLE'}, {text: '未激活', value: 'UNACTIVE'},
            {text: '在线', value: 'ONLINE'}, {text: '离线', value: 'OFFLINE'}]"
            :filter-method="filterHandler">
            <template slot-scope="scope">
                <el-button 
                    :type="scope.row.Status === 'ONLINE' ? 'success': scope.row.Status === 'UNACTIVE' ? 'info': scope.row.Status === 'DISABLE' ? 'danger':'warning'"  
                    circle 
                    size="mini"
                    class="statu-circle">
                </el-button>
                 <span> {{ scope.row.Status | deviceStatusFilter}}</span>
                  <el-switch  v-if="scope.row.Status === 'UNACTIVE'"  value ="scope.row.Status === 'UNACTIVE'"></el-switch>  
            </template>
        </el-table-column>
        <el-table-column
            label="最后上线时间"  >
            <template slot-scope="scope">{{ scope.row.LastOnlineTime  }}</template>
        </el-table-column>
    </el-table>
     <el-row>
              <el-button @click="cancel">取消</el-button>
              <el-button type="primary" @click="submit">确定</el-button>
      </el-row>
  </div>
</template>

<script>
  export default {
    props:{
      products:{
        type:Array,
        default: () => []
      }
    },
    data() {
      return {
        tableData: [],
        multipleSelection: [],
        product:'',
        search:''
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
       },
      toggleSelection(rows) {
        if (rows) {
          rows.forEach(row => {
            this.$refs.multipleTable.toggleRowSelection(row);
          });
        } else {
          this.$refs.multipleTable.clearSelection();
        }
      },
      handleSelectionChange(val) {
        this.multipleSelection = val;
      },
      filterHandler(value, row, column) {
        const property = column['property'];
        return row[property] === value;
      },
      deleteSelection(type){
          console.log(this.multipleSelection)
      },
      cancel(){
          this.$emit('close')
      },
      submit(){
          this.$emit('submit',this.multipleSelection)

      }
    }
  }
</script>

<style scoped>
 .search-input{width: 300px}

</style>