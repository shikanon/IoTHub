import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)


export default new Router({
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/home',
      name: 'home',
      component: () => import( '@/views/Home.vue'),
      meta: {
        title: '首页',
        requireLogin: true
      },
      children:[
        {
          path: '/index',
          name: 'index',
          component: () => import( '@/views/Index.vue'),
          meta: {
            title: '概览',
            leftMenu:true
          },
        },
        {
          path: '/home/system',
          name: 'system',
          component:  () =>import( '@/views/Empty.vue'),
          meta: {
            title: '设备管理',
            leftMenu:true
          },
          redirect: "/home/system/user",
          children:[
               
            {
              path: '/home/product/product-list/:id',
              name: 'product-list',
              component:  () =>import( '@/views/product/ProductList.vue'),
              meta: {
                title: '产品',
                leftMenu:true
              },
              children:[
                {
                  path: '/home/product/add',
                  name: 'product-add',
                  component:  () =>import( '@/views/product/AddProduct.vue'),
                  meta: {
                    title: '创建产品',
                    leftMenu:false
                  },
                } ,
                {
                  path: '/home/product/details/:id',
                  name: 'product-details',
                  component:  () =>import( '@/views/product/ProductDetails.vue'),
                  meta: {
                    title: '产品详情',
                    leftMenu:false
                  },
                } 
              ]
            }     
            ,{
              path: '/home/device/:id',
              name: 'device-list',
              component:  () =>import( '@/views/device/DeviceManage.vue'),
              meta: {
                title: '设备',
                leftMenu:true
              },
              children:[
                {
                  path: '/home/device/details/:id',
                  name: 'device-details',
                  component:  () =>import( '@/views/device/DeviceDetails.vue'),
                  meta: {
                    title: '设备详情',
                    leftMenu:false
                  },
                } 
              ]
            }  
            ,{
              path: '/home/group/:id',
              name: 'group-list',
              component:  () =>import( '@/views/group/GroupList.vue'),
              meta: {
                title: '分组',
                leftMenu:true
              },
              children:[
                {
                  path: '/home/group/details/:id',
                  name: 'group-details',
                  component:  () =>import( '@/views/group/GroupDetails.vue'),
                  meta: {
                    title: '分组详情',
                    leftMenu:false
                  },
                } 
              ]
            } 
            
          ]
        } ,
        {
          path: '/test1',
          name: 'test1',
          component: () => import( '@/views/Empty.vue'),
          meta: {
            title: '测试1',
            leftMenu:true
          }
        },
        {
            path: '/test2',
            name: 'test2',
            component: () => import( '@/views/Empty.vue'),
            meta: {
              title: '测试2',
              leftMenu:true
            }
        }
             
      ]
    
    }, 
    {
      path: '/',
      name: 'Login',
      component: () => import( '@/views/Login.vue')
    }
  ]
  

 
})


