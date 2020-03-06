<template>
    <div class="video-box" >
        <div >
            <el-button @click="test">截图</el-button>
            <div class="opt-btn" >
                <el-button @click="changeType('rtmp')">rtmp</el-button>
                <el-button @click="changeType('m3u8')">m3u8</el-button>
                <el-button @click="changeType('flv')">flv</el-button>
            </div>
        </div>
        <div class="video-item" >


        <video  controls="controls"  width="500" height="300"  ref="videoPlayer" webkit-playsinline>
             <source src="https://logos-channel.scaleengine.net/logos-channel/live/biblescreen-ad-free/playlist.m3u8">
        </video>
            <!-- <video src="../assets/video.mp4" controls autoplay ></video>   -->
<!-- 
             <video-player
                class="video-player vjs-custom-skin"
                ref="videoPlayer"
                :playsinline="true"
                :options="playerOptions"         
             >
            </video-player>  -->
        </div>
        <div class="cut-screen" >
             <canvas  id="canvas"></canvas> 
        </div>
    </div>
</template>

<script>

import videojs from 'video.js'
import 'video.js/dist/video-js.css'
import 'vue-video-player/src/custom-theme.css'
import {videoPlayer} from 'vue-video-player'
import 'videojs-flash'  //引入才可以播放rtmp
import 'videojs-contrib-hls/dist/videojs-contrib-hls'//引入才可以播放m3u8

 import SWF_URL from 'videojs-swf/dist/video-js.swf'
 
 videojs.options.flash.swf = SWF_URL // 设置flash路径，Video.js会在不支持html5的浏览中使用flash播放视频文件

import html2canvas from 'html2canvas'
import Canvas2Image from '../utils/canvas2image.js'


