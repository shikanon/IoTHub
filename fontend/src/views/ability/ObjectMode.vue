<template>
    <div name="object-mode">
        <div class="desc-txt-gray margin-b-20">
            物模型是对设备在云端的功能描述，包括设备的属性、服务和事件。
            物联网平台通过定义一种物的描述语言来描述物模型，
            称之为 TSL (即Thing Specification Language)，
            采用JSON格式，您可以根据TSL组装上报设备的数据。您可以导出完整物模型，
            用于云端应用开发；您也可以只导出精简物模型，配合设备端SDK实现设备开发。
        </div>
        <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
            <el-tab-pane label="完整物模型" name="complete" class="tab-pane">
                <JsonView :data="completeObject"></JsonView>
            </el-tab-pane>
            <el-tab-pane label="精简物模型" name="simplify"  class="tab-pane">
                <JsonView :data="simplifyObject"></JsonView>
            </el-tab-pane>
        </el-tabs>
        
    </div> 
</template>
 

 <script>
 import FileSaver from 'file-saver'
  export default {
    props:{
        productId:{
            type:Number,
            default:0
        }
    },
    data() {
      return {
        activeName: 'complete',
        completeObject: {},
        simplifyObject:{},
        exportData:'',
      };
    },
    created(){
       this.init()
    },
    methods: {
        handleClick(tab, event){
            if(tab.index === '0'){
                this.exportData = this.completeObject               
            }else{
                this.exportData = this.simplifyObject
              
            }
        },

        exportFile(){
            const data = JSON.stringify( this.exportData)
            const blob = new Blob([data], {type: ''})
            FileSaver.saveAs(blob, 'model.json')
        },
        init(){

            this.$API_IOT.getProductMode(this.productId).then((res) => {
                    this.completeObject  = res.data.data.concise_model
                    this.simplifyObject  = res.data.data.intact_model
            })

            //  this.$axios.get('http://localhost:8080/static/complete.json').then((res) => {
            //      this.completeObject  = res.data         
            //  }) 

            // this.$axios.get('http://localhost:8080/static/simplify.json').then((res) => {
            //         this.simplifyObject  = res.data       
            // })
        },
      
    }
  }
</script>


<style scoped>
 .tab-pane{
     height: 240px !important;
     overflow-y: auto !important;
 }
</style>
