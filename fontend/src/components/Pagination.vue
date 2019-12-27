<template>
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage "
      :page-sizes="pageSizes "
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total">
    </el-pagination>
</template>
<script>
  export default {
      props: {
              currentPage: {
                type: Number,
                default: 1
              },
               pageSize: {
                type: Number,
                default: 10
              },
               total: {
                type: Number,
                default: 0
              },
               pageSizes: {
                type: Array,
                default: [10,30,50,100]
              }
      },
   
    data() {
   
        return {
            
        }
    },
    methods: {
      handleSizeChange(val) {
        console.log(`每页 ${val} 条`);
        this.$emit('handleSizeChange',val)
      },
      handleCurrentChange(val) {
        console.log(`当前页: ${val}`);
        this.$emit('handleCurrentChange',val)
     
      }
    },
   
  }
</script>

<style scoped >
  .el-pagination{
    margin-top: 20px;
  }
</style>
