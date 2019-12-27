<template>
    <div name="attribute-list"> 
        <div >
              <p>
                此产品分类包含了 31 项标准功能定义，选择后系统将为您自动创建必选功能，您可以根据需要来添加可选功能或自定义功能
              </p>
        </div>
        <el-table 
          :data="attributes"
          highlight-current-row
          style="width: 100%">
          <el-table-column
            prop="type"
            label="功能类型"
            width="150">
             
          </el-table-column>
           <el-table-column
            prop="name"
            label="功能名称"
            width="150">
          </el-table-column>
           <el-table-column
            prop="symbal"
            label="标识符"
            width="100">
          </el-table-column>
          <el-table-column
            prop="dataType"
            label="数据类型"
            width="100">
          </el-table-column>
        </el-table>
       
      </div>
  </template>

  <script>
  
    export default {
      props:{
        attributes:{
          type:Array,
          default: () => []
        }
      },
      data() {
        return {
        }
      },
   
    
      created(){
      
      },
      methods:{
       

            
      }
    }
  </script>

  <style scoped>
    .search-input{
      width:200px;
    }
  </style>