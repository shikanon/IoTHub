<template>
    <div class="content">
        <el-container>
        
          <el-main>
              <div class="login-box">
                <div  class="login-form">    
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
              </div>  
          </el-main>
          <!-- <el-footer>
              <el-select v-model="lang_value"  @change="changeLanguage">
                  <el-option
                    v-for="item in languageArr"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value">
                  </el-option>
              </el-select>
          </el-footer> -->
        </el-container>
    </div>
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
            { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' }
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
            // this.$API_IOT.login(this.loginForm).then((res) => {
            //    console.log(res)
            //    if(res.data.errorCode === 0){

                 
            //       this.$alert(this.$t('text.login_err_txt'), this.$t('text.system_tips'), {
            //          confirmButtonText: this.$t('btn.sure')
            //       })
            //       return false;
            //    }else{
            //       this.setLoginName(this.loginForm.user_name)
            //       this.setTOKEN(res.data.token)
            //       this.$router.push({ path: '/home' })
            //    }
               
            // })

               this.setLoginName(this.loginForm.user_name)
               this.setTOKEN("res.data.token")
               this.$router.push({ path: '/home' })
                       
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
        height: calc(100vh - 145px);
  }
  .login-form{
    width: 400px;
    padding: 40px 50px 30px 30px;
    border-radius: 30px;
    /* background-image: url(../assets/login-box.png);
    background-repeat: no-repeat;
    background-size: cover; */
    background-color: #fff;
  }
  el-form{
      margin-right: 30px;
      margin-top: 30px
  }

  .el-form-item:last-child{margin-bottom: 0px;}

  .content{
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;
    background: #fff;
    bottom: 0;
    right: 0;
    margin: auto;
    background-image: url(../assets/login-bg.jpg);
    background-repeat: no-repeat;
    background-size: cover;
  }


</style>
