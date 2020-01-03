<template>
    <div name="add-product">
        <div v-if="addProduct">
        <el-page-header @back="goBack" content="创建产品" class="page-header"></el-page-header>
        <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
          <el-form-item label="产品名称" prop="ProductName">
              <el-input v-model="ruleForm.ProductName" placeholder="请输入产品名称"></el-input>
          </el-form-item>
          <el-form-item label="所属品类" prop="type">

            <el-radio v-model="ruleForm.type" label="1">标准品类</el-radio>
            <el-radio v-model="ruleForm.type" label="2">自定义品类</el-radio>
              <!-- <el-radio-group v-model="ruleForm.type">
                <el-radio label="标准品类"  value="1"></el-radio>
                <el-radio label="自定义品类"  value="2"></el-radio>
              </el-radio-group>        -->
          </el-form-item>
          <el-form-item >
               <el-select v-if="ruleForm.type === '1'" v-model="CategoryName" placeholder="请选择标准品类" @visible-change="click" > 

               </el-select>
          </el-form-item>
          <el-form-item label="节点类型" prop="NodeType">
             <el-radio-group v-model="ruleForm.NodeType">
                <el-radio  value="1" label="直连设备" border></el-radio>
                <el-radio  value="2" label="网关子设备" border></el-radio>
                <el-radio  value="3" label="网关设备"  border></el-radio>
             </el-radio-group>
          </el-form-item>
          <p>连网与数据</p>
          <el-form-item label="连网方式" prop="NetType">
              <el-select v-model="ruleForm.NetType" placeholder="请选择连网方式">
                <el-option
                  v-for="item in NetTypeArr"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
          </el-form-item>
          <el-form-item label="数据格式" prop="DataFormat">
              <el-select v-model="ruleForm.DataFormat" placeholder="请选择">
                <el-option
                  v-for="item in DataFormatArr"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
          </el-form-item>
          <el-collapse  @change="handleChange">
            <el-collapse-item title="认证方式" name="1">
              <el-select v-model="ruleForm.DataFormat" placeholder="请选择">
                <el-option
                  v-for="item in arr"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-collapse-item>
          </el-collapse>
          <p>更多信息</p>
          
            <el-collapse @change="handleChange">
            <el-collapse-item title=" 产品描述" name="1">
              <el-input type="textarea" v-model="ruleForm.Desc"></el-input>
            </el-collapse-item>
          </el-collapse>
          <el-form-item>
              <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
              <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-form>
        <el-drawer
          title="选择品类"
          :visible.sync="drawer"
          :before-close="handleClose"
           size='25%'>
            <CategoryList @select="selectCategory"></CategoryList>
        </el-drawer>
        </div>
        <div v-else>
          <AfterAddProduct :productName="ruleForm.ProductName" ></AfterAddProduct>
        </div>
    </div>
</template>
<script>
import CategoryList from './CategoryList'
import AfterAddProduct from './AfterAddProduct'

  export default {
    components:{CategoryList,AfterAddProduct},
    data() {
      return {
        CategoryName:'',
        drawer:false,
        addProduct:true,
        ruleForm: {
          RegionId:"cn-shanghai",
          AliyunCommodityCode:"iothub_senior",
          ServiceCode:"",
          ProductName:"123123123",
          CategoryId:138,
          NodeType:'1',
          Gateway:false,
          NetType:"WIFI",
          DataFormat:'1',
          AuthType:"secret",
          NameSpace:"ICA",
          Desc:'',
          type:'1',
          
        },
        rules: {
          ProductName: [
            { required: true, message: '请输入产品名称', trigger: 'blur' },
            // { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
          type: [
            { required: true, message: '请选择所属品类', trigger: 'change' }
          ],
          date1: [
            { type: 'date', required: true, message: '请选择日期', trigger: 'change' }
          ],
          date2: [
            { type: 'date', required: true, message: '请选择时间', trigger: 'change' }
          ],
       
          resource: [
            { required: true, message: '请选择活动资源', trigger: 'change' }
          ],
          desc: [
            { required: true, message: '请填写活动形式', trigger: 'blur' }
          ]
        },
         NetTypeArr: [{
          value: '1',
          label: 'WIFI'
        }, {
          value: '2',
          label: '蜂窝 (2G / 3G / 4G / 5G)'
        }, {
          value: '3',
          label: '以太网'
        }, {
          value: '4',
          label: 'LoRaWAN'
        }, {
          value: '5',
          label: '其他'
        }],
         DataFormatArr: [{
          value: '1',
          label: 'ICA 标准数据格式 (Alink JSON)'
        }, {
          value: '2',
          label: '透传/自定义'
        }],
         arr: [{
          value: '1',
          label: '设备密钥'
        }, {
          value: '2',
          label: 'ID²'
        },{
          value: '3',
          label: 'X.509证书'
        }],
       

        
      };
    },
    created(){
       
    },
    methods: {
       goBack(){
            this.$router.back(-1)
        },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
            setTimeout(()=>{
                this.addProduct = false
            },1000)
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      handleChange(){

      },
      handleClose(){

      },
      click(value){
        if(value){
            this.drawer = true
        }else{
            this.drawer = false

        }
      },
      selectCategory(id,name){    
        this.CategoryName = name        
      }
    
       
    }
  }
</script>
<style scoped >
     .el-form {
        width:460px;
    } 
</style>