# panic的一些特性测试

### 1.panic是否只能在当前协程内捕获？
验证结果：是的，panic只能在当前协程内recover，且recover到的值是一个interface。
