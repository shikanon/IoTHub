<template>
  <div>
     <el-row >
        <el-input
            v-model="identifier"
            :placeholder="type === 'event' ? '请输入事件标识符' :'请输入服务标识符'"
            class="search-input"
            clearable
            size="medium"
            @keyup.enter.native="init()"
            >
          <i slot="suffix" class="el-input__icon el-icon-search" @click="init()"></i>
        </el-input> 
        <el-select v-if="type === 'event'" v-model="eventType" placeholder="全部类型"  size="medium" @change="changeEventType" >
                <el-option
                v-for="item in eventTypes"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
          </el-select>
          <el-select v-model="time" :placeholder="times[0].label"   size="medium" @change="changeTime" >
                <el-option
                v-for="item in times"
                :key="item.value"
                :label="item.label"
                :value="item.value">
                </el-option>
          </el-select>
           <el-date-picker v-if="time === '4'"
                    v-model="timeArr"
                    @change="changeTimeArr"
                    type="datetimerange"
                    range-separator="至"
                    align="right"
                    start-placeholder="起始日期"
                    end-placeholder="结束日期"
                    size="medium">
           </el-date-picker>
    </el-row> 
    
    <el-table
      :data="tableData"
      tooltip-effect="dark"
      style="width: 100%"
      >
     <el-table-column
            label="时间"
            width="200"  
            >
            <template slot-scope="scope">{{ scope.row.time }}</template>
        </el-table-column>
        <el-table-column
            prop="identifier"
            label="标识符"
            width="200">
        </el-table-column>
        <el-table-column
            :label="type === 'event' ? '事件名称' : '服务名称'"   width="200">
            <template slot-scope="scope">{{ scope.row.name  }}</template>
        </el-table-column>
            
        <el-table-column
            label="事件类型" v-if="type === 'event'" >
            <template slot-scope="scope">{{ scope.row.event_type === 'info' ? '信息':  scope.row.event_type === 'alert' ?'告警':'故障'}}</template>
        </el-table-column>
        <el-table-column
            label="输入参数"  v-else >
            <template slot-scope="scope">{{ scope.row.input_data }}</template>
        </el-table-column>
        <el-table-column
            label="输出参数"  >
            <template slot-scope="scope">{{ scope.row.output_data }}</template>
        </el-table-column>
    </el-table>
    <el-button @click="init" v-if="isBtnShow">加载更多</el-button>
  
  </div>
</template>

<script>
  export default {
    props:{
      deviceId:{
        type:Number,
        default:0
      },
      type:{
        type:String,
        default:''
      }
    },

     watch:{
        //监听type,若发生变化，重新查询设备批次
        deviceId:{  
            handler:function(val,oldval){ 
                if(val!=oldval){
                    this.$nextTick(()=>{
                        this.init()
                    })
                }
            },  
            immediate:true,//关键
            deep:true
          },
      },
  
    data() {
      return {
        tableData: [],
        eventTypes: [{label:'全部类型',value:''},{label:'信息',value:'info'},{label:'告警',value:'alert'},{label:'故障',value:'error'}],
        identifier:'',
        eventType:'',
        times:[{label:'1小时',value:'1'},{label:'24小时',value:'2'},{label:'7天',value:'3'},{label:'自定义',value:'4'}],
        time:'1',
        currentPage:1,
        startTime:0,
        endTime:0,
        timeArr:[],
        isBtnShow:false,
        next_time:0
      }
    },
 

    created(){
        this.changeTime("1")   
        this.init()
    },
    methods: {
      init(){
               
          if(this.type === 'event'){
              this.getDiviceList()
          }else{
              this.getServiceList()
          }
      },
       getDiviceList(){
           this.$API_IOT.getEventList(this.deviceId,this.currentPage,this.startTime,this.next_time === 0 ? this.endTime :this.next_time ,this.eventType,this.identifier).then((res) => {
            
             if(res.data.data.data_list != null){
               this.tableData = this.tableData.concat(res.data.data.data_list)
             }else{
               this.tableData = []
             }
             if(res.data.data.next_time !== 0 ){
               this.isBtnShow = true     
             }else{
               this.isBtnShow = false
             }
              this.next_time = res.data.data.next_time    
 
            })
           
       },
       getServiceList(){
           this.$API_IOT.getServiceList(this.deviceId,this.currentPage,this.startTime,this.next_time === 0 ? this.endTime :this.next_time ,this.identifier).then((res) => {
                if(res.data.data.data_list != null){ 
                  this.tableData = this.tableData.concat(res.data.data.data_list)
                }else{
                  this.tableData = []
                }
                if(res.data.data.next_time !== 0 ){
                  this.isBtnShow = true     
                }else{
                  this.isBtnShow = false
                }

                this.next_time = res.data.data.next_time    

            })
           
       },
      
       changeTime(value){
           let nowDate  = new Date()
           this.endTime = nowDate.getTime().toString().substr(0,10)    
            if(value === "1"){//1小时
                this.startTime = nowDate.setHours(nowDate.getHours() -1).toString().substr(0,10)
            }else if(value === "2"){//24小时
                this.startTime = nowDate.setDate(nowDate.getDate() - 1).toString().substr(0,10)
            }else if(value === "3"){//7天
                this.startTime = nowDate.setDate(nowDate.getDate() - 7).toString().substr(0,10)
            }else {//自定义
              //1-清空
              this.timeArr = []
              //2-赋值 
              let end = new Date().getTime()
              let start  = new Date(new Date().setHours(nowDate.getHours() -1))
              this.timeArr = [start,end]
              this.startTime = start.getTime().toString().substr(0,10)
              this.endTime = end.toString().substr(0,10)
            }
            this.init()
       },
       changeTimeArr(value){
            this.startTime = value[0].getTime().toString().substr(0,10)
            this.endTime = value[1].getTime().toString().substr(0,10)
            this.next_time = 0 
            this.init()
       },

       changeEventType(value ){
         this.eventType = value 
         this.next_time = 0 
         this.init()
       }
      
    }
  }
</script>

<style scoped>
 .search-input{width: 300px}

</style>