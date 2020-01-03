<template>
    <div name="copy-product-ability">
        <div class="desc-txt-blue margin-b-20">注：导入的物模型会覆盖原来的功能。</div>
        <el-tabs v-model="activeName" type="card">
            <el-tab-pane label="拷贝产品" name="form">
                 <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
                    <el-form-item label="选择产品" prop="ProductKey">
                        <el-select v-model="ruleForm.ProductKey" placeholder="请选择产品">
                                <el-option
                                v-for="item in productArr"
                                :key="item.ProductKey"
                                :label="item.ProductName"
                                :value="item.ProductName">
                                </el-option>
                            </el-select>   
                    </el-form-item>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="导入物模型" name="uopload">
                    <el-upload
                        action="https://jsonplaceholder.typicode.com/posts/"
                        :on-success="success"
                        :before-uoload="beforeUpload"
                        :on-preview="handlePreview"
                        :on-remove="handleRemove"
                        :file-list="fileList"
                        :auto-upload="false"
                         :limit="1">
                        <el-button size="small" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传json文件，且不超过256kb</div>
                    </el-upload>
            </el-tab-pane>
     </el-tabs>
       
</div>
 
 
</template>
<script>
  export default {
      
    props:{
          SrcProductKey:{
              type:Array,
              default: () => []
          }
    },
    data() {
      return {
        ruleForm: {       
          ProductKey: '',      
        },
        rules: {
          ProductKey: [
            { required: true, message: '请选择产品', trigger: 'change' },
            {  message: '请选择产品', trigger: 'blur' },
          ]
              
        },
        activeName:'form',
        productArr:[],
        fileList:[]
      };
    },
    created(){
       this.init()

    },
    methods: {
        init(){
            this.ruleForm = {       
                ProductKey: ''      
            }
            this.getProductList()
        },

        getProductList(){
                this.productArr =[{
                    "GmtCreate": "2019-12-19 11:35:01",
                    "ProductName": "4444444",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726501000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1ELejzj0h9"
                },
                {
                    "GmtCreate": "2019-12-19 11:34:50",
                    "ProductName": "123123123",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726490000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1qfUCxdfqg"
                },
                {
                    "GmtCreate": "2019-12-19 11:34:39",
                    "ProductName": "3231323",
                    "AliyunCommodityCode": "iothub_senior",
                    "NodeType": 0,
                    "CreateTime": 1576726479000,
                    "DataFormat": 1,
                    "ProductStatus": "DEVELOPMENT_STATUS",
                    "NetType": 3,
                    "ProductKey": "a1Ibli2tqC2"
                },]
        },
        
        submit(){
            if(this.activeName === 'form'){
                this.submitForm()
            }else{
                this.submitUpload()
            }
        },
        submitForm() {
            this.$refs.ruleForm.validate((valid) => {
            // {"SrcProductKey":"a16cKjxJOx4","TargetProductKey":"a1L6W2BQytT"}
            if (valid) {
                alert('submit!');
                //回调父组件方法关闭弹窗
                this.$emit('close')
             
            } else {
                console.log('error submit!!');
                return false;
            }
            });
        },
        success(response, file, fileList){
            if(file.raw.type !== 'application/json'){
                this.$message.error('文件类型格式只能为json文件');
                return false
            }else if((file.raw.size/1024/1024) > 256){
                this.$message.error('文件大小不能超过256kb');
                return false

            }

        },

        submitUpload() {
            this.$refs.upload.submit();
            this.$emit('close')
         },
        handleRemove(file, fileList) {
            console.log(file, fileList);
        },
        handlePreview(file) {
            alert(2)
            console.log(file);
        }


     
    }
  }
</script>
<style scoped >
  .el-select{
      width: 440px;
  }
</style>