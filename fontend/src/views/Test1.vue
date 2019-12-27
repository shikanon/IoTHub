<template>
   <div>
      <el-table
        ref="singleTable"
        :data="tableData"
        highlight-current-row
        @current-change="handleCurrentChange"
        style="width: 100%">
        <el-table-column
          type="index"
          width="50">
        </el-table-column>
        <el-table-column
          property="id"
          label="序号"
          width="120">
        </el-table-column>
        <el-table-column
          property="name"
          label="姓名"
          width="120">
        </el-table-column>
        <el-table-column
          property="address"
          label="地址">
        </el-table-column>
      </el-table>
      <div style="margin-top: 20px">
        <el-button @click="setCurrent(tableData[1])">选中第二行</el-button>
        <el-button @click="setCurrent()">取消选择</el-button>
        <el-button @click="getHall()">测试</el-button>

      </div>
   </div>
</template>

<script>
  export default {
    data() {
      return {
        tableData: [],
        currentRow: null
      }
    },
    created(){
        this.getHall()
    },

    methods: {
      setCurrent(row) {
        this.$refs.singleTable.setCurrentRow(row);
      },
      handleCurrentChange(val) {
        this.currentRow = val;
      },
      getHall(){
          this.$API.getHall(1,9999).then((res) => {
                console.log(res)
                this.tableData = res.data.objects
            })
      }
    }
  }
</script>