<template>
    <div name="index">
        <el-card class="margin-b-10">
            <div class="hello-msg">{{helloMsg}}</div>
            <div class="card-box margin-b-10">
                <el-card  shadow="hover">     
                    <div class="">当月设备消息量</div>
                    <div>2</div>                  
                </el-card>
                <el-card shadow="hover">     
                    <div class="margin-b-10">当月设备在线时长 (分钟)</div>
                    <div>192</div>                  
                </el-card>
                <el-card shadow="hover">     
                    <div class="margin-b-10">应用项目数</div>
                    <div>1</div>                  
                </el-card>
                <el-card shadow="hover">     
                    <div class="margin-b-10">边缘实例数</div>
                    <div>1</div>                  
                </el-card>
            </div>
        </el-card>
        
        <div class="flex-box">
            <div class="left-box">
                <div class="margin-b-10">
                    <el-card>
                        <p>快速入门</p>
                        <div class="card-box">
                            <el-card shadow="hover" class="card-50">   
                                <div class="flex-box space-around"> 
                                    <div class="img-box">
                                        <img src="https://img.alicdn.com/tfs/TB1C0TpbKH2gK0jSZJnXXaT1FXa-72-72.svg"/> 
                                    </div>
                                    <div class="margin-b-10">
                                        <p>快速连接设备</p>
                                        <div class="desc-div"> 基于该流程，可快速将你的计算机，或者设备连接到阿里云物联网平台。</div>
                                        <el-button size="medium">开始体验</el-button>                                   
                                    </div> 
                                </div>                                             
                            </el-card>
                            <el-card shadow="hover" class="card-50">     
                                <div class="flex-box space-around"> 
                                    <div class="img-box">                               
                                         <img   src="https://img.alicdn.com/tfs/TB1rjvlauL2gK0jSZFmXXc7iXXa-72-72.svg"/> 
                                    </div>
                                    <div class="margin-b-10">
                                        <p>物联网开发者新手文档</p>
                                         <div class="desc-div">
                                           详解阿里云物联网平台的产品架构、和基于用户角色的上手流程指南。
                                        </div>
                                        <el-button size="medium">前往查看</el-button>  
                                    </div>
                                </div>                
                            </el-card>
                        </div>
                    </el-card>
                </div> 
                <div class="margin-b-10">
                    <el-card>
                        <div class="flex-box space-between">
                            <p>使用教学</p>
                            <el-button type="text">更多</el-button>
                        </div>
                        <div class="card-box">
                               <div class="card-box">
                            <a>                               
                                <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1sUzLbxD1gK0jSZFyXXciOVXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">物联网开发全栈课程</div>
                                <div class="list-description margin-b-10">从实际项目维度，深入讲解物联网端到端全栈开发过程。</div>
                                <div class="tag-box">    
                                    <span >Web应用</span>                
                                    <span>STM32</span>    
                                    <span>嵌入式</span>         
                                 </div>                
                            </a>
                            <a>                               
                                 <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1sUzLbxD1gK0jSZFyXXciOVXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">物联网开发全栈课程</div>
                                <div class="list-description margin-b-10">从实际项目维度，深入讲解物联网端到端全栈开发过程。</div>
                                <div class="tag-box">    
                                    <span>小程序</span>               
                                    <span>MQTT协议</span>
                                 </div>                
                            </a>   
                            <a>                                
                                <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1sUzLbxD1gK0jSZFyXXciOVXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">物联网开发全栈课程</div>
                                <div class="list-description margin-b-10">从实际项目维度，深入讲解物联网端到端全栈开发过程。</div>
                                <div class="tag-box">    
                                    <span>嵌入式</span>    
                                 </div>                
                            </a>   
                            <a>                                
                                <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1sUzLbxD1gK0jSZFyXXciOVXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">物联网开发全栈课程</div>
                                <div class="list-description margin-b-10">从实际项目维度，深入讲解物联网端到端全栈开发过程。</div>
                                <div class="tag-box">    
                                    <span>LoRaWAN</span>                             
                                 </div>                
                            </a>
                        </div>
                        </div>
                    </el-card>
                </div>
                  <div>
                    <el-card>
                        <div class="flex-box space-between">
                            <p>解决方案</p>
                            <el-button type="text">更多</el-button>
                        </div>
                        <div class="card-box">
                            <a>                               
                                <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1quHIbp67gK0jSZPfXXahhFXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">RAK LoraWAN 入门套件</div>
                                <div class="list-description margin-b-10">快速搭建基于Lora的物联网端到端 PoC 项目。</div>
                                <div class="tag-box">    
                                    <span >IoT Studio</span>                
                                    <span>LoRaWAN</span>       
                                 </div>                
                            </a>
                            <a>                               
                                 <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1quHIbp67gK0jSZPfXXahhFXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">RAK LoraWAN 入门套件</div>
                                <div class="list-description margin-b-10">快速搭建基于Lora的物联网端到端 PoC 项目。</div>
                                <div class="tag-box">    
                                    <span>IoT Studio</span>               
                                    <span>LoRaWAN</span>
                                 </div>                
                            </a>   
                            <a>                                
                                <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1quHIbp67gK0jSZPfXXahhFXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">RAK LoraWAN 入门套件</div>
                                <div class="list-description margin-b-10">快速搭建基于Lora的物联网端到端 PoC 项目。</div>
                                <div class="tag-box">    
                                    <span>IoT Studio</span>    
                                    <span>LoRaWAN</span>       
                                 </div>                
                            </a>   
                            <a>                                
                                <div class="margin-b-10">
                                    <img  class="image" src="https://img.alicdn.com/tfs/TB1quHIbp67gK0jSZPfXXahhFXa-600-336.jpg"/>
                                </div>
                                <div class="list-title">RAK LoraWAN 入门套件</div>
                                <div class="list-description margin-b-10">快速搭建基于Lora的物联网端到端 PoC 项目。</div>
                                <div class="tag-box">    
                                    <span>IoT Studio</span>                             
                                    <span>LoRaWAN</span>     
                                 </div>                
                            </a>
                        </div>
                    </el-card>
                </div>
            </div>
             <div class="right-box">
                 <el-card>
                    <img src="https://img.alicdn.com/tfs/TB1d0TanHY1gK0jSZTEXXXDQVXa-1440-616.jpg"/>
                    <p>快速入口</p>
                    <el-row class="margin-b-10">
                        <el-button >购买实例</el-button>
                        <el-button >购买资源包</el-button>
                    </el-row>
                    <el-row class="margin-b-10">
                        <el-button >SDK 下载</el-button>
                        <el-button >提交工单</el-button>
                    </el-row>
                     <el-divider></el-divider>
                     <div>
                         <p>平台消息</p>
                         <div class="msg-list">
                            <div class="msg-box margin-b-10">
                                <span>包年包月实例上线</span><br>
                                <span>2019年12月31日</span>
                            </div>
                            <div class="msg-box margin-b-10">
                                <span>12月平台功能发布</span><br>
                                <span>2019年12月09日</span>
                            </div>
                            <div class="msg-box margin-b-10">
                                <span>固件升级计费公告</span><br>
                                <span>2019年11月27日</span>
                            </div>
                            <div class="msg-box margin-b-10">
                                <span>11月平台功能发布</span><br>
                                <span>2019年11月21日</span>
                            </div>
                            <div class="msg-box margin-b-10">
                                <span> 10月平台功能发布-2</span><br>
                                <span>2019年10月31日</span>
                            </div>
                            <div class="msg-box margin-b-10">
                                <span>10月平台功能发布-1</span><br>
                                <span>2019年10月16日</span>
                            </div>
                         </div>
                     </div>
                 </el-card>    
             </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {      
         helloMsg:''
        }
      },
      created(){
          var now = new Date();
          var hours = now.getHours();
          this.helloMsg = hours >= 12 ? "下午好!" : "上午好!" 

      }
}
</script>


