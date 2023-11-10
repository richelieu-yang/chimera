var localStorageKey = "wsUrl";
var channel = null;
var urlInput = document.getElementById("urlInput"),
    connectBtn = document.getElementById("connectBtn"),
    disconnectBtn = document.getElementById("disconnectBtn"),
    clearBtn = document.getElementById("clearBtn"),
    output = document.getElementById("output");

var url = getFromLocalStorage()
urlInput.value = url;

connectBtn.onclick = function () {
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

    println("[建立连接]")

    connect(url);
    setToLocalStorage(url);
    println("url: [" + url + "]")
};

disconnectBtn.onclick = function () {
    println("[断开连接]")

    disconnect();
};

clearBtn.onclick = function () {
    output.value = "";
};

/**
 * PS: EventSource 没有onclose事件.
 */
function connect(url) {
    disconnect();

    channel = new WebSocket(url);
    channel.onopen = function () {
        println("onopen");
    };
    channel.onmessage = function (e) {
        var data = e.data;

        if (data instanceof ArrayBuffer) {
            let blob = new Blob([data]), reader = new FileReader();

            reader.readAsText(blob, "UTF-8");
            reader.onload = () => {
                var text = reader.result;
                println("on binary message: " + text);
            };
        } else if (data instanceof Blob) {
            var reader = new FileReader();

            reader.readAsText(data, "UTF-8");
            reader.onload = () => {
                var text = reader.result;
                println("on binary message: " + text);
            };
        } else if (typeof data === "string") {
            var text = e.data;
            println("on text message: " + text);
        }
    };
    channel.onerror = function (e) {
        println("onerror");
    };
    channel.onclose = function () {
        println("onclose");
    };
}

function disconnect() {
    if (channel == null) {
        return;
    }

    channel.close();
    channel = null;
}

function println(text) {
    output.value += text + "\n";
    console.info(text);
}

function getFromLocalStorage() {
    var url = localStorage[localStorageKey];
    if (!url) {
        url = "";
    }
    return url.trim();
}

function setToLocalStorage(url) {
    localStorage[localStorageKey] = url.trim();
}
