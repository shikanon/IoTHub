<template>
    <div name="step-start">
        <p class="step-desc">此引导将通过4个步骤，来帮助您连接设备到阿里云物联网平台：</p>
            <div class="step-list">
                <div>
                    <div>
                        <img src="https://img.alicdn.com/tfs/TB1xcSFmCzqK1RjSZFLXXcn2XXa-162-144.png">
                    </div>
                    <div>
                        <p>注册设备</p>
                        <p>设备是平台中物理设备的表示形式和记录。任何物理设备都需要注册才能使用阿里云物联网平台。</p>
                    </div>
                </div>
                <div>
                    <div>
                        <img src="https://img.alicdn.com/tfs/TB1XkCGmCzqK1RjSZFpXXakSXXa-159-141.png">
                    </div>
                    <div>
                        <p>选择设备平台</p>
                        <p>选择电脑或者设备运行的平台，自动生成相应的设备开发工具包。</p>
                    </div>
                </div>
                <div>
                    <div>
                        <img src="https://img.alicdn.com/tfs/TB1Vp5GmpzqK1RjSZFCXXbbxVXa-162-141.png">
                    </div>
                    <div>
                        <p>下载开发工具包</p>
                        <p>开发工具包中包括一些重要组件: 设备三元组、您选择的软件开发工具包和示例脚本。</p>
                    </div>
                </div>
                <div >
                    <div>
                        <img src="https://img.alicdn.com/tfs/TB10Be4mxnaK1RjSZFtXXbC2VXa-159-135.png">
                    </div>
                    <div>
                        <p>配置和测试设备</p>
                        <p>通过使用连接工具包，您可以测试设备是否已正确连接到阿里云物联网平台，并且还可以上报数据以及接收数据。</p>
                    </div>
                </div>
            </div> 
    </div>        
</template>

<style  scoped>

.step-list>div{
    display: flex;
    align-items: center;
    margin-top: 30px;
}
.step-list img{
    width: 58px;
    height: 54px;
    margin-right: 30px;
}
.step-list p{
    margin: 0px;
}
.step-list p:first-child{
    color: #333;
    font-size: 16px;
}
.step-list p:last-child{
    font-size: 12px;
    margin-top: 4px;
    color: #999
}
</style>