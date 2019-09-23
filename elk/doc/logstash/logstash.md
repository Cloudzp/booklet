# LogStash
## 一、logstash安装

### 1.1 下载安装包；
```
$ wget https://artifacts.elastic.co/downloads/logstash/logstash-6.5.4.tar.gz
```
### 1.2 解压并修改配置文件；
```
$ tar -xzvf  logstash-6.5.4.tar.gz
$ cd logstash-6.5.4
```

### 1.3 修改配置文件：
因为我们需要通过logstash将kafka中的日志导入到es中，所以需要以下配置：

```
$ vi config/logstash-sample.conf

input {
 kafka{
   bootstrap_servers => "localhost:9092"
   topics_pattern => "topic.*"
   consumer_threads => 50
   decorate_events => true
   codec => json {
       charset => "UTF-8"
   }
  }
}

filter {
    grok {
       match => {"message" => "%{LOGLEVEL:level}"}
    }
}


output {
  elasticsearch {
    hosts => ["http://localhost:9200"]
    index => "app-logs"
    #user => "elastic"
    #password => "changeme"
  }
}

```

### 1.4 启动服务
```
$ nohup bin/logstash -f config/logstash-sample.conf --config.reload.automatic &
```
- 说明：通过<code>--config.reload.automatic<\code>可以在修改配置之后不需要重启服务即可生效；

## 二、常见问题：
 > 参考：
 - https://blog.csdn.net/qq_32292967/article/details/78622647
 - https://www.elastic.co/guide/en/logstash/current/plugins-inputs-kafka.html#plugins-inputs-kafka-topics (input配置参考官方文档)
 - https://www.elastic.co/guide/en/logstash/current/plugins-outputs-elasticsearch.html#plugins-outputs-elasticsearch-action (output配置参考官方文档)
1. 修改配置后，可否不用重启让配置生效？

   - 使用命令<code>kill -1 {pid}</code>可以热加载配置
   - 也可以是使用如下命令，在启动时让logstash支持热重启；（亲测有效）
   ```
   $ ./bin/lagstash -f configfile.conf --config.reload.automatic
   ```

2. logstash如何处理json格式的日志？
   - 直接设置format => json， 配置如下：
   ```
   file {
           type => "voip_feedback"
           path => ["/usr1/data/voip_feedback.txt"]  
           format => json
           sincedb_path => "/home/jfy/soft/logstash-1.4.2/voip_feedback.access"     
       }
    这种方法亲测不行，错误如下：
    [2019-01-23T16:59:00,297][ERROR][logstash.agent ] Failed to execute action {:action=>LogStash::PipelineAction::Create/pipeline_id:main, :exception=>"LogStash::ConfigurationError", :message=>"The setting `format` in plugin `kafka` is obsolete and is no longer available. You should use the newer 'codec' setting instead. If you have any questions about this, you are invited to visit https://discuss.elastic.co/c/logstash and ask.", :backtrace=>["/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/config/mixin.rb:114:in `block in config_init'", "org/jruby/RubyHash.java:1343:in `each'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/config/mixin.rb:96:in `config_init'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/inputs/base.rb:60:in `initialize'", "org/logstash/plugins/PluginFactoryExt.java:233:in `plugin'", "org/logstash/plugins/PluginFactoryExt.java:166:in `plugin'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/pipeline.rb:71:in `plugin'", "(eval):8:in `<eval>'", "org/jruby/RubyKernel.java:994:in `eval'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/pipeline.rb:49:in `initialize'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/pipeline.rb:90:in `initialize'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/pipeline_action/create.rb:42:in `block in execute'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/agent.rb:92:in `block in exclusive'", "org/jruby/ext/thread/Mutex.java:148:in `synchronize'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/agent.rb:92:in `exclusive'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/pipeline_action/create.rb:38:in `execute'", "/home/cloudzp/elk/logstash-6.5.4/logstash-core/lib/logstash/agent.rb:317:in `block in converge_state'"]}
   ```
   - 使用codec => json （亲测有效）
   ```
    file {
           type => "voip_feedback"
           path => ["/usr1/data/voip_feedback.txt"]  
           sincedb_path => "/home/jfy/soft/logstash-1.4.2/voip_feedback.access"
           codec => json {
               charset => "UTF-8"
           }       
       }
   ```
   - 使用filter json （未测试）
   ```
   filter {
       if [type] == "voip_feedback" {
           json {
               source => "message"
               #target => "doc"
               #remove_field => ["message"]
           }        
       }
   }
   ```
   
3. 如何设置从日志信息中截取日志级别？并添加成一个新的字段？
参考如下配置即可：
```
filter {
   grok {
     // grok 解析后会根据你的语法在整个结构体中添加新字段
     match => {"message" => "%{LOGLEVEL:level}"}
   }
}
```

> 参考：
- https://www.elastic.co/guide/en/logstash/current/config-examples.html (官方配置)
- https://discuss.elastic.co/t/split-row-data-to-fields/42586 (logstash社区)
- https://www.jianshu.com/p/443f1ea7b640 (grok语法)
- http://blog.51cto.com/seekerwolf/2110174 (一个老哥的示例)  
