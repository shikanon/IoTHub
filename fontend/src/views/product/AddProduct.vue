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
          <el-form-item v-if="ruleForm.category === 'standard_category'" >

            <el-select  
                  v-model="categoryName" 
                  placeholder="请选择标准品类" 
                  @visible-change="openCategoryList"
                  ref="category"> 
            </el-select>
            <el-button type="text" :disabled="categoryName === ''"  @click="showAttributeList(ruleForm.model_id)">查看功能</el-button>       
          </el-form-item>
          <el-form-item label="节点类型" prop="NodeType">  
             <el-radio-group v-model="ruleForm.node_type_id">
                <el-radio  v-for="(item,key) in nodeTypeArr" :key="key" :label="item.id" border>{{item.name}}</el-radio>
             </el-radio-group>
          </el-form-item>
           <el-divider></el-divider>
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

          <el-button class="el-icon-arrow-down" type="text" @click="authMethodVisiable = true" v-if="!authMethodVisiable">认证方式</el-button><br>
          <el-form-item  label="认证方式" prop="auth_method_id" v-if="authMethodVisiable">
              <el-select v-model="ruleForm.auth_method_id" placeholder="请选择" >
                <el-option
                  v-for="item in authTypeArr"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
                </el-option>
              </el-select><br>
              <el-button class="el-icon-arrow-up" type="text" @click="authMethodVisiable = false" v-if="authMethodVisiable">收起</el-button>
          </el-form-item>
          <el-divider></el-divider>
          <p>更多信息</p>    
          <el-button class="el-icon-arrow-down" type="text" @click="descVisiable = true" v-if="!descVisiable">产品描述</el-button><br>
          <el-form-item  label="产品描述" prop="desc" v-if="descVisiable">
              <el-input type="textarea" v-model="ruleForm.desc" maxlength="100" show-word-limit></el-input>               
              <el-button class="el-icon-arrow-up" type="text" @click="descVisiable = false" v-if="descVisiable">收起</el-button>
          </el-form-item>
          <el-divider></el-divider>
          <el-form-item>
              <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
              <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-form>
        <el-drawer
          :visible.sync="drawer"
           size='25%'>
            <CategoryList @select="selectCategory" @showAttributeList="showAttributeList($event)" :categoryId="ruleForm.model_id"></CategoryList>
            <el-drawer
              :append-to-body="true"
              :visible.sync="innerDrawer"
               direction="rtl"
               style="right:26%">
              <AttributeList :attributes="attributes"></AttributeList>
          </el-drawer>
        </el-drawer>
      </div>
      <div v-else>
        <AfterAddProduct :productId="productId" ></AfterAddProduct>
      </div>
    </div>
</template>
<script>
import CategoryList from './CategoryList'
import AttributeList from './AttributeList'

import AfterAddProduct from './AfterAddProduct'

  export default {
    components:{CategoryList,AttributeList,AfterAddProduct},
    data() {
     var checkProductName = (rule, value, callback) => {    
        setTimeout(() => {
          // 支持中文、英文字母、数字、和特殊字符_-@()，长度限制4~30，中文算2位
          let pattern =/[^a-z|A-Z|\u4E00-\u9FA5|0-9|\-|_|\@|\(|\))]/
          let length =  Number(value.replace(/[^\x00-\xff]/g,"01").length )
          if (!(length > 3 &&  length < 33 && !pattern.test(value))) {
            callback(new Error('支持中文、英文字母、数字、和特殊字符_-@()，长度限制4~30，中文算2位'))
          } else {     
            callback()      
          }
        }, 1000)
      }
      return {
        drawer:false,
        innerDrawer:false,
        addProduct:true,
        ruleForm: {
            name:"",
            category:"standard_category",
            model_id:0,
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
        attributes:[],
   
        rules: {
          name: [
           { required: true, message: '请输入产品名称', trigger: 'blur' },
           { validator: checkProductName, trigger: 'blur' }
          ],
          categoryName: [
            { required: true, message: '请选择所属品类', trigger: 'blur' }
          ],
          desc:[
           { min: 0, max: 100, message: '长度在 0 到 100 个字符', trigger: 'blur' }
          ]
        },   
        authMethodVisiable:false ,
        descVisiable:false   
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
                  this.$message.success( '添加产品成功' )
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

      openCategoryList(){console.log()
        this.drawer = true
        this.$refs.category.blur()
      },

      closeCategoryList(){ 
        this.drawer = false    
      },
  
      selectCategory(category){ 
        this.categoryName = `${category.territory.name}/${category.scene}/${category.name}`
        this.ruleForm.model_id = category.id 
        this.closeCategoryList()
      },

      showAttributeList(id){
          this.innerDrawer = true        
          this.$API_IOT.getModelfuncs(id ).then((res) => {
             this.attributes = res.data.data
           }) 
         
      }
    
    }
  }
</script>
<style scoped >
     .el-form {
        width:600px;
    } 
    .el-drawer{top:50px !important}
    .el-select,.el-input{width: 460px;}
</style>