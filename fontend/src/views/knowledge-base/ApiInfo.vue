<template>
    <div>
        <el-container>
            <el-header>API接口</el-header>
            <el-container>
                <el-aside width="auto">
                      <el-tabs tab-position="left" style="height: 100%;"  @tab-click="clickApiName">
                        <el-tab-pane  v-for="(item,key) in apiList" :key="key" :label="item.name" :apiId="item._id"></el-tab-pane>          
                    </el-tabs>                  
                </el-aside>
                <el-main v-if="apiInfo">
                    <div>
                        <p>{{apiInfo.name}}</p>
                        <p>调用地址：  <el-link type="primary">http(s)://goodsdect.market.alicloudapi.com/goodsdect</el-link></p>
                         <p>请求方式：{{apiInfo.HttpMethod}}</p>
                        <p>返回类型：{{apiInfo.ResultType}}</p> 
                        <p>API 调用： <el-link type="primary">API 简单身份认证调用方法（APPCODE）</el-link>
                            <span @click="flag = true" v-if="!flag"> 展开<i class="el-icon-arrow-down"></i></span>
                            <span @click="flag = false" v-if="flag"> 收起<i class="el-icon-arrow-up"></i></span>
                        </p>
                        <p  v-if="flag"><el-link type="primary" style=" margin-left: 84px;">API 签名认证调用方法（AppKey & AppSecret）</el-link></p>
                        <p>调试工具：<el-button @click="gotoDebugApi">去调试</el-button></p>
                        <el-collapse v-model="activeNames" >
                            <el-collapse-item title="请求参数（Headers）" name="1">
                                <RequestTable :tableData="apiInfo.RequestHeaders.RequestParam"></RequestTable>       
                            </el-collapse-item>
                            <el-collapse-item title="请求参数（Query）" name="2">
                                  <RequestTable :tableData="apiInfo.RequestQueries.RequestParam"></RequestTable>                                                          
                            </el-collapse-item>
                            <el-collapse-item title="请求参数（Body）" name="3">
                                <RequestTable :tableData="apiInfo.RequestBody.RequestParam"></RequestTable>                                                          
                            </el-collapse-item>
                            <el-collapse-item title="请求实例" name="4">
                                 <el-tabs v-model="activeName">
                                    <el-tab-pane v-for="(item,key) in apiInfo.SdkDemos.SdkDemo"  :key="key" :label="item.IdeKey" :name="item.IdeKey">
                                        <div  v-highlight>
                                            <pre><code v-html="item.CallDemo"></code></pre>
                                        </div>
                                    </el-tab-pane>
                                </el-tabs>
                            </el-collapse-item>
                             <el-collapse-item title="正常返回示例" name="5">
                                   <JsonView :data="apiInfo.ResultSample ? JSON.parse(apiInfo.ResultSample):{} "></JsonView>
                            </el-collapse-item>
                             <el-collapse-item title="失败返回示例" name="6">
                                 <JsonView :data="apiInfo.FailResultSample ? JSON.parse(apiInfo.FailResultSample):{}"></JsonView>                              
                            </el-collapse-item>
                            <el-collapse-item title="错误码定义" name="7">
                                 <el-table
                                :data="apiInfo.ErrorCodeSamples.ErrorCodeSample"
                                style="width: 100%">                             		
                                    <el-table-column                                   
                                        label="错误码"
                                        prop="Code"
                                        width="100"
                                        >
                                    </el-table-column>
                                    <el-table-column                                    
                                        label="错误信息"
                                        prop="Message"
                                        width="280">
                                    </el-table-column>
        
                                    <el-table-column                                    
                                        label="描述"
                                        prop="Description"
                                        width="280">
                                    </el-table-column>
                                </el-table>                        
                            </el-collapse-item>                            
                        </el-collapse>
                    </div>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
<script>
    import RequestTable from './RequestTable';

