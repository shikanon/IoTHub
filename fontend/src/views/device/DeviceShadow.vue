<template>
    <div name="device-shadow">
        <el-button type="primary" :disabled="status === 'UNACTIVE'" @click="dialogVisiable = true">更新影子</el-button>
        <div class="shadow-box">
            <div class="title">
                <p>最近更新时间:</p><p class="update-time"> {{updateTime}}</p>
            </div>
            <div class="content">
                <JsonView :data="defaultData"></JsonView>
            </div>
        </div>
        <el-dialog title="更新设备影子" :visible.sync="dialogVisiable" width="30%">
            <div class="shadow-box">
                <div class="title">
                    <p>"reported":</p>
                </div>
                <div class="content">
                     <VueJsonEditor v-model="reportedJson"
                     :showBtns="false"
                     :mode="'code'" 
                     lang="zh"
                      @json-change="onReportedJsonChange"
                      @has-error="onError('reported')"></VueJsonEditor>

                </div>
            </div>
            <div class="shadow-box">
                <div class="title">
                    <p>"desired":</p>
                </div>
                <div class="content">
                      <VueJsonEditor v-model="desiredJson"
                     :showBtns="false"
                     :mode="'code'" 
                      @json-change="ondesiredJsonJsonChange"
                      @has-error="onError('desired')">
                    </VueJsonEditor>

                </div>
            </div>
            <span slot="footer" class="dialog-footer">
                <el-button @click="updateShadow" type="primary"> 确定</el-button>
                <el-button @click="dialogVisiable = false">关 闭</el-button>
            </span>
        </el-dialog>   

    </div>
</template>

<script>
import { create } from 'domain'
export default {
    props:{
        status:{
            type:String,
            default:''
        }
    },
    data(){
        return{
            defaultData:{
                    "state": {
                        "reported": {},
                        "desired": {}
                    },
                    "metadata": {
                        "reported": {},
                        "desired": {}
                    },
                    "timestamp": 0,
                    "version": 0
            },
            updateTime:'2020/01/10 10:38:18',
            dialogVisiable:false,
            codeEditor:{},
            reportedJson:{},
            desiredJson:{},
            errMsg:""

        }
    },
     mounted: function () {
        let container = document.getElementById('jsoneditor')
        
        let options = {
            mode: 'code',
            modes: ['code'],
            onError: function(err) {
                alert(err.toString());
            }
        }
        this.codeEditor = new this.$jsoneditor(container, options);
        this.codeEditor.set(this.defaultData)
      },

    create(){
        this.init()
    },
    methods:{
        init(){
            this.reportedJson = this.defaultData.state.reported
            this.desiredJson = this.defaultData.state.desired
        },
   
        onReportedJsonChange(value){
            this.reportedJson = value
        },

        ondesiredJsonJsonChange(value){
            this.desiredJson = value
        },

        onError(type){
           // this.errMsg = `${type}数据格式不正确`
        },
        updateShadow(){
            
            if( Object.keys(this.reportedJson).length === 0 ){
                this.$message.error('reported数据为空或数据格式不正确');
                return false
            }else if(Object.keys(this.desiredJson).length === 0){
                this.$message.error('desired数据为空或数据格式不正确');
                return false
            }

            this.defaultData.state.reported = this.reportedJson 
            this.defaultData.state.desired =  this.desiredJson 
        }
    }
}
</script>

<style scoped>
.shadow-box{
    margin-top: 10px;
}
.title{
    background-color: rgb(249, 249, 249);
    border: 1px solid rgb(235, 236, 236);
    display: flex;
    align-self: center;
}

.title p{
    text-indent: 16px;
    color: #373d41;
    font-weight: 700;
}

.title .update-time{
    color: #666;
    font-weight: 100;
}

.content{
    border: 1px solid rgb(235, 236, 236);
    border-top: 0px;
    min-height: 150px;
}

.jsoneditor-menu{
    display: none;
}


</style>