<style scoped>
.hello-msg{
    font-size: 20px;
}
 .card-box{
     display: flex;
     margin-top: 20px;
 }

.card-50{
    width: 50%;
}
 .card-box .el-card{
    box-sizing: border-box;
    min-width: 229px;
    margin-right: 8px;
    border: 1px solid #e8ebee;
    border-radius: 3px;
 }
 .margin-b-10{
     margin-bottom: 10px;
 }

.flex-box{
    display: flex;
}

.space-around{
    justify-content: space-around;

}
.space-between{
      justify-content: space-between;
}
.img-box{
    display: flex;
    align-items: center;
}
.img-box img{
    width: 60px;
    height: 60px
}
.desc-div{
    width: 270px;
    font-size: 14px;
    color: #999;
    margin-bottom: 15px;
}


.left-box{
 width: 80%;
}

.right-box{
    width: 20%;
    margin-left: 20px;
}
.right-box img{
    width: 100%;
}

 .list-description{
     font-size: 14px;
     color:#999
 }

 .msg-box span{
    font-size: 14px
 }
  .msg-box span:last-child{
      color:#999;
  }

.no-padding .el-card__body{
      padding: 0px;
  }

.image{
    width: 100%;
}
.tag-box span{
    color:#999999;
    background-color:#f9f9f9;
    font-size: 12px;
    padding: 5px;
}


</style>