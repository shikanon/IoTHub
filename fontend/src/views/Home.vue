<template>
  <el-container>
    <el-header>    
      <div class="user-box"> 
        <el-radio-group v-model="themecolor">
          <el-radio :label="'00'">00</el-radio>
          <el-radio :label="'01'">01</el-radio>
          <el-radio :label="'02'">02</el-radio>
        </el-radio-group>    
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
        // return this.$store.state.themecolor === '00' ? '409EFF':'ffa666'

        },
        set(val){  
            this.setThemeColor(val)
       
           // this.setThemeColor(val === '409EFF' ? '00' : '01')
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
  // pickColor(val){
  //   this.setThemeColor(val === '409EFF' ? '00' : '01')
  // }


   
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped >
  .el-header{
    position: fixed;
    width: 100%;
    margin-right: 20px;
  }
  .el-header, .el-footer { 
    background-color: #B3C0D1;
    color: #333;
    text-align: center;
    line-height: 60px;
  }
  
  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 200px;
    position:fixed;
    margin-top: 60px;
  }
  
  .el-main {
    background-color: #E9EEF3;
    color: #333;
    text-align: left;
    margin-left:200px ;
    margin-top: 60px;
  }
  
  .el-container {
    min-height: calc(100vh - 76px);
  }
  
  .el-container:nth-child(5) .el-aside,
  .el-container:nth-child(6) .el-aside {
    line-height: 260px;
  }
  
  .el-container:nth-child(7) .el-aside {
    line-height: 320px;
  }

  .user-box{
    display: flex;
    align-items: center;
    justify-content: flex-end;
    height: 60px;
  }
</style>
