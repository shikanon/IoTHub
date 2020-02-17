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
                     <el-button type="text" size="small" @click="queryDataByName(card)" >查看数据</el-button>
                  </div>                
                    <p v-if="card.DataType === 'struct'" class="value-json" >
                        <el-tooltip placement="top" effect="light"> 
                          <div slot="content"> {{ card.Value  |  cardValFilter(card.Unit) }} </div>
                           <span>{{ card.Value  |  cardValFilter(card.Unit) }} </span> 
                        </el-tooltip>                                      
                    </p>    
                     <p v-else class="value-span" >
                        {{ card.Value  |  cardValFilter(card.Unit) }}    
                        <el-tooltip placement="top" effect="light" v-if="hopeData.filter(item => item.Identifier === card.Identifier).length > 0">
                          <div slot="content">期望值{{hopeData.filter(item => item.Identifier === card.Identifier)[0].Value ? hopeData.filter(item => item.Identifier === card.Identifier)[0].Value : '--'}}</div>
                          <span class="el-icon-info"></span>  
                        </el-tooltip>                                      
                    </p>                                                   
                  <p style="color: #999">{{card.Time ? card.Time : '--' }}</p>
            </el-card>
          </el-col>
         <el-dialog title="查看数据" :visible.sync="dialogTableVisible" width="40%">
                <StateDetailsl :deviceId="deviceId" :identifier="identifier" :name="name"></StateDetailsl>
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
          identifier:'',
          name:'',
          hopeData:[]
        
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

           let that = this 
           let promise = new Promise(function(resolve, reject) {
              //期望
              that.$API_IOT.getRunState(that.deviceId,'des').then((res) => {
                  if(res.data.status  === 'Y'){
                    that.hopeData = res.data.data                 
                  }else{
                    that.$message.error(res.message);
                  }               
                 resolve()    
              })  
               
            })

           promise.then(function() {
               //实际
              that.$API_IOT.getRunState(that.deviceId,'pro').then((res) => {
                  if(res.data.status  === 'Y'){
                     that.cards = res.data.data  
                  }else{
                    that.$message.error(res.message);
                  }  
              })
                                            
          })
                 
        },

     
        queryDataByName(card){
          this.dialogTableVisible = true 
          this.identifier = card.Identifier
          this.name = card.Name

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

    .value-span{
      font-size: 18px;
      color: rgb(51, 51, 51);
      height:21px;
    }


    .value-json{
      font-size: 14px;
      color: rgb(51, 51, 51);
      height:21px;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }

    .el-icon-info{
      color:#ccc;
    }

    
 

  </style>