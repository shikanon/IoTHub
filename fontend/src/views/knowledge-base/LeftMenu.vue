<template>
    <el-menu :default-active="'0'"  @select="handleSelect">
         <!-- <template v-for="(item,i) in menuList.data"  >
            <el-menu-item :key="i" :index="i" > {{item.name}}</el-menu-item>
        </template>     -->
        <template v-for="(item,i) in menuList"  >
            <el-menu-item :key="i" :index="i" > {{item}}</el-menu-item>
        </template>          
     </el-menu>
</template>

<script>
export default {
    props:{
        index:{
            type:Number,
            default:1
        }
    },
    data(){
        return{
            menuList:{}
        }
    },
    created(){
        this.getMenus()
    },
    methods:{
        getMenus(){

            this.$API_SOTA.getIndexDtlList(this.index).then((res) => {
                this.menuList =  res.data 
                this.handleSelect(0)
            })


            // this.menuList =  {"data":[{"id":"8a0ace81-a1e8-44ec-962f-dbe351a26e37","name":"自然语言处理"},
            //  {"id":"7b34c318-efc0-4f0c-b180-fe75771b9ee6","name":"计算机视觉"},
            //  {"id":"3d358464-c9bc-40fa-9ba8-9e519e3a71c4","name":"机器人技术"},
            //  {"id":"72a098ee-4f1d-4b87-a745-c0fcee0cd3eb","name":"基础理论"},
            //  {"id":"046371d4-7e3f-4b2f-8c25-c5ec17e004cd","name":"运筹优化"},
            //  {"id":"200115c2-e410-4dd8-b8ce-2b24cb3825a5","name":"因果推理"},
            //  {"id":"89daa72e-a534-4710-8057-1656607de8e6","name":"知识工程"},
            //  {"id":"da44ca52-ca81-4069-8501-281af4aac977","name":"数据科学"},
            //  {"id":"9fabf495-478d-473d-a964-6780812064e1","name":"模型优化"},
            //  {"id":"9e5dbab7-ac4f-4ed9-8396-8a8cfdca1817","name":"信息检索"},
            //  {"id":"bb25eb9c-1499-4bb0-b329-bed31ad99ae4","name":"推荐系统"},
            //  {"id":"1a0cec26-30a2-420e-adfa-d198186995a9","name":"自动驾驶技术"},
            //  {"id":"8731d902-b40a-43ff-9e8b-b0c555027e83","name":"生成对抗网络"},
            //  {"id":"e46bc8e6-77b2-4f2f-a4ed-ad0137193fbc","name":"机器学习"},
            //  {"id":"817eda25-04bc-4a2b-b78f-867a696996da","name":"自动化机器学习"},
            //  {"id":"3d2e2a3e-2640-46bb-8380-b0e5a693a10c","name":"深度学习"},
            //  {"id":"6ecb9e8c-cbee-4eeb-a2f3-ee2f82a6c234","name":"强化学习"},
            //  {"id":"a7a585d7-4339-4d4c-adbb-5075a72429a2","name":"多模态学习"},
            //  {"id":"a8b41c8d-f8e5-4d0a-942e-80e030c1626e","name":"迁移学习"},
            //  {"id":"002cdfc9-fd09-4ce4-86fc-894c76117df6","name":"元学习"},
            //  {"id":"93a45904-064b-423b-abaa-c216d52470d3","name":"联邦学习"},
            //  {"id":"825becba-0b0c-414a-bc5a-e3415986aff8","name":"演化学习"},
            //  {"id":"75163eaf-dd52-4360-a788-a44a1b91a052","name":"监督学习"},
            //  {"id":"126a71a0-68a9-41ff-8b13-843aae4198e6","name":"自监督学习"},
            //  {"id":"0e530241-eae8-4931-8ba8-b3ba257e01c2","name":"半监督学习"},
            //  {"id":"5751242d-2d00-429b-9c91-be71143903ea","name":"无监督学习"},
            //  {"id":"2759cc33-65ba-44c5-b7ef-c745cc6aefee","name":"大数据技术"},
            //  {"id":"e24ac996-e295-4e32-8788-ad873f6e79d9","name":"集成学习"},
            //  {"id":"007f5c11-3b53-4725-bcbd-8a82afafebab","name":"分布式计算技术"},
            //  {"id":"17ef7592-2156-4c3d-a739-9e795eafdb91","name":"物联网技术"},
            //  {"id":"1762ea61-c37d-46ab-823e-271b9fa2c9d0","name":"沉浸式技术"},
            //  {"id":"3e605f0a-1711-4570-9441-c6628b2912e1","name":"AI 芯片技术"},
            //  {"id":"7475de02-edab-4e8d-b32c-b141738e4a0c","name":"神经计算"},
            //  {"id":"388e7e26-7734-466d-9225-fdb789a5fcc8","name":"量子计算"},
            //  {"id":"14926caa-acff-4222-be94-1d3754da2212","name":"群体智能"},
            //  {"id":"52eb3167-dbba-4043-9feb-d4f3668bddf1","name":"人机共生"},
            //  {"id":"a245c45c-1ee7-4dbf-8b29-488dac26e121","name":"AI 基础设施"},
            //  {"id":"31119361-6399-4726-97ec-5d87199054a2","name":"人工智能安全"},
            //  {"id":"6b821322-b681-4408-8f6d-9350a0e2a961","name":"通用人工智能"},
            //  {"id":"04288a16-2433-4926-8c5c-13a05bfe78a8","name":"其他"}],"total_count":40}

            
         },

        handleSelect(index){
             let name = this.menuList[index]
            //  let name = this.menuList.data[index].name
            //  let id = this.menuList.data[index].id
            //根据id查询
            this.$emit('getInfoById',{id:index,name:name})
         },
    }
}
</script>
<style  scoped>
    .el-menu{
        width:150px;
        overflow-y: auto;
        max-height: calc(100vh - 250px);
    }
    /* ::-webkit-scrollbar {
        width: 5px;
        height: 5px;
    }
    ::-webkit-scrollbar-thumb {
        background: #ccc;
        border-radius: 5px;
    }
    ::-webkit-scrollbar-track {
        background: transparent;
        border-radius: 5px;
    } */
</style>