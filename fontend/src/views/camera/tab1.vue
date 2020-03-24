<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="12">
                <div>
                    <el-divider content-position="left">温度采图</el-divider>
                    <el-card>
                        <el-form label-position="left"  size="mini">
                            <el-form-item label="">
                                <el-col :span="8">                                 
                                    <el-checkbox v-model="checked">ThreshAuto</el-checkbox>                               
                                </el-col>
                                <el-col :span="4">
                                        ir
                                </el-col>
                                <el-col :span="8">
                                    <el-input v-model="form.laserTime"></el-input>    
                                </el-col>
                                    <el-col :span="4">
                                    ℃
                                </el-col>
                            </el-form-item>
                            <el-form-item label="">
                                
                                <el-col :span="4" :offset="8">
                                    ldmp
                                </el-col>
                                <el-col :span="8">
                                    <el-input v-model="form.laserTime"></el-input>    
                                </el-col>
                                <el-col :span="4">
                                    ℃
                                </el-col>
                            </el-form-item>
                            <el-form-item label="">
                                <el-col :span="8">                                 
                                    <el-checkbox v-model="checked">ThreshAuto</el-checkbox>                               
                                </el-col>
                                <el-col :span="4">
                                        Wait
                                </el-col>
                                <el-col :span="8">
                                    <el-input v-model="form.laserTime"></el-input>    
                                </el-col>
                                    <el-col :span="4">
                                        ms
                                </el-col>
                            </el-form-item>   
                            <el-form-item label="">  
                                <el-col :span="12"> 
                                    <el-button size="mini">手动采图</el-button> 
                                </el-col> 
                            </el-form-item>                 
                        </el-form>
                    </el-card>  
                    <el-divider content-position="left">对齐采图</el-divider>
                        <el-card>
                            <el-form  >
                                <el-form-item >
                                        <el-button size="mini">对齐彩图</el-button>
                                </el-form-item>
                                <el-form-item >
                        
                                </el-form-item>
                                  <el-form-item >
                                        <el-button size="mini">生成imageParis</el-button>
                                </el-form-item>
                            </el-form>
                    </el-card>   
                </div>   
            </el-col>    
            <el-col :span="12">
                <div>
                    <el-divider content-position="left">TEC采图</el-divider>
                    <el-card>
                        <el-form :inline="true" size="mini">
                            <el-form-item >
                                 <el-col :span="6">                                 
                                    <span>TecCount</span>                                          
                                </el-col>
                                <el-col :span="6">                                 
                                    <el-input v-model="form.Tea_value"  placeholder="120"></el-input>                                  
                                </el-col>
                                <el-col :span="6">                                 
                                    <span>UseTec</span>                                           
                                </el-col>
                                <el-col :span="6">
                                    <el-input v-model="form.Tea_value"  placeholder="1"></el-input>                                  
                                </el-col>
                            </el-form-item>
                            <el-form-item >
                                 <el-col :span="6">                                 
                                    <span>TecInterval</span>                                          
                                </el-col>
                                <el-col :span="6">                                 
                                    <el-input v-model="form.Tea_value"  placeholder="1000"></el-input>                                  
                                </el-col>
                                <el-col :span="6" style="padding:0px;">                                 
                                    <span>UseTecCount</span>                                           
                                </el-col>
                                <el-col :span="6">
                                    <el-input v-model="form.Tea_value"  placeholder="130"></el-input>                                  
                                </el-col>
                            </el-form-item>
                            <el-form-item >
                                 <el-col :span="10">                                 
                                    <span>TecTheashold</span>                                          
                                </el-col>
                                <el-col :span="14">                                 
                                    <el-input v-model="form.Tea_value"  placeholder="1025"></el-input>                                  
                                </el-col>
                            </el-form-item>
                            <el-form-item >
                                <el-col :span="6">                                 
                                    <span>TecRange</span>                                          
                                </el-col>
                                 <el-col :span="6">                                 
                                    <el-input v-model="form.Tea_value"  placeholder="25"></el-input>                                  
                                </el-col>
                                <el-col :span="6" :offset="6">                                 
                                    <el-button size="mini">Set</el-button>                                  
                                </el-col>
                            </el-form-item>
                        </el-form>
                        <el-divider content-position="left">Tec value</el-divider>
                        <el-card>
                            <el-form label-position="left" label-width="100px"  size="mini">
                                <el-form-item label="转换激光温度">
                                    <el-col :span="20">                                 
                                        <el-input v-model="form.Tea_value"></el-input>                                  
                                    </el-col>
                                </el-form-item>
                                <el-form-item label="Tea调控值">
                                    <el-col :span="20">                                 
                                        <el-input v-model="form.Tea_value"></el-input>                                  
                                    </el-col>
                                    <el-col :span="4">
                                        <el-button size="mini">get</el-button>
                                    </el-col>
                                </el-form-item>                                                                
                            </el-form>
                        </el-card>  
                        <el-form  size="mini" style="margin-top:10px;">
                            <el-form-item >
                                <el-col :span="12" >
                                    <el-checkbox v-model="checked">自动采图</el-checkbox>    
                                </el-col>   
                                 <el-col :span="12">                        
                                    <el-button size="mini">手动采图</el-button>
                                </el-col>  
                            </el-form-item>                                                                
                        </el-form>
                    </el-card>                    
                </div>  
            </el-col>  
        </el-row>
    </div>                
</template>

<script>
export default {
    data(){
        return{
            form:{},
            checked:false
        }
    }
}
</script>