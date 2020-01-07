<template>
  <el-container>
    <el-header>    
      <el-dropdown  @command="pickThemeColor">
          <i class="el-icon-setting"></i>主题设置
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item command="00" style="color:#409eff">蓝色</el-dropdown-item>
            <el-dropdown-item command="01" style="color:#ff9040">橙色</el-dropdown-item>
            <el-dropdown-item command="02" style="color:#ff4080">粉色</el-dropdown-item>
          </el-dropdown-menu>
      </el-dropdown>    
      <div class="flex-center">
        <el-avatar> {{loginName}} </el-avatar>
        <span @click="loginOut">退出</span>  
      </div>     
    </el-header>
    <el-container>
      <el-aside width="200px">
            <Submenu :menuNam="menuNam"></Submenu>
      </el-aside>
      <el-main>    
            <Breadcrumb></Breadcrumb>  
        <el-card >
            <router-view /> 
        </el-card>        
      </el-main>
    </el-container>
  </el-container>

</template>

<script>
import Submenu from '@/components/Submenu.vue'
import Breadcrumb from '@/components/Breadcrumb.vue'
import {mapState,mapGetters, mapMutations} from 'vuex'
export default {

  components:{
        Submenu,
        Breadcrumb
  },
    data () {
      return {
        menuNam:'',
      }
    },
    computed:{
      ...mapState({ 
          loginName:state=>state.loginName 
        }),
      themecolor:{
          get(){
            return this.$store.state.themecolor
          },
          set(val){  
              this.setThemeColor(val)    
          }
        }  
    },
    watch:{
      themecolor:{
        handler(){
          this.toggleClass(document.body,'custom-' + this.themecolor)
        }
      }
    },
    mounted(){
      this.toggleClass(document.body,'custom-' + this.themecolor)
    },
    created(){
        this.menuNam = 'home'

      },
    methods:{
      ...mapMutations([
        'setLoginName',
        'setThemeColor'
      ]),
    
    loginOut(){
        this.setLoginName("")
        this.$router.push({ path: '/' })
    }
    ,
    // 换肤加class函数
    toggleClass(element, className) {
      if (!element || !className) {
        return;
      }
      element.className = className;
    },
    pickThemeColor(common){
      this.themecolor = common
    }


    
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped >
  

  .el-header{
    background-color: #B3C0D1;
    color: #333;
    line-height: 60px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    position: fixed;
    width: 100%;
    margin-right: 20px;
    z-index: 9999;
  }
  

  .el-container .el-container{
    min-height: calc(100vh - 60px);
    margin-top: 60px;
  }
    .el-aside {
    background-color: #D3DCE6;
    color: #333;
    position:fixed;
  }
  
  .el-main {
    background-color: #E9EEF3;
    color: #333;
    text-align: left;
    margin-left:200px ;
    min-height: calc(100vh - 60px);
  }
  .el-card{
        min-height: calc(100vh - 140px);
  }

  .flex-center{
    display: flex;
    align-items: center
  }


</style>
