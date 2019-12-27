<template>
    <div name="add-label"> 
        <div><span>{{type === 'product' ? '产品':'设备'}}标签</span></div>
        <div v-for="(item,index) in label" :key="item.key" class="label-box">
            <el-input placeholder="请输入标签key"  v-model="item['key']"></el-input>
            <el-input placeholder="请输入标签value"  v-model="item['value']"></el-input>
            <el-button type="text" size="small" @click="delLabel(index)">删除</el-button>
         </div>
         <el-button type="text" size="small" @click="addLabel">+新增标签</el-button>
    </div>
</template>

<script>
    export default {
        props:{
            labelArr:{
                type:Array,
                default:() => []
            },
            type:{
                type:String,
                default:''
            }
        },
        data() {
            return {
                label:[]
            }
        },
        created(){
            if(this.labelArr.length === 0){
                this.addLabel()
            }else{
                this.label = this.labelArr
            }
        },
        methods:{
            addLabel(){
                this.label.push({key:'',value:''})
            },
            delLabel(index){
                this.label = this.label.filter((item,key) => key !== index)
            },
            save(){
                this.$emit('save',this.label )
            }
        }
    }
    </script>
<style scoped>
    .label-box{
        display: flex;
        padding-top: 10px;
    }

    .label-box .el-input{
        padding-right: 10px;
    }
</style>