export default {
    components: {
        videoPlayer
    },

    data() {
        return {
            playerOptions: {
                live: true,
                autoplay: true, // 如果true，浏览器准备好时开始播放
                muted: false, // 默认情况下将会消除任何音频
                loop: false, // 是否视频一结束就重新开始
                preload: 'auto', // 建议浏览器在<video>加载元素后是否应该开始下载视频数据。auto浏览器选择最佳行为,立即开始加载视频（如果浏览器支持）
                aspectRatio: '16:9', // 将播放器置于流畅模式，并在计算播放器的动态大小时使用该值。值应该代表一个比例 - 用冒号分隔的两个数字（例如"16:9"或"4:3"）
                fluid: true, // 当true时，Video.js player将拥有流体大小。换句话说，它将按比例缩放以适应其容器。
                controlBar: {
                  timeDivider: false,//是否显示已播放时间
                  durationDisplay: false,//是否显示总时长
                  remainingTimeDisplay: false,//是否显示剩余时间
                  currentTimeDisplay: false, // 当前时间
                  volumeControl: true, // 声音控制键
                  playToggle: true, // 暂停和播放键
                  progressControl: false, // 进度条
                  fullscreenToggle: true // 全屏按钮
                 },
                controls: true,
                poster: "https://github.surmon.me/vue-quill-editor/static/images/surmon-9.jpg",
                techOrder: ['flash', 'html5'], // 兼容顺序    
                flash: { hls: {withCredentials: false },swf: SWF_URL },   
                html5: { hls: { withCredentials: false } },
                sources: [
                    {
                        src: "http://demo.vcxiaohan.com/source/video.mp4", // 路径
                        type: "video/mp4" // 类型
                    },
                //     {
                //         type: "rtmp/mp4",
                //         src:"rtmp://10.10.6.85:1935/live/livestream",// 视频地址-改变它的值播放的视频会改变
                //     }, 

                    //  {
                    //     withCredentials: false,
                    //     type: "application/x-mpegURL",
                    //     src: "http://10.10.6.85:7002/live/livestream.m3u8",
                    //     src: "https://logos-channel.scaleengine.net/logos-channel/live/biblescreen-ad-free/playlist.m3u8"  
                    //  },
                    // {
                    //  type: 'video/x-flv',
                    //  src: 'http://10.10.6.85:7001/live/livestream.flv'
                    // }
                                
                 ],           
                notSupportedMessage: '此视频暂无法播放，请稍后再试', // 允许覆盖Video.js无法播放媒体源时显示的默认信息。
                showLoading:true 
             }

        }
    },
  

    created(){
      
    },

    computed: {
        player() {
            return this.$refs.videoPlayer.player
        }
    },

    // watch:{
    //      $route:function(){
    //          this.$destroy('video-player')        
    //      }     
    // },

    // mounted:{
    //     监测浏览器是否安装了flash播放器
    //     this.isflashFn()
    // },

    methods:{

        changeType(type){
            if(type === 'rtmp'){
                 this.playerOptions.sources = [{
                    type: "rtmp/mp4",
                    src:"rtmp://10.10.6.85:1935/live/livestream",// 视频地址-改变它的值播放的视频会改变
                }]
            }else if(type ==='m3u8'){
                this.playerOptions.sources = [{
                          
                    withCredentials: false,
                    type: "application/x-mpegURL",
                 //   src: "http://10.10.6.85:7002/live/livestream.m3u8",
                    src: "https://logos-channel.scaleengine.net/logos-channel/live/biblescreen-ad-free/playlist.m3u8"  
                     
                }]
            }else if(type === 'flv'){
                this.playerOptions.sources = [{
                     type: 'video/x-flv',
                    src: 'http://10.10.6.85:7001/live/livestream.flv'
                }]
            }
               
        },

        fullScreen() {
            const player = this.$refs.videoPlayer.player
            player.requestFullscreen()//调用全屏api方法
            player.isFullscreen(true)
            player.play()
        },

        onPlayerPlay(player) {
            player.play()
        },

        onPlayerPause(player) {
             alert("pause");
        },
       
    
        test(){
 
            let video =  document.querySelector('video')//获取前台要截图的video对象，
          //  video.crossOrigin = 'anonymous'

            let canvas = document.querySelectorAll('canvas')[0]//获取前台的canvas对象，用于作图
           // let canvas = document.createElement('canvas') //创建
            let ctx = canvas.getContext('2d')//设置canvas绘制2d图，
            let width = parseInt(video.offsetWidth)
            let height = parseInt(video.offsetHeight)
            canvas.width = width;
            canvas.height = height;
            ctx.drawImage(video, 0, 0, width, height);//将video视频绘制到canvas中
          //  let dataURL = canvas.toDataURL('image/png');//canvas的api中的toDataURL（）保存图像
           //canvas绘制图片，由于浏览器的安全考虑，如果在使用canvas绘图的过程中，使用到了外域的图片资源，那么在toDataURL()时会抛出安全异常
          //  console.log(dataURL)
    

            //let image = Canvas2Image.convertToPNG(canvas, canvas.width, canvas.height)
            let image = Canvas2Image.saveAsPNG(canvas, canvas.width, canvas.height)
            console.log(image)
           //  let dataURL = img.getAttribute('src')
          
            // let aLink = document.createElement("a")
            // let timestamp = Date.parse(new Date());
            // aLink.setAttribute("download",`截图${timestamp}.png`)
            // aLink.style.display = "none"
            // aLink.href = dataURL
            // // 触发点击-然后移除
            // document.body.appendChild(aLink)
            // aLink.click()
            // document.body.removeChild(aLink)
                       
        },

       //截取图片
        html2canvas () {
            let that = this
            html2canvas(this.$refs.canvas).then(canvas => {
                let data = canvas.toDataURL("image/jpeg", 1)
                that.url = data 
                console.log(data)          
            });
        },

        getBlob (canvas) { //获取blob对象
        var data = canvas.toDataURL("image/jpeg", 1);
        data = data.split(',')[1];
        data = window.atob(data);
        var ia = new Uint8Array(data.length);
        for (var i = 0; i < data.length; i++) {
            ia[i] = data.charCodeAt(i);
        }
        return new Blob([ia], {
            type: "image/jpeg"
        });
        }

    

    },
    beforeDestroy(){
        //this.$refs.videoPlayer.dispose()
    },
    destroyed(){
       // this.$refs.videoPlayer.dispose()
    }
}
</script>

<style scoped >
  .video-box{
      display: flex;
      flex-direction: row;
      width: 70%;
     justify-content: flex-start
  }

  .video-item{
      width: 500px;
  }

  .opt-btn{
      margin-top: 50px;
      
  }

/* .cut-screen{
      width: 500px;
} */

</style>
