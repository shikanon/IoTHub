<template>
    <el-menu :default-active="'0'"  @select="handleSelect">
         <!-- <template v-for="(item,i) in menuList.data"  >
            <el-menu-item :key="i" :index="i" > {{item.name}}</el-menu-item>
        </template>     -->
        <template v-for="(item,i) in menuList"  >
            <el-menu-item :key="i" :index="i + ''" > {{item}}</el-menu-item>
        </template>          
     </el-menu>
</template>

<script>
export default {
    props:{
        index:{
            type:String,
            default:"1"
        }
    },
    data(){
        return{
            menuList:{}
        }
    },
    created(){
        this.getMenus()
    },
    methods:{
        getMenus(){

            this.$API_SOTA.getIndexDtlList(this.index).then((res) => {
                this.menuList =  res.data 
                this.handleSelect(0)
            })           
         },

        handleSelect(index){
             let name = this.menuList[index]
            //根据id查询
            this.$emit('getInfoById',{id:index ,name:name})
         },
    }
}
</script>
<style  scoped>
    .el-menu{
        width:150px;
        overflow-y: auto;
        max-height: calc(100vh - 250px);
    }
    /* ::-webkit-scrollbar {
        width: 5px;
        height: 5px;
    }
    ::-webkit-scrollbar-thumb {
        background: #ccc;
        border-radius: 5px;
    }
    ::-webkit-scrollbar-track {
        background: transparent;
        border-radius: 5px;
    } */
</style>