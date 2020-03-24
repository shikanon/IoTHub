<template>
    <div id="camera-calibration">
        <div class="top_item">
            <el-row :gutter="20">
                <el-col :span="4">
                     <el-divider content-position="left">Show</el-divider>
                    <el-card >
                        <el-checkbox-group  v-model="form.show" size="small">
                            <el-checkbox v-for="item in showOptions" :label="item" :key="item"  size="small" border>{{item}}</el-checkbox>
                        </el-checkbox-group>
                    </el-card>
                </el-col>
                <el-col :span="6">
                    <el-divider content-position="left">Save param</el-divider>
                    <el-card >
                        <el-checkbox-group v-model="form.sava" size="small">
                            <el-checkbox v-for="item in savaOptions" :label="item" :key="item" border>{{item}}</el-checkbox>
                        </el-checkbox-group>
                    </el-card>
                </el-col>
                <el-col :span="8">
                   <el-divider content-position="left">Camera</el-divider>
                    <el-card class="card-8" >              
                        <el-radio v-model="form.camera" :label="item" v-for="item in cameraOptions" :key = "item" border size="small">{{item}}</el-radio>
                       
                        <el-checkbox-group v-model="form.option" size="small">
                            <el-checkbox v-for="option in options" :label="option" :key="option" border>{{option}}</el-checkbox>
                        </el-checkbox-group> 
                    </el-card>
                </el-col>
                <el-col :span="6">
                     <el-divider content-position="left">real time temprature</el-divider>
                    <el-card >
                        <el-form label-position="left" label-width="60px"  size="mini" >
                            <el-form-item label="ir">
                                  <!-- <el-col :span="20">   -->
                                    <el-input v-model="form.ir"></el-input>  
                                  <!-- </el-col>  -->
                                   <!-- <el-col :span="4"> 
                                       ℃
                                   </el-col>                            -->
                            </el-form-item>
                            <el-form-item label="ldmp">
                                <el-input v-model="form.ldmp"></el-input>                             
                            </el-form-item>
                            <el-form-item label="z0">
                                <el-input v-model="form.z0"></el-input>
                            </el-form-item>
                            <el-form-item label="ZPPS">
                                <el-input v-model="form.ZPPS"></el-input>
                            </el-form-item>
                            <el-form-item label="基线">
                                <el-input v-model="form.type"></el-input>
                            </el-form-item>
                        </el-form>
                    </el-card>
                </el-col>
             </el-row>         
        </div>    
        <div class="middle_item">
            <el-row :gutter="8">
                <el-col :span="14">
                    <el-divider content-position="left">Parsm</el-divider>
                    <el-card >
                        <el-form label-position="left" label-width="60px"  size="mini">
                            <el-form-item label="SN">
                                <el-col :span="16">                                 
                                     <el-input v-model="form.SN"></el-input>                                  
                               </el-col>
                                <el-col :span="8">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>
                                </el-col>
                            </el-form-item>
                            <el-form-item label="IR">
                                <el-col :span="16">                                 
                                     <el-input v-model="form.IR"></el-input>                                  
                               </el-col>
                                <el-col :span="8">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>                                
                                </el-col>
                            </el-form-item>
                            <el-form-item label="Depth">
                                 <el-col :span="16">                                 
                                     <el-input v-model="form.Depth"></el-input>                                  
                               </el-col>
                                <el-col :span="8">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>                                
                                 </el-col>
                             </el-form-item>
                            <el-form-item label="Color">
                                 <el-col :span="16">                                 
                                     <el-input v-model="form.Color"></el-input>                                  
                               </el-col>
                                <el-col :span="8">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>                                
                                </el-col>
                             </el-form-item>    
                        </el-form>
                    </el-card>
                </el-col>
                          
                <el-col :span="10">
                    <el-divider content-position="left">gain/exposure/cerrent/pluse</el-divider>
                    <el-card >
                        <el-form label-position="left" label-width="75px"  size="mini">
                            <el-form-item label="gain">
                                  <el-col :span="12">                                 
                                     <el-input v-model="form.gain"></el-input>                                  
                               </el-col>
                                <el-col :span="12">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>                                
                                </el-col>
                            </el-form-item>
                            <el-form-item label="exposure">
                                  <el-col :span="12">                                 
                                     <el-input v-model="form.exposure"></el-input>                                  
                               </el-col>
                                <el-col :span="12">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>                                
                                </el-col>
                            </el-form-item>
                            <el-form-item label="current">                                                      
                                 <el-col :span="12">                                 
                                     <el-input v-model="form.current"></el-input>                                  
                                </el-col>
                                <el-col :span="12">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>                                
                                </el-col>                             
                            </el-form-item>
                            <el-form-item label="laserTime">
                                <el-col :span="12">                                 
                                     <el-input v-model="form.laserTime"></el-input>                                  
                                </el-col>
                                <el-col :span="12">
                                    <el-button size="mini">get</el-button>
                                    <el-button size="mini">set</el-button>                                
                                </el-col>
                            </el-form-item>                         
                        </el-form>
                    </el-card>
                </el-col>
             </el-row> 
        </div> 
        <div class="bottom_item">
            <el-tabs v-model="activeName" >
                <el-tab-pane label="温度/对齐/TEC采图" name="first">
                  <tab1></tab1>
                </el-tab-pane>
                <el-tab-pane label="计算标定参数" name="second">
                    <tab2></tab2>
                </el-tab-pane>
                <el-tab-pane label="标定对齐" name="third">
                    <tab3></tab3>
                </el-tab-pane>
                <el-tab-pane label="读写内外参数" name="fourth">
                       <tab4></tab4>
                </el-tab-pane>
                <el-tab-pane label="参考图" name="five">
                        <tab5></tab5>
                </el-tab-pane>
            </el-tabs>
        </div> 
    </div>    
</template>

<script>
 import tab1 from './tab1'
 import tab2 from './tab2'
 import tab3 from './tab3'
 import tab4 from './tab4'
 import tab5 from './tab5'
export default {
    components:{tab1,tab2,tab3,tab4,tab5},
    data(){
        return {
            showOptions:['ir','Depth','color','uvc'],
            savaOptions:['ir','depth','color','raw','png','bmg','offLine'],
            cameraOptions:['单目','双目'],
            options:['Laser','floodlight','LDP','IrMirror','IsSwitch','DepthMirror'],
            form:{show:['ir'],sava:['png','bmg'],camera:'单目',option:['floodlight','ldp']},
            activeName:'first'
        }
    }
}
</script>

<style scoped>
#camera-calibration{
    display: flex;
    flex-direction: column;
    height:500px;
    overflow-y:auto;
    overflow-x: hidden;

}
.top_item,.middle_item{

    margin-bottom: 20px;
}
.el-form-item--mini.el-form-item{
        margin-bottom: 5px;
}
.top_item .el-checkbox{
    margin-bottom: 10px;
    width: 80px;
    margin-right: 5px;
}
.top_item .el-checkbox.is-bordered+.el-checkbox.is-bordered{
    margin-left: 0px;
}
.top_item .card-8 .el-checkbox-group{
       margin-top: 10px;
}
.top_item .card-8 .el-checkbox{
    width: 110px;
    margin-right: 14px;
 
}
    
   


</style>