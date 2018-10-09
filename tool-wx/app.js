//app.js
var wxRequest = require('utils/wxRequest')
var Session = require('utils/session')
App({
  onLaunch: function () {
    // 展示本地存储能力
    var logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
         var user = {};
          // 这里是登陆流程
          // @todo
           wx.getUserInfo({
            success: function (result){
              var userinfo = result.userInfo;
              wxRequest.postRequest("login?code="+res.code, userinfo).then(res1 =>{
                console.log(res1)
                console.log("res1.data.code", res1.data.code);
                // if(res1.data.code==0){
                   console.log("res1===>",res1.data.data.session.id);
                  Session.set(res1.data.data.session);
                 
                  console.log(Session.get().id);
                // }
                console.log("222");
              })
            }
          })
      }
    })
    // 获取用户信息
    wx.getSetting({
      success: res => {
        if (res.authSetting['scope.userInfo']) {
          // 已经授权，可以直接调用 getUserInfo 获取头像昵称，不会弹框
          wx.getUserInfo({
            success: res => {
              // 可以将 res 发送给后台解码出 unionId
              this.globalData.userInfo = res.userInfo

              // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
              // 所以此处加入 callback 以防止这种情况
              if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
              }
            }
          })
        }
      }
    })
  },
  getUserInfo:function(cb){
    var that = this
  },
  globalData: {
    userInfo: null
  }
})