<template>
    <div class="block">
        <el-timeline>
            <el-timeline-item
            v-for="(activity, index) in activities"
            :key="index"
            :icon="activity.icon"
            :type="activity.type"
            :color="activity.color"
            :size="activity.size"
            :timestamp="activity.timestamp">
            {{activity.content}}
            </el-timeline-item>
        </el-timeline>
    </div>
</template>


<script>
  export default {
    props:{
     activities:{
            type:Array,
            default: () => []
        }
    },
    data() {
      return {
       
      }
    },

    create(){
    },

  };
</script>