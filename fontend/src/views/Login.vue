<template>

    <el-container>
    
    <el-main>
        <div class="login-box">
          <el-form :label-position="labelPosition" label-width="80px" ref="loginForm" :model="loginForm" :rules="rules" >
              <el-form-item :label="$t('text.userName')" prop="user_name">
                <el-input v-model="loginForm.user_name" auto-complete="off" ></el-input>
              </el-form-item>
              <el-form-item :label="$t('text.password')" prop="password">
                <el-input v-model="loginForm.password" auto-complete="off" ></el-input>
              </el-form-item>
            
              <el-form-item>
                <el-button type="primary" @click="submitForm('loginForm')">{{$t('btn.login')}}</el-button>
                <el-button @click="resetForm('loginForm')">{{$t('btn.reset')}}</el-button>
              </el-form-item>
            </el-form>     
        </div>  
    </el-main>
    <el-footer>
        <el-select v-model="lang_value"  @change="changeLanguage">
            <el-option
              v-for="item in languageArr"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
        </el-select>
    </el-footer>
  </el-container>
    
</template>

<script>
import {mapState,mapMutations} from 'vuex'
export default {
   data() {
      return {
        labelPosition: 'right',
        loginForm: {
          user_name: '',
          password: '',
        },
         rules: {
          user_name: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
            { min: 3, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
          ],
          password: [
            { required: true, message:  '请输入密码', trigger: 'blur' }
          ]        
        },
        languageArr:[{label:"中文",value:"zh"},{label:"English",value:"en"}],
        lang_value:'中文',
      };
    },
   computed:{

    ...mapState({  
         language:state=>state.language                                                 
      }),

  },

    methods: {
      ...mapMutations([
        'setLoginName',
        'setTOKEN',
      ]),
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            console.log('login start!!',this.loginForm.user_name);
            // this.$post('/login',this.loginForm).then((res) => {
            //   console.log(res)
            //   this.setLoginName(this.loginForm.user_name)
            //   this.$router.push({ path: '/home' })
            // });
            this.$API.login(this.loginForm).then((res) => {
               console.log(res)
               if(res.data.errorCode === 0){

                 
                  this.$alert(this.$t('text.login_err_txt'), this.$t('text.system_tips'), {
                     confirmButtonText: this.$t('btn.sure')
                  })
            console.log('error submit!!');
            return false;
               }else{
                  this.setLoginName(this.loginForm.user_name)
                  this.setTOKEN(res.data.token)
                  this.$router.push({ path: '/home' })
               }
               
            })
                       
          } 
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },

      changeLanguage(value){
         this.$i18n.locale = value      
          if (value === 'en') {          
            this.rules.user_name[0].message = '请输入用户名'
            this.rules.user_name[1].message = '长度在 3 到 5 个字符'
            this.rules.password[0].message = '请输入密码'
          } else {
             this.rules.user_name[0].message = 'please enter the username'
             this.rules.user_name[1].message = 'length must betwwen 3 and 5'
             this.rules.password[0].message = 'please enter the password'

          }
     }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped >
  .login-box{
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;
        height: calc(100vh - 145px);
  }
  el-form{
      width: 300px;
  }
</style>
