import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('./../views/Login.vue')
    },
    {
      path: '/home',
      name: 'Home',
      component: () => import('./../views/Home.vue')
    },
    {
      path: '/',
      name: 'Login',
      component: () => import('./../views/Login.vue')
    }
  ]
})
