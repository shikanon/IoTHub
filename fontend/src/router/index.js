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
            title: '系统管理',
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
              name: 'device-details',
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
        }
        ,{
          path: '/home/system1',
          name: 'system1',
          component:  () =>import( '@/views/Test1.vue'),
          meta: {
            title: '营业厅管理',
            leftMenu:true
          },
        }
      
      
        ,{
          path: '/home/system2',
          name: 'system2',
          component:  () =>import( '@/views/Empty.vue'),
          redirect: "/home/system2/user",
          meta: {
            title: '测试管理2',
            leftMenu:true
          },
          children:[
            {
              path: '/home/system2/user',
              name: 'user',
              component:  () =>import( '@/views/Empty.vue'),
              redirect:'/home/system/user/query',
              meta: {
                title: '用户管理',
                leftMenu:true
              },
              children:[
                {
                  path: '/home/system2/user/query',
                  name: 'user-list',
                  component:  () =>import( '@/views/UserList.vue'),
                  meta: {
                    title: '用户查询',
                    leftMenu:true
                  }
               },
                {
                  path: '/home/system2/user/add',
                  name: 'user-add',
                  component:  () =>import( '@/views/AddForm.vue'),
                  meta: {
                    title: '用户新增',
                    leftMenu:false
                  },
                }  
                ,{
                  path: '/home/system2/user/edit/:id',
                  name: 'user-update',
                  component:  () =>import( '@/views/AddForm.vue'),
                  meta: {
                    title: '用户编辑',
                    leftMenu:false
                  },
                }                                                         
              ]
            } ,
            {
              path: '/home/system2/children1',
              name: 'system2-1',
              component:  () =>import( '@/views/Test2-1.vue'),
              meta: {
                title: '测试管理2-1',
                leftMenu:true
              },
             
           },
            {
              path: '/home/system2/children2',
              name: 'system2-2',
              component:  () =>import( '@/views/Tabs.vue'),
              meta: {
                title: '测试管理2-2',
                leftMenu:true
              },
            }                                                                           
          ]
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


