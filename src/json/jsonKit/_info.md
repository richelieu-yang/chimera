## 获取指定字段的值
结合gjson轻松获取json字段
    https://req.cool/zh/docs/tutorial/handle-response/#%e7%bb%93%e5%90%88-gjson-%e8%bd%bb%e6%9d%be%e8%8e%b7%e5%8f%96-json-%e5%ad%97%e6%ae%b5

如果返回的响应体是 json 格式，想要获取指定字段的值，但又不想定义 struct 去 Unmarshal 来获取字段的值，可以结合 gjson 来获取指定字段的值。
