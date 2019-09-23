# filebeat
## 一、部署filebeat
1. 下载软件包
```
$ wget https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-6.5.4-linux-x86_64.tar.gz
```
2. 修改配置（参考1.1中的配置介绍）
```
$ tar -xzvf filebeat-6.5.4-linux-x86_64.tar.gz
$ cd filebeat-6.5.4-linux-x86_64
$ vi filebeat.yml

# Paths that should be crawled and fetched. Glob based paths.
  paths:
    - /home/data1/lizean/upayserver/Logs/*.log
    - /home/data1/lightyagami/logs/*.log
    #- c:\programdata\elasticsearch\logs\*
```
3. 启动
```
$ nohup nohup ./filebeat -e -c filebeat.yml &
```

### 1.1 filebeat 配置介绍
> 参考 ：
- http://blog.51cto.com/seekerwolf/2110174
- https://www.elastic.co/guide/en/beats/filebeat/current/kafka-output.html (官网介绍)

配置文件简介：

1. 配置分为输入配置，输出配置
1.1 输入配置主要配置日志来源（filebeat.inputs），这里主要使用log日志采集的形式，filebeat.inputs就是一块日志；
```
filebeat.inputs:
- type: log
  enabled: true // 次块配置是否生效，
  paths：
    - /logpath/*.log      //日志目录，支持正则匹配
  tags: ["testlog"]   // 支持为日志增加标签，以便筛选
  exclude_lines: ['DEBUG'] // 根据正则过滤筛选日志每行的内容，如：exclude_lines: ['DEBUG']表示筛选掉包含DEBUG的日志行
  include_lines: ['INFO'] // 与exclude_lines功能一样，表示包含筛选出包含某个字段的那行日志
  exclude_files: ['.gz$'] // 表示不采集以.gz结尾的压缩日志
  fields: // 可以在此新增字段，比如topic名称
    level: INFO   // 为日志设置级别
  // multiline 有的时候日志不是一行输出的，如果不用multiline的话，会导致一条日志被分割成多条收集过来，形成不完整日志，这样的日志对于我们来说是没有用的！通过正则匹配语句开头，这样multiline 会在匹配开头之后，一直到下一个这样开通的语句合并成一条语句。
  //有以下属性：
  //#pattern：多行日志开始的那一行匹配的pattern
  //#negate：是否需要对pattern条件转置使用，不翻转设为true，反转设置为false
  //#match：匹配pattern后，与前面（before）还是后面（after）的内容合并为一条日志
  //#max_lines：合并的最多行数（包含匹配pattern的那一行 默认值是500行
  //#timeout：到了timeout之后，即使没有匹配一个新的pattern（发生一个新的事件），也把已经匹配的日志事件发送出去  
  multiline.pattern: '^\d{4}/\d{2}/\d{2}'
  
#-------------------------- Kafka output ------------------------------
output.kafka:
  hosts: ["localhost:9092"]
  topic: '%{[fields.log_topic]}'
  keep_alive: 3000 
   
```
   