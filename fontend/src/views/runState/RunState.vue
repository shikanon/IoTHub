<template>
    <div name="run-state">
        <div class="state-type">
            <el-radio-group v-model="radio">
              <el-radio :label="'chart'">图表</el-radio>
              <el-radio :label="'table'">表格</el-radio>
            </el-radio-group>
            <span>实时刷新</span>
            <el-switch v-model="autoRefresh">
            </el-switch>
        </div>
          <div v-if="radio === 'table'">
            <el-table
                ref="singleTable"
                :data="cards"
                highlight-current-row
                style="width: 100%">
                
                <el-table-column
                  property="Name"
                  label="属性名称"
                  width="300">
                </el-table-column>
                <el-table-column
                  property="DataType"
                  label="数据类型"
                  width="200">
                </el-table-column>
                 <el-table-column
                  property="Time"
                  width="200">
                </el-table-column>
                <el-table-column
                  label="最新值">
                </el-table-column>
                  <el-table-column
                  property="address"
                  label="期望值">
                    <template slot-scope="scope">
                      <span>{{scope.Value}}</span>
                    </template>
                </el-table-column>
                <el-table-column
                  fixed="right"
                  label="操作"
                  width="100">
                  <template slot-scope="scope">
                     <el-button type="text" size="small" @click="queryDataByName(scope.row.Name)">查看数据</el-button>
                  </template>
              </el-table-column>
            </el-table>
         </div>
          <el-col :span="6" v-for="(card,index)  in cards" :key="index" v-else>
            <el-card  shadow="hover" class="card  ">
                  <div style="display: flex;justify-content: space-between;">
                     {{card.Name}}
                     <el-button type="text" size="small" @click="queryDataByName" >查看数据</el-button>
                  </div>
                  <p style="font-size: 18px;color: rgb(51, 51, 51);height:21px;" >{{ card.val | cardValFilter(card.Unit) }}  </p>
                    <!-- <el-tooltip class="item" effect="dark" :content="card.Value" placement="top-start">
                        <p style="color: #999">{{card.Value ? card.Value : '--' }}</p>
                    </el-tooltip>            -->
                     <p style="color: #999">{{card.Value ? card.Value : '--' }}</p>
            </el-card>
          </el-col>
         <el-dialog title="查看数据" :visible.sync="dialogTableVisible" width="40%">
                <StateDetailsl></StateDetailsl>
          </el-dialog>     
    </div>  
 </template>  
<script>
import StateDetailsl from './StateDetailsl'
  export default {
    components:{StateDetailsl},
      props:{
        deviceId:{
          type:Number,
          default:0
        }
      },
      data() {
        return {
          cards:[],
          radio:'chart',
          autoRefresh:false,
          tableData:[],
          dialogTableVisible:false,
          intervalId:null,
        
        }
      },
      watch:{    
          autoRefresh:function(){             
              if(this.autoRefresh){
                this.intervalId =  setInterval(() => {
                    this.getCardData()
                }, 1*1000);  
              }else{
                clearInterval(this.intervalId)
              }
                      
          }
         
      },
      created(){  
        this.getCardData()        
      },
      destoryed() {
        clearInterval(this.intervalId)
      },
      methods:{
       
       
  
        handleClick(tab, event) {
          //console.log(tab, event);
        },
        
    
        getCardData(){
            this.$API_IOT.getRunState(this.deviceId).then((res) => {
                console.log(res.data)
                if(res.data.status  === 'Y'){
                  this.cards = res.data.data    
                }else{
                  this.$message.error(res.message);
                }              
             })
        },
        queryDataByName(){
          this.dialogTableVisible = true 
        },

       
         
         
      }
    }
  </script>

  <style scoped>
    .page-header{
      margin-bottom: 20px;
    }
    .card{
      height: 150px;
    }
    .state-type{
      padding-bottom: 20px;
    }

    
 

  </style>