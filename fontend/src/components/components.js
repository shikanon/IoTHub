
import Submenu from './Submenu'
import Breadcrumb from './Breadcrumb'
import Pagination from './Pagination'
import CopyBtn from './CopyBtn'
import Topic from '@/views/topic/Topic'
import AddLabel from './AddLabel'
import JsonView from 'vue-json-views'


function plugin(Vue){

    if(plugin.installed){
        return 
    }
    
    Vue.component('Pagination',Pagination)
    Vue.component('Submenu',Submenu)
    Vue.component('Breadcrumb',Breadcrumb)
    Vue.component('CopyBtn',CopyBtn)
    Vue.component('Topic',Topic)
    Vue.component('AddLabel',AddLabel)
    Vue.component('JsonView',JsonView)

  
}

export default plugin