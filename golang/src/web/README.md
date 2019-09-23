# web-go
   由于近期经常与一些第三方的服务对接，一般的企业大多用java语言开发，大量使用web form表单来提交数据，上传文见，
而我们的所有服务都是用go语言开发的，所以对go web相关的一些功能进行一次总结：
## 关于form表单
form表单是web服务常用来传输数据的一种形式，主要分为get、和post两中，一般：
- get主要用来获取数据，数据会被串联在url中以'?k=v&k2=v2'的形式，在服务端如果需要获取提交的参数可以参考以下代码：
[Get方法](client/main.go)
- post将数据编码到body体中
[Post方法](client/main.go)
- 上传文件同时上传数据
[upload方法](client/main.go)

