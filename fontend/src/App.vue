<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  created(){
    //解决刷新页面时store重置的问题
    //在页面加载时读取sessionStorage里的状态信息
    if(sessionStorage.getItem('store')){
      this.$store.replaceState(Object.assign(this.$store.state,JSON.parse(sessionStorage.getItem("store"))))

    }

    //在页面刷新时将vuex里的信息保存到sessionStorage里
    window.addEventListener("beforeunload",()=>{
      sessionStorage.setItem('store',JSON.stringify(this.$store.state))
    })

  }


}
</script>

<style >
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.el-message-box__message p{
    word-wrap: break-word;
    word-break: normal;
}
.search_input{
  width:200px;
}

.align_right{
  text-align: right;
}

</style>>




