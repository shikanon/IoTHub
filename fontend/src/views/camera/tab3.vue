<template>
    <div>
        <el-row :gutter="20">
            <el-divider content-position="left">Aligment Config</el-divider>
            <el-card>
                <el-row :gutter="20">
                    <el-col :span="4"><span>对齐配置参数</span></el-col>
                    <el-col :span="16"><el-input size="mini"></el-input></el-col>
                    <el-col :span="4"><el-button size="mini">SelectConfig</el-button></el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="12" class="left_item">
                        <el-checkbox-group  v-model="form.show" size="middle">
                            <el-checkbox label="RectifyLeftImg"   border></el-checkbox>
                            <el-checkbox label="RectifyRightImg"  border></el-checkbox>
                            <el-checkbox label="IsCalibMulti"  border></el-checkbox>
                            <el-checkbox label="IsDeviceMulti"  border></el-checkbox>
                            <el-checkbox label="DisableDistortion"  border></el-checkbox>
                            <el-checkbox label="EnabkeDebug"  border></el-checkbox>
                            <el-checkbox label="IsLdmpOnRight"  border></el-checkbox>

                        </el-checkbox-group>
                    </el-col>
                    <el-col :span="6" >
                        <el-divider content-position="left">Depth</el-divider>
                        <el-card>
                            <el-form label-position="left" label-width="75px" size="mini">
                                <el-form-item label="ClipLeftW">
                                    <el-input></el-input>
                                </el-form-item>
                                <el-form-item label="ClipTopH">
                                    <el-input></el-input>
                                </el-form-item>
                                <el-form-item label="Scale">
                                    <el-input></el-input>
                                </el-form-item>
                            </el-form>   
                        </el-card>
                    </el-col>
                    <el-col :span="6">
                        <el-divider content-position="left">RGB</el-divider>
                        <el-card>
                            <el-form label-position="left" label-width="75px" size="mini">
                                <el-form-item label="ClipLeftW">
                                    <el-input></el-input>
                                </el-form-item>
                                <el-form-item label="ClipTopH">
                                    <el-input></el-input>
                                </el-form-item>
                                <el-form-item label="Scale">
                                    <el-input></el-input>
                                </el-form-item>
                            </el-form>  
                        </el-card>
                    </el-col>
                </el-row>
            </el-card>  
         </el-row>  
        <el-row :gutter="20">
            <el-col :span="3"><span>标定参数文件</span></el-col>
            <el-col :span="16"><el-input size="mini"></el-input></el-col>
            <el-col :span="5"><el-button size="mini">SelectCaliParam</el-button></el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="24">
                <el-form :inline="true" size="mini">
                    <el-form-item label="标定参数文件类型" >
                        <el-select v-model="form.type" placeholder="活动区域">
                            <el-option label="init" valueinit="init"></el-option>
                        </el-select>
                    </el-form-item>
                    <el-form-item label="标定对齐结果文件">
                        <el-input></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button @click="onSubmit">查询</el-button>
                    </el-form-item>
                </el-form>
            </el-col> 
           
        </el-row>
    </div>                
</template>

<script>
export default {
    data(){
        return{
            form:{show:"RectifyLeftImg"}
        }
    }
}
</script>

<style  scoped>
.el-row {
    margin-bottom: 20px;
  }
.left_item{
    margin-top:30px;
}
.left_item .el-checkbox{
    margin-bottom: 10px;
    width: 150px;
    margin-right: 50px;
}
.left_item .el-checkbox.is-bordered+.el-checkbox.is-bordered{
    margin-left: 0px;
}

.flex_item{
    display: flex;
}

</style>