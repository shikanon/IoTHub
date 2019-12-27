<template>
  <div>
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
        label="日期"
        width="120"
        :filters="[{text: '2016-05-01', value: '2016-05-01'}, {text: '2016-05-02', value: '2016-05-02'}, {text: '2016-05-03', value: '2016-05-03'}, {text: '2016-05-04', value: '2016-05-04'}]"
        :filter-method="filterHandler"
        >
        <template slot-scope="scope">{{ scope.row.date }}</template>
      </el-table-column>
      <el-table-column
        prop="name"
        label="姓名"
        width="120">
      </el-table-column>
      <el-table-column
        prop="address"
        label="地址"
        show-overflow-tooltip>
      </el-table-column>
       <el-table-column
        label="状态/启用状态"
        show-overflow-tooltip>
          <template slot-scope="scope">
            <el-badge  class="item" type="warning">
            </el-badge>
              <el-switch v-model="scope.row.active" active-text=""
               inactive-text="已禁用"></el-switch>
          </template>
      </el-table-column>
      <el-table-column
            fixed="right"
            label="操作"
            width="200">
            <template slot-scope="scope">
              <el-button @click="handleClick(scope.row)" type="text" size="small">查看 </el-button><el-divider direction="vertical"></el-divider>                
              <el-button type="text" size="small" @click="deleteClick(scope.$index, scope.row)">删除1</el-button><el-divider direction="vertical"></el-divider>     
              <el-popconfirm
                confirmButtonText='确定'
                cancelButtonText='取消'
                icon="el-icon-info"
                iconColor="red"
                title="确定删除吗"
              >
                <el-button slot="reference" type="text" size="small" >删除2</el-button>
              </el-popconfirm>
            </template>
          </el-table-column>
    </el-table>
    <div style="margin-top: 20px" >
      <el-checkbox v-model="selectFlg" ></el-checkbox>
      <el-button @click="handleSelection('1')" :disabled="!selectFlg">删除</el-button>
      <el-button @click="handledSelection('2')" :disabled="!selectFlg">禁用</el-button>
      <el-button @click="handleSelection('3')" :disabled="!selectFlg">启用</el-button>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        tableData: [{
          date: '2016-05-03',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄',
          state:0,
          active:true
        }, {
          date: '2016-05-02',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄',
          state:0,
          active:true

        }, {
          date: '2016-05-04',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄',
          state:1,
          active:true
        }, {
          date: '2016-05-01',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄',
          state:1,
          active:false 
        }, {
          date: '2016-05-08',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄',
          state:0,
          active:false 
        }, {
          date: '2016-05-06',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄上海市普陀区金沙江路 1518 弄上海市普陀区金沙江路 1518 弄',
          state:0
        }, {
          date: '2016-05-07',
          name: '王小虎',
          address: '上海市普陀区金沙江路 1518 弄',
          state:1,
          active:true 
        }],
        multipleSelection: []
      }
    },
    computed:{
      selectFlg:function(){
          return this.multipleSelection.length > 0
      }
    },

    methods: {
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
      }
    }
  }
</script>