# Kafka
kafka在日志系统中只做日志存储及日志消息转发，对日志消息没有任何更改；

## 一、搭建kafka单机版
######  依赖：
- Java：openjdk version "1.8.0_191" 
- zooleeper： zookeeper-3.4.13
- kafka： kafka_2.11-2.1.0 

### 1.1  安装步骤：

#### 1.1.1 准备工作：
- 检查是否安装java，并确认java版本；
- 安装java
  1. 下载安装包安装(下载网址：https://www.oracle.com/technetwork/java/javase/downloads/index.html)
     1. 解压下载好的tar包；
     2. 配置安装包的bin目录到/etc/pforile文件中；
  
  2. yum源安装
  ```
  $ yum install java-1.8.0-openjdk.x86_64
  ```


#### 1.1.2 安装zookeeper
> zookeeper配置简介参考：
https://blog.csdn.net/lifupingcn/article/details/78327609

1. 下载安装包；
```
$ wget http://mirror.bit.edu.cn/apache/zookeeper/zookeeper-3.4.13/zookeeper-3.4.13.tar.gz
```
2. 解压,修改配置,启动：
```
$ tar -xzvf zookeeper-3.4.13.tar.gz
$ cd zookeeper-3.4.13
$ cp conf/zoo_sample.cfg conf/zoo.cfg
$ vi conf/zoo.cfg

dataDir=/tmp/zookeeper

clientPort=2182

```
3.  启动：
```
$ ./bin/zkServer.sh start conf/zoo_sample.cfg
```


#### 1.1.3 安装kafka

1.下载安装包：
```
$ wget http://mirror.bit.edu.cn/apache/kafka/2.1.0/kafka_2.11-2.1.0.tgz
```

2.解压启动：
```
$ tar -xzvf kafka_2.11-2.1.0.tgz
$ cd kafka_2.11-2.1.0
$ nohup bin/kafka-server-start.sh config/server.properties &
```

#### 1.1.4 测试kafka
1.创建topic

```
$ bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic cloudzp-test
```

2.查看topic
```
$ bin/kafka-topics.sh --list --zookeeper localhost:2181
```

3.创建生产者发送消息到主题中

```
$ bin/kafka-console-producer.sh --broker-list localhost:9092 --topic cloudzp-test
```

4.创建消费者从主题中消费消息；
```
$ bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic cloudzp-test  --from-beginning
```
- Note：这里提示： 旧版本与新版本的命令稍有差距 新版本中<code>--zookeeper localhost:2181</code>，
被替换为<code>--bootstrap-server localhost:9092</code>,其中localhost:9092为kafka的地址信息；


## Kafka文件存储机制
> 参考：https://www.cnblogs.com/jun1019/p/6256514.html
