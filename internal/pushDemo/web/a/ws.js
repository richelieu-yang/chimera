var ws = null;
var urlInput = document.getElementById("urlInput"),
    connectBtn = document.getElementById("connectBtn"),
    disconnectBtn = document.getElementById("disconnectBtn"),
    clearBtn = document.getElementById("clearBtn"),
    console = document.getElementById("console");

var url = localStorage["url"];
url = url.trim();
if (!url) {
    url = "";
}
urlInput.value = url;

connectBtn.onclick = function () {
    println("[建立连接]")

    var url = urlInput.value;
    url = url.trim();
    if (!url) {
        alert("Please enter url!!!");
        return;
    }
    if (!url.startsWith("ws://") && !url.startsWith("wss://")) {
        alert("Invalid url!!!");
        return;
    }
    localStorage["url"] = url;

    if (ws != null) {
        ws.close();
        ws = null;
    }

    // if (source != null) {
    //     source.close()
    // }
    // source = new EventSource(url);
    //
    // source.onopen = function (event) {
    //     println("onopen")
    // };
    //
    // source.onmessage = function (e) {
    //     var origin = e.origin,
    //         id = e.lastEventId,
    //         event = e.type,
    //         data = e.data;
    //
    //     println("onmessage: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")")
    // };
    //
    // // 自定义事件
    // source.addEventListener('test', function (e) {
    //     var origin = e.origin,
    //         id = e.lastEventId,
    //         event = e.type,
    //         data = e.data;
    //
    //     println("ontest: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")")
    // }, false);
    //
    // source.onmessage = function (e) {
    //     var origin = e.origin,
    //         id = e.lastEventId,
    //         event = e.type,
    //         data = e.data;
    //
    //     println("onmessage: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")")
    // };
    //
    // // 如果发生通信错误（比如连接中断），就会触发error事件
    // source.onerror = function (event) {
    //     println("onerror")
    // };

    println("url: [" + url + "]")
};

disconnectBtn.onclick = function () {
    println("[断开连接]")

    // if (!source) {
    //     return
    // }
    // source.close();
    // source = null;
};

clearBtn.onclick = function () {
    console.value = "";
};

function println(text) {
    console.value += text + "\n";
}
