<template>
  <div name="ability-list">
      <el-row>
        <el-button type="primary" @click="addAbility">{{type === 'standard' ? '添加标准功能':'添加自定义功能'}}</el-button>
        <el-button @click="copyVisible = true">导入物理模式</el-button>
        <el-button @click="objectModeVisible = true">查看物流模型</el-button>
        <el-button @click="exportDeviceEndCode">生成设备端代码</el-button>
     </el-row>
    
    <el-table
      :data="tableData"
      >
       <el-table-column
          label="功能类型"
           prop="type"
          width="150"  
         >         
        </el-table-column>
        <el-table-column
            label="功能名称"
            width="200">
            <template slot-scope="scope">{{ scope.row.name  }} 
                <el-tag  class="tag">{{ scope.row.required === 'true' ? '必选':'可选' }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column
            prop="identifier"
            label="标识符"  
             width="200"
        >
        </el-table-column>
            
        <el-table-column
            label="数据类型" >
            <template slot-scope="scope">{{ scope.row.data_type  }}</template>
        </el-table-column>
        <el-table-column
            label="数据定义" >
            <template slot-scope="scope">           
              <p  v-if="scope.row.data_type === 'struct'"></p> 
              <div  v-else-if="scope.row.data_type === 'bool'">
                <span>布尔值：</span>
                <div class="bool-data-list">
                  <p v-for="(value,key,index) in scope.row.data_condition" :key="index">
                    <span  class="bool-data">{{key}} - {{value}} </span>
                  </p>
                </div>
              </div> 
              <p v-else >取值范围:{{scope.row.data_condition.min}}~{{scope.row.data_condition.max}}</p>
             </template>                       
        </el-table-column>
         <el-table-column    
            label="操作"
            >
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="edit(scope.row)">编辑</el-button>
              <el-divider v-if="scope.row.required === 'false'" direction="vertical"></el-divider>
              <el-button v-if="scope.row.required === 'false'" type="text" size="small" @click="delete(scope.$index, scope.row)">删除</el-button>
            </template>
          </el-table-column>
    </el-table>
    <el-dialog title="查看标准功能" :visible.sync="editVisible" width="27%">
          <EditAbility ref="editAbility" :ability="ability" @close="editVisible = false"></EditAbility>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="editSubmit">确 定</el-button>
                <el-button @click="editVisible = false">取 消</el-button>
            </span>
      </el-dialog>
    <el-dialog title="导入物模型" :visible.sync="copyVisible" width="25%">
          <CopyProductAbility ref="copy"  @close="copyVisible = false"></CopyProductAbility>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="copySubmit">确 定</el-button>
                <el-button @click="copyVisible = false">取 消</el-button>
            </span>
      </el-dialog>
      <el-dialog title="查看物模型" :visible.sync="objectModeVisible" width="35%">
          <ObjectMode ref="objectMode"  @close="objectModeVisible = false" :productId="productId"></ObjectMode>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="exportFile">导出模型文件</el-button>
            </span>
      </el-dialog>
  </div>
</template>

<script>
import CopyProductAbility from './CopyProductAbility'
import ObjectMode from './ObjectMode'
import EditAbility from './EditAbility'
import FileSaver from 'file-saver'
import JSZip from 'jszip'

  export default {
    components:{CopyProductAbility,ObjectMode,EditAbility},
     props:{
       type:{
         type:String,
         default:'standard'
       },
        productKey:{
          type:String,
          default:''
        },
         productId:{
          type:Number,
          default:0
        }
     },

  
    data() {
      return {
        tableData: [],
        copyVisible:false,
        objectModeVisible:false,
        editVisible:false,
        ability:{}
       
      }
    },
 

    created(){
        this.getAbilityList()
    },
    methods: {
       getAbilityList(){
          this.tableData =  []
           this.$API_IOT.getAbilityList(this.productId).then((res) => {


                if(res.data.data.property.length > 0){
                  res.data.data.property.map(item => {item['type'] = '属性'; return item}).forEach(item => {
                    this.tableData.push(item)
                  });
                
                }
                 if(res.data.data.events.length  > 0){
                     res.data.data.events.map(item => {item['type'] = '事件'; return item}).forEach(item => {
                      this.tableData.push(item)
                    });
                }
                 if(res.data.data.services.length > 0){
                    res.data.data.services.map(item => {item['type'] = '服务'; return item}).forEach(item => {
                      this.tableData.push(item)
                    });
                }

                
            })
                        
       },

       edit(data){
          this.editVisible = true
          this.ability = data 
       },

       editSubmit(){

       },
       addAbility(){
         if(this.type === 'standard'){

         }else{

         }
       },
       copySubmit(){
          this.$refs.copy.submit()
       },

       exportFile(){
          this.$refs.objectMode.exportFile()
       },

       exportDeviceEndCode(){
  
           this.$axios.get('http://localhost:8080/static/test.c').then((res) => {
              //用axios的方法引入地址

              const zip = new JSZip()
            
              // let txtData = ""
              // res.data.forEach((row) => {
              //   let tempStr = ''
              //   tempStr = row.toString()
              //   txtData += `${tempStr}\r\n`
              // })
              zip.file(`${this.productKey}.c`, res.data)
              zip.generateAsync({
                type: "blob"
              }).then((blob) => {
                saveAs(blob,`${this.productKey}.zip`)
              }, (err) => {
                alert('导出失败')
              })
              //console.log(res.data)
              // const blob = new Blob([res.data], {type: ''})
              // FileSaver.saveAs(blob, `${this.productKey}.zip`)
            })          
       }
      
    }
  }
</script>

<style scoped>
 .tag{
        border-radius: 20px;
        margin-left: 10px;
        height: 24px;
        line-height: 24px;
    }
    .bool-data-list{
      display: flex;
      flex-direction:column;
    }
    .bool-data-list p {
      margin: 0px
    }
    .bool-data{
      background-color: #ecedee;
      padding: 0px 8px;
      margin: 8px 0 ; 
    }
</style>