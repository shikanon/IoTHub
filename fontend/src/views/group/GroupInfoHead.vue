<template>
    <div name="group-inf-head">
        <div class="title-inf" >  
            <div>
                <span class="text-gray title">分组层级</span> 
                <span class="text-gray">分组/{{group.GroupName}}</span>
            </div> 
            <div>
                <span class="text-gray title">分组ID</span> 
                <span class="text-gray">{{group.GroupId}}</span>
                <CopyBtn :content="group.GroupId" size="small"></CopyBtn>        
            </div>  
            <div> </div>
        </div>
         <div class="title-inf" >  
            <div>
                <span class="text-gray title">设备总数</span> 
                <span class="text-gray">{{group.DeviceCount}}</span>
            </div> 
            <div>
                <span class="text-gray title">激活设备</span> 
                <span class="text-gray">{{group.DeviceActive}}</span>
            </div>  
            <div> </div>
        </div>
          <div  class="title-inf" >
            <div>
                <span class="text-gray title">当前在线</span> <span class="text-gray">{{group.DeviceOnline}}</span>
            </div>     
        </div>
  
    </div>    
</template>
<script>
export default {
    props:{
        group:{
            type:Object,
            default: () =>({}),
        },
      
    },
    data(){
        return {
           
        }
    },
    created(){
    },

    methods:{
     
  
    }
}
</script>
<style  scoped>
    
</style>