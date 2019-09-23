# Elasticsearch

### 一、部署elasticSearch

1. 下载安装包

```
$ wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-6.5.4.tar.gz
```

2. 解压并修改参数：

```
$ tar -xzvf elasticsearch-6.5.4.tar.gz
$ vi config/elasticsearch.yml


# Path to directory where to store the data (separate multiple locations by comma):
# 修改数据地址
path.data: /home/data1/cloudzp/elaticsearch/data
#
# Path to log files:
# 修改日志地址
path.logs: /home/data1/cloudzp/elaticsearch/logs
# 修改IP
network.host: 172.16.65.220
#
# Set a custom port for HTTP:
# 端口
#http.port: 9200
```

3. 启动
```
$ nohup bin/elasticsearch &
```

Note: 如果启动遇到以下问题，可按照如下方法解决：

问题一:
```
ERROR: [4] bootstrap checks failed
[1]: max file descriptors [4096] for elasticsearch process is too low, increase to at least [65536]
```

- 解决方法: 修改可创建的文件数目：

```
$ su root
$ vi /etc/security/limits.conf

hadoop          soft    nofile  100000   # soft表示为超过这个值就会有warnning,如果之前很小将这个值改大；
hadoop          hadr    nofile  100000  # hard则表示不能超过这个值,如果之前很小将这个值改大；
```

问题二：
```
[1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]
```
- 解决方法: 修改该虚拟内存的使用限制

增加虚拟内存的大小：
```
sudo sysctl -w vm.max_map_count=262144
```
### 二、DSL语法学习
#### 常用语句
a. 删除数据
```
POST /app-logs/_delete_by_query
{
          "query": {
            "bool": {
              "filter": {
                "range": {
                  "@timestamp": {
                    "gte": "now/m",
                    "format": "epoch_millis"
                  }
                }
              }
            }
          }
        }
```
b. 查找一个时间段内的满足条件的日志
```
{
          "query": {
            "bool": {
              "must": [
                {
                   "match": {
                      "level": "ERROR"
                    }
                },
                {
                   "match": {
                      "tags": "upayserver"
                    }
                }
                
              ], 
              "filter": {
                "range": {
                  "@timestamp": {
                    "gte": "now-15m/m",
                    "lte": "now/m",
                    "format": "epoch_millis"
                  }
                }
              }
            }
          },
          "size": 0,
          "aggs": {
            "dateAgg": {
              "date_histogram": {
                "field": "@timestamp",
                "time_zone": "Europe/Amsterdam",
                "interval": "1m",
                "min_doc_count": 1
              }
            }
          }
        }
```
c. 创建模板
```
PUT /_template/app-logs-template
{
    "order" : 1,
    "index_patterns" : [
      "app-logs-*"
    ],
    "settings" : {
      "index" : {
        "number_of_shards" : 1,
        "auto_expand_replicas" : "0-1",
        "refresh_interval" : "5s"
      }
    },
    "mappings" : {
      "_default_" : {
        "dynamic_templates" : [
          {
            "message_field" : {
              "path_match" : "message",
              "match_mapping_type" : "string",
              "mapping" : {
                "type" : "text",
                "norms" : false
              }
            }
          },
          {
            "string_fields" : {
              "match" : "*",
              "match_mapping_type" : "string",
              "mapping" : {
                "type" : "text",
                "norms" : false,
                "fields" : {
                  "keyword" : {
                    "type" : "keyword",
                    "ignore_above" : 256
                  }
                }
              }
            }
          }
        ],
        "properties" : {
          "@timestamp" : {
            "type" : "date"
          },
          "@version" : {
            "type" : "keyword"
          },
          "geoip" : {
            "dynamic" : true,
            "properties" : {
              "ip" : {
                "type" : "ip"
              },
              "location" : {
                "type" : "geo_point"
              },
              "latitude" : {
                "type" : "half_float"
              },
              "longitude" : {
                "type" : "half_float"
              }
            }
          }
        }
      }
    },
    "aliases" : { }
}
```

d. 删除指定的索引
```

```


# 性能调优及使用经验
> 1. https://www.cnblogs.com/ibook360/archive/2013/03/15/2961141.html (使用经验分享，包括如何对日志转储)
> 2. 30GB堆内存的节点最多可以有600-750个分片
> 3. https://www.jianshu.com/p/1f67e4436c37 (介绍es模板信息)
> 4. https://www.elastic.co/guide/cn/elasticsearch/guide/current/index.html (es权威指南)
> 5. https://blog.csdn.net/laoyang360/article/details/78080602/ (关于分片的介绍及优化建议)
> 6. https://www.cnblogs.com/sunxucool/p/3799190.html (es配置详解)


集群搭建：
问题1：  failed to flush export bulks ？
解决：https://elasticsearch.cn/question/1915

问题2： 如何驱逐集群中的一个节点？
解决参考：https://blog.csdn.net/laoyang360/article/details/83218266

问题3： 集群状态为red
解决： https://www.jianshu.com/p/ff03bf296de9

问题4： 为什么分片、索引不能太多？
每一个分片就是一个独立自治的lucene

问题5：如何设置分片分配策略？
https://www.cnblogs.com/bonelee/p/7443727.html
```yaml
# 是否开启基于硬盘的分发策略
cluster.routing.allocation.disk.threshold_enabled: true
# 不会分配分片到硬盘使用率高于这个值的节点
cluster.routing.allocation.disk.watermark.low: "85%"
# 如果硬盘使用率高于这个值，则会重新分片该节点的分片到别的节点
cluster.routing.allocation.disk.watermark.high: "90%"
#当前硬盘使用率的查询频率
cluster.info.update.interval: "30s"
# 计算硬盘使用率时，是否加上正在重新分配给其他节点的分片的大小
cluster.routing.allocation.disk.include_relocations: true

```
