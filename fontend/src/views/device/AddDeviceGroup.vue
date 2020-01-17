<template>
    <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
        <el-form-item v-if="productId" label="产品" >        
           <el-input v-model="productName"  disabled></el-input>
        </el-form-item>
        <el-form-item  v-else  label="产品" prop="pid">
            <el-select v-model="ruleForm.pid" placeholder="请选择产品">
                <el-option
                    v-for="item in productArrTemp"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                </el-option>
              </el-select>   
          </el-form-item>

  
        <el-form-item label="添加方式" prop="type">
              <el-radio v-model="ruleForm.type" label="1"  border size="medium">自动生成</el-radio>
              <el-radio v-model="ruleForm.type" label="2"  border size="medium">批量上传</el-radio>
        </el-form-item>
    
        <el-form-item  v-if="ruleForm.type === '1'" label="设备数量" prop="count">
            <el-input-number v-model="ruleForm.count" controls-position="right"  :min="1"  label="请输入设备数量"></el-input-number>
        </el-form-item>  
         
          <el-form-item v-if="ruleForm.type === '2'" label="批量上传文件" prop="fileList">      
            <el-upload
                action="https://jsonplaceholder.typicode.com/posts/"
                :on-success="success"
                :before-uoload="beforeUpload"
                :on-preview="handlePreview"
                :on-remove="handleRemove"
                :file-list="fileList"
                :auto-upload="false"
                :limit="1"
               >
               <el-button size="small" type="primary">上传文件</el-button>
               
                <download-excel :fields = "json_fields"  :data="json_data" name="Template.xls"> 
                  <el-button size="small" type="text">下载.csv模板</el-button>
                </download-excel> 
                 <div slot="tip" class="el-upload__tip">只能上传excel文件，单次最多添加 1000 台</div>
            </el-upload>
         </el-form-item>  
    </el-form>
 
</template>
<script>
  export default {
    props:{
      productArr:{
          type:Array,
          default: () => []
      },
      productId:{
        type:Number,
        default:0
      }
    },
     watch:{
        //监听productId,若发生变化，重新查询设备列表
        productId:{  
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
        productName:"",
        ruleForm: {},
        productArrTemp:[],
        rules: {
          pid: [
            { required: true, message: '请选择产品', trigger: 'blur' },
           ],
          //   type: [
          //       { required: true, message: '请选择添加方式', trigger: 'change' },
          //   ],
          //  count: [
          //   { required: true,type:'number',message: '请输入设备数量', trigger: 'blur' }
          // ]
              
        }
      };
    },
    created(){
       this.init()
    },
    methods: {
        init(){
          this.productArrTemp = JSON.parse(JSON.stringify(this.productArr))
          this.productArrTemp.shift();
          //选择了产品id，固定写死产品名称，不能选择
          if(this.productId && this.productId !== 0){
               this.productName =  this.productArr.filter(item => item.id === this.productId)[0].name
                this.ruleForm =  {
                  pid: this.productId,
                  type: '1',
                  count: 1 ,     
                }
          }else{
            this.ruleForm= {    
              product: null,
              type: '1',
              count: 1 ,
            }
          }

            
        },    
    
      submit(){
        if(this.ruleForm.type === "1"){
          this.submitForm()
        }else{
            this.submitUpload()
        }
      },
      submitForm() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {
            alert('submit!');
            //回调父组件方法
            this.$emit('close')
            //回调父组件方法显示结果
             this.$emit('showAddResult');
            
            setTimeout(()=>{
                this.$emit('addSuccess',this.ruleForm.count)
             },3*1000)
             
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },  
      submitUpload() {
          this.$refs.upload.submit();
          this.$emit('close')
      }, 

      downCsvTemplate(){

      }
    }
  }
</script>
<style scoped >
  .el-select{
      width: 440px;
  }
  .el-upload{
    display: flex !important;
  }
</style>