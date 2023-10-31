var source;

window.onload = function () {
    var urlInput = document.getElementById("urlInput"),
        connectBtn = document.getElementById("connectBtn"),
        disconnectBtn = document.getElementById("disconnectBtn"),
        console = document.getElementById("console");

    // var sseUrl = window.location.href;
    // var index = sseUrl.indexOf("/s");
    // if (index !== -1) {
    //     sseUrl = sseUrl.substring(0, index);
    //     sseUrl += "/sse";
    // }
    // urlInput.value = sseUrl;

    var url = localStorage["url"]
    if (!url) {
        url = "";
    }
    urlInput.value = url;

    connectBtn.onclick = function () {
        var url = urlInput.value;
        if (!url) {
            alert("请输入url!!!");
            return;
        }

        localStorage["url"] = url;

        if (source != null) {
            source.close()
        }

        source = new EventSource(url);
        source.onopen = function (event) {
            println("onopen")
        };
        source.onmessage = function (e) {
            var origin = e.origin,
                id = e.lastEventId,
                event = e.type,
                data = e.data;

            println("onmessage: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")")
        };
        // 自定义事件
        source.addEventListener('test', function (e) {
            var origin = e.origin,
                id = e.lastEventId,
                event = e.type,
                data = e.data;

            println("ontest: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")")
        }, false);
        source.onmessage = function (e) {
            var origin = e.origin,
                id = e.lastEventId,
                event = e.type,
                data = e.data;

            println("onmessage: origin(" + origin + "), id(" + id + "), event(" + event + "), data(" + data + ")")
        };

        // 如果发生通信错误（比如连接中断），就会触发error事件
        source.onerror = function (event) {
            println("onerror")
        };

        println("url: [" + url + "]")
    };

    disconnectBtn.onclick = function () {
        if (!source) {
            return
        }
        source.close();
        source = null;
    };

    function println(text) {
        console.value += text + "\n";
    }
};
