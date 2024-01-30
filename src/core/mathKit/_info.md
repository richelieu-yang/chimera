## 注意点
#### shopspring/decimal 的 decimal.NewFromFloat()
- 传入的参数是 float64 类型，不能是 NaN, +/-inf（会panic）!!!

