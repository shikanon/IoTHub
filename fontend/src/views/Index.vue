<template>
    <div name="index">
        <div>{{helloMsg}}</div>
        <div class="card-box">
            <el-card  shadow="hover">     
                <div class="margin-b-10">当月设备消息量</div>
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
 .card-box{
     display: flex;
     margin-top: 20px;
 }

 .card-box .el-card{
    width: 229px;
    box-sizing: border-box;
    min-width: 229px;
    margin-right: 8px;
    border: 1px solid #e8ebee;
    border-radius: 3px;
 }
 .margin-b-10{
     margin-bottom: 10px;
 }

</style>