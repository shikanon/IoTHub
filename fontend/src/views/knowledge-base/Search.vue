<template>
   <div name="search">
        <div class="search_head">
            <p v-if="!result.knowledge">人工智能<span>知识库</span></p>
            <div class="search_content">
                <el-input placeholder="请输入内容" v-model="search_input" class="input-with-select">>
                        <el-button slot="append" icon="el-icon-search" @click="search"></el-button>
                </el-input>
            </div>
        </div>
       <el-divider v-if="!result.knowledge"></el-divider>
       <div class="search_body"  v-if="!result.knowledge">
            <div class="knowledge">
                <p>知识</p>
                <div class="card-list">   
                    <el-card class="box-card" v-for="(item,key) in data.knowledge" :key="key" > 
                        <div @click="clickTab('knowledge',key)">
                            <div class="card-title">
                                    <img src="https://pro.jiqizhixin.com/assets/packs/sites/dashboard/images/wiki_icons/icon-concept-1f5a8b1f2ada42c3a02dfb94f39d1e32.svg" alt="">
                                    <span>{{item.type}}</span>
                            </div>
                            <div class="card-boby">
                                    <span>共记录<span class="number">{{item.total_count}}</span>条</span>
                                    <div>
                                        <div>{{item.data.name}}</div>
                                        <span>最新纪录</span>
                                    </div>
                            </div> 
                        </div>      
                    </el-card>   
                </div>
            </div>
            <div class="resource">
                <p>资源</p>
                <div class="card-list">
                    <el-card class="box-card" v-for="(item,key) in data.resource" :key="key"  >
                        <div @click="clickTab('resource',key)">
                            <div class="card-title" >
                                <img src="https://pro.jiqizhixin.com/assets/packs/sites/dashboard/images/wiki_icons/icon-concept-1f5a8b1f2ada42c3a02dfb94f39d1e32.svg" alt="">
                                <span>{{item.type}}</span>
                            </div>
                            <div class="card-boby" >
                                <span>共记录<span class="number">{{item.total_count}}</span>条</span>
                                <div>
                                    <div>{{item.data.name}}</div>
                                    <span>最新纪录</span>
                                </div>
                            </div>   
                        </div>    
                    </el-card>  
                    <div class="box-card-hidden" v-if="data.resource.length % 2  === 1"></div>
                </div>
            </div>
            <div class="people">
                <p>人物</p>
                <div class="card-list" >
                    <el-card class="box-card"  v-for="(item,key) in data.people" :key="key" >
                         <div @click="clickTab('people',key)">

                            <div class="card-title">
                                    <img src="https://pro.jiqizhixin.com/assets/packs/sites/dashboard/images/wiki_icons/icon-concept-1f5a8b1f2ada42c3a02dfb94f39d1e32.svg" alt="">
                                    <span>{{item.type}}</span>
                            </div>
                            <div class="card-boby">
                                <span>共记录<span class="number">{{item.total_count}}</span>条</span>
                                <div>
                                    <div>{{item.data.name}}</div>
                                    <span>最新纪录</span>
                                </div>
                            </div>   
                        </div>    
                    </el-card>         
                </div>
            </div>
       </div>
        <div class="search_result" v-if="result.knowledge">
                <SearchResult :result="result"></SearchResult>
        </div>
  </div>
</template>

