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
                    @change="changeTimeArr"
                    type="datetimerange"
                    range-separator="至"
                    align="right"
                    start-placeholder="开始日期"
                    end-placeholder="结束日期"
                    size="medium">
                </el-date-picker>
            </div> 
            <div>
                <el-tabs v-model="activeName" type="card"  size="medium"  @tab-click="handleClick">
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
                height="400"
                style="width: 100%">
                    <el-table-column
                        prop="Time"
                        label="时间"
                        >
                        
                    </el-table-column>
                    <el-table-column
                        prop="Value"
                        label="原始值"
                        width="150">
                    </el-table-column>                  
                </el-table>
                <el-button @click="query" v-if="isBtnShow">加载更多</el-button>
            </div>
            <div v-else>
                <div ref="myChart" :style="{width: '100', height: '300px'}"></div>
            </div>


        </div>    
        
    </div>   
</template>

<script>
export default {
    props:{
        deviceId:{
            type:Number,
            default:0
        },
        identifier:{
            type:String,
            default:''
        },
           name:{
            type:String,
            default:''
        },
    },
    data(){
        return {
            type:"1",
            timeArr:[],
            startTime:0,
            endTime:0,
            activeName:'table',
            data:[],
            next_time:0,
            isBtnShow:false
          
        }
    },

    watch:{               
        //identifier,若发生变化，重新查询设备批次
        identifier:{  
            handler:function(val,oldval){ 
                if(val!=oldval){
                    this.$nextTick(()=>{
                        this.query()
                    })
                }
            },  
            immediate:true,//关键
            deep:true
          },
    },
   
    created(){

        this.changeQueryTime("1")
    },
    methods:{

       
         changeQueryTime(value){

           let nowDate  = new Date()
           this.endTime = nowDate.getTime().toString().substr(0,10)
            if(value === "1"){
                this.startTime = nowDate.setHours(nowDate.getHours() -1).toString().substr(0,10)
            }else if(value === "2"){
                this.startTime = nowDate.setDate(nowDate.getDate() - 1).toString().substr(0,10)
            }else if(value === "3"){
                this.startTime = nowDate.setDate(nowDate.getDate() - 7).toString().substr(0,10)
            }else {
                 //1-清空
              this.timeArr = []
              //2-赋值 
              let end = new Date().getTime()
              let start  = new Date(new Date().setHours(nowDate.getHours() -1))
              this.timeArr = [start,end]
              this.startTime = start.getTime().toString().substr(0,10)
              this.endTime = end.toString().substr(0,10)
            }
            
            this.query()
        },

       
        changeTimeArr(value){
          this.startTime = value[0].getTime().toString().substr(0,10)
          this.endTime = value[1].getTime().toString().substr(0,10)
          this.query()
        },


        query(){
            //console.log(`查询数据,${this.deviceId},${this.identifier}开始时间${this.startTime}结束时间${this.endTime}`)
            this.identifier = 'LightAdjustLevel'

            this.$API_IOT.getRunStateDtl(this.deviceId,this.identifier,this.startTime,this.next_time !== 0 ?this.next_time : this.endTime,'dev').then((res) => {
                if(res.data.data.data_list != null ){
                     this.data = this.data.concat(res.data.data.data_list)
                }else{
                    this.data = []
                }

                if(res.data.data.next_time !== 0 ){
                  this.isBtnShow = true     
                }else{
                  this.isBtnShow = false
                }
                this.next_time = res.data.data.next_time 
            
             })
        },
        handleClick(tab, event) {
            if(tab.name === 'echart'){            
                this.drawLine()   
            }
        },

        drawLine(){

            this.$nextTick(()=>{

                // 基于准备好的dom，初始化echarts实例
                let myChart =this.$echarts.init(this.$refs.myChart)
                let yAxis_data= this.data.map(item => { return item.Value })
                let xAxis_data= this.data.map(item => {  return  item.Time})
    
                //赋值
                let option = {
                    // title: {
                    //     text: 'ECharts 入门示例'
                    // },
                    tooltip: {},
                    legend: {
                        data:[this.name]
                    },
                    xAxis: {
                        data: xAxis_data
                    },
                    yAxis: {},
                    series: [{
                        name: this.name,
                        type: 'line',
                        data: yAxis_data
                    }]
                }
                // 绘制图表
                myChart.setOption(option) 

                })
            
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