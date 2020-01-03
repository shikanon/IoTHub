<template>
  <div class="breadcrumb">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item v-for="(item,index) in routerList" :key="index" :to="item.path">{{item.title}}</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
</template>

<script>

export default {
  data() {
    return {
      routerList:[]
    };
  },

  created(){
      window.addEventListener("beforeunload",()=>{
        sessionStorage.setItem('store',JSON.stringify(this.$store.state))
      })
  },
  watch: {
    $route:function(){
      this.routerList = [];    
      for (const routeInfo of this.$route.matched) {
          if (typeof routeInfo.name === 'string' && routeInfo.name !== 'Home') {
              this.routerList.push({path:routeInfo.path,title:routeInfo.meta.title});
          }             
      }         
    }
  },
  methods:{

    
  }

}
</script>

<style scoped>
   .breadcrumb{
    margin-bottom: 15px;
    margin-left: 5px;
   }

</style>

