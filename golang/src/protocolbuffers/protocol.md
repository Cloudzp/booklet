# Protocol Buffers
## 一、 Protocol Buffers 是什么？protoc是什么？ 
 - Protocol Buffers (ProtocolBuffer/ protobuf )是Google公司开发的一种数据描述语言，类似于XML能够将结构化数据序列化，可用于数据存储、通信协议等方面
 - protoc是Protocol Buffers的一个工具，用来支持将Protocol Buffers定义的文件转换为各种语言的客户端代码。
 
## 二、 为什么要有Protocol Buffers？
1、通过它，可以定义我们的数据的结构，并生成基于各种语言的代码。这些你定义的数据流可以轻松地在传递并不破坏我们原有的程序。并且也可以更新这些数据而现有的程序也不会受到任何的影响。
2、而且同XML相比，Protocol buffers在序列化结构化数据方面有许多优点：
- （1）更简单
- （2）数据描述文件只需原来的1/10至1/3
- （3）解析速度是原来的20倍至100倍
- （4）减少了二义性
- （5）生成了更容易在编程中使用的数据访问类
- （6）支持多种编程语言
> 参考：
// 阿里技术文档
1. https://yq.aliyun.com/ziliao/580483 
2. https://www.ibm.com/developerworks/cn/linux/l-cn-gpb/

### 三、如何安装使用？
1.protor的安装 

>参考： https://blog.csdn.net/JustinSeraph/article/details/70171331

2.使用
2.1 使用命令生成proto文件的go语言代码，然后根据go语言代码去做开发即可；
```
protoc --go_out=./go/ ./proto/helloworld.proto
```
[点击查看示例代码](src/main.go)

> 参考：
1. https://www.cnblogs.com/chenyangyao/p/5422044.html (protobuf 介绍)
2. https://github.com/golang/protobuf (go protobuf库)