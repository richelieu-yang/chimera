## 多次关闭连接
gorilla/websocket 中的 Conn.Close()：
    可以多次调用，不会panic，但从第二次关闭开始，返回非nil的error（可以直接忽略）.
