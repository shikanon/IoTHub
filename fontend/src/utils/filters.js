



export function menuFilter(value) {
        return  `${value}菜单`
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
        if(unit){
              return value ? value: '--' + unit
        }else{
                return ''
        }
}


export function deviceStatusFilter(value) {
        switch (value) {
                case 'DISABLE':
                        value = "已禁用";
                        break;
                case 'UNACTIVE':
                        value = "未激活";
                        break;
                case 'ONLINE':
                        value = "在线";
                        break;
                case 'OFFLINE':
                        value = "离线";
                        break;                                   
                } 
        return value
}




