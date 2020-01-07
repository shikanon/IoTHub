<template>
    <el-tabs   @tab-click="handleClick" >  
        <el-tab-pane :label="`知识(${result.knowledge.total_count})`" name="0">                         
            <div class="card-list">
                <el-card v-for="(item,key) in result.knowledge.data" :key="key" class="box-card" >
                    <div @click="gotoApiDtl(item)">
                        <p>{{item.name}}</p>
                        <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                    </div>
                </el-card>                   
            </div> 
        </el-tab-pane>   
            <el-tab-pane :label="`资源(${result.resource.total_count})`" name="1">
            <div class="card-list">
                <el-card v-for="(item,key) in result.resource.data" :key="key" class="box-card" >
                    <div @click="gotoApiDtl(item)">
                        <p>{{item.name}}</p>
                        <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                    </div>
                </el-card>                   
            </div> 
        </el-tab-pane> 
            <el-tab-pane :label="`人物(${result.people.total_count})`" name="2">
            <div class="card-list">
                <el-card v-for="(item,key) in result.people.data" :key="key" class="box-card" >
                    <div @click="gotoApiDtl(item)">
                        <p>{{item.name}}</p>
                        <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                    </div>
                </el-card>                   
            </div> 
        </el-tab-pane> 
    </el-tabs>
</template>
<script>
export default {
    props:{
        result:{
            type:Object,
            default:() => ({})
        }
    }
}
</script>

<style scoped>
 .card-list {
    display: flex;
    flex-wrap: wrap;
}

.box-card {
    width: 295px;
    margin: 0 20px 20px 0;
    position: relative;
    height: 120px;
}
.box-card:hover{
     background: #4676aa;
     color: #fff;
 }
.box-card span{
    font-size: 14px;
    color: #999;
} 
.box-card:after {
    content: "";
    position: absolute;
    right: 0;
    bottom: 0;
    width: 0;
    height: 0;
    border-color: transparent #4676aa #4676aa transparent;
    border-style: solid;
    border-width: 6px;
    border-radius: 0 0 6px 0;
}
</style>