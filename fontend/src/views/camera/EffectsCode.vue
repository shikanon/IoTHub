<template>
    <div name="effects-code">
        <el-row>
            <el-col :span="12">
                <el-row>
                    <el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <div>
                                <el-select v-model="language" placeholder="请选择" @change="changeLanguage">
                                    <el-option
                                    v-for="item in options"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                                    </el-option>
                                </el-select>
                            </div>
                            <div>
                                <el-button icon="el-icon-refresh-left" type="text" @click="reset"></el-button>
                                <el-button icon="el-icon-setting" type="text" @click="settingVisiable = true"></el-button>
                                <el-button icon="el-icon-full-screen" type="text"></el-button>
                            </div>
                        </div>
                        <div class="box">
                            <monacoEditor v-if="editType === 'Monaco'" height="300"
                                :language="typescript"
                                :code="content" 
                                :theme="theme"	
                                :editorOptions="editOptions">
                                <!-- @mounted="onMounted" @codeChange="onCodeChange" > -->
                            </monacoEditor>
                            <codemirror v-else ref="cmRef" v-model="content" :options="cmOptions" class="codesql"></codemirror>
                        </div>
                    </el-card>
                </el-row>
                <el-row v-if="consoleItemVisiable">
                    <el-card class="box-card">
                        <el-tabs v-model="activeName" >
                            <el-tab-pane label="测试用例" name="first">
                                <div class="result-item">
                                    <el-input type="textarea" v-model="test_case_value"></el-input>
                                </div> 
                            </el-tab-pane>
                            <el-tab-pane label="代码执行结果" name="second">
                                <div class="result-item">
                                    请先运行您的代码
                                </div>    
                            </el-tab-pane>             
                        </el-tabs>
                    </el-card>
                </el-row>   
                <el-row> 
                    <el-card class="box-card " >
                        <div class="clearfix">
                            <div>
                                <span @click="consoleItemVisiable = !consoleItemVisiable" >
                                    控制台
                                <i :class="['el-icon--right',consoleItemVisiable ? 'el-icon-caret-top':'el-icon-caret-bottom' ]"></i>
                                </span>                       
                                <span>如何创建一个测试用例</span>
                            </div>
                            <div>
                                <el-button>执行代码</el-button>
                                <el-button type="success">提交</el-button>
                            </div>
                        </div>
                    </el-card>
                </el-row>     
            </el-col>
            <el-col :span="12">
                <el-row>
                    <el-card class="box-card">
                        <div slot="header" >
                            <span>输出</span>
                        </div>
                        <div class="box">
                        </div>  
                    </el-card>  
                </el-row>  
                <el-row>
                    <el-card class="box-card">
                        <div slot="header" >
                            <span>stdin
                                    <i class="el-icon-caret-bottom"></i>
                            </span>
                        </div>
                        <div class="">
                        </div>    
                    </el-card> 
                </el-row>
            </el-col>
        </el-row>  
        <el-dialog title="代码编辑器设置" :visible.sync="settingVisiable" width="35%">
            <div class="setting-box">
                <div class="setting-item">
                    <el-col :span="16">
                        <span>代码编辑器</span>
                        <p>你可以随意在 Monaco 和 Codemirror 编辑器之间切换</p>
                    </el-col>
                    <el-col :span="8">
                        <el-select v-model="editType">
                            <el-option label="Monaco" value="Monaco"></el-option>
                            <el-option label="Codemirror" value="Codemirror"></el-option>
                        </el-select>  
                    </el-col> 
                </div>
                 <div class="setting-item">
                    <el-col :span="16">
                        <span>字体设置</span>
                        <p>调整适合你的字体大小。</p>
                    </el-col>
                    <el-col :span="8">
                        <el-select v-model="fontSize">
                            <el-option label="12px" value="1"></el-option>
                            <el-option label="14px" value="2"></el-option>
                            <el-option label="16px" value="3"></el-option>
                            <el-option label="18px" value="4"></el-option>
                            <el-option label="2opx" value="5"></el-option>
                        </el-select> 
                    </el-col>   
                </div>
                 <div class="setting-item">
                    <el-col :span="16">
                        <span>主题设置</span>
                        <p>切换不同的代码编辑器主题，选择适合你的语法高亮。</p>
                    </el-col>
                    <el-col :span="8">
                        <el-select v-model="theme" v-if="editType === 'Monaco'">
                            <el-option label="Visual Studio" value="vs"></el-option>
                            <el-option label="Visual Studio Dark" value="vs-dark"></el-option>
                            <el-option label="GitHub" value="3"></el-option>
                            <el-option label="xCode" value="4"></el-option>
                            <el-option label="Material" value="hc-black"></el-option>
                        </el-select> 
                        <el-select  v-model="theme" v-else>
                            <el-option label="Textmate" value="1"></el-option>
                            <el-option label="GitHub" value="3"></el-option>
                            <el-option label="xCode" value="4"></el-option>
                            <el-option label="Monokai" value="4"></el-option>
                        </el-select>    
                    </el-col>   
                </div>
                 <div class="setting-item">
                    <el-col :span="16">
                        <span>键位绑定</span>
                        <p>想要练习使用 Vim 或者 Emacs？你可以在这里切换这些预设的键位绑定。</p>
                    </el-col>
                    <el-col :span="8">
                        <el-select v-model="fontSize">
                            <el-option label="Standard" value="1"></el-option>
                            <el-option label="Vim" value="2"></el-option>
                            <el-option label="Emacs" value="3"></el-option>                     
                        </el-select> 
                    </el-col>   
                </div>
                 <div class="setting-item">
                    <el-col :span="16">
                        <span>Tab 长度</span>
                        <p>选择你想要的 Tab 长度，默认设置为4个空格。</p>
                    </el-col>
                    <el-col :span="8">
                        <el-select v-model="tabSize">
                            <el-option label="2个空格" value="1"></el-option>
                            <el-option label="4个空格" value="2"></el-option>
                            <el-option label="8个空格" value="3"></el-option>
                        </el-select> 
                    </el-col>   
                </div>
            </div>   
            <!-- <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="settingSubmit">确 定</el-button>
                <el-button @click="settingVisiable = false">取 消</el-button>
            </span> -->
        </el-dialog>      
    </div> 
