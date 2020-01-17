import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)


export default new Router({
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/Login.vue')
    },
    {
      path: '/home',
      name: 'home',
      component: () => import( '@/views/Home.vue'),
      redirect:'/index',
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
            icon:'el-icon-s-home',
            leftMenu:true

          },
        },
        {
          path: '/home/system',
          name: 'system',
          component:  () =>import( '@/views/Empty.vue'),
          meta: {
            title: '设备管理',
            icon:'el-icon-menu',
            leftMenu:true
          },
          redirect: "/home/product/product-list",
          children:[
               
            {
            //  path: '/home/product/product-list/:id',
              path: '/home/product/product-list',

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
                  // path: '/home/product/details/:id',
                  path: '/home/product/details',
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
              path: '/home/device',
              // path: '/home/device/:id',
              name: 'device-list',
              component:  () =>import( '@/views/device/DeviceManage.vue'),
              meta: {
                title: '设备',
                leftMenu:true
              },
              children:[
                {
                 // path: '/home/device/details/:id',
                  path: '/home/device/details',
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
             // path: '/home/group/:id',
              path: '/home/group',
              name: 'group-list',
              component:  () =>import( '@/views/group/GroupList.vue'),
              meta: {
                title: '分组',
                leftMenu:true
              },
              children:[
                {
                  path: '/home/group/details',
                //  path: '/home/group/details/:id',

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
            icon:'el-icon-s-help',
            leftMenu:true
          }
        },
        {
            path: '/test2',
            name: 'test2',
            component: () => import( '@/views/Empty.vue'),
            meta: {
              title: '测试2',
              icon:'el-icon-s-help',
              leftMenu:true
            }
        },
        {
          path: '/home/knowledge-base',
          name: 'knowledge-base',
          component: () => import( '@/views/knowledge-base/Base.vue'),
          meta: {
            title: '知识库',
            icon:'el-icon-document',
            leftMenu:true
          },
          children:[
            {
              path: '/home/knowledge-base/api-info/:id',
              name: 'api-info',
              component:  () =>import( '@/views/knowledge-base/ApiInfo.vue'),
              meta: {
                title: 'API接口信息',
                leftMenu:false
              },
            } 
          ]
      },
   
             
      ]
    
    }, 
    {
      path: '/',
      name: 'home',
      component: () => import( '@/views/Home.vue')
    }
  ]
  

 
})


