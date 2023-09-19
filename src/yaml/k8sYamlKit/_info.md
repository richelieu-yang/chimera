## sigs.k8s.io/yaml
github: https://github.com/kubernetes-sigs/yaml

优点:
(1) 支持json tag（即不需要yaml tag）.
缺点:
(2) 同一层级的字段，顺序与结构体定义的不一致（按照json tag字典排序）.
