<template>
    <el-button  :type="type" :size="size"
        v-clipboard:copy="content"
        v-clipboard:success="onCopy"
        >{{btnTxt}}</el-button>

</template>
 <script>

export default {
    props: {
        type:{
            type: String,
            default: 'text'
        },
        size:{
            type: String,
            default: 'medium'
        },
        content:{
            type: String,
            default: ''
        },
        btnTxt:{
            type:String,
            default:'复制'
        }
   
    },
    data () {
        return {
        
        }
   },

    methods:{
        onCopy(){
            this.$message({
                message: '复制成功',
                type: 'success'
            });
        } 
    }


}
</script>             
   
 <style scoped>
              
</style>