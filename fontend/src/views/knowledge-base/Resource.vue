<template>
    <el-container name="knowledge">
        <el-aside width="160px">
            <div class="sub-title">
                <i class="el-icon-document"></i>
                <span slot="title">技术领域</span>
            </div>
            <LeftMenu @getInfoById="getInfoById($event)"></LeftMenu>
        </el-aside>
        <el-main>  
           <p class="name">{{name}}</p>                       
            <div>
                 <el-input placeholder="输入关键字进行搜索" v-model="search_input" class="input-with-select" size="medium">
                        <el-button slot="append" icon="el-icon-search"></el-button>
                </el-input>
                <el-tabs v-model="activeName"  @tab-click="handleClick" >
                    <el-tab-pane :label="`数据集(${info.data_sets.total_count})`" name="0">
                         <div class="card-list">
                            <div v-for="(item,key) in info.data_sets.data" :key="key" class="box-card" @click="gotoApiDtl">
                                <div>                             
                                   <img :src="item.avatar_url" class="small-img"> 
                                </div>
                                <div>
                                    <p>{{item.name}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                    <i class="el-icon-paperclip"></i>
                                </div>        
                            </div>    
                            <el-divider></el-divider>               
                        </div>  
                    </el-tab-pane>
                    <el-tab-pane :label="`开发工具(${info.tools.total_count})`" name="1">
                         <div class="card-list">
                            <div v-for="(item,key) in info.tools.data" :key="key" class="box-card" @click="gotoApiDtl">
                                <div>                             
                                   <img :src="item.avatar_url" class="small-img"> 
                                </div>
                                <div>
                                    <p>{{item.name}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                    <i class="el-icon-paperclip"></i>
                                </div>        
                            </div>    
                            <el-divider></el-divider>               
                        </div> 
                    </el-tab-pane>    
                    <el-tab-pane :label="`教程(${info.tutorials.total_count})`" name="2">
                         <div class="card-list">
                            <div v-for="(item,key) in info.tutorials.data" :key="key" class="box-card" @click="gotoApiDtl">
                                <div>                             
                                   <img :src="item.avatar_url" class="small-img"> 
                                </div>
                                <div>
                                    <p>{{item.name}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                    <i class="el-icon-paperclip"></i>
                                </div>        
                            </div>    
                            <el-divider></el-divider>               
                        </div> 
                    </el-tab-pane>    
                    <el-tab-pane :label="`活动及会议(${info.events.total_count})`" name="3">
                          <div class="card-list">
                            <div v-for="(item,key) in info.events.data" :key="key" class="box-card" @click="gotoApiDtl">
                                <div>                             
                                   <img :src="item.cover_image_url" class="big-img"> 
                                </div>
                                <div>
                                    <p>{{item.name}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                    <i class="el-icon-paperclip"></i>
                                </div>        
                            </div>    
                            <el-divider></el-divider>               
                        </div> 
                    </el-tab-pane>  
                    <el-tab-pane  :label="`书籍期刊(${info.publications.total_count})`" name="4">
                          <div class="card-list">
                            <div v-for="(item,key) in info.publications.data" :key="key" class="box-card" @click="gotoApiDtl">
                                <div>                             
                                   <img :src="item.cover_image_url" class="big-img"> 
                                </div>
                                <div>
                                    <p>{{item.name}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                    <i class="el-icon-paperclip"></i>
                                </div>        
                            </div>    
                            <el-divider></el-divider>               
                        </div> 
                    </el-tab-pane>                                
                </el-tabs>
            </div>
        </el-main> 
    </el-container>
 
</template>

<script>
export default {
     props:{
        index:{
            type:String,
            default:'0'
        }
    },
     computed:{

        activeName:{
            get:function () {
                return this.index
            },
            set:function () {

            }
        }      
    },

     data() {
        return { 
            name:'',
            search_input:'',
            info:{}
        };
      },
  
    created(){  
        this.getInfoById("123123213")
    },
     methods:{      
         getInfoById(event){
             console.log(event)
             this.name = event.name
             this.info= {
                "id": "7b34c318-efc0-4f0c-b180-fe75771b9ee6",
                "name": "计算机视觉",
                "data_sets":{
                    "total_count":2,
                    "data":[
                        {
                        "id": "468d12b3-3e81-43b5-a449-88992622cdab",
                        "name": "LSTM-CF",
                        "avatar_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/Z/135_149_169/160.png",
                        "summary": "",
                        "reference_urls": {
                        "LSTM-CF": "https://github.com/icemansina/LSTM-CF"
                        },
                        "institutions": [],
                        "tags": []
                        },
                        {
                        "id": "90f2c04c-f4ce-4fd6-9d6b-be0d568046fc",
                        "name": "Eigenfaces－python",
                        "avatar_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/E/113_82_86/160.png",
                        "summary": "",
                        "reference_urls": {
                        "Eigenfaces－python": "https://github.com/agyorev/Eigenfaces"
                        },
                        "institutions": [],
                        "tags": []
                        },
                    ]
                  
                },
                "tools":{
                    "total_count":8,
                      "data":[
                        {
                        "id": "d5d78e38-5f71-41aa-b27b-730243e58e92",
                        "name": "zinc",
                        "avatar_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/Z/135_149_169/160.png",
                        "summary": null,
                        "reference_urls": {},
                        "tags": []
                        },
                        {
                        "id": "cfe3ae87-a6cd-404b-b19a-c0c442387187",
                        "name": "wireframe dataset",
                        "avatar_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/W/203_185_164/160.png",
                        "summary": null,
                        "reference_urls": {},
                        "tags": []
                        },
                    ]
                },
                "tutorials":{
                    "total_count":6,
                    "data":[
                        {
                            "id": "d5f04a7f-922b-4c90-9137-bf6935556b50",
                            "name": "Image inpainting 相关资料",
                            "avatar_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/I/160_80_84/160.png",
                            "summary": "",
                            "reference_urls": {
                            "Introduction to Image Inpainting|Coursera": "https://www.coursera.org/lecture/image-processing/1-introduction-to-image-inpainting-duration-08-16-oIocq",
                            "An example of how to do image inpainting": "https://sites.google.com/site/rexstribeofimageprocessing/Home/image-inpainting",
                            "心中无码：这是一个能自动脑补漫画空缺部分的AI项目|机器之心": "https://www.jiqizhixin.com/articles/102901?from=synced&keyword=图像修复",
                            "用自注意力GAN为百年旧照上色：效果惊艳，多图预警！|机器之心": "https://www.jiqizhixin.com/articles/2018-11-03-3?from=synced&keyword=图像修复"
                            },
                            "institutions": [],
                            "tags": [
                            "资料集"
                            ]
                            },
                            {
                            "id": "44e27b53-fec6-4f12-b539-e21fa239bef3",
                            "name": "语义分割相关资料",
                            "avatar_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/Y/105_99_105/160.png",
                            "summary": "",
                            "reference_urls": {
                            "鱼与熊掌可以兼得：何恺明等人一统语义分割与实例分割": "https://www.jiqizhixin.com/articles/2019-01-11-13?from=synced&keyword=%E8%AF%AD%E4%B9%89%E5%88%86%E5%89%B2",
                            "AI所有领域最优论文+代码查找神器：966个ML任务、8500+论文任你选": "https://www.jiqizhixin.com/articles/020201?from=synced&keyword=%E8%AF%AD%E4%B9%89%E5%88%86%E5%89%B2",
                            "Deep Learning in Computer Vision": "https://www.coursera.org/lecture/deep-learning-in-computer-vision/image-segmentation-o5e7h",
                            "CS 231n detection and segmentation": "http://cs231n.stanford.edu/2017/syllabus"
                            },
                            "institutions": [],
                            "tags": [
                            "资料集"
                            ]
                            }
                    ]
                },
                "events":{
                    "total_count":13,
                    "data":[
                        {
                        "id": "d6601f31-842f-45bb-bdc2-81d5f5188e8f",
                        "name": "International Conference on Awareness Science and Technology (iCAST)",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/I/129_112_135/160.png",
                        "summary": "",
                        "reference_urls": {},
                        "institutions": [],
                        "begin_at": null,
                        "end_at": null,
                        "tags": [
                        "学术会议"
                        ]
                        },
                        {
                        "id": "fc779ab6-cb1d-40c8-8247-f55f40cc6a4a",
                        "name": "ICDAR (International Conference on Document Analysis and Recognition)",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/I/135_149_169/160.png",
                        "summary": null,
                        "reference_urls": {},
                        "institutions": [],
                        "begin_at": null,
                        "end_at": null,
                        "tags": [
                        "学术会议"
                        ]
                        },
                        {
                        "id": "e77b0f06-eb12-43fe-b49b-385c4d550db0",
                        "name": "NeurIPS (Conference on Neural Information Processing Systems)",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/N/183_178_165/160.png",
                        "summary": "NeurIPS，全称神经信息处理系统大会(Conference and Workshop on Neural Information Processing Systems)，是一个关于机器学习和计算神经科学的国际顶级会议。该会议固定在每年的12月举行，由NIPS基金会主办 。在中国计算机学会的国际学术会议排名中，NeurIPS为人工智能领域的A类会议。曾用缩写NIPS，2018年，因 NIPS 缩写存在的「暧昧的内涵」，带有性别歧视的意味，更名为NeurIPS。2018年 NeurIPS会议于12月3日-8日在加拿大蒙特利尔举办。",
                        "reference_urls": {},
                        "institutions": [],
                        "begin_at": null,
                        "end_at": null,
                        "tags": [
                        "学术会议"
                        ]
                        },
                        {
                        "id": "e44fab32-8f8d-43a7-891b-9443d8e01e95",
                        "name": "ICLR (International Conference on Learning Representations)",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/I/105_99_105/160.png",
                        "summary": "ICLR是深度学习领域的盛会，每年举办一次。2013 年，深度学习巨头 Yoshua Bengio、Yann LeCun 主持举办了第一届 ICLR 大会。经过几年的发展，在深度学习火热的今天，ICLR 已经成为人工智能领域不可错过的盛会之一。",
                        "reference_urls": {},
                        "institutions": [],
                        "begin_at": null,
                        "end_at": null,
                        "tags": [
                        "学术会议"
                        ]
                        },
                        {
                        "id": "eb4d4903-29e9-4c57-90c7-eb2c169c626e",
                        "name": "ACL (Annual Meeting of the Association for Computational Linguistics)",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/A/113_82_86/160.png",
                        "summary": null,
                        "reference_urls": {},
                        "institutions": [],
                        "begin_at": null,
                        "end_at": null,
                        "tags": [
                        "学术会议"
                        ]
                        },
                    ]
                },
                "publications":{
                    "total_count":23,
                    "data":[
                        {
                        "id": "e9ef43b8-b3ea-40e9-9bbe-3970b8ef203d",
                        "name": "Computer Graphics",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/C/129_112_135/160.png",
                        "summary": "",
                        "reference_urls": {},
                        "tags": [
                        "学术期刊"
                        ]
                        },
                        {
                        "id": "f96702c7-fd6c-44ae-8aa0-1e5a124366e3",
                        "name": "TPAMI (IEEE Transactions on Pattern Analysis and Machine Intelligence)",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/T/135_149_169/160.png",
                        "summary": null,
                        "reference_urls": {},
                        "tags": [
                        "学术期刊"
                        ]
                        },
                        {
                        "id": "f95df68b-517a-4b5e-a879-62b5027bbe18",
                        "name": "Proceedings of the IEEE",
                        "cover_image_url": "https://www.jiqizhixin.com/synced/letter_avatars/2/P/105_99_105/160.png",
                        "summary": null,
                        "reference_urls": {},
                        "tags": [
                        "学术期刊"
                        ]
                        }
                    ]
                }
             }
         },
 
         handleClick(tab, event) {
            //console.log(tab, event);
           // this.getCards(tab.index)
        },

         gotoApiDtl(item){
            this.$emit('click',item.id)
        },
     }
}
</script>

<style scoped>
.sub-title{    
    height: 50px;
    line-height: 50px;
}

 .head{
    display: flex;
    justify-content: space-between;
 }
 .head img{margin-left:90px}
 .name{
    font-size: 28px;
    font-weight: 500;
    color: #333;
    margin-top: 0px;
 }
 
.card-list{
    display: flex;
    flex-wrap: wrap;
}
 .box-card{
     width: 100%;
     padding: 10px 20px;
     position: relative;
     display: flex;
    align-items: center;
    border-bottom: 1px solid #ccc;
 }
.box-card .el-avatar{
    margin-right: 20px;
}


 .box-card p{
     margin-top: 0px;
} 
.box-card span{
    font-size: 14px;
    color: #999;
}   
.el-icon-paperclip{
    position: absolute;
    right: 20px;
    top: 40%;
}

.input-with-select{
    width: 300px;
    position: absolute;
    right: 35px;
}

.small-img{
    width: 50PX;
    height: 50px;
    margin-right: 20px;
}

.big-img{
    width: 176px;
    height: 110px;
    margin-right: 20px;
}
</style>