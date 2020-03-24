<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="12">
                <div>
                    <el-divider content-position="left">标定图像</el-divider>
                    <el-card>
                        <el-form >
                            <el-form-item>
                                <el-button>选择标定对齐图像列表文件</el-button>
                            </el-form-item>
                            <el-form-item>
                                <el-checkbox-group v-model="checkList1">
                                    <el-checkbox label="Mirror_IR"></el-checkbox>
                                    <el-checkbox label="Mirror_IR2"></el-checkbox>
                                    <el-checkbox label="Mirror_color"></el-checkbox>                       
                                </el-checkbox-group>
                            </el-form-item>
                             <el-form-item>
                                 <el-input type="textarea" v-model="form.desc"></el-input>
                             </el-form-item>
                        </el-form>
                    </el-card>  
                    <el-divider content-position="left">结果保存样式</el-divider>
                     <el-card>
                         <el-row>
                            <el-col :span="20">
                                <el-checkbox-group v-model="checkList2">
                                    <el-checkbox label="ini"></el-checkbox>
                                    <el-checkbox label="bin"></el-checkbox>
                                    <el-checkbox label="ply"></el-checkbox>   
                                    <el-checkbox label="json"></el-checkbox>                             
                                </el-checkbox-group>
                             </el-col> 
                            <el-col :span="4">
                                <el-button size="small">结果导出</el-button>
                            </el-col> 
                         </el-row>
                    </el-card>   
                     <el-divider content-position="left">标定</el-divider>
                     <el-card>
                         <el-row>
                             <el-col :span="20">
                                <el-checkbox-group v-model="checkList3">
                                    <el-checkbox label="启用K3"></el-checkbox>
                                    <el-checkbox label="使用控制点"></el-checkbox>
                                    <el-checkbox label="FOV全覆盖"></el-checkbox>                       
                                </el-checkbox-group>
                                <!-- <p>标定进度</p> -->
                                <el-progress :text-inside="true" :stroke-width="26" :percentage="70"></el-progress>
                             </el-col>
                            <el-col :span="4">
                                <el-button>标定</el-button>                            
                             </el-col>
                         </el-row>
                    </el-card> 
                </div>   
            </el-col>    
            <el-col :span="12">
                <el-divider content-position="left">参数配置</el-divider>
                <el-card>
                    <el-form label-position="left" label-width="100px"  size="mini">
                        <el-form-item label="标志点">
                            <el-select v-model="value" placeholder="编码位数">
                                <el-option label="73_71_17_13_33_25" value="1"> </el-option>
                                <el-option label="73_71_17_13_33_25" value="2"> </el-option>
                            </el-select>
                        </el-form-item> 
                        <el-form-item label="中间圆半径比">
                             <el-input v-model="form.Tea_value"></el-input>                                  
                        </el-form-item>
                        <el-form-item label="外圆半径比">
                            <el-input v-model="form.Tea_value"></el-input>                                                                               
                        </el-form-item>   
                        <el-form-item label="椭圆最小半径">
                            <el-input v-model="form.Tea_value"></el-input>                                                                               
                        </el-form-item>   
                        <el-form-item label="椭圆最大半径">
                            <el-input v-model="form.Tea_value"></el-input>                                                                               
                        </el-form-item>    
                        <el-form-item label="标志点">
                            <el-select v-model="value" placeholder="请选择">
                                <el-option label="黑色" value="1"> </el-option>
                            </el-select>
                        </el-form-item>  
                        <el-form-item label="拟合程度">
                            <el-input v-model="form.Tea_value"></el-input>                                                                               
                        </el-form-item>   
                         <el-form-item >
                            <el-button>保存参数</el-button>  
                            <el-button>编码点检测</el-button>                                                                                                                                                           
                        </el-form-item>   
                        <el-form-item >
                            <el-button>上一张</el-button>  
                            <el-button>下一张</el-button>                                                                                                                                                           
                        </el-form-item>   
                        <el-form-item >
                            <el-button>保存当前帧</el-button>  
                        </el-form-item>                                                         
                    </el-form>
                </el-card>
            </el-col>  
        </el-row>
    </div>                
</template>

<script>
export default {
    data(){
        return {
            checkList1:[],
            checkList2:[],
            checkList3:[],
            form:{},
            value:"1"
        }
    }
}
</script>