<script>
import SearchResult from './SearchResult';
export default {
    components:{SearchResult},
    data() {
    return {
        search_input:'',
        data:{},
        result:{},
    }
    },
    created(){
        this.init()         
    },
    methods:{

        init(){

            this.$API_SOTA.getIndexList().then((res) => {
                console.log(res.data)          
                this.data = res.data
            })

            // [
            //     {
            //         "id": 1,
            //         "name": "基础模型",
            //         "APIType": 0,
            //         "Image": "icon-concept.svg"
            //     },
            //     {
            //         "id": 2,
            //         "name": "SOTA模型",
            //         "APIType": 0,
            //         "Image": "icon-concept.svg"
            //     },
            //     {
            //         "id": 3,
            //         "name": "行业解决方案",
            //         "APIType": 1,
            //         "Image": "icon-concept.svg"
            //     },
            //     {
            //         "id": 4,
            //         "name": "数据集",
            //         "APIType": 2,
            //         "Image": "icon-concept.svg"
            //     }
            // ]
            this.data = 
            {"knowledge":[
                {"total_count":694,"type_en":"concept","type":"基础概念","icon":"icon-concept.svg","data":{"id":"ce06fe06-610e-4a82-8af3-a7e1c0d5d351","name":"用户画像"}},
                {"total_count":4202,"type_en":"tech_method","type":"技术方法","icon":"icon-method.svg","data":{"id":"c4a5d672-d51e-451f-83cf-dc86b3ea2adb","name":"Γ-model"}},
                {"total_count":516,"type_en":"tech_task","type":"技术任务","icon":"icon-task.svg","data":{"id":"d43c21c2-131f-4166-abe9-ec02e2efeb21","name":"图像聚类"}}
                ],
                
            "resource":[
                    {"total_count":1423,"type_en":"data_set","type":"数据集","icon":"icon-dataset.svg","data":{"id":"32796975-9890-4966-af95-5f5309f78336","name":"SQuAD"}},
                    {"total_count":956,"type_en":"tool","type":"开发工具","icon":"icon-tool.svg","data":{"id":"3feca874-32fb-4694-9ca1-82c37e4fcec6","name":"nullabor"}},
                    {"total_count":778,"type_en":"tutorial","type":"教程","icon":"icon-tutorial.svg","data":{"id":"778e86c6-c5f2-4928-b7f0-9b2f6f6f3112","name":"图神经网络相关资料"}},
                    {"total_count":344,"type_en":"event","type":"活动及会议","icon":"icon-event.svg","data":{"id":"955766d6-8a05-41e6-91e8-e3ecab344739","name":"SEKE  International Conference on Software Engineering \u0026 Knowledge Engineering"}},
                    {"total_count":366,"type_en":"publication","type":"书籍期刊","icon":"icon-publication.svg","data":{"id":"fb957a10-70fe-4ba2-b854-9ad9d7ab0bb4","name":"IEEE transactions of industrial informatics"}}
                    ],
            "people":[
                    {"total_count":2087,"type_en":"expert","type":"专家","icon":"icon-expert.svg","data":{"id":"454ea7d2-f5d2-4e1f-bac6-dd2d74400490","name":"陈文亚"}}
                    ]
            }
            
            
        },      
        clickTab(name,index){
            this.$emit('click',{name:name,index:index})
        },

        search(){
            if(this.search_input){
                this.result ={
                    "knowledge":{
                        "total_count":148,
                        "data":[
                            {
                            "model_type": "tech_method",
                            "name": "深度学习",
                            "id": "01946acc-d031-4c0e-909c-f062643b7273",
                            "tags": [
                            "技巧",
                            "技术方法",
                            "深度学习"
                            ]
                            },
                            {
                            "model_type": "tech_method",
                            "name": "自动化深度学习",
                            "id": "7d404f13-b4d8-4bee-952a-46c596a865d3",
                            "tags": [
                            "技巧",
                            "技术方法",
                            "自动化机器学习"
                            ]
                            },
                            {
                            "model_type": "tech_method",
                            "name": "移动端深度学习",
                            "id": "d484e2f3-bfd1-47c8-a430-db148416b574",
                            "tags": []
                            },
                            {
                            "model_type": "tech_task",
                            "name": "3D目标分类",
                            "id": "06ac415b-1306-4b39-9cb5-d4fc3db34f7d",
                            "tags": [
                            "技术任务",
                            "计算机视觉"
                            ]
                            },
                            {
                            "model_type": "tech_method",
                            "name": "深度神经网络",
                            "id": "f82b7976-b182-40fa-b7d8-a3aad9952937",
                            "tags": [
                            "模型",
                            "技术方法",
                            "深度学习"
                            ]
                            }, 
                        ]
                    },
                    "resource":{
                        "total_count":32,
                        "data":[
                            {
                        "model_type": "tutorial",
                        "name": "移动深度学习相关资料",
                        "id": "7b4a88bd-f64f-434d-97c9-dc8b5fd52429",
                        "tags": [
                        "资料集",
                        "教程与实现",
                        "深度学习"
                        ]
                        },
                        {
                        "model_type": "tool",
                        "name": "一文概览深度学习中的激活函数",
                        "id": "97c094ce-7e33-4c2b-8431-6891dc2b9442",
                        "tags": [
                        "工具包",
                        "深度学习"
                        ]
                        },
                        ]
                    },
                    "people":{
                        "total_count":41,
                        "data":[
                            {
                        "model_type": "expert",
                        "name": "Guilin Liu",
                        "id": "9259c26c-1f2b-4c1b-985c-9f4606712f00",
                        "tags": [
                        "人物",
                        "计算机视觉"
                        ]
                        },
                        ]
                    }
                }
            }else{
                this.result = {}
             }
        }
    }
}
</script>

<style scoped>
    #search{
            display: flex;
            flex-direction: column;
    }
    .search_head{
            text-align: center
        }
        .search_head p{
            font-size: 40px;
            font-weight: 700;
            color: #333;
            margin: 10px 0 30px;
            text-align: center;
        }
        .search_head span{
            color: #4676aa;
        }

    .search_content{
            display: flex;
            width: 50%;
            background: #f2f2f2;
            margin: 0 auto;
            border-radius: 25px;
            align-items: center;
            font-size: 16px;
    }
    .el-input__inner{
            height: 50px;
    }
    .el-divider{
            height: 2px;
            margin: 50px 0 0 0;
    }
    .search_body{
            display: flex;
            justify-content: space-between;
            margin: 0 auto;
            max-width: 1140px;
            padding: 32px 30px 0;
            height: 60%;

    }
    .search_body>div>p{
        margin-left:20px;
    }

    .knowledge,.people{
            flex:1;
    }

    .resource{
            flex:2;
            display: flex;
            flex-direction: column;
    }
    .card-list{
            display: flex;
            flex-wrap: wrap;
            justify-content: space-around;
    }
    .box-card,.box-card-hidden{
            margin: 20px 0px;
            width: 263px;
    }

    .box-card:hover{
        background: #4676aa;
        color: #fff;
    }
    .card-title{
            display: flex;
    }
    .card-title span{
            line-height: 35px;
            padding-left: 15px;
    }

    .card-boby{
            display: flex;
            align-items: flex-end;
            justify-content: space-between;
            text-align: right;
    }

    .card-boby span{
            font-size: 12px;
                color: #bbb;
    }
    .card-boby span .number{color: #4676aa;}
    .card-boby div{
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        width: 120px;
    }
    .box-card:hover span .number{
        color: #fff;
    } 
</style>