import Vue from 'vue';
import Vuex from 'vuex';
Vue.use(Vuex);
 const state={   //要设置的全局访问的state对象
     loginName:'',
     TOKEN:'',
     isUserEdit:false,
     cancelTokenArr: [], // 取消请求token数组
     themecolor:'00'

   };

const getters = {   //实时监听state值的变化(最新状态)
    
    getLoginName(){  
       return state.loginName
    },
    getTOKEN(){  
        return state.TOKEN
     },
    
};


const mutations = {
 
   
    setLoginName(state,value){  
        state.loginName = value
    },
    setTOKEN(state,token){  
        state.TOKEN = token
    },
  
    pushPending (state, params) {    //增加要取消的请求
        state.cancelTokenArr.push(params)
    },
    deletePending (state, config = {}) {//删除要取消的请求
        if(Object.keys(config).length > 0){       
            let pending = state.cancelTokenArr.find((item) => item.u === config.url + '&' + config.method )
            if(pending && pending.f){
                pending.f() //执行取消操作
                state.cancelTokenArr =  state.cancelTokenArr.filter((item) => item.u !== config.url + '&' + config.method) //把这条请求移除
            }
        }else{
            //跳转路由，取消所有请求
            state.cancelTokenArr.forEach(e => { e && e.f && e.f() });
            state.cancelTokenArr = []
        }
    },
    //更新主题颜色
    setThemeColor(state,curcolor){
        state.themecolor = curcolor
    }

};
const actions = {   //实时监听state值的变化(最新状态)
    

    
};
 const store = new Vuex.Store({
       state,
       getters,
       actions,
       mutations
});


export default store