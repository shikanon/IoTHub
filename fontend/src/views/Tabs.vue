
<template>
  <div>
    <el-page-header @back="goBack" :content="pageName" class="page-header">
      <span>复制</span>
    </el-page-header>
      
      <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
        <el-tab-pane label="用户管理" name="first">
              <Submenu mode="horizontal"></Submenu>
        </el-tab-pane>
        <el-tab-pane label="运行状态" name="second">
          <el-col :span="24" >
            <el-radio-group v-model="radio">
              <el-radio :label="'chart'">图表</el-radio>
              <el-radio :label="'table'">表格</el-radio>
            </el-radio-group>
            <span>实时刷新</span>
            <el-switch v-model="autoRefresh">
            </el-switch>
          </el-col>
          <el-col :span="24" v-if="radio === 'table'">
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
                  property="address"
                  label="最新值">
                </el-table-column>
                  <el-table-column
                  property="address"
                  label="期望值">
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
         </el-col>
          <el-col :span="6" v-for="(card,index)  in cards" :key="index" v-else>
            <el-card  shadow="hover" class="card">
                  <div style="display: flex;justify-content: space-between;">
                     {{card.Name}}
                   <el-button type="text" size="small" @click="queryDataByName" >查看数据</el-button>
                  </div>
                  <p style="font-size: 18px;color: rgb(51, 51, 51);height:21px;" >{{ card.val | cardValFilter(card.Unit) }}  </p>
                  <p style="color: #999">--</p>
            </el-card>
          </el-col>
         
        </el-tab-pane>
        <el-tab-pane label="角色管理" name="third">
            <!-- <Pagination 
           ></Pagination> -->
        </el-tab-pane>
        <el-tab-pane label="复制粘贴" name="fourth">
            <el-button  type="text" size="medium"
              v-clipboard:copy="pageName"
              v-clipboard:success="onCopy"
              v-clipboard:error="onError">点击复制名称</el-button>
        </el-tab-pane>
      </el-tabs>
      <el-dialog title="查看数据" :visible.sync="dialogTableVisible">
          详细数据
     </el-dialog>
  </div>
</template>
  <script>
  import Submenu from '@/components/Submenu.vue'
  import Breadcrumb from '@/components/Breadcrumb.vue'
  import Pagination from '@/components/Pagination'
  
    export default {
      components: { Pagination, Submenu,Breadcrumb },
      data() {
        return {
          activeName:'first',
          pageName:'',
          cards:[],
          radio:'chart',
          autoRefresh:false,
          tableData:[],
          dialogTableVisible:false,
          intervalId:null
        }
      },
      watch:{
          radio:function(){
              
          },
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
        this.pageName = this.$route.params.username
        this.getCardData()        
      },
      destoryed() {
        clearInterval(this.intervalId)
      },
      methods:{
        clearInterval(){

        },
        goBack(){
            this.$router.back(-1)
        },
       
        handleClick(tab, event) {
          console.log(tab, event);
        },
        onCopy(e){
          alert(e.text)
            console.log(e)
        },
        onError(){

        },
        getTableData(){
            this.tableData = [{
                date: '2016-05-02',
                name: '王小虎',
                address: '上海市普陀区金沙江路 1518 弄'
              }, {
                date: '2016-05-04',
                name: '王小虎',
                address: '上海市普陀区金沙江路 1518 弄'
              }, {
                date: '2016-05-01',
                name: '王小虎',
                address: '上海市普陀区金沙江路 1518 弄'
              }, {
                date: '2016-05-03',
                name: '王小虎',
                address: '上海市普陀区金沙江路 1518 弄'
              }]
        },
        getCardData(){
            this.cards =[
              {"Identifier":"LightStatus","DataType":"bool","Unit":"","Name":"工作状态"},
              {"Identifier":"LightAdjustLevel","DataType":"int","Unit":"%","Name":"调光等级"},
              {"Identifier":"LightVolt","DataType":"float","Unit":"V","Name":"工作电压"},
              {"Identifier":"LightCurrent","DataType":"float","Unit":"A","Name":"工作电流"},
              {"Identifier":"ActivePower","DataType":"float","Unit":"W","Name":"有功功率值"},
              {"Identifier":"PowerRatio","DataType":"float","Unit":"pF","Name":"功率因数"},
              {"Identifier":"PowerConsumption","DataType":"float","Unit":"kW·h","Name":"用电量"},
              {"Identifier":"DrainVoltage","DataType":"float","Unit":"V","Name":"漏电压"},
              {"Identifier":"TiltValue","DataType":"int","Unit":"°","Name":"倾斜角度值"},
              {"Identifier":"ErrorPowerThreshold","DataType":"int","Unit":"W","Name":"故障功率门限"},
              {"Identifier":"ErrorCurrentThreshold","DataType":"float","Unit":"A","Name":"故障电流门限"},
              {"Identifier":"TiltThreshold","DataType":"int","Unit":"°","Name":"倾斜阈值"},
              {"Identifier":"UnderVoltThreshold","DataType":"int","Unit":"V","Name":"欠压阈值"},
              {"Identifier":"OverCurrentThreshold","DataType":"int","Unit":"A","Name":"过流阈值"},
              {"Identifier":"OverVoltThreshold","DataType":"int","Unit":"V","Name":"过压阈值"},
              {"Identifier":"LightErrorEnable","DataType":"bool","Unit":"","Name":"灯具故障使能"},
              {"Identifier":"OverCurrentEnable","DataType":"bool","Unit":"","Name":"过流告警使能"},
              {"Identifier":"OverVoltEnable","DataType":"bool","Unit":"","Name":"过压告警使能"},
              {"Identifier":"UnderVoltEnable","DataType":"bool","Unit":"","Name":"欠压告警使能"},
              {"Identifier":"LeakageEnable","DataType":"bool","Unit":"","Name":"漏电告警使能"},
              {"Identifier":"OverTiltEnable","DataType":"bool","Unit":"","Name":"倾斜告警使能"},
              {"Identifier":"LampError","DataType":"bool","Unit":"","Name":"灯具故障告警"},
              {"Identifier":"OverCurrentError","DataType":"bool","Unit":"","Name":"过流告警"},
              {"Identifier":"OverVoltError","DataType":"bool","Unit":"","Name":"过压告警"},
              {"Identifier":"UnderVoltError","DataType":"bool","Unit":"","Name":"欠压告警"},
              {"Identifier":"OverTiltError","DataType":"bool","Unit":"","Name":"倾斜告警"},
              {"Identifier":"LeakageError","DataType":"bool","Unit":"","Name":"漏电告警"},
              {"Identifier":"GeoLocation","DataType":"struct","Unit":"","Name":"地理位置"}

            ]
        },
        queryDataByName(){
          this.dialogTableVisible = true 
        }
         
         
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
   
 

  </style>