</template>

<script>
  import monacoEditor from 'vue-monaco-editor'
  import { codemirror } from 'vue-codemirror'  
  export default {
      components: {  monacoEditor,codemirror  },
    data() {
      return {
        options: [{
          value: '选项1',
          label: 'C++'
        }, {
          value: '选项2',
          label: 'JAVA'
        }, {
          value: '选项3',
          label: 'Python'
        }, {
          value: '选项4',
          label: 'C'
        }, {
          value: '选项5',
          label: 'C#'
        },{
          value: '选项6',
          label: 'JavaScript'
        },{
          value: '选项7',
          label: 'Go'
        }],
        language: 'Java',
        activeName:'first',
        test_case_value:"123123",
        consoleItemVisiable:true,
        demo:"",
        settingVisiable:false,
        content:"select * from user",
        cmOptions: {
            mode: 'Java',
            indentWithTabs: true,
            smartIndent: true,
            lineNumbers: true,
            matchBrackets: true,
            styleActiveLine: true,
            cursorHeight:1, // 光标高度
            autoRefresh: true,
            abSize:4
        },
        editOptions:{        
            //quickSuggestionsDelay: 500,   //代码提示延时

            selectOnLineNumbers: true,
            roundedSelection: false,
            readOnly: false,// 只读
            cursorStyle: 'line',//光标样式
            automaticLayout: false, //自动布局
            glyphMargin: true,//字形边缘
            useTabStops: false,
            fontSize: 28,       //字体大小
            autoIndent:true,//自动布局
        },
        editType:'Monaco',
        fontSize:'2',
        tabSize:'2',
        theme:'1'
      }
  
    },
    created(){
        this.initEditor()
    },

    methods:{
        changeLanguage(val){
            this.cmOptions.mode = val
        },

         initEditor(){
           
        },

        RunResult(){
            console.log(this.monacoEditor.getValue());
        },

        themeChange(val){
            this.initEditor();
        },
    
         //重置
        reset(){
             this.$confirm('您将放弃所有更改并初始化代码, 是否继续?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
            }); 
        },
      

 
    }
  }
</script>

<style scoped>
.clearfix{
    display: flex;
    justify-content: space-between;
}

.result-item{
    display: flex;
    align-items: center;
    justify-content: center;
    color: #e0e0e0;
    height: 120px;
    font-size: 18px;
    font-weight: 500;
  
}

.box{
    min-height: 350px;
}

.setting-box{
    display: flex;
    flex-direction: column;
}

.setting-item{
    margin-bottom: 15px;
}

.setting-item p{
    color: #9e9e9e;
    margin:5px 9px 0px;
}
</style>