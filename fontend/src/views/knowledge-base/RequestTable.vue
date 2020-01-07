<template>
    <el-table
        :data="tableData"
            style="width: 100%">
                <el-table-column                                    
                    label="名称"
                    prop="ApiParameterName"
                    width="180"
                    >
                </el-table-column>
                <el-table-column                                    
                    label="类型"
                    prop="ParameterType"
                    width="120">
                </el-table-column>
                <el-table-column                                    
                    label="是否必须"
                    width="100">
                    <template slot-scope="scope">{{scope.row.Required === 'REQUIRED'? '是':'否'}}</template>
                </el-table-column>
                <el-table-column                                    
                    label="描述"
                    prop="Description"
                    >
                </el-table-column>`
    </el-table>
</template>
<script>
export default {
    props:{
        tableData:{
            type:Array,
            default:() => ([])
        }
    }
}
</script>