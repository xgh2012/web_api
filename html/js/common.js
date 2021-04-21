const baseHost = "http://app.longguanjia.so:1051/";

var layer = null;

layui.use(['layer','jquery'], function () {
    layer = layui.layer;
    layer.config(
        {
            // extend: 'wmhd/style.css', //加载您的扩展样式
            // skin: 'layer-ext-wmhd',
        }
    );
});

$.ajaxSetup({
    headers: {'Authorization': getAuthToken()}
});

//获取真实请求地址
var getUrl = function (uri) {
    return "/api/" + uri
};

//退出登录
var checkIsLoginOut = function(res){
    if (res.authorization !== undefined && res.authorization !== ""){
        if (setAuthToken(res.authorization)){
            //表格查询渲染 TODO
        }
    }
    if (res.is_login !== undefined && res.is_login === -1){
        loginOut(res.msg !== undefined ? res.msg + ",请重新登录":"请重新登录")
        return false;
    }
    if (res.code !== 0){
        layer.alert(res.msg,{icon:2});
        return false;
    }
    return true;
};

//退出登录
var loginOut = function(msg){
    layer.alert(msg,{icon:2});
    clearCookie("Authorization")
    setTimeout(function () {
        window.location = '/web/login.html';
    },1000);
};

function getAuthToken() {
    return getCookie("Authorization");
}
//设置验证token
function setAuthToken(newAuth) {
    var Auth = getCookie("Authorization");
    if (Auth !== newAuth){
        setCookie("Authorization",newAuth,1);
        return true;
    }
    return false;
}

// 设置cookie
function setCookie(c_name, value, expiredays) {
    var exDate = new Date();
    exDate.setDate(exDate.getDate() + expiredays);
    document.cookie = c_name + "=" + escape(value) + ";expires=" + exDate.toGMTString() + ";path=/";
}

// 读取cookie
function getCookie(c_name) {
    if (document.cookie.length > 0)     {
        c_start = document.cookie.indexOf(c_name + "=")
        if (c_start !== -1){
            c_start = c_start + c_name.length + 1
            c_end = document.cookie.indexOf(";", c_start)
            if (c_end === -1)
                c_end = document.cookie.length
            return unescape(document.cookie.substring(c_start, c_end))
        }
    }
    return ""
}

// 清除cookie
function clearCookie(name) {
    setCookie(name, "", -1);
}

//rt true 必须返回数据
function post(url,data,callback,doing,rt) {
    $.ajax({
        url: "/api/" + url,
        method: 'post',
        data:data,
        success: function (res) {
            if (checkIsLoginOut(res) === false && rt !== true){
                return false;
            }
            if (callback === undefined){
                layer.alert(res.msg);
            }else{
                callback(res.data)
            }
        },
        error: function (xhr, textStatus, errorThrown) {
        },
        beforeSend:function (request) {
            if (rt !== true){
                request.setRequestHeader("Authorization", getAuthToken());
            }
            if (doing !== undefined){
                doing()
            }
        }
    });
}

function success(msg) {
    layer.alert(msg, {
        icon :"1",
        title: '提交信息'
    })
}
function wrong(msg) {
    layer.alert(msg, {
        icon :"2",
        title: '错误提示'
    })
}

/**
 * 获取hash参数
 */
function getHashParameter(key){
    var params = getHashParameters();
    return params[key];
}

function getHashParameters(){
    var arr = (location.hash || "").replace(/^\#/,'').split("&");
    var params = {};
    for(var i=0; i<arr.length; i++){
        var data = arr[i].split("=");
        if(data.length == 2){
            params[data[0]] = data[1];
        }
    }
    return params;
}
