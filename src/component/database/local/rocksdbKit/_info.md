## 注意点!!!
DB.Get()返回的 *grocksdb.Slice 实例，需要手动释放（Slice.Free()），否则会造成内存泄漏。
