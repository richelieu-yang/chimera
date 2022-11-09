function get(url) {
    let xmlHttp;

    if (window.ActiveXObject) {											//IE6, IE5 浏览器执行代码
        xmlHttp = new ActiveXObject("Microsoft.XMLHTTP");
    } else if (window.XMLHttpRequest) {									//IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        xmlHttp = new XMLHttpRequest();
    }
    xmlHttp.open("GET", url, false);
    xmlHttp.send();
    //回调函数，监听response消息事件
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState === 4 && xmlHttp.status === 200) {
            console.log("responseText: " + xmlHttp.responseText);
        }
    };
}