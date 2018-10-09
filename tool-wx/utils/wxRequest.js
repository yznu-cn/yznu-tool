var Promise = require('../plugins/es6-promise.js')
var Session = require('./session');
function wxPromisify(fn) {
  return function (obj = {}) {
    return new Promise((resolve, reject) => {
      obj.success = function (res) {
        //成功
        resolve(res)
      }
      obj.fail = function (res) {
        //失败
        reject(res)
      }
      fn(obj)
    })
  }
}
//无论promise对象最后状态如何都会执行
Promise.prototype.finally = function (callback) {
  let P = this.constructor;
  return this.then(
    value => P.resolve(callback()).then(() => value),
    reason => P.resolve(callback()).then(() => { throw reason })
  );
};

var buildHeader = function buildHeader() {
    var header = {'Content-Type': 'application/json'};
    var session = Session.get();
    if (session && session.id) {
        header['Cookie'] ='sessionID='+session.id;
    }
    return header;
};

/**
 * 微信请求get方法
 * url
 * data 以对象的格式传入
 */
function getRequest(url, data) {
  var getRequest = wxPromisify(wx.request)
  return getRequest({
    url: "http://localhost:8000/api/v1/" + url,
    method: 'GET',
    data: data,
    header: buildHeader(),
  })
}

/**
 * 微信请求post方法封装
 * url
 * data 以对象的格式传入
 */
function postRequest(url, data) {
  var postRequest = wxPromisify(wx.request)
  return postRequest({
    url: "http://localhost:8000/api/v1/" + url,
    method: 'POST',
    data: data,
    header: buildHeader(),
  })
}

/**
 * 微信请求put方法封装
 * url
 * data 以对象的格式传入
 */
function putRequest(url, data) {
  var putRequest = wxPromisify(wx.request)
  return putRequest({
    url: "http://localhost:8000/api/v1/" + url,
    method: 'PUT',
    data: data,
    header: buildHeader(),
  })
}


/**
 * 微信请求DELETE方法封装
 * url
 * data 以对象的格式传入
 */
function deleteRequest(url) {
  var deleteRequest = wxPromisify(wx.request)
  return deleteRequest({
    url: "http://localhost:8000/api/v1/" + url,
    method: 'DELETE',
    header: buildHeader(),
  })
}


module.exports = {
  postRequest: postRequest,
  getRequest: getRequest,
  putRequest: putRequest,
  deleteRequest: deleteRequest
}