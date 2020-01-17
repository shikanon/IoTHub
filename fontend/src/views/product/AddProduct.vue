<template>
    <div name="add-product">
      <div v-if="addProduct">
        <el-page-header @back="goBack" content="创建产品" class="page-header"></el-page-header>
        <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
          <el-form-item label="产品名称" prop="name">
              <el-input v-model="ruleForm.name" placeholder="请输入产品名称"></el-input>
          </el-form-item>
          <el-form-item label="所属品类" prop="category">

            <el-radio v-model="ruleForm.category" label="standard_category">标准品类</el-radio>
            <el-radio v-model="ruleForm.category" label="custom_category">自定义品类</el-radio>
      
          </el-form-item>
          <el-form-item >

            <el-select  v-if="ruleForm.category === 'standard_category'" 
              v-model="categoryName" 
              placeholder="请选择标准品类" 
              @visible-change="drawer = true"
              no-data-text="未选择" > 
                </el-select>
            <el-button type="text" :disabled="categoryName === ''"  @click="queryAttributeList">查看功能</el-button>
              <!-- <el-input
               v-if="ruleForm.category === 'standard_category'"
              placeholder="请选择标准品类"
              suffix-icon="el-icon-arrow-down"
              @focus="focusCategoryName"
              class="search-input"
              size="medium"
              v-model="categoryName">
            </el-input> -->
         
          </el-form-item>
          <el-form-item label="节点类型" prop="NodeType">  
             <el-radio-group v-model="ruleForm.node_type_id">
                <el-radio  v-for="(item,key) in nodeTypeArr" :key="key" :label="item.id" border>{{item.name}}</el-radio>
             </el-radio-group>
          </el-form-item>
          <p>连网与数据</p>
          <el-form-item label="连网方式" prop="NetType">
              <el-select v-model="ruleForm.network_id" placeholder="请选择连网方式">
                <el-option
                  v-for="item in netTypeArr"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
                </el-option>
              </el-select>
          </el-form-item>
          <el-form-item label="数据格式" prop="DataFormat">
              <el-select v-model="ruleForm.data_format_id" placeholder="请选择">
                <el-option
                  v-for="item in dataFormatArr"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
                </el-option>
              </el-select>
          </el-form-item>
          <el-collapse  @change="handleChange">
            <el-collapse-item title="认证方式" >
              <el-select v-model="ruleForm.auth_method_id" placeholder="请选择">
                <el-option
                  v-for="item in authTypeArr"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
                </el-option>
              </el-select>
            </el-collapse-item>
          </el-collapse>
          <p>更多信息</p>
          
            <el-collapse @change="handleChange">
            <el-collapse-item title=" 产品描述" >
              <el-input type="textarea" v-model="ruleForm.desc"></el-input>
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
           size='26%'>
            <CategoryList @select="selectCategory"></CategoryList>
        </el-drawer>
      </div>
      <div v-else>
        <AfterAddProduct :productId="productId" ></AfterAddProduct>
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
        drawer:false,
        addProduct:true,
        ruleForm: {
            name:"",
            category:"standard_category",
            model_id:1,
            node_type_id:1,
            network_id:1,
            data_format_id:1,
            auth_method_id:1,
            desc:"",         
        },
        nodeTypeArr:[],
        netTypeArr:[],
        dataFormatArr:[],
        authTypeArr:[],
        categoryName:'',
        productId:0,
        // ruleForm: {
        //   RegionId:"cn-shanghai",
        //   AliyunCommodityCode:"iothub_senior",
        //   ServiceCode:"",
        //   ProductName:"",
        //   CategoryId:138,
        //   NodeType:'1',
        //   Gateway:false,
        //   NetType:"WIFI",
        //   DataFormat:'1',
        //   AuthType:"secret",
        //   NameSpace:"ICA",
        //   Desc:'',
        //   type:'1',
          
        // },
        rules: {
          name: [
            { required: true, message: '请输入产品名称', trigger: 'blur' },
            ],
          categoryName: [
            { required: true, message: '请选择所属品类', trigger: 'blur' }
          ],
        },       
      };
    },
    created(){
        this.getNodeTypeList()
        this.getNetTypeList()
        this.getDataFormatList()
        this.getAuthTypeList()
        
    },
    methods: {
       goBack(){
            this.$router.back(-1)
      },

      getNodeTypeList(){
         this.$API_IOT.getNodeTypeList(this.currentPage,this.pageSize,this.search).then((res) => {
            this.nodeTypeArr = res.data.data       
          })
      },

      getNetTypeList(){
         this.$API_IOT.getNetTypeList(this.currentPage,this.pageSize,this.search).then((res) => {
              this.netTypeArr = res.data.data
         
            })
      },

      getDataFormatList(){
         this.$API_IOT.getDataFormatList(this.currentPage,this.pageSize,this.search).then((res) => {
              this.dataFormatArr = res.data.data      
          })
      },

      getAuthTypeList(){
         this.$API_IOT.getAuthTypeList(this.currentPage,this.pageSize,this.search).then((res) => {
              this.authTypeArr = res.data.data
          })
      },      
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$API_IOT.addProduct(this.ruleForm).then((res) => {
              if(res.data.status === 'Y'){
                  this.$message({
                    message: '添加产品成功',
                    type: 'success'
                  })
                  this.addProduct = false

                  this.productId = res.data.data.product_id 

              }else{
                  this.$message.error(res.message);
              } 
           })         
          } else {
            return false;
          }
          })
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      handleChange(){

      },

      focusCategoryName(){
        if(this.drawer){
          this.drawer = false
        }else{
          this.drawer =  true
        }
       
      },

      selectCategory(category){ 
        this.categoryName = `${category.territory.name}/${category.scene}/${category.name}`
        this.ruleForm.model_id = category.id
        this.drawer = false
      },

      queryAttributeList(){
          console.log(this.ruleForm.model_id );
      }
    
    }
  }
</script>
<style scoped >
     .el-form {
        width:460px;
    } 
</style>