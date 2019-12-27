<template>
    <el-form label-position="top" :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" >
     <el-form-item label="父组" prop="SuperGroup">
          <el-cascader :props="group"></el-cascader>
    </el-form-item>
    <el-form-item label="分组名称" prop="GroupName">
        <el-input v-model="ruleForm.GroupName"></el-input>
    </el-form-item>
   
    <el-form-item label="分组描述" prop="Desc">
        <el-input type="textarea" v-model="ruleForm.Desc"></el-input>

    </el-form-item>
    
    </el-form>
 
</template>
<script>
  export default {
    
    data() {
      return {
        ruleForm: {
          SuperGroup:'',
          GroupName: '',
          Desc: '',
       
        },
        group:{},
        rules: {
          GroupName: [
            { required: true, message: '请输入分组名称', trigger: 'blur' },

          ]
              
        }
      };
    },
    created(){
       this.getSuperGroup()

    },
    methods: {
        
        getSuperGroup(){
             let id = 0;
             this.group= {
                  lazy: true,
                   lazyLoad (node, resolve) {
                    const { level } = node;
                    setTimeout(() => {
                    const nodes = Array.from({ length: level + 1 })
                        .map(item => ({
                        value: ++id,
                        label: `选项${id}`,
                        leaf: level >= 2
                        }));
                    // 通过调用resolve将子节点数据返回，通知组件数据加载完成
                    resolve(nodes);
                    }, 1000);
                }
            }
            console.log(group)
        },
        
        
      submitForm() {
        this.$refs.ruleForm.validate((valid) => {
          if (valid) {
            alert('submit!');
            //回调父组件方法关闭弹窗
            this.$emit('close')
            //回调父组件方法
             this.$emit('getGroupList');
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },


     
    }
  }
</script>
<style scoped >
  .el-select{
      width: 440px;
  }
</style>