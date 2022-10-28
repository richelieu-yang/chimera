var url = "ws://127.0.0.1:8080/ping",
    ws = new WebSocket(url);

ws.onopen = function () {
    console.info("onopen");
};
ws.onerror = function () {
    console.error("onerror");
};
ws.onclose = function () {
    console.warn("onclose");
};