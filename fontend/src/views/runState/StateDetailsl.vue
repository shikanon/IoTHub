<template>
    <div name="state-detailsl">
        <div class="query-params">
            <div class="time-type">
                <el-select v-model="type" placeholder="请选择" @change="changeQueryTime"  size="medium">
                    <el-option label="1小时" value="1"> </el-option>
                    <el-option label="24小时" value="2"> </el-option>
                    <el-option label="7天" value="3"> </el-option>
                    <el-option label="自定义" value="4"> </el-option>           
                </el-select>
                <el-date-picker v-if="type === '4'"
                    v-model="timeArr"
                    @change="changeTime"
                    type="datetimerange"
                    range-separator="至"
                    align="right"
                    start-placeholder="开始日期"
                    end-placeholder="结束日期"
                    size="medium">
                </el-date-picker>
            </div> 
            <div>
                <el-tabs v-model="activeName" type="card"  size="medium">
                    <el-tab-pane label="表格" name="table"></el-tab-pane>
                    <el-tab-pane label="图标" name="echart"></el-tab-pane>
                </el-tabs>
            </div>
        </div>
        <div>
            <div v-if="activeName === 'table'">
                  <el-table 
                    :data="data"
                    highlight-current-row
                    style="width: 100%">
                        <el-table-column
                            prop="function_type"
                            label="时间"
                            width="120">
                            
                        </el-table-column>
                        <el-table-column
                            prop="function_name"
                            label="原始值"
                            width="150">
                        </el-table-column>                  
                    </el-table>
            </div>
            <div v-else>
                
            </div>
        </div>    
        
    </div>   
</template>

<script>
export default {
    data(){
        return {
            type:"1",
            timeArr:[],
            startTime:"",
            endTime:"",
            activeName:'table',
            data:[],
            startFlg:false,
            endFlg:false
        }
    },

    watch:{     
       startTime: {
            handler: function(val, oldVal){ 
                if(val !== oldVal){
                  this.startFlg = true
                  
                }else{
                    this.startFlg = false
                }   

                if(this.startFlg  || this.endFlg){
                     this.refreshData()
                }
            },
            // 深度观察监听
            deep: true
        } ,

        endTime: {
            handler: function(val, oldVal){ 
                if(val !== oldVal){
                  this.endFlg = true
                  
                }else{
                    this.endFlg = false
                }  
                
                if(this.startFlg  || this.endFlg){
                    this.refreshData()
                }
            },
            // 深度观察监听
            deep: true
        },
        refreshFlg:function(){
            this.refreshData()
        }
      
    },
    computed:{
        refreshFlg:function(){
            console.log('1111111111')
            return this.startFlg && this.endFlg
        }
    },
    created(){
        this.changeQueryTime("1")
    },
    methods:{
         changeQueryTime(value){

           let nowDate  = new Date()
           this.endTime = nowDate.getTime()
            if(value === "1"){
                this.startTime = nowDate.setHours(nowDate.getHours() -1)
            }else if(value === "2"){
                this.startTime = nowDate.setDate(nowDate.getDate() - 1)
            }else if(value === "3"){
                this.startTime = nowDate.setDate(nowDate.getDate() - 7)
            }else {
                this.getDefaultTime()
            }
        },

        getDefaultTime(){
            //1-清空
            this.timeArr = []
            //2-赋值 
            let end = new Date()
            let start  = new Date(new Date().setDate(end.getDate() - 7))
            this.timeArr = [start,end]
            console.log(this.timeArr)
        },
        changeTime(value){
          this.startTime = value[0].getTime()
          this.endTime = value[1].getTime()
        },
        refreshData(){
           // alert(`查询数据开始时间${this.startTime}结束时间${this.endTime}`)
        }
    }
}
</script>


<style scoped>
    .query-params{
      display: flex;
      justify-content:space-between;
    }
    .el-select {width: 100px;}
   
</style>