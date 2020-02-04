<template>
    <el-container name="knowledge">
        <el-aside width="160px">
            <div class="sub-title">
                <i class="el-icon-document"></i>
                <span slot="title">技术领域</span>
            </div>
            <LeftMenu @getInfoById="getInfoById($event)" :index="index"></LeftMenu>                  
        </el-aside>
        <el-main>  
            <div class="head">
                <div>
                    <p class="name">{{name }}</p>
                    <div class="desc">{{info.APIClassDescription}}</div>
                </div> 
                <div>
                 <img src="https://pro.jiqizhixin.com/assets/packs/sites/dashboard/images/comingSoon-5535ea724d5f84a2e07383ca4b8132d0.svg"> 
                </div> 
            </div>
            <div v-if="info.data">
                 <el-tabs> 
                     <el-tab-pane value="0"  label="相关API"> 
                        <div class="card-list">
                                <el-card v-for="(item,key) in info.data" :key="key" class="box-card" >
                                    <div @click="gotoApiDtl(item)">
                                        <p>{{item.APIGroup}}</p>
                                        <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                    </div>
                                </el-card>                   
                            </div> 
                     </el-tab-pane>
                 </el-tabs>       
            </div>
            <!-- <div v-if="info.tech_tasks">
                
                <el-tabs v-model="index"  @tab-click="handleClick" >  
                    <el-tab-pane :label="`技术任务(${info.data.length})`" name="2">                         
                        <div class="card-list">
                            <el-card v-for="(item,key) in info.data" :key="key" class="box-card" >
                                <div @click="gotoApiDtl(item)">
                                    <p>{{item.APIGroup}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                </div>
                            </el-card>                   
                        </div> 
                    </el-tab-pane>   
                     <el-tab-pane :label="`技术方法(${info.tech_methods.total_count})`" name="1">
                        <div class="card-list">
                            <el-card v-for="(item,key) in info.tech_methods.data" :key="key" class="box-card" >
                                <div @click="gotoApiDtl(item)">
                                    <p>{{item.name}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                </div>
                            </el-card>                   
                        </div> 
                    </el-tab-pane> 
                      <el-tab-pane :label="`基础概念(${info.concepts.total_count})`" name="0">
                        <div class="card-list">
                            <el-card v-for="(item,key) in info.concepts.data" :key="key" class="box-card" >
                                <div @click="gotoApiDtl(item)">
                                    <p>{{item.name}}</p>
                                    <span><i class="el-icon-price-tag"></i>{{item.tags | formatTags}}</span>
                                </div>
                            </el-card>                   
                        </div> 
                    </el-tab-pane> 
                </el-tabs> 
            </div> -->
        </el-main>    
    </el-container>
</template>

<script>
export default {
    props:{
        index:{
            type:Number,
            default:1
        },
        arrData:{
            type:Array,
            default:() => []
        }
    },
     data() {
        return { 
            name:'',
            info:{},
            selectedItem:{}
        };
      },
  
    created(){
    },
     methods:{
         getInfoById(event){
            this.name = event.name
            this.$API_SOTA.getApiList(this.name ).then((res) => {
                this.info =  res.data              
            })
            
            //  this.info = {
            //  "id":"8a0ace81-a1e8-44ec-962f-dbe351a26e37",
            //  "name":"自然语言处理",
            //  "summary":"自然语言处理（英语：natural language processing，缩写作 NLP）是人工智能和语言学领域的分支学科。此领域探讨如何处理及运用自然语言；自然语言认知则是指让电脑“懂”人类的语言。自然语言生成系统把计算机数据转化为自然语言。自然语言理解系统把自然语言转化为计算机程序更易于处理的形式。",
            //  "tech_tasks":{
            //     "total_count":4,
            //     "data":[
            //         {"id":"bdee18a9-b091-4ea8-a934-5da1220389dd",
            //         "name":"语音身份分离",
            //         "summary":"在人机（电脑）交互（HCI）和人机（机器人）交互（HRI）中，通常需要解决多方对话问题。例如，如果两个或更多个人参与对话，则在自动语音识别（ASR）和自然语言处理（NLP）之前要解决的一个重要任务是将某一时间的语音段正确地分配给相应的说话者。在语音和语言处理文献中，这个问题被称为语音身份分离，或“who speaks when？”。",
            //         "completeness":"level3",
            //         "tags":["专业技术任务"]},
            //         {"id": "12cbd40e-d5ae-4170-86f7-4f02b8d60c06",
            //             "name": "（文本）情感分析",
            //             "summary": "文本情感分析是指用自然语言处理、文本挖掘以及计算机语言学等方法来识别和提取原素材中的主观信息。 通常来说，情感分析的目的是为了找出说话者/作者在某些话题上或者针对一个文本两极的观点的态度。这个态度或许是他或她的个人判断或是评估，也许是他当时的情感状态，或是作者有意向的情感交流。",
            //             "completeness": "level3",
            //             "tags": [
            //             "专业技术任务"
            //             ]
            //         },
            //         {"id":"bdee18a9-b091-4ea8-a934-5da1220389dd",
            //         "name":"自然语言处理",
            //         "summary":"在人机（电脑）交互（HCI）和人机（机器人）交互（HRI）中，通常需要解决多方对话问题。例如，如果两个或更多个人参与对话，则在自动语音识别（ASR）和自然语言处理（NLP）之前要解决的一个重要任务是将某一时间的语音段正确地分配给相应的说话者。在语音和语言处理文献中，这个问题被称为语音身份分离，或“who speaks when？”。",
            //         "completeness":"level3",
            //         "tags":["专业技术任务"]},
            //         {"id":"bdee18a9-b091-4ea8-a934-5da1220389dd",
            //         "name":"文本分类",
            //         "summary":"在人机（电脑）交互（HCI）和人机（机器人）交互（HRI）中，通常需要解决多方对话问题。例如，如果两个或更多个人参与对话，则在自动语音识别（ASR）和自然语言处理（NLP）之前要解决的一个重要任务是将某一时间的语音段正确地分配给相应的说话者。在语音和语言处理文献中，这个问题被称为语音身份分离，或“who speaks when？”。",
            //         "completeness":"level3",
            //         "tags":["专业技术任务"]},
            //     ]
            //     },
            //  "tech_methods": {
            //     "total_count":9,
            //     "data":[
            //         {
            //         "id": "d0913904-a3ce-4316-8a0c-c5ddb39ebf08",
            //         "name": "隐含狄利克雷分布",
            //         "summary": "隐含狄利克雷分布简称LDA(Latent Dirichlet allocation)，是一种主题模型，它可以将文档集中每篇文档的主题按照概率分布的形式给出。同时它是一种无监督学习算法，在训练时不需要手工标注的训练集，需要的仅仅是文档集以及指定主题的数量k即可。此外LDA的另一个优点则是，对于每一个主题均可找出一些词语来描述它。LDA首先由Blei, David M.、吴恩达和Jordan, Michael I于2003年提出，目前在文本挖掘领域包括文本主题识别、文本分类以及文本相似度计算方面都有应用。",
            //         "tags": [
            //         "模型"
            //         ]
            //         },
            //         {
            //         "id": "c61ba3b9-40e2-4864-a941-9adc19e6792e",
            //         "name": "word2vec",
            //         "summary": "Word2vec，为一群用来产生词向量的相关模型。这些模型为浅而双层的神经网络，用来训练以重新建构语言学之词文本。网络以词表现，并且需猜测相邻位置的输入词，在word2vec中词袋模型假设下，词的顺序是不重要的。\r\n训练完成之后，word2vec模型可用来映射每个词到一个向量，可用来表示词对词之间的关系。该向量为神经网络之隐藏层。\r\nWord2vec依赖skip-grams或连续词袋（CBOW）来建立神经词嵌入。Word2vec为托马斯·米科洛夫（Tomas Mikolov）在Google带领的研究团队创造。该算法渐渐被其他人所分析和解释。",
            //         "tags": [
            //         "模型"
            //         ]
            //         },
            //         {
            //         "id": "e49b21d8-935a-4da6-910d-504c79b9785f",
            //         "name": "主题模型",
            //         "summary": "主题模型（Topic Model）在机器学习和自然语言处理等领域是用来在一系列文档中发现抽象主题的一种统计模型。直观来讲，如果一篇文章有一个中心思想，那么一些特定词语会更频繁的出现。比方说，如果一篇文章是在讲狗的，那“狗”和“骨头”等词出现的频率会高些。如果一篇文章是在讲猫的，那“猫”和“鱼”等词出现的频率会高些。而有些词例如“这个”、“和”大概在两篇文章中出现的频率会大致相等。但真实的情况是，一篇文章通常包含多种主题，而且每个主题所占比例各不相同。因此，如果一篇文章10%和猫有关，90%和狗有关，那么和狗相关的关键字出现的次数大概会是和猫相关的关键字出现次数的9倍。一个主题模型试图用数学框架来体现文档的这种特点。主题模型自动分析每个文档，统计文档内的词语，根据统计的信息来断定当前文档含有哪些主题，以及每个主题所占的比例各为多少。",
            //         "tags": [
            //         "模型"
            //         ]
            //         },
            //         {
            //         "id": "b08fead2-c1b9-4476-9219-d73172b54851",
            //         "name": "Transformer-XL",
            //         "summary": "Transformer-XL 预训练模型是对 Transformer 及语言建模的修正，这项前沿研究是2019年1月份公布。一般而言，Transformer-XL 学习到的长期依赖性比标准 Transformer 学到的长 450%，无论在长序列还是短序列中都得到了更好的结果，而且在评估时比标准 Transformer 快 1800 多倍。",
            //         "tags": [
            //         "模型"
            //         ]
            //         },
            //         {
            //         "id": "dac652ba-7bac-4d1b-8885-fc935974ba23",
            //         "name": "MT-DNN",
            //         "summary": "MT-DNN 是微软提出的在多种自然语言理解任务上学习表征的多任务深度神经网络。与 BERT 模型类似，MT-DNN 分两个阶段进行训练：预训练和微调。与 BERT 不同的是，MT-DNN 在微调阶段使用多任务学习，在其模型架构中具有多个任务特定层。",
            //         "tags": [
            //         "算法"
            //         ]
            //         },
            //         {
            //         "id": "20ed7263-df85-4418-8838-ad1b39c5efcd",
            //         "name": "GloVe",
            //         "summary": "Stanford开发的用于词向量表示的一个库/工具",
            //         "tags": [
            //         "常用"
            //         ]
            //         },
            //         {
            //         "id": "e8e4513e-695d-49e6-9155-31ae453de372",
            //         "name": "GPT-2",
            //         "summary": "GPT-2是OpenAI于2019年2月发布的基于 transformer 的大型语言模型，包含 15 亿参数、在一个 800 万网页数据集上训练而成。据介绍，该模型是对 GPT 模型的直接扩展，在超出 10 倍的数据量上进行训练，参数量也多出了 10 倍。在性能方面，该模型能够生产连贯的文本段落，在许多语言建模基准上取得了 SOTA 表现。而且该模型在没有任务特定训练的情况下，能够做到初步的阅读理解、机器翻译、问答和自动摘要。",
            //         "tags": [
            //         "模型"
            //         ]
            //         },
            //         {
            //         "id": "57bc17b0-c1e5-48de-be3b-c7bf80131f51",
            //         "name": "ELMo",
            //         "summary": "ELMO 是“Embedding from Language Models”的简称， ELMO 本身是个根据当前上下文对 Word Embedding 动态调整的思路。ELMO 采用了典型的两阶段过程，第一个阶段是利用语言模型进行预训练；第二个阶段是在做下游任务时，从预训练网络中提取对应单词的网络各层的 Word Embedding 作为新特征补充到下游任务中。",
            //         "tags": [
            //         "算法"
            //         ]
            //         },
            //         {
            //         "id": "99d58474-7f04-44a0-a926-d043b7906dec",
            //         "name": "上下文无关文法",
            //         "summary": "上下文无关文法，在计算机科学中，若一个形式文法 G = 的产生式规则都取如下的形式：V -&gt; w，则谓之。其中 V∈N，w∈* 。上下文无关文法取名为“上下文无关”的原因就是因为字符V 总可以被字串w 自由替换，而无需考虑字符V 出现的上下文。",
            //         "tags": [
            //         "技巧"
            //         ]
            //         },
            //     ]
            //     },
            //  "concepts":{
            //     "total_count":10,
            //     "data":[
            //       {
            //         "id": "ce06fe06-610e-4a82-8af3-a7e1c0d5d351",
            //         "name": "用户画像",
            //         "summary": "用户画像（persona）的概念最早由交互设计之父Alan Cooper提出:“Personas are a concrete representation of target users.” 是指真实用户的虚拟代表，是建立在一系列属性数据之上的目标用户模型。",
            //         "completeness": "level2",
            //         "tags": []
            //         },
            //         {
            //         "id": "7d69b898-2d93-4af8-b075-24a15b95e217",
            //         "name": "Word Error Rate (WER)",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "0c64bb6c-00f4-4d40-97f2-d6bbfabfa24a",
            //         "name": "UAS (accuracy)",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "d8607502-1adb-4f5a-ae28-db096a1d798d",
            //         "name": "Smatch",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "abcb4b41-7c7b-4c87-9d6b-e0eb2171fda9",
            //         "name": "SICK-R",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "7ca28cdb-65a0-422f-9409-4c2fc837d28e",
            //         "name": "SICK-E",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "8109119f-3c8e-49a2-9c2b-0e9d6a36148f",
            //         "name": "ROUGE (Recall-Oriented Understudy for Gisting Evaluation)",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "dcde6150-5fb3-44e1-b402-f81c035d7b02",
            //         "name": "Precision",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "f020faec-f391-4e3a-9ab4-ddeb61017a8f",
            //         "name": "Pearson Correlation",
            //         "summary": null,
            //         "completeness": "level1",
            //         "tags": []
            //         },
            //         {
            //         "id": "6d0d9bd5-d920-4adc-bf07-ada50b925307",
            //         "name": "Articulatory phonology",
            //         "summary": "",
            //         "completeness": "level1",
            //         "tags": [
            //         "其它"
            //         ]
            //         }
            //     ]
            //     }
            //  }
         },
        
        gotoApiDtl(item){
            this.$emit('click',item.APIGroup)
        },       
         handleClick(tab, event) {
        }
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
 .desc{
     font-size: 15px;
    color: #666;
    opacity: .9;
    line-height: 1.87em;
 }
.card-list{
    display: flex;
    flex-wrap: wrap;
}
 .box-card{
     width: 420px;
     margin: 0 20px 20px 0;
     position: relative;
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
 .box-card:hover{
     background: #4676aa;
     color: #fff;
 }
.box-card span{
    font-size: 14px;
    color: #999;
}   

</style>