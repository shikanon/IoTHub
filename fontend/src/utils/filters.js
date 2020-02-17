



export function menuFilter(value) {
        return  `${value}管理`
}
 
//格式化日期
export function dateFilter(value) {
        return  `${value}管理`
}

//格式化时间
export function timeFilter(value) {
        return  `${value}菜单`
}

export function cardValFilter(value,unit) {
        if(value){
              return  value  + unit
        }else{
                return '--' + unit
        }
}


export function deviceStatusFilter(value) {
        switch (value) {
          
                case 1:
                        value = "未激活";
                        break;
                case 2:
                        value = "禁用";
                        break;        
                case 3:
                        value = "在线";
                        break;   
                case 4:
                        value = "离线";
                        break;                        
                } 
        return value
}


export function topicOperationFilter(value) {
        switch (value) {
                case 1:
                        value = "发布";
                        break;
                case 2:
                        value = "订阅";
                        break;
                case 3:
                        value = "发布和订阅";
                        break;
                                                
                } 
        return value
}


export function formatTags(tags) {
        if(tags){
                return tags.map((item) => item).join("·")

        }else{
                return ""
        }
        
}


export function nodeTypeFilter(nodeType) {
        switch (nodeType) {
                case 1:
                        nodeType = "直连设备";
                        break;
                case 2:
                        nodeType = "网关子设备";
                        break;
                case 3:
                        nodeType = "网关设备";
                        break;
                                              
                } 
        return nodeType
}





