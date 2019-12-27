
import Submenu from './Submenu'
import Breadcrumb from './Breadcrumb'
import Pagination from './Pagination'
import CopyBtn from './CopyBtn'
import TopicList from '@/views/topic/TopicList'
import AddLabel from './AddLabel'



function plugin(Vue){

    if(plugin.installed){
        return 
    }
    
    Vue.component('Pagination',Pagination)
    Vue.component('Submenu',Submenu)
    Vue.component('Breadcrumb',Breadcrumb)
    Vue.component('CopyBtn',CopyBtn)
    Vue.component('TopicList',TopicList)
    Vue.component('AddLabel',AddLabel)
  
}

export default plugin