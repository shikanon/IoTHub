<template>
    <div class="message-box" v-show="isShowMessageBox">
        <div class="mask" @click="cancel"></div>
        <div class="message-content">
        <svg class="icon" aria-hidden="true" @click="cancel">
          <!-- <use xlink:href="#icon-delete"></use> -->
          <span>X</span>
        </svg>
        <p class="title">{{ title }}</p>
        <div v-if="msgState === 'warning'">
          <div class="msgState" ></div>
          <p class="content warning">{{ content }}</p>
        </div>
        <p v-else class="content" >{{ content }}</p>
        <div>
          <input type="text" v-model="inputValue" v-if="isShowInput" ref="input">
        </div>
            <div class="btn-group">
                <button class="btn-default" @click="cancel" v-show="isShowCancelBtn">{{ cancelBtnText }}</button>
                <button class="btn-primary btn-confirm" @click="confirm" v-show="isShowConfimrBtn">{{ confirmBtnText }}</button>
            </div>
        </div>
    </div>
    </template>
    
    <script>
    export default {
      props: {
        title: {
          type: String,
          default: '标题'
        },
        content: {
          type: String,
          default: '这是弹框内容'
        },
        isShowInput: false,
        inputValue: '',
        msgState: {
          type: String,
          default: 'none'
        },
        isShowCancelBtn: {
          type: Boolean,
          default: true
        },
        isShowConfimrBtn: {
          type: Boolean,
          default: true
        },
        cancelBtnText: {
          type: String,
          default: '取消'
        },
        confirmBtnText: {
          type: String,
          default: '确定'
        }
      },
      data () {
        return {
          isShowMessageBox: false,
          resolve: '',
          reject: '',
          promise: '' // 保存promise对象
        };
      },
      methods: {
        // 确定,将promise断定为resolve状态
        confirm: function () {
          this.isShowMessageBox = false;
          if (this.isShowInput) {
            this.resolve(this.inputValue);
          } else {
            this.resolve('confirm');
          }
          this.remove();
        },
        // 取消,将promise断定为reject状态
        cancel: function () {
          this.isShowMessageBox = false;
          this.reject('cancel');
          this.remove();
        },
        // 弹出messageBox,并创建promise对象
        showMsgBox: function () {
          this.isShowMessageBox = true;
          this.promise = new Promise((resolve, reject) => {
            this.resolve = resolve;
            this.reject = reject;
          });
          // 返回promise对象
          return this.promise;
        },
        remove: function () {
          setTimeout(() => {
            this.destroy();
          }, 300);
        },
        destroy: function () {
          this.$destroy();
          document.body.removeChild(this.$el);
        }
      }
    };
    </script>
    <style  scoped >
      .message-box{
        position: fixed;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        text-align: center;
        z-index: 999;
     
      }

      .message-box:after {
          content: "";
          display: inline-block;
          height: 100%;
          width: 0;
          vertical-align: middle;
      }
      .mask{
          position: fixed;
          left: 0;
          top: 0;
          width: 100%;
          height: 100%;
          opacity: .5;
          background: #000;
      }
      .message-content{
        position: relative;
        display: inline-block;
        width: 420px;
        padding: 15px;
        vertical-align: middle;
        background-color: #FFF;
        border-radius: 4px;
        border: 1px solid #EBEEF5;
        font-size: 18px;
        -webkit-box-shadow: 0 2px 12px 0 rgba(0,0,0,.1);
        box-shadow: 0 2px 12px 0 rgba(0,0,0,.1);
        text-align: left;
        overflow: hidden;
        -webkit-backface-visibility: hidden;
        backface-visibility: hidden;
      }

      p{
        padding-bottom: 15px;
        width: 100%;
        height: auto;
        word-wrap: break-word;
        word-break: break-all;
        overflow: hidden;
      }

      .title{
        padding-left: 0;
        margin-bottom: 0;
        font-size: 18px;
        line-height: 20px;
        color: #303133;
        position: relative;
        left: 0;
        margin-top: 0;
        height: 20px;
      }
      
      .content{
        margin: 0;
        line-height: 24px;
        font-size:16px;
      }
      .msgState{
        position: absolute;
        top: 50%;
        -webkit-transform: translateY(-50%);
        transform: translateY(-50%);
        font-size: 24px!important;
        color: #E6A23C;
      }
      .msgState:before {
        padding-left: 1px;
        content: "\E7A3";
      }

      .warning{
         padding-left: 36px;

      }
      input{
        position: relative;
        font-size: 14px;
        display: inline-block;
        width: 100%;
        -webkit-appearance: none;
        background-color: #fff;
        background-image: none;
        border-radius: 4px;
        border: 1px solid #dcdfe6;
        box-sizing: border-box;
        color: #606266;
        display: inline-block;
        font-size: inherit;
        height: 40px;
        line-height: 40px;
        outline: none;
        padding: 0 15px;
        transition: border-color .2s cubic-bezier(.645,.045,.355,1);
        width: 100%;
        margin-bottom: 15px;
      }
      .btn-group{
        text-align: right;     
      }
      button{     
          display: inline-block;
          line-height: 1;
          white-space: nowrap;
          cursor: pointer;
          background: #FFF;
          border: 1px solid #DCDFE6;
          color: #606266;
          -webkit-appearance: none;
          text-align: center;
          -webkit-box-sizing: border-box;
          box-sizing: border-box;
          outline: 0;
          margin: 0;
          -webkit-transition: .1s;
          transition: .1s;
          font-weight: 500;
          padding: 9px 15px;
          font-size: 14px;
          border-radius: 4px;
      }
      .btn-primary{
        background: #66b1ff;
        border-color: #66b1ff;
        color: #FFF;
      }
      
      .icon{
        position: absolute;
        top: 15px;
        right: 15px;
        padding: 0;
        border: none;
        outline: 0;
        background: 0 0;
        font-size: 16px;
        cursor: pointer;
        height: 20px;
        width: 20px;
      }
    </style>