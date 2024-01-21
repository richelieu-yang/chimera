## 结构体中有"匿名字段"
(1) 匿名字段有json tag（json tag将作为key）
e.g.
代码:
type user struct {
    gorm.Model `json:"model"`

    Name string `json:"name"`
}

u := &user{
    Model: gorm.Model{
        ID: 1,
    },
    Name: "test",
}
fmt.Println(MarshalToString(u, WithIndent("    ")))
输出:
{
    "model": {
        "ID": 1,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null
    },
    "name": "test"
} <nil>

(2) 匿名字段无json tag（内联）
代码:
type user struct {
    gorm.Model

    Name string `json:"name"`
}

u := &user{
    Model: gorm.Model{
        ID: 1,
    },
    Name: "test",
}
fmt.Println(MarshalToString(u, WithIndent("    ")))
输出:
{
    "ID": 1,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "name": "test"
} <nil>

## map是无序键值对
e.g.
由于map是无序的键值对集合，多次序列化一个map实例为json字符串，得到的结果可能不一.如果希望多次得到的结果一致，
可以通过 jsonKit.WithApi(jsoniter.ConfigCompatibleWithStandardLibrary) 进行配置.