 <template>
    <el-table
      :data="tableData"
      style="width: 100%">
      <el-table-column
        prop="date"
        label="订阅类型"
        width="180">
      </el-table-column>
      <el-table-column
        prop="name"
        label="订阅消息"
        >
      </el-table-column>
      <el-table-column
        prop="address"
        label="创建时间"
        width="380">
      </el-table-column>
      <el-table-column
      
        label="操作"
        width="200"
        >
            <template slot-scope="scope" >
                 <el-button type="text" size="small" @click="edit(scope.row)">编辑</el-button>

            </template>
          </el-table-column>
    </el-table>
  </template>

  <script>
    export default {
      data() {

        return {
          tableData: [{
            date: 'AMQP',
            name: '设备上报消息,设备状态变化通知,设备生命周期变更,物模型历史数据上报,固件升级状态通知',
            address: '2020/02/06 13:20:11'
          }, {
            date: 'AMQP',
            name: '设备上报消息,设备状态变化通知,设备生命周期变更,物模型历史数据上报,固件升级状态通知',
            address: '2020/02/06 13:20:11'
          }, {
           date: 'AMQP',
            name: '设备上报消息,设备状态变化通知,设备生命周期变更,物模型历史数据上报,固件升级状态通知',
            address: '2020/02/06 13:20:11'
          }]
        }
      }
    }
  </script>