export default {
    components:{RequestTable},
    data(){
        return {
            flag:false,
            activeNames:["0"],
            groupId:'',
            APIGroup:'',
            apiList:[],

           apiInfo: {SdkDemos: {SdkDemo:[]},
            Description: "",
            RequestQueries: {RequestParam: []},
            PathParameters: {PathParameter: []},
            ErrorCodeSamples: {ErrorCodeSample: []},
            HttpProtocol: "",
            ResultDescriptions: {ResultDescription: []},
            RequestHeaders: {RequestParam: []},
            RequestBody: {RequestParam:[]},
            Path: "",
            ResultSample: "",
            StageName: "",
            ResultType: "",
            FailResultSample: "",
            ServiceTimeout: 0,
            BodyFormat: "",
            HttpMethod: "POT"
            },
            activeName:'',
    
        }
    },
    created(){
       
       // this.groupId = this.$route.params.groupId
        this.APIGroup = this.$route.params.APIGroup
        this.getAPIGroupList()
       // this.getApiList(this.groupId)
    },
    methods:{

        clickApiName(tab,event){

            this.getApiInfById(tab.$attrs.apiId)
        },
        // getApiList(groupId){
        //     console.log(groupId)
        //     this.apiList = [
        //         {"GroupName":"ali_gateway","GroupDescription":"云市场api网关","DeployedTime":"2019-12-05T09:21:00Z","Description":"可上传多种格式的文档进行翻译","StageName":"RELEASE","ApiName":"文档上传翻译","RegionId":"cn-hangzhou","ApiId":"5666d53dfb9845568c25d1ca1a2e9fcf","GroupId":"805c3653c8eb4e9cbbf54727e5ca188e"},
        //         {"GroupName":"ali_gateway","GroupDescription":"云市场api网关","DeployedTime":"2019-12-05T09:20:38Z","Description":"根据上传接口返回的tid查询翻译进度，返回status=314时即翻译完成，可以进行下载译文，该接口同一tid调用间隔在3s及以上","StageName":"RELEASE","ApiName":"查询翻译进度","RegionId":"cn-hangzhou","ApiId":"5b2e740a504f4d6094959f54afa936ed","GroupId":"805c3653c8eb4e9cbbf54727e5ca188e"},
        //         {"GroupName":"ali_gateway","GroupDescription":"云市场api网关","DeployedTime":"2019-12-05T09:20:29Z","Description":"下载可选格式的译文，仅翻译完成的tid可以下载","StageName":"RELEASE","ApiName":"下载译文","RegionId":"cn-hangzhou","ApiId":"c241742953024607959799b4dd9cc4c6","GroupId":"805c3653c8eb4e9cbbf54727e5ca188e"}
        //     ]
        //     this.getApiInfById(this.apiList[0].ApiId)
                
        //   },
        getAPIGroupList(){   
             this.$API_SOTA.getApiGroup(this.APIGroup).then((res) => {
                 this.apiList = res.data 
                 if(this.apiList.length > 0){
                     this.getApiInfById(this.apiList[0]._id)
                 }
                                
            })
            // this.$axios.get('http://localhost:8080/static/apiInfo.json').then((res) => {
            //      this.apiInfo  = res.data.result    
            //      this.apiInfo.ApiName = this.apiList.filter(item => item.ApiId === apiId)[0].ApiName
            //      this.activeName   = this.apiInfo.SdkDemos.SdkDemo[0].IdeKey
            //  }) 
                
         },

         getApiInfById(apiId){
            this.apiInfo = this.apiList.filter(item => item._id === apiId)[0].SDK
            this.apiInfo.name = this.apiList.filter(item => item._id === apiId)[0].name
            this.activeName   = this.apiInfo.SdkDemos.SdkDemo[0].IdeKey
         },

         gotoDebugApi(){
            this.$router.push({name :'debug-api',params: {groupName:groupName}})      
         }
    }
    
}
</script>


<style  scoped>

  .el-header, .el-footer {
    color: #333;
    text-align: center;
    line-height: 60px;
    border-bottom: 1px solid #ccc;
  }
  
  .el-aside {
    color: #333;
    text-align: center;
    line-height: 200px;
  }
  
  .el-main {
    color: #333;
    text-align: center;
    line-height: 160px;
    overflow: hidden;
  }
  
  p{
      height: 30px;
      line-height: 30px;
      text-align: left;
  }

  p:first-child{
      margin-top: 0px;
  }
  .el-collapse{
      width: 50%;
  }


  .el-collapse-item{
      text-align: left;
  }
  .el-tab-pane,.el-collapse-item__content{
      max-height: 300px;
      overflow: auto
  }
  

</style>