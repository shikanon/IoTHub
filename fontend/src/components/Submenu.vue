<template>
    <el-menu class="el-menu-vertical-demo"
        @select="handleSelect"
         :collapse="isCollapse"
         :mode="menuMode"
         :router="isRouter"
         :unique-opened ="isUniqueOpended"
         :default-active ="defaultActive"
     >
        <template v-for="(item,i) in navList">
            <!-- 一级菜单（有子菜单） -->
            <el-submenu v-if="item.children && item.children.length "  :key="i" :index="item.path">
                <template slot="title">
                <i class="el-icon-location"></i>
                <span>{{ item.meta.title  | menuFilter  }}</span>
                </template> 
                <!-- 存在二级菜单 -->
                <template v-if="item.children">
                    <template v-for="(children,n) in item.children">
                         <!-- 二级菜单（有子菜单,并且是需要显示的） -->
                        <el-submenu  v-if="children.children  && children.children.filter(item => item.meta.leftMenu === true).length > 0" :key="`${i}-${n}`" :index="children.path">
                            <template slot="title">{{children.meta.title   | menuFilter }}</template>
                            <!-- 三级菜单 -->
                            <el-menu-item  v-for="(childrenItem,p) in children.children.filter(item => item.meta.leftMenu === true)" :key="`${i}-${n}-${p}`" :index="childrenItem.path">
                                {{ childrenItem.meta.title  }}
                            </el-menu-item>
                        </el-submenu>
                        <!-- 二级菜单（无子菜单） -->
                        <el-menu-item v-else :key="`${i}-${n}`" :index="children.path">{{  children.meta.title  }}</el-menu-item>
                    </template>
                </template>      
            </el-submenu>               
            
            <!-- 一级菜单（无子菜单） -->
            <el-menu-item v-else  :key="i" :index="item.path">
                <i v-if="!item.children" class="el-icon-setting"></i>
                <span v-if="!item.children" slot="title">{{ item.meta.title   }}</span>
            </el-menu-item>  
        </template>                   
    </el-menu>         
</template>

<script>
export default {
 props: {
     menuNam:{
        type: String,
        default: 'home'
     },
     menuMode:{
        type: String,
        default: 'vertical'
     },
     
    //  backgroundColor:{
    //     type: String,
    //     default: '#ffffff'
    //  },
    //  textColor:{
    //     type: String,
    //     default: '#303133'
    //  },
    //  activeTextColor:{
    //     type: String,
    //     default: '#409EFF'
    // },
     isUniqueOpended:{
        type: Boolean,
        default: true
     },
     isRouter:{
        type: Boolean,        
        default: false
     },
    isCollapse:{
        type: Boolean,
        default: false
     }, 
    },
  data () {
    return {
      navList:[],
      defaultActive:''
    }
  },
  watch: {
   
    $route:function(){   
        //匹配路由
        this.defaultActive = this.$route.matched.filter(route => route.name === this.$route.name)[0].path
    }
  },
  created(){
      console.log("生成菜单",this.menuNam)
        //获取路由信息
      let router =   this.$router.options.routes;
      this.navList = router.filter(item => item.name === this.menuNam)[0].children     
    },
    methods:{
        //跳转路由
        handleSelect(key, keyPath){    
            if(!this.routerFlg){
                this.$router.push(key)
            }
            
        }
    }

}
</script>

<style scoped>

</style>
