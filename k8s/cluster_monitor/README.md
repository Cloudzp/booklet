# k8s集群监控

## 一、监控方案：
 各种采集器 +  Prometheus + Grafana 方案,具体如下：
- 使用 kube-state-metrics 采集k8s集群中各种资源对象的状态数据，如 Deployment，Daemonset，StatefulSet 等；
- 使用kubelet自带的cdvisor 采集容器性能数据；
- 使用node-exporter采集k8s节点性能数据；
- 使用blackbox-exporter采集容器应用网络数据；
## 二、 步骤
  采集数据 --> 汇总 --》 展示 -->告警
  
## 三、监控指标采集实现：
> 参考：
- https://blog.csdn.net/liukuan73/article/details/78881008
- https://prometheus.io/docs/prometheus/latest/configuration/configuration/ (prometheus配置集)
## 四、了解prometheus
[点击学习Prometheus](../../Prometheus/Prometheus.md)
