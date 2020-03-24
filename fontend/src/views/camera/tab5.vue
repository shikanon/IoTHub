<template>
    <div>
        <el-row >
            <el-divider content-position="left">RefPicture</el-divider>
            <el-card>
                <el-row>
                    <el-col :span="18">
                        <el-row>
                            <el-col :span="4"><span>配置文件</span></el-col>
                            <el-col :span="14"><el-input size="mini"></el-input></el-col>
                            <el-col :span="4" :offset="2"><el-button size="mini">SelectConfig</el-button></el-col>
                        </el-row>
                         <el-row>
                             <el-divider content-position="left">Ref maker</el-divider>
                             <el-card>
                                 <el-form :inline="true" size="mini"  class="inner-from">
                                    <el-row>
                                        <el-form-item label="ws" >
                                            <el-input class="input_78"></el-input>
                                        </el-form-item>
                                        <el-form-item label="le">
                                            <el-input class="input_78"></el-input>
                                        </el-form-item>
                                        <el-form-item label="re">
                                            <el-input class="input_78"></el-input>
                                        </el-form-item>
                                        <el-form-item label="te">
                                            <el-input class="input_78"></el-input>
                                        </el-form-item>
                                        <el-form-item label="be">
                                            <el-input class="input_78"></el-input>
                                        </el-form-item>
                                    </el-row>
                                    <el-row>
                                        <el-form-item label="ro" >
                                            <el-input class="input_78"></el-input>
                                        </el-form-item>
                                        <el-form-item label="去畸变模型">
                                              <el-select v-model="form.type" placeholder="活动区域">
                                                <el-option label="Brown_K2T2" value="1"></el-option>
                                            </el-select>
                                        </el-form-item>
                                        
                                    </el-row>
                                </el-form> 
                             </el-card>    
                        </el-row>
                    </el-col>
                    <el-col :span="4" :offset="2">
                          <el-row class="margin_top_20">
                            <el-checkbox v-model="checked">硬件去畸变</el-checkbox>
                         </el-row>
                         <el-row class="margin_top_20">
                            <el-button>转换参考图</el-button>
                         </el-row>
                         <el-row class="margin_top_20">
                            <el-button>DownLoadBtn</el-button>
                         </el-row>
                    </el-col>
                </el-row>
                 <el-row class="margin_top_20">
                    <el-col :span="3"><span>Bin文件</span></el-col>
                    <el-col :span="10"><el-input size="mini"></el-input></el-col>
                    <el-col :span="4" :offset="2">
                        <el-button size="mini">Select</el-button>
                    </el-col>
                </el-row>
                <el-row class="margin_top_20">
                    <el-col :span="3"><span>参考图</span></el-col>
                    <el-col :span="10"><el-input size="mini"></el-input></el-col>
                    <el-col :span="4" :offset="2"><el-button size="mini">SelectConfig</el-button></el-col>
                </el-row>
            </el-card>    
        </el-row>
        <el-row >
            <el-divider content-position="left">RefBin</el-divider>
            <el-card>
                <el-row>
                    <el-col :span="4"><span>参考图Bin文件路径</span></el-col>
                    <el-col :span="12"><el-input size="mini"></el-input></el-col>
                    <el-col :span="6" :offset="2">
                        <el-button size="mini">Select</el-button>
                        <el-button size="mini">DownLoadRefBin</el-button>
                    </el-col>
                </el-row>
            </el-card>
        </el-row>
    </div>                
</template>

<script>
export default {
    data(){
        return{
            checked:true,
            form:{}
        }
    }
}
</script>


<style scoped>
  .input_78{
        width: 78px;
    }
    .margin_top_20{
        margin-top: 20px;
    }
</style>