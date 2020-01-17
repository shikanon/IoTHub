
import Submenu from './Submenu'
import Breadcrumb from './Breadcrumb'
import Pagination from './Pagination'
import CopyBtn from './CopyBtn'
import Topic from '@/views/topic/Topic'
import AddLabel from './AddLabel'
import JsonView from 'vue-json-views'
import JsonExcel from 'vue-json-excel'
import UseStep from '@/views/useStep/UseStep'
import AceEditor from 'vue-ace-editor'
import VueJsonEditor from 'vue-json-editor'


import LeftMenu from '@/views/knowledge-base/LeftMenu'
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
    Vue.component('LeftMenu',LeftMenu)
    Vue.component('downloadExcel', JsonExcel)
    Vue.component('UseStep', UseStep)
    Vue.component('AceEditor', AceEditor)
    Vue.component('VueJsonEditor', VueJsonEditor)

    
  
}

export default plugin