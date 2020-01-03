<template>
 <div name="object-mode">
        <div class="desc-txt-gray margin-b-20">
            物模型是对设备在云端的功能描述，包括设备的属性、服务和事件。
            物联网平台通过定义一种物的描述语言来描述物模型，
            称之为 TSL (即Thing Specification Language)，
            采用JSON格式，您可以根据TSL组装上报设备的数据。您可以导出完整物模型，
            用于云端应用开发；您也可以只导出精简物模型，配合设备端SDK实现设备开发。
        </div>
        <el-tabs v-model="activeName" type="card">
            <el-tab-pane label="完整物模型" name="complete">
                 <json-view :data="completeObject"/>
            </el-tab-pane>
            <el-tab-pane label="精简物模型" name="simplify">
                 <json-view :data="simplifyObject"/>
            </el-tab-pane>
     </el-tabs>
       
</div>
    
</template>
 
<script>
    import jsonView from '@/components/json-view';
 
    export default {
        components: {
            jsonView
        },
        data() {
            return {
                completeObject: {},
                simplifyObject:{},
                 activeName:'complete',

            }
        }
    }
