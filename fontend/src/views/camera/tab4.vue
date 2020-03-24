<template>
    <div>
        <el-row >
            <el-divider content-position="left">Depth Param</el-divider>
            <el-card>
                 <el-form :inline="true" size="mini" label-width="30px" class="inner-from">
                     <el-row>
                        <el-form-item label="fx" >
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="fy">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="cx">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="cy">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                     </el-row>
                    <el-row>
                        <el-form-item label="k1" >
                             <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="k2">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="p1">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="p2">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                         <el-form-item label="k3">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                     </el-row>
                </el-form>
            </el-card>    
        </el-row>
        <el-row >
            <el-col :span=20>
                <el-divider content-position="left">RGB param</el-divider>
                <el-card>
                    <el-form :inline="true" size="mini" label-width="30px" class="inner-from">
                        <el-row>
                            <el-form-item label="fx" >
                                <el-input class="input_120"></el-input>
                            </el-form-item>
                            <el-form-item label="fy">
                                <el-input class="input_120"></el-input>
                            </el-form-item>
                             <el-form-item label="cx">
                                <el-input class="input_120"></el-input>
                            </el-form-item>
                            <el-form-item label="cy">
                                 <el-input class="input_120"></el-input>
                            </el-form-item>
                        </el-row>
                        <el-row>
                            <el-form-item label="k1" >
                               <el-input class="input_120"></el-input>
                            </el-form-item>
                            <el-form-item label="k2">
                               <el-input class="input_120"></el-input>
                            </el-form-item>
                            <el-form-item label="p1">
                               <el-input class="input_120"></el-input>
                            </el-form-item>
                            <el-form-item label="p2">
                               <el-input class="input_120"></el-input>
                            </el-form-item>
                            <el-form-item label="k3">
                               <el-input class="input_120"></el-input>
                            </el-form-item>
                        </el-row>
                    </el-form>             
                </el-card> 
            </el-col>
             <el-col :span="4">
                 <el-row>
                     <el-button size="small">read</el-button>
                 </el-row>
                 <el-row >
                    <el-button size="small">write</el-button>
                 </el-row>
             </el-col>
        </el-row>
         <el-row >
            <el-divider content-position="left">Depth to RGB param</el-divider>
            <el-card>
                <el-form :inline="true" size="mini" label-width="40px" >
                    <el-row>
                        <el-form-item label="r00" >
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="r01">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="r02">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="t1">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                    </el-row>
                    <el-row>
                        <el-form-item label="r10" >
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="r11">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="r12">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="t2">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                
                    </el-row>
                    <el-row>
                        <el-form-item label="r20" >
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="r21">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="r22">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                        <el-form-item label="t3">
                            <el-input class="input_120"></el-input>
                        </el-form-item>
                
                    </el-row>
                </el-form>
            </el-card> 
        </el-row>
        <el-row class="margin_top_20">
            <el-col :span="4"><span>内外参数配置文件</span></el-col>
            <el-col :span="14"><el-input size="mini"></el-input></el-col>
            <el-col :span="4" :offset="2"><el-button size="mini">Select</el-button></el-col>
        </el-row>
    </div>                
</template>


<script>
export default {
    data(){
        return{
            form:{}
        }
    }
}
</script>

<style scoped>
    .input_120{
        width: 120px;
    }
    .margin_top_20{
        margin-top: 20px;
    }
    
</style>>

